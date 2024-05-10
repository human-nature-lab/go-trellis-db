package go_trellis_db

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/wyattis/z/zsql"
)

type Geo struct {
	Id                string           `db:"id" json:"id"`
	GeoTypeId         string           `db:"geo_type_id" json:"geo_type_id"`
	ParentId          zsql.NullString  `db:"parent_id" json:"parent_id"`
	AssignedId        zsql.NullString  `db:"assigned_id" json:"assigned_id"`
	NameTranslationId zsql.NullString  `db:"name_translation_id" json:"name_translation_id"`
	Latitude          zsql.NullFloat64 `db:"latitude" json:"latitude"`
	Longitude         zsql.NullFloat64 `db:"longitude" json:"longitude"`
	Altitude          zsql.NullFloat64 `db:"altitude" json:"altitude"`
	CreatedAt         LaravelNullTime  `db:"created_at" json:"created_at"`
	UpdatedAt         LaravelNullTime  `db:"updated_at" json:"updated_at"`
	DeletedAt         LaravelNullTime  `db:"deleted_at" json:"deleted_at"`

	NameTranslation *Translation `json:"name_translation"`
	Ancestors       []Geo        `json:"ancestors"`
}

func (g *Geo) Load(db *sqlx.DB) (err error) {
	if g.NameTranslationId.Valid && g.NameTranslation == nil {
		ts, err := LoadTranslations(db, g.NameTranslationId.String)
		if err != nil {
			return err
		}
		g.NameTranslation = ts[g.NameTranslationId.String]
	}
	if g.ParentId.Valid {
		parent := Geo{}
		if err = db.Get(&parent, `SELECT * FROM geo WHERE id = ?`, g.ParentId.String); err != nil {
			return
		}
		if err = parent.Load(db); err != nil {
			return
		}
		g.Ancestors = append(parent.Ancestors, parent)
	}
	return
}

func (g Geo) EnglishAncestry() []string {
	names := []string{}
	for _, ancestor := range g.Ancestors {
		names = append(names, ancestor.NameTranslation.English())
	}
	names = append(names, g.NameTranslation.English())
	return names
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
	Id                 string          `db:"id" json:"id"`
	SourceRespondentId string          `db:"source_respondent_id" json:"source_respondent_id"`
	TargetRespondentId string          `db:"target_respondent_id" json:"target_respondent_id"`
	Note               zsql.NullString `db:"note" json:"note"`
	CreatedAt          LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt          LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt          LaravelNullTime `db:"deleted_at" json:"deleted_at"`

	Datum *Datum `json:"datum"`
}

func (e *Edge) Load(db *sqlx.DB) (err error) {
	if e.Datum == nil {
		datum := Datum{}
		err = db.Get(&datum, `SELECT * FROM datum WHERE edge_id = ?`, e.Id)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		if err != nil {
			return
		}
		e.Datum = &datum
	}
	return
}

type Datum struct {
	Id               string          `db:"id" json:"id"`
	Name             zsql.NullString `db:"name" json:"name"`
	Val              zsql.NullString `db:"val" json:"val"`
	ChoiceId         zsql.NullString `db:"choice_id" json:"choice_id"`
	SurveyId         zsql.NullString `db:"survey_id" json:"survey_id"`
	QuestionId       zsql.NullString `db:"question_id" json:"question_id"`
	ParentDatumId    zsql.NullString `db:"parent_datum_id" json:"parent_datum_id"`
	DatumTypeId      zsql.NullString `db:"datum_type_id" json:"datum_type_id"`
	SortOrder        int             `db:"sort_order" json:"sort_order"`
	CreatedAt        LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt        LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt        LaravelNullTime `db:"deleted_at" json:"deleted_at"`
	RosterId         zsql.NullString `db:"roster_id" json:"roster_id"`
	EventOrder       int             `db:"event_order" json:"event_order"`
	QuestionDatumId  string          `db:"question_datum_id" json:"question_datum_id"`
	GeoId            zsql.NullString `db:"geo_id" json:"geo_id"`
	EdgeId           zsql.NullString `db:"edge_id" json:"edge_id"`
	PhotoId          zsql.NullString `db:"photo_id" json:"photo_id"`
	RespondentGeoId  zsql.NullString `db:"respondent_geo_id" json:"respondent_geo_id"`
	RespondentNameId zsql.NullString `db:"respondent_name_id" json:"respondent_name_id"`
	ActionId         string          `db:"action_id" json:"action_id"`
	RandomSortOrder  int             `db:"random_sort_order" json:"random_sort_order"`
}

type ActionPayload struct {
	EdgeId string `json:"edge_id"`
}

