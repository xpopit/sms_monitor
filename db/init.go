package db

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Ctx context.Context
	Rdc *redis.Client
	Pdb *gorm.DB
}

func Init() *DB {
	var ctx = context.Background()
	var redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
	})
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &DB{ctx, redisClient, db}
}
