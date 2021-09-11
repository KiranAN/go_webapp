package main

import (
    "fmt"
    "io/ioutil"
    "os"
	"net/http"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/assets/view", handler).Methods("GET")
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return router
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	
    // Open jsonFile
    jsonFile, err := os.Open("assets/data.json")

    // handle error
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened ")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()
    // read as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteValue)
}
