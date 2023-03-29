package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	// Database config
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`

	// Elastic Search config
	ESUser            string `mapstructure:"ELASTIC_USER"`
	ESPassword        string `mapstructure:"ELASTIC_PASSWORD"`
	ESCertFingerPrint string `mapstructure:"CERT_FINGERPRINT"`

	// Url client
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app_example")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
