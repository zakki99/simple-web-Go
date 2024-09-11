package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"a21hc3NpZ25tZW50/model"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		cookie, err := c.Cookie("session_token")
		if err != nil {
			if c.GetHeader("Content-Type") == "application/json" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			} else {
				c.Redirect(http.StatusSeeOther, "/login")
			}

			c.Abort()
			return
		}
		tokenClaims := &model.Claims{}
		token, err := jwt.ParseWithClaims(cookie, tokenClaims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("email", tokenClaims.Email) // Menggunakan Email sebagai contoh, sesuaikan dengan field yang ada pada Claims
		c.Next()
	})
}