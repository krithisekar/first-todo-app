package main

import (

	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))

}


func params1(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")

	APILINK := "http://localhost:8081/api/v1/Acc/krithi/Task/buy-vegetables"
	req, err := http.NewRequest(http.MethodGet, APILINK, nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
    w.Write(body)
    
}
func params2(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")

	APILINK := "http://localhost:8083/api/v1/acc/{name}"
	req, err := http.NewRequest(http.MethodPost, APILINK, nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	
	if resp.StatusCode==http.StatusCreated{
		w.WriteHeader(http.StatusOK)
	}	else{
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/account/{YourName}", params1).Methods(http.MethodGet)
	api.HandleFunc("/account/{YourName}", params2).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8080", r))
}
