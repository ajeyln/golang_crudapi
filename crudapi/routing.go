package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/customer", CreateCustomer).Methods("POST")
	r.HandleFunc("/customers", GetCustomers).Methods("GET")
	r.HandleFunc("/customer/{id}", GetCustomerByID).Methods("GET")
	r.HandleFunc("/customer/{id}", UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customer/{id}", DeleteCustomer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5555", r))
}
