package go_trellis_db

import (
	"database/sql"
	"strings"
	"time"
)

const LARAVEL_FORMAT = "2006-01-02 15:04:05"

type LaravelNullTime sql.NullTime

func (t *LaravelNullTime) Scan(val interface{}) (err error) {
	if val == nil {
		t.Valid = false
		return
	}
	s := string(val.([]byte))
	if strings.HasPrefix(s, "0000") {
		t.Valid = false
		return
	}
	t.Time, err = time.Parse(LARAVEL_FORMAT, s)
	return
}
