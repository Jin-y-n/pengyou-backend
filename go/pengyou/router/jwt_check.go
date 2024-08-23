package router

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"pengyou/constant"
	"pengyou/model/common/response"
	"pengyou/utils/common"
	"strings"
)

// Define a simple JWT validation function.
func jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header.
		authHeader := c.GetHeader(constant.TokenName)
		if authHeader == "" {
			response.NoAuth(constant.NoAuthority, c)
			return
		}

		//tokenParts := strings.Split(authHeader, " ")
		//if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		//	response.NoAuth(constant.NoAuthority, c)
		//	return
		//}

		tokenParts := strings.Split(authHeader, ".")
		if len(tokenParts) != 3 {
			response.NoAuth(constant.NoAuthority, c)
			return
		}

		// Validate the token.
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			// Check if the token uses the expected signing method.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Replace 'yourSecretKey' with your actual secret key used for signing the tokens.
			return []byte("yourSecretKey"), nil
		})

		if err != nil {
			response.NoAuth(constant.NoAuthority, c)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Token is valid
			userID := claims["user_id"].(string)

			common.SetTokenInContext(context.Background(), userID)

			// Continue processing the request.
			c.Next()
		} else {
			response.NoAuth(constant.NoAuthority, c)
			return
		}
		return
	}
}
