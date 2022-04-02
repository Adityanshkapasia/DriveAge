// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/gorilla/sessions"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// 	"net/http"
// )

// var cookie = "Driveage"
// var cookieStore = sessions.NewCookieStore([]byte(cookie))

// type userRegister struct {
// 	Email    string `json:"email"`
// 	Name     string `json:"name"`
// 	Password string `json:"password"`
// }

// type userLogin struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // bcrypt strength
// var strength = 11

// // GetUser returns the user from the session info
// func GetUser(r *http.Request) (*User, error) {
// 	user := User{}
// 	session, err := cookieStore.Get(r, cookie)

// 	if err != nil {
// 		return nil, err
// 	}

// 	if email, ok := session.Values["id"].(string); ok {
// 		db.Where(&User{Email: email}).Find(&user)

// 		if user.Email == email {
// 			return &user, nil
// 		}
// 		return nil, fmt.Errorf("No session found for user %s", email)
// 	}
// 	return nil, fmt.Errorf("No session found")
// }

// // GetAdmin ensures the user exists and it is an admin
// func GetAdmin(r *http.Request) (*User, error) {
// 	user, err := GetUser(r, db)
// 	if err == nil && user.Role != "admin" {
// 		return nil, fmt.Errorf("User %s is not an admin", user.Email)
// 	}
// 	return user, err
// }

// // GetUserEmail recoveres user email from session, nil if error
// // and fills in the reply. This works even for token based sessions
// func GetUserEmail(w http.ResponseWriter, r *http.Request) string {
// 	// Try first to find a token

// 	user, err := GetUser(r, db)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusForbidden)
// 		return ""
// 	}
// 	return user.Email
// }

// // NoSession replies with standard answer if no session
// // Usage: if NoSession(w,r) return
// func NoSession(w http.ResponseWriter, r *http.Request) bool {
// 	// First, determine if this a token-based request

// 	_, err := GetUser(r, db)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusForbidden)
// 		return true
// 	}
// 	return false
// }

// // NoAdmin replies with standard answer if user not an admin or no session
// func NoAdmin(w http.ResponseWriter, r *http.Request) bool {
// 	_, err := GetAdmin(r, db)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusForbidden)
// 		return true
// 	}
// 	return false
// }

// // HandleRegister registers a new user based on an invitation
// func HandleRegister(w http.ResponseWriter, r *http.Request) {
// 	// convert request to registration data
// 	var registration userRegister
// 	err := json.NewDecoder(r.Body).Decode(&registration)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	hash, err := bcrypt.GenerateFromPassword([]byte(registration.Password), strength)

// 	if err != nil {
// 		http.Error(w, "Password hashing failed", http.StatusBadRequest)
// 		return
// 	}

// 	// Count admins. If none, make this user an admin
// 	var count int64
// 	db.Model(&User{}).Where("role = ?", "admin").Count(&count)

// 	var user = User{}
// 	db.FirstOrCreate(&user, User{Email: registration.Email})
// 	user.Name = registration.Name
// 	user.Hash = string(hash)
// 	if count == 0 {
// 		user.Role = "admin"
// 	} else {
// 		user.Role = "user"
// 	}

// 	db.Save(&user)

// 	// Delete registration

// 	// Redirect to main page
// 	http.Error(w, "Registration successfull", http.StatusOK)

// }

// // HandleWhoAmI returns information about logged in user
// func HandleWhoAmI(w http.ResponseWriter, r *http.Request) {
// 	user, err := GetUser(r, db)
// 	if err == nil {
// 		res, err := json.Marshal(user)
// 		w.Write(res)
// 	} else {
// 		res, err := json.Marshal(map[string]string{
// 			"error": "Not logged in",
// 		})
// 		w.Write(res)
// 	}
// }

// // HandleLogin loggs in the user attaches a session COOKIE to the reply. Returns WhoAmI info
// func HandleLogin(w http.ResponseWriter, r *http.Request) {
// 	// convert request to registration data
// 	var login userLogin
// 	err := json.NewDecoder(r.Body).Decode(&login)

// 	if err != nil {
// 		http.Error(w, "Difformed login request", http.StatusBadRequest)
// 		return
// 	}

// 	// Bring the user from the database
// 	user := User{}
// 	db.Where(&User{Email: login.Email}).Find(&user)

// 	if user.Email != login.Email {
// 		http.Error(w, fmt.Sprintf("User %s does not exist", login.Email), http.StatusForbidden)
// 		return
// 	}

// 	// check the password
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(login.Password))
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Password for user %s is incorrect", login.Email), http.StatusForbidden)
// 		return
// 	}

// 	// setup the session and tell user that everything is fine
// 	session, err := cookieStore.New(r, cookie)
// 	if err == nil {
// 		session.Values["id"] = login.Email
// 		err = session.Save(r, w)
// 	}
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		http.Error(w, fmt.Sprintf("Could not setup session for %s user", login.Email), http.StatusConflict)
// 		return
// 	}

// 	res, err := json.Marshal(user)
// 	w.Write(res)
// }

// // HandleLogout closess current session
// func HandleLogout(w http.ResponseWriter, r *http.Request) {
// 	session, err := cookieStore.Get(r, cookie)
// 	if err == nil {
// 		// delete the cookie
// 		session.Options.MaxAge = -1
// 		session.Save(r, w)

// 		http.Error(w, "Successfull logout", http.StatusOK)
// 	} else {
// 		http.Error(w, "No session found", http.StatusNotFound)
// 	}
// }
