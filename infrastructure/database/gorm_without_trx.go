package database

import (
	"context"

	"gorm.io/gorm"
)

// GormWithoutTrxImpl ...
type GormWithoutTrxImpl struct {
	DB *gorm.DB
}

// NewGormWithoutTrxImpl ...
func NewGormWithoutTrxImpl(db *gorm.DB) *GormWithoutTrxImpl {
	return &GormWithoutTrxImpl{DB: db}
}

// GetDatabase ...
func (r *GormWithoutTrxImpl) GetDatabase(ctx context.Context) (context.Context, error) {
	trxCtx := context.WithValue(ctx, contextDBValue, r.DB)
	return trxCtx, nil
}
