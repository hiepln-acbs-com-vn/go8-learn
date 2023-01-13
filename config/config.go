package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Api           Api
	Database      Database
	Cache         Cache
	Elasticsearch Elasticsearch
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	return &Config{
		Api:           API(),
		Database:      DataStore(),
		Cache:         NewCache(),
		Elasticsearch: ElasticSearch(),
	}
}

func Load(file string) *Config {
	err := godotenv.Load(file)
	if err != nil {
		log.Println(err)
	}

	return &Config{
		Api:           API(),
		Database:      DataStore(),
		Cache:         NewCache(),
		Elasticsearch: ElasticSearch(),
	}
}
