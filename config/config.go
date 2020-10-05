package config

// Config: Model for DB Config
type Config struct {
	DB *DBConfig
}

// DBConfig: Model for DB Config
type DBConfig struct {
	Dialect  string `default:"mysql"`
	Host     string `envconfig:"DB_HOST"`
	Port     int    `default:"3306"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
	Charset  string `default:"utf8"`
}
