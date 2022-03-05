package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Posts struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Car struct {
	gorm.Model
	Brand string
	Type  string
	Price float64
	MPG   int
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//      db.DropTableIfExists(&Posts{})
	db.AutoMigrate(&Posts{})
	// db.Create(&Posts{Id: 0, Title: "abc", Desc: "qwer"})
	handleRequests()
}

func handleRequests() {
	log.Println("Starting server at http://127.0.0.1:8000/")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/newpost", createNewPost).Methods("POST")
	myRouter.HandleFunc("/allpost", returnAllPost)
	myRouter.HandleFunc("/post/{id}", returnSinglePost)
	myRouter.HandleFunc("/delete/{id}", deletePost)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
func createNewPost(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read the body: %v\n", err)
	}
	var post Posts
	json.Unmarshal(reqBody, &post)
	if e := db.Create(&post).Error; e != nil {
		log.Println("Unable to create new post")
	}
	fmt.Println("EndPoint Hit! Create New Post")
	json.NewEncoder(w).Encode(post)
}

func returnAllPost(w http.ResponseWriter, r *http.Request) {
	posts := []Posts{}
	if result := db.Find(&posts).Error; result != nil {
		log.Println("Unable to delete post: %v\n", result)
	}
	fmt.Println("EndPoint Hit: returnAllpost")
	json.NewEncoder(w).Encode(posts)
}

func returnSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	posts := []Posts{}
	if err := db.Find(&posts).Error; err != nil {
		log.Println("Unable to find post with id: %d", key)
	}
	for _, posts := range posts {
		//string to int
		s, err := strconv.Atoi(key)
		if err == nil {
			if posts.Id == s {
				fmt.Println(posts)
				fmt.Println("Encoding Hit: Booking No:", key)
				json.NewEncoder(w).Encode(posts)
			}
		}
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post Posts
	if err := db.First(&post, params["id"]).Error; err != nil {
		log.Println("Unable to find id %v\n", err)
	}
	if result := db.Delete(&post).Error; result != nil {
		log.Println("Unable to delete post")
	}
	var posts []Posts
	db.Find(&posts)
	fmt.Println("User Deleted")
	json.NewEncoder(w).Encode(&posts)
}
