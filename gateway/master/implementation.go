package master

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/raismaulana/my-gogen-rest-setup/gateway/rdbms"
	redisGateway "github.com/raismaulana/my-gogen-rest-setup/gateway/redis"
	"github.com/raismaulana/my-gogen-rest-setup/gateway/shared"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/auth"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/envconfig"
	"gorm.io/gorm"
)

type masterGateway struct {
	rdbms.RDBMSGateway
	redisGateway.RedisGateway
	shared.SharedGateway
}

func NewMasterGateway(ctx context.Context, env *envconfig.EnvConfig, db *gorm.DB, rdb *redis.Client, jwtToken *auth.JWTToken) (*masterGateway, error) {
	rdbmsG, err := rdbms.NewRDBMSGateway(ctx, env, db)
	if err != nil {
		return nil, err
	}

	redisG := redisGateway.NewRedisGateway(rdb)
	sharedG := shared.NewSharedGateway(env, jwtToken)
	return &masterGateway{
		RDBMSGateway:  *rdbmsG,
		RedisGateway:  *redisG,
		SharedGateway: *sharedG,
	}, nil
}
