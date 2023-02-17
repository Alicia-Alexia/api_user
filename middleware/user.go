package middleware

import (
	"fmt"
	
	"net/http"
	 "github.com/gorilla/mux"
	 "encoding/json"
	"github.com/Alicia-Alexia/api_user/models"

)
var users = []models.Users{}
func indexBy(users []models.Users, username string) int {
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			return i
		}
	}
	return -1
}


func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
	    fmt.Println(err)
	    http.Error(w, "Error encoding response object", http.StatusInternalServerError)
    }
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	index := indexBy(users, username)
	if index < 0 {
		http.Error(w, "Usuario não existe", http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users[index]); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	index := indexBy(users, username)
	if index < 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	u := models.Users{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		return
	}
	users[index] = u
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&u); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	index := indexBy(users, username)
	users = append(users[:index], users[index+1:]...)
	w.WriteHeader(http.StatusOK)
 }

func CreateUserArray(w http.ResponseWriter, r *http.Request) {
	u := []models.Users{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(err)
		http.Error(w, "Request não possui body", http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(&u)
	users = append(u)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.Users{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		return
	}
	users = append(users, u)
	response, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}