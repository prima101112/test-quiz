package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	m "test-quiz/models"
)

func showQuestion(w http.ResponseWriter, r *http.Request) {
	log.Print("enter show question function")
	var qs m.Questions
	m.DB.Select("ID, question").Order("RANDOM()").Find(&qs)
	GetRandomAns(&qs)
	log.Print(qs)

	//write json response
	w.Header().Add("Content-Type", "application/json")
	by, _ := json.Marshal(qs)
	fmt.Fprintf(w, string(by))
}

func GetRandomAns(qs *m.Questions) {
	for i := 0; i < len(*qs); i++ {
		log.Print(i)
		qid := (*qs)[i].ID
		err := m.DB.Select("ID, field_answer").Order("RANDOM()").Where("question_id = ?", qid).Find(&(*qs)[i].Answers).Error
		if err != nil {
			log.Print("no answers available")
		}
	}
}
func showAnswer(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query()
	qid := a["qid"][0]
	var as []m.Answer
	err := m.DB.Select("ID, field_answer").Order("RANDOM()").Where("question_id = ?", qid).Find(&as).Error
	if err != nil {
		fmt.Fprintf(w, "no qid")
	}
	//write json response
	w.Header().Add("Content-Type", "application/json")
	by, _ := json.Marshal(as)
	fmt.Fprintf(w, string(by))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/question", showQuestion)
	http.HandleFunc("/answers", showAnswer)
	log.Print("server run on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Print(err)
	}
}
