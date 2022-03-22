package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlerequest() {
	router := mux.NewRouter()
	router.HandleFunc("/student", InsertStudent).Methods(http.MethodPost)
	router.HandleFunc("/student", GetAllStudent).Methods(http.MethodGet)
	router.HandleFunc("/student/{id}", GetStudentById).Methods(http.MethodGet)
	router.HandleFunc("/student/{id}", EditUser).Methods(http.MethodPut)
	router.HandleFunc("/student/{id}", DeleteStudentById).Methods(http.MethodDelete)
	fmt.Println("connecting.....")
	log.Fatal(http.ListenAndServe(":8000", router))

}
