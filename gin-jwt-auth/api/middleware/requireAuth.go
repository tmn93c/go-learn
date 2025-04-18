package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-jwt-auth/db/initializers"
	"gin-jwt-auth/internal/models"
	"gin-jwt-auth/logger"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthUser struct {
	ID    uint   `json:"ID"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

func RequireAuth(c *gin.Context) {
	// Get the cookie from the request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// validate with redis cache
	email, errorRedis := initializers.RedisClient.Get(c, tokenString).Result()

	if errorRedis != nil || email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("request_id", uuid.New().String())

	// anywhere in your code
	logger.Info(c, "User request", zap.String("user_id", email))

	// Decode and validate it
	// Parse and takes the token string and a function for look
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration time
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		SetUserByEmail(c, email, claims["sub"].(float64))
		// Find the user with token sub

		// authUser := AuthUser{
		// 	ID:    user.ID,
		// 	Name:  user.Name,
		// 	Email: user.Email,
		// }

		// Attach the user to request
		// c.Set("authUser", authUser)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// Cache aside pattern (pseudo-code)

func SetUserByEmail(c *gin.Context, email string, id float64) {
	key := email

	// 1. Try cache
	userJson, _ := initializers.RedisClient.Get(c, key).Result()

	if userJson == "" {
		// 2. Fallback to DB
		var user models.User
		initializers.DB.Find(&user, id)

		// 3. Cache result
		jsonBytes, _ := json.Marshal(user)
		initializers.RedisClient.Set(context.TODO(), key, jsonBytes, time.Minute*10)
	}

}
