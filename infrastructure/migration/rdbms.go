package migration

import (
	"context"

	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/envconfig"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/log"
	"gorm.io/gorm"
)

func RDBMSMigration(ctx context.Context, db *gorm.DB, env *envconfig.EnvConfig) error {
	log.Info(ctx, "Migrate RDBMS")

	if err := db.AutoMigrate(); err != nil {
		return err
	}
	// this transaction will always make user default super user is exsist
	if err := db.Transaction(func(tx *gorm.DB) error {
		return nil
	}); err != nil {
		log.Error(ctx, err.Error())
	}

	return nil
}
