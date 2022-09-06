package go_trellis_db

import "database/sql"

type Geo struct {
	Id        string `db:"id"`
	GeoTypeId string `db:"geo_type_id"`
	ParentId  string `db:"parent_id"`
}

type Edge struct {
	Id                 string          `db:"id"`
	SourceRespondentId string          `db:"source_respondent_id"`
	TargetRespondentId string          `db:"target_respondent_id"`
	Note               sql.NullString  `db:"note"`
	CreatedAt          LaravelNullTime `db:"created_at"`
	UpdatedAt          LaravelNullTime `db:"updated_at"`
	DeletedAt          LaravelNullTime `db:"deleted_at"`
}

type ActionPayload struct {
	EdgeId string `json:"edge_id"`
}

type PreloadAction struct {
	Id           string          `db:"id"`
	ActionType   string          `db:"action_type"`
	Payload      string          `db:"payload"`
	RespondentId string          `db:"respondent_id"`
	QuestionId   string          `db:"question_id"`
	CreatedAt    LaravelNullTime `db:"created_at"`
	DeletedAt    LaravelNullTime `db:"deleted_at"`
}

type Snapshot struct {
	Id        string          `db:"id"`
	FileName  string          `db:"file_name"`
	Hash      string          `db:"hash"`
	CreatedAt LaravelNullTime `db:"created_at"`
	UpdatedAt LaravelNullTime `db:"updated_at"`
	DeletedAt LaravelNullTime `db:"deleted_at"`
}
