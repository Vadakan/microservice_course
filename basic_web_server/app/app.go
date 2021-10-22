package app

import "net/http"

func Start(){

	mux := http.NewServeMux()

	mux.HandleFunc("/greet", Greet)

	mux.HandleFunc("/customer", GetAllCustomers)

	http.ListenAndServe(":8080",mux)
}
