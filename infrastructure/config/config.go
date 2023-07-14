package config

import (
	"os"

	"github.com/spf13/viper"
)

type(
	AppConfig struct{
		Application Application
		Redis Redis
		Aws Aws
	}

	Application struct{
		Name string
		LevelLog string
		Server Server
	}

	Redis struct{
		Host string
	}

	Aws struct{
		Endpoint string
		Region string
		Dynamodb Dynamodb
	}

	Server struct{
		Port string
		Timeout string
	}

	Dynamodb struct{
		TableName string
	}
)

func NewViperConfig() *viper.Viper {
	os.Setenv("ENVIRONMENT","local")
	var env = os.Getenv("ENVIRONMENT")
	viper.SetConfigName("config."+env)
	viper.AddConfigPath("./infrastructure/config/env")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	viper.SetDefault("log_level","debug")
	viper.SetDefault("log_format","json")

	return viper.GetViper()
}

func LoadConfig() (*AppConfig, error){
	var appConfig AppConfig

	viper := NewViperConfig()
	if err:= viper.ReadInConfig(); err!= nil{
		return &AppConfig{},err
	}

	if err := viper.Unmarshal(&appConfig); err != nil{
		return &AppConfig{},err
	}

	return &appConfig, nil
}


