package middlewares

import (
	"banter/responses"
	"banter/utils/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.Unauthorized(c, "Missing auth header", "Authorization header is missing")
			c.Abort()
			return
		}

		// Ensure the token is prefixed with 'Bearer '
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			responses.Unauthorized(c, "Invalid auth header format", "Authorization header must be in 'Bearer <token>' format")
			c.Abort()
			return
		}

		tokenString := authParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token uses the expected signing method (HMAC in this case)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.Configs.Jwt.Secret), nil
		})

		if err != nil || !token.Valid {
			responses.Unauthorized(c, "Invalid token", err.Error())
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			responses.Unauthorized(c, "Invalid token claims", "Unable to parse token claims")
			c.Abort()
			return
		}

		// Extract user information or other claims from the token
		userID, ok := claims["user_id"].(string) // Use user_id instead of username for generalization
		if !ok {
			responses.Unauthorized(c, "Invalid user claim", "User ID claim in token is invalid")
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", userID)

		// Continue with the request processing
		c.Next()
	}
}
