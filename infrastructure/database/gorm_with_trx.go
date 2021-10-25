package database

import (
	"context"

	"example/infrastructure/log"

	"gorm.io/gorm"
)

// GormWithTrxImpl ...
type GormWithTrxImpl struct {
	DB *gorm.DB
}

// NewGormWithTrxImpl ...
func NewGormWithTrxImpl(db *gorm.DB) *GormWithTrxImpl {
	return &GormWithTrxImpl{DB: db}
}

// BeginTransaction ...
func (r *GormWithTrxImpl) BeginTransaction(ctx context.Context) (context.Context, error) {

	dbTrx := r.DB.Begin()

	trxCtx := context.WithValue(ctx, contextDBValue, dbTrx)

	return trxCtx, nil
}

// CommitTransaction ...
func (r *GormWithTrxImpl) CommitTransaction(ctx context.Context) error {
	log.Info(ctx, "Commit")

	db, err := ExtractDB(ctx)
	if err != nil {
		return err
	}

	return db.Commit().Error
}

// RollbackTransaction ...
func (r *GormWithTrxImpl) RollbackTransaction(ctx context.Context) error {
	log.Info(ctx, "Rollback")

	db, err := ExtractDB(ctx)
	if err != nil {
		return err
	}

	return db.Rollback().Error
}
