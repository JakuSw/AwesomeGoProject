package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB
var err error

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	json.NewDecoder(r.Body).Decode(&article)
	db.Create(&article)
	json.NewEncoder(w).Encode(article)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var articles []Article
	db.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}

func main() {
	db, err = gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Starting migration")
	db.AutoMigrate(&Article{})

	router := mux.NewRouter()
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")
	fmt.Println("Listen at :8000")

	http.ListenAndServe(":8000", router)
}
