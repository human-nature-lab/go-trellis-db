package go_trellis_db

import "database/sql"

type Geo struct {
	Id        string `db:"id"`
	GeoTypeId string `db:"geo_type_id"`
	ParentId  string `db:"parent_id"`
}

type User struct {
	Id       string
	Username string
	Name     string
}

type Token struct {
	Id   string
	Hash string
	Exp  string
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
	FileName  string          `db:"file_name" json:"file_name"`
	Hash      string          `db:"hash"`
	CreatedAt LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type RespondentConditionTag struct {
	Id             string          `db:"id"`
	RespondentId   string          `db:"respondent_id"`
	ConditionTagId string          `db:"condition_tag_id"`
	CreatedAt      LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt      LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt      LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type RespondentName struct {
	Id           string          `db:"id"`
	RespondentId string          `db:"respondent_id"`
	Name         string          `db:"name"`
	CreatedAt    LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt    LaravelNullTime `db:"updated_at" json:"updated_at"`
}

type RespondentGeo struct {
	Id           string          `db:"id"`
	RespondentId string          `db:"respondent_id"`
	GeoId        string          `db:"geo_id"`
	CreatedAt    LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt    LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt    LaravelNullTime `db:"deleted_at" json:"deleted_at"`

	Geo Geo `json:"geo"`
}

type Survey struct {
	Id           string          `db:"id"`
	FormId       string          `db:"form_id"`
	RespondentId string          `db:"respondent_id"`
	CreatedAt    LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt    LaravelNullTime `db:"updated_at" json:"updated_at"`
	CompletedAt  LaravelNullTime `db:"completed_at" json:"completed_at"`
	DeletedAt    LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type Form struct {
	Id           string `db:"id"`
	FormMasterId string `db:"form_master_id"`
}
