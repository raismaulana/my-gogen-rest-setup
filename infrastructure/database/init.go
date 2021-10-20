package database

import (
	"fmt"

	"github.com/raismaulana/digilibP/infrastructure/env"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func NewGormPostgres(config *gorm.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.Var().DBHost,
		env.Var().DBUser,
		env.Var().DBPassword,
		env.Var().DBName,
		env.Var().DBPort,
	)
	return initGorm(postgres.Open(dsn), config)
}
