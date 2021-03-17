package controllers

import (
	"log"
	"time"
	"to-do-list/models"
	logger "to-do-list/sentry"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Email    string
	FullName string
	Role     string
}

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func (StrDB *StrDB) MiddleWare() (mw *jwt.GinJWTMiddleware) { // the jwt middleware

	if err := godotenv.Load(".env"); err != nil {
		log.Println("ENV File Not Found!")
	}
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: "",
		IdentityKey:      identityKey,
		Key:              []byte("dimdim1223"),
		Timeout:          time.Minute * 15,
		MaxRefresh:       time.Minute * 11,
		Authenticator: func(c *gin.Context) (interface{}, error) {

			var (
				loginVals login
				user      models.User
			)
			if err := c.ShouldBind(&loginVals); err != nil {
				logger.Sentry(err)
				return "", jwt.ErrMissingLoginValues
			}

			StrDB.DB.Where("email = ? ", loginVals.Email).First(&user)

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password)); err != nil {
				log.Println(user)
				log.Println(user.Password)
				log.Println(loginVals.Password)
				log.Println("Password does not match!")
			} else {
				return &user, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			method := c.Request.Method
			if v, ok := data.(*user); ok && v.Role == "admin" {
				return true
			}

			if v, ok := data.(*user); ok && v.Role == "guest" {
				if method != "GET" {
					return false
				}
				return true
			}

			return false
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*login); ok {
				return jwt.MapClaims{identityKey: v.Email}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"code": code, "message": message})
		},
		TokenLookup:       "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:     "Bearer",
		TimeFunc:          time.Now,
		PrivKeyFile:       "",
		PubKeyFile:        "",
		SendCookie:        false,
		CookieMaxAge:      0,
		SecureCookie:      false,
		CookieHTTPOnly:    false,
		CookieDomain:      "",
		SendAuthorization: false,
		DisabledAbort:     false,
		CookieName:        "",
		CookieSameSite:    0,
	}) // close authMidlleware

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	// fmt.Println("ini param valuenya : ", mw)
	// fmt.Println("ini returnya : ", authMiddleware)
	return authMiddleware
}
