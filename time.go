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
	t.Valid = err == nil
	return
}

func (t *LaravelNullTime) UnmarshalJSON(d []byte) (err error) {
	if d == nil {
		t.Valid = false
		return
	}
	s := strings.ReplaceAll(string(d), `"`, "")
	if s == "null" {
		t.Valid = false
		return
	}
	if strings.HasPrefix(s, "0001") {
		t.Valid = false
		return
	}
	t.Time, err = time.Parse(LARAVEL_FORMAT, s)
	t.Valid = err == nil
	return
}

func (t LaravelNullTime) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Time.Format(LARAVEL_FORMAT) + `"`), nil
}
