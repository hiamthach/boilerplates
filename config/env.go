package config

type AppEnv struct {
	App struct {
		Name string `env:"APP_NAME"`
	}

	Server struct {
		Port       string `env:"HTTP_PORT"`
		Cors       string `env:"HTTP_CORS"`
		CorsHeader string `env:"HTTP_CORS_HEADER"`
		CorsMethod string `env:"HTTP_CORS_METHOD"`
	}

	MongoDb struct {
		Connection string `env:"MONGODB_CON"`
		DbName     string `env:"MONGODB_NAME"`
	}

	Redis struct {
		Host     string `env:"REDIS_HOST"`
		Port     string `env:"REDIS_PORT"`
		Password string `env:"REDIS_PASSWORD"`
	}

	Log struct {
		MinLevel int    `env:"LOG_MIN_LEVEL"`
		Path     string `env:"LOG_PATH"`
		Type     string `env:"LOG_TYPE"`
	}
}
