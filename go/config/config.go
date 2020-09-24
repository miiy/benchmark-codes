package config

type Config struct {
	Redis Redis
}

type Redis struct {
	Host string
	Port string
	Password string
	Database int
}
