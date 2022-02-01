package model

type Answer struct {
	// gorm.Model
	ID          uint   `gorm:"primarykey"`
	JawabanSatu string `json:"jawaban_satu"`
	JawabanDua  string `json:"jawaban_dua"`
	JawabanTiga string `json:"jawaban_tiga"`
}
