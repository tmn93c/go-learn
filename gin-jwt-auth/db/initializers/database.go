package initializers

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
)

// ConnectDB connects to Postgres
func ConnectDB() {
	var err error
	dsn := os.Getenv("DNS")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("❌ Failed to connect to database: %v", err))
	}
	fmt.Println("✅ Connected to Postgres")
}

// ConnectRedis connects to Redis and pings to ensure it's up
func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"), // e.g. "localhost:6379"
	})

	// Check Redis connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("❌ Failed to connect to Redis: %v", err))
	}
	fmt.Println("✅ Connected to Redis")
}
