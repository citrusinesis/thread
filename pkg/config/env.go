package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type DBConfig struct {
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	Database string `mapstructure:"DB_DATABASE"`
}

func NewDBConfig() (value DBConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read app.env file: %v", err)
	}
	if err := viper.Unmarshal(&value); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return
}

func (db *DBConfig) CreateURL() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=True",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Database,
	)
}
