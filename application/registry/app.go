package registry

import (
	"github.com/raismaulana/digilibP/application"
	"github.com/raismaulana/digilibP/controller"
	"github.com/raismaulana/digilibP/infrastructure/auth"
	"github.com/raismaulana/digilibP/infrastructure/database"
	"github.com/raismaulana/digilibP/infrastructure/env"
	"github.com/raismaulana/digilibP/infrastructure/log"
	"github.com/raismaulana/digilibP/infrastructure/server"
	"gorm.io/gorm"
)

type app struct {
	server.GinHTTPHandler
	controller.BaseController
	// TODO Another controller will added here ... <<<<<<
}

func NewApp() func() application.RegistryContract {
	return func() application.RegistryContract {
		// setup logger write file
		log.UseRotateFile(".log", "default", 30)

		// setup db
		db := database.NewGormPostgres(&gorm.Config{})

		enforcer := auth.NewCasbinEnforcerByDB(db)

		// setup server
		httpHandler := server.NewGinHTTPHandler(env.Var().AppPort)

		return &app{
			GinHTTPHandler: httpHandler,
			BaseController: controller.BaseController{
				Enforcer: enforcer,
			},
			// TODO another controller will added here ... <<<<<<
		}

	}
}

func (r *app) SetupController() {

	// TODO another router call will added here ... <<<<<<
}
