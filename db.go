package go_trellis_db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string `default:"localhost"`
	Port     int    `default:"3306"`
	Database string
	Username string
	Password string
}

func Connect(config DBConfig) (*sqlx.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	return sqlx.Connect("mysql", url)
}
