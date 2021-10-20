package auth

import (
	"context"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/raismaulana/digilibP/infrastructure/log"
	"gorm.io/gorm"
)

// NewCasbinEnforcer is
func NewCasbinEnforcer() *casbin.Enforcer {

	e, err := casbin.NewEnforcer("infrastructure/auth/casbin_model.conf", "infrastructure/auth/casbin_policy.csv")
	if err != nil {
		log.Error(context.Background(), "%v", err.Error())
		os.Exit(1)
	}
	return e
}

func NewCasbinEnforcerByDB(db *gorm.DB) *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Error(context.Background(), "%v", err.Error())
		os.Exit(1)
	}
	e, err := casbin.NewEnforcer("infrastructure/auth/casbin_model.conf", a)
	if err != nil {
		log.Error(context.Background(), "%v", err.Error())
		os.Exit(1)
	}

	e.LoadPolicy()

	e.SavePolicy()
	return e
}
