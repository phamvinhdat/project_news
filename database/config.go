package database

// Config database connection
type Config struct {
	Dialect      string
	Username     string
	Password     string
	DatabaseName string
	IP           string
	Port         string
}

func DefaultConfig() *Config {
	return &Config{
		Dialect:      "mysql",
		Username:     "root",
		Password:     "root",
		DatabaseName: "news",
		IP:           "127.0.0.1",
		Port:         "32815",
	}
}
