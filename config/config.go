package config

import (
	"QuisAPI/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	//                              username:password@tcp(host:port)/nama_database
	DB, err = gorm.Open(mysql.Open("root:ayamjantan@tcp(127.0.0.1:3306)/quis"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB connected")
	}
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(
		&model.Answer{},
		&model.Question{},
		&model.User{},
	)

}
