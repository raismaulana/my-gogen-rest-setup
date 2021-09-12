package envconfig

import (
	"context"

	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/log"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/util"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	SuperUsername string `mapstructure:"super_username"` //
	SuperPassword string `mapstructure:"super_password"` //
	DBHost        string `mapstructure:"db_host"`        //
	DBPort        string `mapstructure:"db_port"`        //
	DBName        string `mapstructure:"db_name"`        //
	DBUser        string `mapstructure:"db_user"`        //
	DBPassword    string `mapstructure:"db_password"`    //
	AppName       string `mapstructure:"app_name"`       //
	AppBaseURL    string `mapstructure:"base_url"`       //
	AppBaseURLV1  string `mapstructure:"base_url_v1"`    //
	AppPort       string `mapstructure:"port"`           //
	SecretKey     string `mapstructure:"secretkey"`      //
	SMTPHost      string `mapstructure:"smtp_host"`      //
	SMTPPort      int    `mapstructure:"smtp_port"`      //
	SMTPSender    string `mapstructure:"smtp_sender"`    //
	SMTPEmail     string `mapstructure:"smtp_email"`     //
	SMTPPassword  string `mapstructure:"smtp_password"`  //
	RedisHost     string `mapstructure:"redis_host"`     //
	RedisPort     string `mapstructure:"redis_port"`     //
	RedisPassword string `mapstructure:"redis_password"` //
	RedisDB       int    `mapstructure:"redis_db"`       //
}

func NewEnvConfig() (*EnvConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var env EnvConfig
	err = viper.Unmarshal(&env)
	if err != nil {
		return nil, err
	}

	log.Info(context.Background(), util.MustJSON(env))
	return &env, nil
}
