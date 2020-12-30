package internal

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	Port int `default:"8772"`
}

var conf Configuration

func init() {
	err := envconfig.Process("SETTINGS", &conf)
	if err != nil {
		panic(err)
	}
}
