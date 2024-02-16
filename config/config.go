package config

type Config struct {
	GRPCListen        string
	GRPCGatewayListen string
	LogLevel          string
}

func Default() *Config {
	return &Config{
		GRPCListen:        "127.0.0.1:50051",
		GRPCGatewayListen: "127.0.0.1:8080",
		LogLevel:          "info",
	}
}
