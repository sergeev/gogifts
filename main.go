package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func routes() {
	http.HandleFunc("/", LoginPageHandler)
	http.HandleFunc("/register", RegisterPageHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
}

var router = mux.NewRouter()

func main() {
	routers()
	/*
		// Загрузка стилей проекта
		router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

		router.HandleFunc("/", Logger(common.LoginPageHandler)) // GET

		router.HandleFunc("/index", Logger(common.IndexPageHandler)) // GET
		router.HandleFunc("/login", Logger(common.LoginHandler)).Methods("POST")

		router.HandleFunc("/register", Logger(common.RegisterPageHandler)).Methods("GET")
		router.HandleFunc("/register", Logger(common.RegisterHandler)).Methods("POST")

		router.HandleFunc("/logout", Logger(common.LogoutHandler)).Methods("POST")

		// Manage page (admin)
		router.HandleFunc("/manage", Logger(common.ManageHandler)) // GET
	*/
	http.Handle("/", router)
	log.Print("Server: start")
	log.Print("Server: port:8080")
	http.ListenAndServe(":8080", nil)
}
