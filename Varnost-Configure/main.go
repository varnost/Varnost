package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
)

func logInit(
	infoHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["app_id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func main() {
	logInit(os.Stdout, os.Stderr)
	Info.Println("Starting Varnost Configuration Service")

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/config/{app_id}", GetConfig).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}