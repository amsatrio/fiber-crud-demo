package nomen

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Masculine  Gender = "masculine"
	Feminine   Gender = "feminine"
	Neuter     Gender = "neuter"
	PluralOnly Gender = "plural_only"
)

type Level string

const (
	A1 Level = "A1"
	A2 Level = "A2"
	B1 Level = "B1"
	B2 Level = "B2"
	C1 Level = "C1"
	C2 Level = "C2"
)

type Noun struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	Singular          string         `gorm:"column:singular;size:100;not null" json:"singular" validate:"required"`
	Gender            Gender         `gorm:"column:gender;type:enum('masculine','feminine','neuter','plural_only');not null" json:"gender" validate:"required"`
	Plural            *string        `gorm:"column:plural;size:100" json:"plural"`
	GenitiveSingular  *string        `gorm:"column:genitive_singular;size:100" json:"genitiveSingular"`
	IsNDeklination    bool           `gorm:"column:is_n_deklination;not null;default:false" json:"isNDeklination"`
	TranslationEn     string         `gorm:"column:translation_en;size:255;not null" json:"translationEn" validate:"required"`
	ExampleSentenceDe *string        `gorm:"column:example_sentence_de;type:text" json:"exampleSentenceDe"`
	ExampleSentenceEn *string        `gorm:"column:example_sentence_en;type:text" json:"exampleSentenceEn"`
	Level             Level          `gorm:"column:level;type:enum('A1','A2','B1','B2','C1','C2');not null;default:'A1'" json:"level" validate:"required"`
	CreatedAt         time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

type CreateNounRequest struct {
	Singular          string  `json:"singular" validate:"required"`
	Gender            Gender  `json:"gender" validate:"required"`
	Plural            *string `json:"plural"`
	GenitiveSingular  *string `json:"genitiveSingular"`
	IsNDeklination    bool    `json:"isNDeklination"`
	TranslationEn     string  `json:"translationEn" validate:"required"`
	ExampleSentenceDe *string `json:"exampleSentenceDe"`
	ExampleSentenceEn *string `json:"exampleSentenceEn"`
	Level             Level   `json:"level" validate:"required"`
}

type UpdateNounRequest struct {
	Singular          string  `json:"singular"`
	Gender            Gender  `json:"gender"`
	Plural            *string `json:"plural"`
	GenitiveSingular  *string `json:"genitiveSingular"`
	IsNDeklination    *bool   `json:"isNDeklination"`
	TranslationEn     string  `json:"translationEn"`
	ExampleSentenceDe *string `json:"exampleSentenceDe"`
	ExampleSentenceEn *string `json:"exampleSentenceEn"`
	Level             Level   `json:"level"`
}
