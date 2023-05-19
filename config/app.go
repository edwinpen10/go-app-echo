package config

import (
	"os"
)

var (
	AppName = "golang-echo"
	AppEnv  = os.Getenv("APP_ENV")
	AppPort = os.Getenv("APP_PORT")
)
