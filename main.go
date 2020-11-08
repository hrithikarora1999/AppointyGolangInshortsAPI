package main

import (

	"context"
	"os"
	"encoding/json"
	"log"
	"net/http"
	"github.com/hrithikarora1999/API/helper"
	"github.com/hrithikarora1999/API/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = helper.ConnectDB()

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var articles []models.Article

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var article models.Article
		err := cur.Decode(&article) 
		if err != nil {
			log.Fatal(err)
		}

		
		articles = append(articles, article)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(articles) 
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article models.Article
	var params = mux.Vars(r)

	id, _ := (params["id"])

	filter := bson.M{"id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&article)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func SearchArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var articles []models.Article
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	
	defer cur.Close(context.TODO())
	var articles1 []models.Article
	for cur.Next(context.TODO()) {

		var article models.Article
		err := cur.Decode(&article) 
		if err != nil {
			log.Fatal(err)
		}

		
		articles = append(articles, article)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	params := mux.Vars(r) 
	for _, item := range articles {
		if (item.Title == params["title"] || item.SubTitle == params["title"] || item.Content == params["title"] ) {
			articles1=append(articles1,item)
			
		}
	}	

	json.NewEncoder(w).Encode(articles1) 

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article models.Article
	
	_ = json.NewDecoder(r.Body).Decode(&article)

	result, err := collection.InsertOne(context.TODO(), article)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/article", getBooks).Methods("GET")
	r.HandleFunc("/article/{id}", getBook).Methods("GET")
	r.HandleFunc("/article", createBook).Methods("POST")
	r.HandleFunc("/article/search/q={title}", SearchArticle).Methods("GET")

	
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), r))


}