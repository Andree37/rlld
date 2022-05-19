package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct

func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err := config.ReadInConfig()

	fmt.Printf("%v\n", env)

	path, err := os.Getwd()
	// handle err
	fmt.Printf("%v\n", path)

	if err != nil {
		log.Fatalf("error on parsing configuration file")
	}
}

func Getconfig() *viper.Viper {
	return config
}
