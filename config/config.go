package config

import "os"

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

func Get() *Config {
	conf := &Config{}

	conf.DBUser = os.Getenv("MONGODB_USERNAME")
	conf.DBPass = os.Getenv("MONGODB_PASSWORD")
	conf.DBHost = os.Getenv("MONGODB_HOST")
	conf.DBPort = os.Getenv("MONGODB_PORT")
	conf.DBName = os.Getenv("MONGODB_DBNAME")

	return conf
}
