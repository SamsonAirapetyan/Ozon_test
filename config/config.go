package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	PostgresDB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbname"`
		SSLmode  string `yaml:"sslmode"`
		MaxConns string `yaml:"maxconns"`
	}
	Grpc struct {
		Network string `yaml:"network"`
		Address string `yaml:"address"`
	}
}

func ConfigViper() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("Can not read config", err.Error())
		os.Exit(1)
	}
	return v
}

func ParseConfig(v *viper.Viper) *Config {
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		fmt.Println("Can not unmarshal config")
	}
	return cfg
}
