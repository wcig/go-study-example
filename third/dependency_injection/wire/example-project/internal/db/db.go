package db

import (
	"context"
	"go-app/third/dependency_injection/wire/example-project/internal/config"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var ProviderSet = wire.NewSet(NewDatabase, NewRedis, NewData)

type Data struct {
	DB       *gorm.DB
	RedisCli *redis.Client
}

func NewData(c *config.Data, logger *zap.SugaredLogger, db *gorm.DB, redisCli *redis.Client) *Data {
	return &Data{
		DB:       db,
		RedisCli: redisCli,
	}
}

func NewDatabase(c *config.Data, logger *zap.SugaredLogger) *gorm.DB {
	cfg := &gorm.Config{Logger: gormLogger.Discard}
	if c.Database.ShowSql {
		cfg.Logger = gormLogger.Default
	}
	db, err := gorm.Open(mysql.Open(c.Database.Dsn), cfg)
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(c.Database.MaxIdle)
	sqlDB.SetMaxOpenConns(c.Database.MaxOpen)
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
	return db
}

func NewRedis(c *config.Data, logger *zap.SugaredLogger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return client
}
