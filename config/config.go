package config

import (
	"log"
	"os"
)

var (
	config map[string]string
)

func init() {
	config = make(map[string]string)
	config["driver"] = "postgres"
	config["spec"] = "dbname=smswebproxy sslmode=disable"
	config["PORT"] = "2020"
}

func Get(key string) string {
	ret := os.Getenv(key)
	log.Printf("env for %s was %s\n", key, ret)
	if ret == "" {
		ret = config[key]
	}
	return ret
}
