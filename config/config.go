package config

// Config: Model for DB Config
type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string `default:"mysql"`
	Host     string `envconfig:"DB_HOST"`
	Port     int `default:3306`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
	Charset  string `default:"utf8"`
}

// func GetConfig() *Config {
// 	return &Config{
// 		DB: &DBConfig{
// 			Dialect:  "mysql",
// 			Host:     "127.0.0.1",
// 			Port:     3306,
// 			Username: "guest",
// 			Password: "Guest0000!",
// 			Name:     "todoapp",
// 			Charset:  "utf8",
// 		},
// 	}
// }
