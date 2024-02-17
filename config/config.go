package config

import (
	"encoding/json"
	"fmt"
	"os"

	"sternx/infrastructure/util"
)

const (
	FileName = "config.json"
)

type Config struct {
	GRPCListen        string `json:"grpc_listen"`
	GRPCGatewayListen string `json:"grpc_gateway_listen"`
	LogLevel          string `json:"log_level"`
	DSN               string `json:"dsn"`
	// it would be better to generate random salt for each user and store it in db
	PasswordSalt []byte `json:"-"`
}

func Default() (*Config, error) {
	cfg := &Config{
		GRPCListen:        "127.0.0.1:50051",
		GRPCGatewayListen: "127.0.0.1:8080",
		LogLevel:          "error",
		DSN: "host=localhost user=postgres password=postgres dbname=postgres port=5432" +
			" sslmode=disable TimeZone=UTC",
		PasswordSalt: []byte("sternx"),
	}

	data, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf("could not read config file error: %s", err.Error())
	}

	tempCfg := &Config{}
	if err := json.Unmarshal(data, tempCfg); err != nil {
		return nil, fmt.Errorf("could not unmarshal config file error: %s", err.Error())
	}

	if !util.IsStringEmpty(tempCfg.GRPCListen) {
		cfg.GRPCListen = tempCfg.GRPCListen
	}

	if !util.IsStringEmpty(tempCfg.GRPCGatewayListen) {
		cfg.GRPCGatewayListen = tempCfg.GRPCGatewayListen
	}

	if !util.IsStringEmpty(tempCfg.LogLevel) {
		cfg.LogLevel = tempCfg.LogLevel
	}

	if !util.IsStringEmpty(tempCfg.DSN) {
		cfg.DSN = tempCfg.DSN
	}

	return cfg, nil
}
