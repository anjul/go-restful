package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Id string `json:"id,omitempty"`
	Title string `json:"Title,omitempty"`
	Desc string `json:"desc,omitempty"`
	Content string `json:"content,omitempty"`
}

var Articles []Article

func main()  {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

func handleRequests(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/all",returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article",createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}",returnSingleArticle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "HomePage Welcomes You!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint Hit: returnArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request)  {
	reqVar := mux.Vars(r)
	key := reqVar["id"]

	for _,article := range Articles{
		if article.Id == key{
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request){

	var reqBody,_ = ioutil.ReadAll(r.Body)
	var article Article

	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}
