package main

import (
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/wyattis/z/zconfig"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Driver   string `default:"mysql"`
	Host     string `default:"localhost"`
	Port     int    `default:"3306"`
	Database string
	Username string
	Password string
}

func Connect(config DBConfig) (db *sqlx.DB, err error) {
	var url string
	switch config.Driver {
	case "mysql":
		url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?sql_mode=''", config.Username, config.Password, config.Host, config.Port, config.Database)
	case "sqlite3":
		url = config.Database
	}
	db, err = sqlx.Connect(config.Driver, url)
	if err != nil {
		return
	}
	return
}

func AutoConnect() (db *sqlx.DB, err error) {
	envLoc := flag.String("env", ".env", "")
	flag.Parse()
	config := Config{}
	if err = zconfig.New(zconfig.Env(*envLoc, "/var/www/trellis-api/.env"), zconfig.Defaults()).Apply(&config); err != nil {
		return
	}
	return Connect(config.DB)
}

func LoadTranslations(db *sqlx.DB, ids ...string) (translations map[string]*Translation, err error) {
	if len(ids) == 0 {
		return
	}
	query, args, err := sqlx.In(`SELECT * FROM translation WHERE id IN (?)`, ids)
	if err != nil {
		return
	}
	ts := []Translation{}
	query = db.Rebind(query)
	err = db.Select(&ts, query, args...)
	texts := []TranslationText{}
	query, args, err = sqlx.In(`SELECT * FROM translation_text WHERE translation_id IN (?)`, ids)
	if err != nil {
		return
	}
	if err = db.Select(&texts, query, args...); err != nil {
		return
	}
	translations = make(map[string]*Translation, len(ts))
	for _, tt := range texts {
		for i := range ts {
			if ts[i].Id == tt.TranslationId {
				ts[i].TranslationTexts = append(ts[i].TranslationTexts, tt)
			}
		}
	}
	for _, t := range ts {
		translations[t.Id] = &t
	}
	return
}
