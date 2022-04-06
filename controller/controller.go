package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Configuration struct {
	Port             string // port no
	ConnectionString string // connection string
	Database         string // database name
	Collection       string // collection
}

func ReadConfig() Configuration {
	var configfile = "utility/config.properties"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Configuration
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func Handlerequest() {
	//log.Println("Started.....................")
	config := ReadConfig()
	var port = ":" + config.Port

	router := mux.NewRouter()

	corsObj := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/students", InsertStudent).Methods(http.MethodPost)
	router.HandleFunc("/students", GetAllStudent).Methods(http.MethodGet)
	router.HandleFunc("/students/{id}", GetStudentById).Methods(http.MethodGet)
	router.HandleFunc("/students/{id}", EditUser).Methods(http.MethodPut)
	router.HandleFunc("/students/{id}", DeleteStudentById).Methods(http.MethodDelete)
	router.HandleFunc("/home", Home).Methods(http.MethodGet)
	fmt.Printf("application listening on port %s", port)

	//log.Fatal(http.ListenAndServe(port, router))
	http.ListenAndServe(port, handlers.CORS(corsObj)(router))

}
