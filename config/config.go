package config

type Config struct {
	GRPCListen        string
	GRPCGatewayListen string
	LogLevel          string
	DSN               string
	PasswordSalt      []byte // it would be better to generate random salt for each user and store it in db
}

func Default() *Config {
	return &Config{
		GRPCListen:        "127.0.0.1:50051",
		GRPCGatewayListen: "127.0.0.1:8080",
		LogLevel:          "trace",
		DSN: "host=localhost user=postgres password=postgres dbname=postgres port=5432" +
			" sslmode=disable TimeZone=UTC",
		PasswordSalt: []byte("sternx"),
	}
}
