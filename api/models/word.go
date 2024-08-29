package models

type Word struct {
	EnglishWord string `json:"english_word"`
	Translate string `json:"translate"`
	Options []string `json:"options"`
}
