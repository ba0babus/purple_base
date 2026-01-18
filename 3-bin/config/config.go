package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	return &Config{
		Key: os.Getenv("KEY"),
	}
}
