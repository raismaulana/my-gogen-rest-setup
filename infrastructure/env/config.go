package env

import (
	"context"
	"os"
	"sync"

	"github.com/raismaulana/digilibP/infrastructure/log"
	"github.com/spf13/viper"
)

var (
	singleton sync.Once
	config    Config
	path      string = "."
	file      string = "default-config"
)

// Config is struct from env variables
type Config struct {
	Production    bool   `mapstucture:"production"`      //
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

func SetFile(filePath, fileName string) {
	path = filePath
	file = fileName
}

// InitConfig is fuction to initialize viper from reading env variables
func Var() Config {
	singleton.Do(func() {

		viper.SetConfigName(file)
		viper.AddConfigPath(path)

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Error(context.Background(), err.Error(), err)
			os.Exit(1)
		}

		if err := viper.Unmarshal(&config); err != nil {
			log.Error(context.Background(), err.Error(), config)
			os.Exit(1)
		}

		log.Info(context.Background(), "env", config)
	})
	return config
}

// GetString is function to read string env value
func GetString(key string, defaultValue string) string {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return viper.GetString(key)
}

// GetInt is function to read int env value
func GetInt(key string, defaultValue int) int {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return viper.GetInt(key)
}

// GetBool is function to read bool env value
func GetBool(key string, defaultValue bool) bool {
	if !viper.IsSet(key) {
		return defaultValue
	}
	return viper.GetBool(key)
}
