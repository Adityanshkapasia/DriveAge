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
	Id int  `json:"id"`
	Brand string `json:"brand"`
	Type  string `json:"type"`
	Price float64 `json:"price"`
	MPG   int `json:"mpg"`
	Name  string `json:"name"`
	TankCapacity  int `json:"tankcapacity"`
	Color string `json:"color"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//      db.DropTableIfExists(&Posts{})
	db.AutoMigrate(&Posts{}, &Car{})
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
	myRouter.HandleFunc("/allcar", returnAllCar)
	myRouter.HandleFunc("/car/{id}", returnSingleCar)
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
func returnAllCar(w http.ResponseWriter, r *http.Request) {
	car := []Car{}
	if result := db.Find(&car).Error; result != nil {
		log.Println("Unable to delete post: %v\n", result)
	}
	fmt.Println("EndPoint Hit: returnAllpost")
	json.NewEncoder(w).Encode(car)
}

func returnSingleCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	car := []Car{}
	if err := db.Find(&car).Error; err != nil {
		log.Println("Unable to find car with id: %d", key)
	}
	for _, car := range car {
		//string to int
		s, err := strconv.Atoi(key)
		if err == nil {
			if car.Id == s {
				fmt.Println(car)
				fmt.Println("Encoding Hit: Booking No:", key)
				json.NewEncoder(w).Encode(car)
			}
		}
	}
}
