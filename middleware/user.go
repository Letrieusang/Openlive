package middleware

import (
	"api-openlive/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Authenticate middleware
func Authenticate(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		bearerToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenStr = strArr[1]
		}

		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"hasError": 1,
				"message":  "Token could not found!",
			})
			return
		}

		token, err := utils.TokenValid(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"hasError": 1,
				"message":  "Invalid Token",
			})
			return
		}

		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			now := time.Now().Unix()
			fmt.Println("user_id", claims["user_id"])
			id, err := strconv.Atoi(fmt.Sprint(claims["user_id"]))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": http.StatusUnauthorized,
					"message":  "can not get user id",
				})
				return
			}

			c.Set("id", id)
			ut := claims["user_type"]
			c.Set("user_type", userType)
			if ut != userType {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": "1",
					"message":  "user type invalid",
				})
				return
			}
			exp := cast.ToInt64(claims["exp"])
			if now > exp {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": "1",
					"message":  "Token has expired",
				})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"hasError": http.StatusUnauthorized,
				"message":  "Invalid Token",
			})
		}

	}
}

func AuthenticateOrNot(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		bearerToken := c.GetHeader("Authorization")
		if len(bearerToken) == 0 {
			return
		}
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenStr = strArr[1]
		}

		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"hasError": 1,
				"message":  "Token could not found!",
			})
			return
		}

		token, err := utils.TokenValid(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"hasError": 1,
				"message":  "Invalid Token",
			})
			return
		}

		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			now := time.Now().Unix()
			fmt.Println("user_id", claims["user_id"])
			id, err := strconv.Atoi(fmt.Sprint(claims["user_id"]))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": http.StatusUnauthorized,
					"message":  "can not get user id",
				})
				return
			}

			c.Set("id", id)
			ut := claims["user_type"]
			c.Set("user_type", userType)
			if ut != userType {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": "1",
					"message":  "user type invalid",
				})
				return
			}
			exp := cast.ToInt64(claims["exp"])
			if now > exp {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": "1",
					"message":  "Token has expired",
				})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"hasError": http.StatusUnauthorized,
				"message":  "Invalid Token",
			})
		}

	}
}

// AuthenticateFalse middleware
func AuthenticateFalse(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string
		bearerToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenStr = strArr[1]
		}
		if tokenStr == "" {
			return
		}
		token, err := utils.TokenValid(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"hasError": 1,
				"message":  "Invalid Token",
			})
			return
		}

		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			now := time.Now().Unix()
			fmt.Println("user_id", claims["user_id"])
			id, err := strconv.Atoi(fmt.Sprint(claims["user_id"]))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": http.StatusUnauthorized,
					"message":  "can not get user id",
				})
				return
			}

			c.Set("id", id)
			ut := claims["user_type"]
			c.Set("user_type", userType)
			if ut != userType {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": "1",
					"message":  "user type invalid",
				})
				return
			}
			exp := cast.ToInt64(claims["exp"])
			if now > exp {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"hasError": "1",
					"message":  "Token has expired",
				})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"hasError": http.StatusUnauthorized,
				"message":  "Invalid Token",
			})
		}

	}
}
