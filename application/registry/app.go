package registry

import (
	"context"
	"fmt"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/raismaulana/my-gogen-rest-setup/application"
	"github.com/raismaulana/my-gogen-rest-setup/gateway/master"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/auth"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/envconfig"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/log"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type app struct {
	server.GinHTTPHandler
	// TODO Another controller will added here ... <<<<<<
}

var (
	ctx = context.Background()
)

func NewApp() func() application.RegistryContract {
	return func() application.RegistryContract {

		env := setupEnv()
		jwtToken := setupJWTToken(env)
		db := setupDB(env)
		rdb := setupRedis(env)
		httpHandler := setupHTTPHandler(env)
		enforcer := setupCasbinEnforcer()

		datasource, err := master.NewMasterGateway(ctx, env, db, rdb, jwtToken)
		if err != nil {
			log.Error(ctx, "%v", err.Error())
			os.Exit(1)
		}

		return &app{
			GinHTTPHandler: httpHandler,
			// TODO another controller will added here ... <<<<<<
		}

	}
}

func (r *app) SetupController() {
	// TODO another router call will added here ... <<<<<<
}

func setupEnv() *envconfig.EnvConfig {
	env, err := envconfig.NewEnvConfig()
	if err != nil {
		log.Error(ctx, "Config Problem %v", err.Error())
		os.Exit(1)
	}
	return env
}

func setupJWTToken(env *envconfig.EnvConfig) *auth.JWTToken {
	jwtToken, err := auth.NewJWTToken(env)
	if err != nil {
		log.Error(context.Background(), "Secret Key Problem %v", err.Error())
		os.Exit(1)
	}
	return jwtToken
}

func setupDB(env *envconfig.EnvConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func setupRedis(env *envconfig.EnvConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost + ":" + env.RedisPort,
		Password: env.RedisPassword,
		DB:       env.RedisDB,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error(ctx, "%v", err.Error())
		os.Exit(1)
	}
	return rdb
}

func setupHTTPHandler(env *envconfig.EnvConfig) server.GinHTTPHandler {
	httpHandler, err := server.NewGinHTTPHandler(":" + env.AppPort)
	if err != nil {
		log.Error(ctx, "%v", err.Error())
		os.Exit(1)
	}
	return httpHandler
}

func setupCasbinEnforcer() *casbin.Enforcer {
	e, err := casbin.NewEnforcer("infrastructure/auth/casbin_model.conf", "infrastructure/auth/casbin_policy.csv")
	if err != nil {
		log.Error(ctx, "%v", err.Error())
		os.Exit(1)
	}
	return e
}