type PreloadAction struct {
	Id           string          `db:"id" json:"id"`
	ActionType   string          `db:"action_type" json:"action_type"`
	Payload      string          `db:"payload" json:"payload"`
	RespondentId string          `db:"respondent_id" json:"respondent_id"`
	QuestionId   string          `db:"question_id" json:"question_id"`
	CreatedAt    LaravelNullTime `db:"created_at" json:"created_at"`
	DeletedAt    LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type Snapshot struct {
	Id        string          `db:"id" json:"id"`
	FileName  string          `db:"file_name" json:"file_name"`
	Hash      string          `db:"hash" json:"hash"`
	CreatedAt LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type RespondentConditionTag struct {
	Id             string          `db:"id" json:"id"`
	RespondentId   string          `db:"respondent_id" json:"respondent_id"`
	ConditionTagId string          `db:"condition_tag_id" json:"condition_tag_id"`
	CreatedAt      LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt      LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt      LaravelNullTime `db:"deleted_at" json:"deleted_at"`

	ConditionTag *ConditionTag `json:"condition_tag"`
}

func (rct *RespondentConditionTag) Load(db *sqlx.DB) (err error) {
	if rct.ConditionTagId != "" && rct.ConditionTag == nil {
		ct := ConditionTag{}
		if err = db.Get(&ct, `SELECT * FROM condition_tag WHERE id = ?`, rct.ConditionTagId); err != nil {
			return
		}
		rct.ConditionTag = &ct
	}
	return
}

type ConditionTag struct {
	Id        string          `db:"id" json:"id"`
	Name      string          `db:"name" json:"name"`
	CreatedAt LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type Respondent struct {
	Id                     string          `db:"id" json:"id"`
	Name                   string          `db:"name" json:"name"`
	AssignedId             zsql.NullString `db:"assigned_id" json:"assigned_id"`
	GeoId                  zsql.NullString `db:"geo_id" json:"geo_id"`
	Notes                  zsql.NullString `db:"notes" json:"notes"`
	GeoNotes               zsql.NullString `db:"geo_notes" json:"geo_notes"`
	AssociatedRespondentId zsql.NullString `db:"associated_respondent_id" json:"associated_respondent_id"`
	CreatedAt              LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt              LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt              LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type RespondentName struct {
	Id                       string          `db:"id" json:"id"`
	IsDisplayName            bool            `db:"is_display_name" json:"is_display_name"`
	LocaleId                 zsql.NullString `db:"locale_id" json:"locale_id"`
	RespondentId             string          `db:"respondent_id" json:"respondent_id"`
	PreviousRespondentNameId zsql.NullString `db:"previous_respondent_name_id" json:"previous_respondent_name_id"`
	Name                     string          `db:"name" json:"name"`
	CreatedAt                LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt                LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt                LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type RespondentGeo struct {
	Id                      string          `db:"id" json:"id"`
	IsCurrent               bool            `db:"is_current" json:"is_current"`
	RespondentId            string          `db:"respondent_id" json:"respondent_id"`
	GeoId                   string          `db:"geo_id" json:"geo_id"`
	Notes                   zsql.NullString `db:"notes" json:"notes"`
	PreviousRespondentGeoId zsql.NullString `db:"previous_respondent_geo_id" json:"previous_respondent_geo_id"`
	CreatedAt               LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt               LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt               LaravelNullTime `db:"deleted_at" json:"deleted_at"`

	Geo Geo `json:"geo"`
}

type Survey struct {
	Id             string          `db:"id" json:"id"`
	FormId         string          `db:"form_id" json:"form_id"`
	RespondentId   string          `db:"respondent_id" json:"respondent_id"`
	StudyId        string          `db:"study_id" json:"study_id"`
	LastQuestionId zsql.NullString `db:"last_question_id" json:"last_question_id"`
	CreatedAt      LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt      LaravelNullTime `db:"updated_at" json:"updated_at"`
	CompletedAt    LaravelNullTime `db:"completed_at" json:"completed_at"`
	DeletedAt      LaravelNullTime `db:"deleted_at" json:"deleted_at"`

	Form Form `json:"form"`
}

type Form struct {
	Id                string          `db:"id" json:"id"`
	FormMasterId      string          `db:"form_master_id" json:"form_master_id"`
	NameTranslationId zsql.NullString `db:"name_translation_id" json:"name_translation_id"`
	Version           int             `db:"version" json:"version"`
	IsPublished       bool            `db:"is_published" json:"is_published"`
	CreatedAt         LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt         LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt         LaravelNullTime `db:"deleted_at" json:"deleted_at"`

	NameTranslation *Translation `json:"name_translation"`
}

func (f *Form) Load(db *sqlx.DB) (err error) {
	if f.NameTranslationId.Valid && f.NameTranslation == nil {
		ts, err := LoadTranslations(db, f.NameTranslationId.String)
		if err != nil {
			return err
		}
		f.NameTranslation = ts[f.NameTranslationId.String]
	}
	return
}

type Translation struct {
	Id        string          `db:"id" json:"id"`
	CreatedAt LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt LaravelNullTime `db:"deleted_at" json:"deleted_at"`

	TranslationTexts []TranslationText `json:"translation_texts"`
}

func (t Translation) English() string {
	for _, tt := range t.TranslationTexts {
		if tt.LocaleId == "48984fbe-84d4-11e5-ba05-0800279114ca" {
			return tt.TranslatedText
		}
	}
	return ""
}

type TranslationText struct {
	Id             string          `db:"id" json:"id"`
	TranslationId  string          `db:"translation_id" json:"translation_id"`
	LocaleId       string          `db:"locale_id" json:"locale_id"`
	TranslatedText string          `db:"translated_text" json:"translated_text"`
	CreatedAt      LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt      LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt      LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}

type Asset struct {
	Id        string          `db:"id" json:"id"`
	FileName  string          `db:"file_name" json:"file_name"`
	Md5Hash   string          `db:"md5_hash" json:"md5_hash"`
	Size      int             `db:"size" json:"size"`
	Type      string          `db:"type" json:"type"`
	MimeType  string          `db:"mime_type" json:"mime_type"`
	CreatedAt LaravelNullTime `db:"created_at" json:"created_at"`
	UpdatedAt LaravelNullTime `db:"updated_at" json:"updated_at"`
	DeletedAt LaravelNullTime `db:"deleted_at" json:"deleted_at"`
}
