package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Posts struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Username string `json:"username"`
}

type Car struct {
	Id           int     `json:"id"`
	Brand        string  `json:"brand"`
	Type         string  `json:"type"`
	Price        float64 `json:"price"`
	MPG          int     `json:"mpg"`
	Name         string  `json:"name"`
	TankCapacity int     `json:"tankcapacity"`
	Color        string  `json:"color"`
}
type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Hash  string `json:"hash"`
	Role  string `json:"role"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//      db.DropTableIfExists(&Posts{})
	db.AutoMigrate(&Posts{}, &Car{}, &User{})
	// db.Create(&Posts{Id: 0, Title: "abc", Desc: "qwer"})
	handleRequests()
}

func extractTags(text string) []string {
	re := regexp.MustCompile("#\\S+")
	// TODO: Strip out #
	return re.FindAllString(text, -1)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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
	myRouter.HandleFunc("/login", HandleLogin)
	myRouter.HandleFunc("/register", HandleRegister)
	myRouter.HandleFunc("/logout", HandleLogout)
	myRouter.HandleFunc("/whomi", HandleWhoAmI)
	myRouter.HandleFunc("/news", HandleNews)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})
	log.Fatal(http.ListenAndServe(":8080", c.Handler(myRouter))) //handlers.CORS(originsOk, headersOk, methodsOk)(a.r)))

}
func HandleNews(w http.ResponseWriter, r *http.Request) {
	posts := []Posts{}
	if result := db.Find(&posts).Error; result != nil {
		log.Println("error: %v\n", result)
	}
	posts2 := []Posts{}
	for _, p := range posts {
		tags := extractTags(p.Desc)
		exists := stringInSlice("#news", tags)
		if exists {
			posts2 = append(posts2, p)
		}
	}
	json.NewEncoder(w).Encode(posts2)

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

// -------------- Auth --------------
var cookie = "Driveage"
var cookieStore = sessions.NewCookieStore([]byte(cookie))

type userRegister struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type userLogin struct {
	Email    string `json:"email" gorm:"UNIQUE_INDEX"`
	Password string `json:"password"`
}

// bcrypt strength
var strength = 11

// GetUser returns the user from the session info
func GetUser(r *http.Request) (*User, error) {
	user := User{}
	session, err := cookieStore.Get(r, cookie)

	if err != nil {
		return nil, err
	}

	if email, ok := session.Values["id"].(string); ok {
		db.Where("role = ?").Find(&user)

		if user.Email == email {
			return &user, nil
		}
		return nil, fmt.Errorf("No session found for user %s", email)
	}
	return nil, fmt.Errorf("No session found")
}

// GetAdmin ensures the user exists and it is an admin
func GetAdmin(r *http.Request) (*User, error) {
	user, err := GetUser(r)
	if err == nil && user.Role != "admin" {
		return nil, fmt.Errorf("User %s is not an admin", user.Email)
	}
	return user, err
}

// GetUserEmail recoveres user email from session, nil if error
// and fills in the reply. This works even for token based sessions
func GetUserEmail(w http.ResponseWriter, r *http.Request) string {
	// Try first to find a token

	user, err := GetUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return ""
	}
	return user.Email
}

// NoSession replies with standard answer if no session
// Usage: if NoSession(w,r) return
func NoSession(w http.ResponseWriter, r *http.Request) bool {
	// First, determine if this a token-based request

	_, err := GetUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return true
	}
	return false
}

// NoAdmin replies with standard answer if user not an admin or no session
func NoAdmin(w http.ResponseWriter, r *http.Request) bool {
	_, err := GetAdmin(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return true
	}
	return false
}

// HandleRegister registers a new user based on an invitation
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	// convert request to registration data
	var registration userRegister
	err := json.NewDecoder(r.Body).Decode(&registration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	print(registration.Password)
	hash, err := bcrypt.GenerateFromPassword([]byte(registration.Password), strength)

	if err != nil {
		http.Error(w, "Password hashing failed", http.StatusBadRequest)
		return
	}

	// Count admins. If none, make this user an admin
	var count int64
	db.Model(&User{}).Where("role = ?", "admin").Count(&count)

	var user = User{}
	user.Name = registration.Name
	user.Hash = string(hash)
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}
	db.FirstOrCreate(&user, User{Email: registration.Email})

	// db.Create(&user)

	// Delete registration

	// Redirect to main page
	http.Error(w, "Registration successfull", http.StatusOK)

}

// HandleWhoAmI returns information about logged in user
func HandleWhoAmI(w http.ResponseWriter, r *http.Request) {
	user, err := GetUser(r)
	if err == nil {
		res, _ := json.Marshal(user)
		w.Write(res)
	} else {
		res, _ := json.Marshal(map[string]string{
			"error": "Not logged in",
		})
		w.Write(res)
	}
}

// HandleLogin loggs in the user attaches a session COOKIE to the reply. Returns WhoAmI info
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// convert request to registration data
	var login userLogin
	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		http.Error(w, "Difformed login request", http.StatusBadRequest)
		return
	}

	// Bring the user from the database
	user := User{}
	db.Where("email = ?", login.Email).Find(&user)

	if user.Email != login.Email {
		http.Error(w, fmt.Sprintf("User %s does not exist", login.Email), http.StatusForbidden)
		return
	}

	// check the password
	println()
	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(login.Password))
	println(login.Password)
	println(user.Hash)
	println(user.Email)
	if err != nil {
		print(err.Error())
		http.Error(w, fmt.Sprintf("Password for user %s is incorrect", login.Email), http.StatusForbidden)
		return
	}

	// setup the session and tell user that everything is fine
	session, err := cookieStore.New(r, cookie)
	if err == nil {
		session.Values["id"] = login.Email
		err = session.Save(r, w)
	}
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, fmt.Sprintf("Could not setup session for %s user", login.Email), http.StatusConflict)
		return
	}

	res, err := json.Marshal(user)
	w.Write(res)
}

// HandleLogout closess current session
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, err := cookieStore.Get(r, cookie)
	if err == nil {
		// delete the cookie
		session.Options.MaxAge = -1
		session.Save(r, w)

		http.Error(w, "Successfull logout", http.StatusOK)
	} else {
		http.Error(w, "No session found", http.StatusNotFound)
	}
}
