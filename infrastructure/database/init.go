package database

import (
	"fmt"

	"example/infrastructure/env"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initGorm(dialector gorm.Dialector, config *gorm.Config) *gorm.DB {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// NewGormDefault ...
func NewGormDefault() *gorm.DB {
	return initGorm(sqlite.Open("wallet.db"), &gorm.Config{})
}

func NewGormPostgres() *gorm.DB {
	var config gorm.Config
	if env.Var().Production {
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.Var().DBHost,
		env.Var().DBUser,
		env.Var().DBPassword,
		env.Var().DBName,
		env.Var().DBPort,
	)
	return initGorm(postgres.Open(dsn), &config)
}
