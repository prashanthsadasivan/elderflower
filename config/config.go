package config

import (
	"os"
)

var (
	config map[string]string
)

func init() {
	config = make(map[string]string)
	config["driver"] = "postgres"
	config["spec"] = "dbname=smswebproxy sslmode=disable"
}

func Get(key string) string {
	ret := os.Getenv(key)
	if ret == "" {
		ret = config[key]
	}
	return ret
}
