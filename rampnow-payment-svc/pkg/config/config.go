package config

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD",
}

func LoadConfig() (Config, error) {

	// TODO: read from .env variable
	config :=  Config{
		DBHost: "postgres",
		DBName: "rampnowdb",
		DBUser: "postgres",
		DBPort: "5432",
		DBPassword: "postgres",
	}
	
	return config, nil
}
