package registry

import (
	"example/application"
	"example/controller"
	"example/infrastructure/auth"
	"example/infrastructure/database"
	"example/infrastructure/env"
	"example/infrastructure/log"
	"example/infrastructure/server"
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
		db := database.NewGormPostgres()

		enforcer := auth.NewCasbinEnforcerByDB(db)

		// setup server
		httpHandler := server.NewGinHTTPHandler(env.Var().AppPort)

		return &app{
			GinHTTPHandler: httpHandler,
			BaseController: controller.BaseController{
				Router:   httpHandler.Router,
				Enforcer: enforcer,
			},
			// TODO another controller will added here ... <<<<<<
		}

	}
}

func (r *app) SetupController() {

	// TODO another router call will added here ... <<<<<<
}
