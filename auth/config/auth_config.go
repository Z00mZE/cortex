package config

type AuthConfig struct {
	Http struct {
		Port string `envconfig:"HTTP_PORT" required:"true"`
	}
	Database struct {
		Host string `envconfig:"DB_HOST" required:"true"`
	}
}
