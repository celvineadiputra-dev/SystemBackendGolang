package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"mvcGolang/app/Helpers"
	"mvcGolang/app/Models/Users"
	"mvcGolang/config"
	"strings"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := Helpers.ApiResponse(401, "Unauthorized", nil)
			c.AbortWithStatusJSON(401, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := config.ValidateToken(tokenString)

		if err != nil {
			response := Helpers.ApiResponse(401, "Unauthorized", nil)
			c.AbortWithStatusJSON(401, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := Helpers.ApiResponse(401, "Unauthorized", nil)
			c.AbortWithStatusJSON(401, response)
			return
		}

		userIdEncrypt := claim["user_id"].(string)
		userIdDecrypt := Helpers.Decrypt(userIdEncrypt)

		var user Users.User

		err = db.Where("id = ?", userIdDecrypt).Find(&user).Error

		if err != nil {
			response := Helpers.ApiResponse(401, "Unauthorized", nil)
			c.AbortWithStatusJSON(401, response)
			return
		}

		c.Set("currentUser", user)
	}
}
