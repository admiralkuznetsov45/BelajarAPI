package model

type Question struct {
	ID            uint   `gorm:"primarykey"`
	QuestionType  string `json:"QuestionType"`
	QuestionPoint uint   `json:"QuestionPoint"`
	Text          string `json:"Text"`
}
