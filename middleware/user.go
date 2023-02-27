package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Alicia-Alexia/api_user/models"
	"github.com/gorilla/mux"
)
var users = []models.Users{}
func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var user models.Users
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "request não possui body", http.StatusBadRequest)
        return
    }
    user.ID = len(users) + 1
    users = append(users, user)
    json.NewEncoder(w).Encode(user)
}

var lastID int

func CreateUsersArray(w http.ResponseWriter, r *http.Request) {
    var newUsers []models.Users
    if err := json.NewDecoder(r.Body).Decode(&newUsers); err != nil {
        http.Error(w,"request não possui body", http.StatusBadRequest)
        return
    }
    for _, newUser := range newUsers {
        lastID++
        newUser.ID = lastID
        users = append(users, newUser)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    username := params["username"]
    for _, user := range users {
        if user.Username == username {
            json.NewEncoder(w).Encode(user)
            return
		}
    }
	for _, user := range users {
        if user.Username != username{
            http.Error(w,"username inválido", http.StatusBadRequest)
            return
        } 
    }
    http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    username := params["username"]
    var updatedUser models.Users
    err := json.NewDecoder(r.Body).Decode(&updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    for i, user := range users {
        if user.Username == username {
            updatedUser.ID = user.ID
            users[i] = updatedUser
            json.NewEncoder(w).Encode(updatedUser)
            return
        }
    }
	for _, user := range users {
        if user.Username != username{
            http.Error(w,"username inválido", http.StatusBadRequest)
            return
        } 
    }
    http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    username := params["username"]

    for i, user := range users {
        if user.Username == username {
            users = append(users[:i], users[i+1:]...)
            json.NewEncoder(w).Encode(user)
            return
        }
    }
	for _, user := range users {
        if user.Username != username{
            http.Error(w,"username inválido", http.StatusBadRequest)
            return
        } 
    }
    http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}