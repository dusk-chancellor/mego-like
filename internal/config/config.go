package config

import (
	"log"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBUser     string `env:"DB_USER" required:"true"`
	DBPassword string `env:"DB_PASSWORD" required:"true"`
	DBHost     string `env:"DB_HOST" required:"true"`
	DBPort     string `env:"DB_PORT" required:"true"`
	DBName     string `env:"DB_NAME" required:"true"`

	RedisHost     string `env:"REDIS_HOST" required:"true"`
	RedisPort     string `env:"REDIS_PORT" required:"true"`
	RedisPassword string `env:"REDIS_PASSWORD" required:"true"`

	GRPCPort string `env:"LIKE_SERVICE_GRPC_PORT" required:"true"`

	UserServiceAddress     string `env:"USER_SERVICE_ADDRESS" required:"true"`
	AuthServiceAddress     string `env:"AUTH_SERVICE_ADDRESS" required:"true"`
	PostServiceAddress     string `env:"POST_SERVICE_ADDRESS" required:"true"`
	StorageServiceAddress  string `env:"STORAGE_SERVICE_ADDRESS" required:"true"`
	LikeServiceAddress     string `env:"LIKE_SERVICE_ADDRESS" required:"true"`
	CommentServiceAddress  string `env:"COMMENT_SERVICE_ADDRESS" required:"true"`
	FavoriteServiceAddress string `env:"FAVORITE_SERVICE_ADDRESS" required:"true"`
}

func LoadConfig(path string) *Config {
	cfg := &Config{} // Create a pointer to Config
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return cfg
}
