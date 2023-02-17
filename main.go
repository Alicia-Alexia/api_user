package main


import (
	"fmt"
	
	"net/http"
	 "github.com/gorilla/mux"
	 "encoding/json"
)

type Users struct {
	Username        string   `json:"username"`
	Firstname       string   `json:"firstname"`
	Lastname        string   `json:"lastname"`
	Email           string   `json:"email"`
	Password        string   `json:"password"`
	Phone           string   `json:"phone"`
	Userstatus      int      `json:"userStatus"`
}
var users = []Users{}
func indexBy(users []Users, username string) int {
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			return i
		}
	}
	return -1
}


func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
	    fmt.Println(err)
	    http.Error(w, "Error encoding response object", http.StatusInternalServerError)
    }
}
func getUser(w http.ResponseWriter, r *http.Request) {
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
func updateUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	index := indexBy(users, username)
	if index < 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	u := Users{}
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
func deleteUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	index := indexBy(users, username)
	users = append(users[:index], users[index+1:]...)

	//! Falta só a mensagem usuário não encontrado
	w.WriteHeader(http.StatusOK)

 }
func createUserArray(w http.ResponseWriter, r *http.Request) {
	u := []Users{}
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
	//! Falta só a mensagem Todos usuários criados
	
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	u := Users{}
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

func main(){
    r := mux.NewRouter()
    usersR := r.PathPrefix("/users").Subrouter()
	usersR.Path("").Methods(http.MethodGet).HandlerFunc(getAllUsers)
    usersR.Path("").Methods(http.MethodPost).HandlerFunc(createUser)
    usersR.Path("/array").Methods(http.MethodPost).HandlerFunc(createUserArray)
    usersR.Path("/{username}").Methods(http.MethodGet).HandlerFunc(getUser)
    usersR.Path("/{username}").Methods(http.MethodPut).HandlerFunc(updateUser)
    usersR.Path("/{username}").Methods(http.MethodDelete).HandlerFunc(deleteUser)
    fmt.Println("Start listening")
    fmt.Println(http.ListenAndServe(":8080", r))
}

