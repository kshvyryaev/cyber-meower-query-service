package pkg

import (
	"flag"
)

type Config struct {
	Port           string
	ElasticAddress string
}

func ProvideConfig() *Config {
	return &Config{
		Port:           *flag.String("port", "8090", "Server port"),
		ElasticAddress: *flag.String("elasticAddress", "http://127.0.0.1:9200", "Elastic search address"),
	}
}
