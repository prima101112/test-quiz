package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Answer struct {
	ID          uint   `gorm:"primary_key"`
	FieldAnswer string `json:"fieldanswer"`
	QuestionID  uint
}

type Question struct {
	ID       uint     `gorm:"primary_key"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

type Questions []Question

var DB *gorm.DB

//migrate and seed
func init() {
	if _, err := os.Stat("test.db"); os.IsNotExist(err) {
		ConnectDB()
		// DB.LogMode(true)
		// Migrate the schema
		DB.CreateTable(&Question{})
		DB.CreateTable(&Answer{})

		// Seed data
		var q Question
		var ans []Answer

		DB.NewRecord(q)
		ans = []Answer{{FieldAnswer: "jawaban 1"}, {FieldAnswer: "jawaban 2"}, {FieldAnswer: "jawaban 3"}, {FieldAnswer: "jawaban 4"}}
		q = Question{Question: "1. is this first question num?", Answers: ans}
		DB.Save(&q)

		DB.NewRecord(q)
		ans = []Answer{{FieldAnswer: "jawaban satu"}, {FieldAnswer: "jawaban dua"}, {FieldAnswer: "jawaban tiga"}, {FieldAnswer: "jawaban lima"}}
		q = Question{Question: "2. is this second question string?", Answers: ans}
		DB.Save(&q)

		DB.NewRecord(q)
		ans = []Answer{{FieldAnswer: "8"}, {FieldAnswer: "9"}, {FieldAnswer: "10"}, {FieldAnswer: "11"}}
		q = Question{Question: "3. how mane 3 + 5", Answers: ans}
		DB.Save(&q)

		var ques []Question
		DB.Find(&ques)
		log.Print(ques)

		var anss []Answer
		DB.Find(&anss)
		log.Print(anss)
		fmt.Println("=========================")
		log.Print("FINISHING MIGRATE SEED DATA")
	} else {
		ConnectDB()
		fmt.Println("=========================")
		log.Print("YOU ALREADY HAVE THE DATA")
	}
}

func ConnectDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	// DB.LogMode(true)
}
