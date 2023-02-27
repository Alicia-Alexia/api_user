package router

import (
    "github.com/Alicia-Alexia/api_user/middleware"

    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/api/user/{username}", middleware.GetUser).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/newuser/array", middleware.CreateUsersArray).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/user/{username}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteuser/{username}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

    return router
}