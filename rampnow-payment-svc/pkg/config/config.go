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
	config :=  Config{
		DBHost: "postgres",
		DBName: "rampnowdb",
		DBUser: "postgres",
		DBPort: "5432",
		DBPassword: "postgres",
	}
	

	// viper.AddConfigPath("./")
	// viper.SetConfigFile(".env")
	// viper.ReadInConfig()

	// for _, env := range envs {
	// 	if err := viper.BindEnv(env); err != nil {
	// 		return config, err
	// 	}
	// }

	// if err := viper.Unmarshal(&config); err != nil {
	// 	return config, err
	// }

	// if err := validator.New().Struct(&config); err != nil {
	// 	return config, err
	// }

	return config, nil
}
