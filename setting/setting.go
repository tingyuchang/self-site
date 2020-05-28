package setting

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server ServerConfiguration
	DB     DatabaseConfiguration
	APP    AppConfiguration
}

type DatabaseConfiguration struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
	Driver string
}

type ServerConfiguration struct {
	RunMode string
	Port    int
	Timeout int
}

type AppConfiguration struct {
	Version   int
	PageSize  int
	JwtSecret string
}

var Config Configuration

func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal(err)
	}
}
