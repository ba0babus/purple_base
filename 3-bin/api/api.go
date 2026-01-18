package api

import (
	"fmt"
	"purple_base/bin/config"
)

func GetKey() {
	cfg := config.NewConfig()
	fmt.Println(cfg.Key)
}
