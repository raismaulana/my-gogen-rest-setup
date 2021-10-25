package repository

import "context"

// WithoutTransactionDB used to get database object from any database implementation.
// For consistency reason both WithTransactionDB and WithoutTransactionDB will seek database object under the context params
type WithoutTransactionDB interface {
	GetDatabase(ctx context.Context) (context.Context, error)
}

// WithTransactionDB used for common transaction handling
// all the context must use the same database session.
type WithTransactionDB interface {
	BeginTransaction(ctx context.Context) (context.Context, error)
	CommitTransaction(ctx context.Context) error
	RollbackTransaction(ctx context.Context) error
}

// WithoutTrx is helper function that simplify the readonly db
func WithoutTrx(ctx context.Context, trx WithoutTransactionDB, trxFunc func(dbCtx context.Context) error) error {
	dbCtx, err := trx.GetDatabase(ctx)
	if err != nil {
		return err
	}
	return trxFunc(dbCtx)
}

// WithTrx is helper function that simplify the transaction execution handling
func WithTrx(ctx context.Context, trx WithTransactionDB, trxFunc func(dbCtx context.Context) error) error {
	dbCtx, err := trx.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			err = trx.RollbackTransaction(dbCtx)
			panic(p)

		} else if err != nil {
			err = trx.RollbackTransaction(dbCtx)

		} else {
			err = trx.CommitTransaction(dbCtx)

		}
	}()

	err = trxFunc(dbCtx)
	return err
}
