package configs

type Config struct {
	Port        string `env:"PORT"`
	PostgresURL string `env:"POSTGRES_URL"`
}
