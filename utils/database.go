package utils

import (
	"context"
	"os"

	// "github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pawdot.app/models"
)

type DatabaseConnection interface {
	GetDB() *gorm.DB
	// GetRDB() *redis.Client
}

type databaseConnection struct {
	DB *gorm.DB
	// RDB *redis.Client
}

// GetRDB implements DatabaseConnection
// func (d *databaseConnection) GetRDB() *redis.Client {
// 	return d.RDB
// }

// GetDB implements DatabaseConnection
func (d *databaseConnection) GetDB() *gorm.DB {
	return d.DB
}

func InitDatabaseConnection() DatabaseConnection {
	ctx := context.Background()
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		AppError(err, "there was an error connecting to database")
	}

	// auto migrating tables
	db.AutoMigrate(
		&models.User{},
		&models.Pet{},
	)

	defer ctx.Done()

	// redisUrl := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	//
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: redisUrl,
	// })
	//
	// if err := rdb.Ping(ctx).Err(); err != nil {
	// 	AppError(err, "there was an error connecting to redis")
	// }

	return &databaseConnection{
		DB: db,
		// RDB: rdb,
	}
}
