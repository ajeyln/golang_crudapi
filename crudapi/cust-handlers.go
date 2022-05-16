package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Create Customer
func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "appliation/json")
	var cust Customer
	json.NewDecoder(r.Body).Decode(&cust)
	Database.Create(&cust)
	json.NewEncoder(w).Encode(cust)

}

//Get all Customer Details from Database
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "appliation/json")
	var custs []Customer
	Database.Find(&custs)
	json.NewEncoder(w).Encode(custs)

}

// Get Customer Information by Customer ID
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "appliation/json")
	var custs Customer
	Database.First(&custs, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(custs)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "appliation/json")
	var custs Customer
	Database.First(&custs, mux.Vars(r)["id"])
	json.NewDecoder(r.Body).Decode(&custs)
	Database.Save(&custs)
	json.NewEncoder(w).Encode(custs)

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "appliation/json")
	var custs Customer
	Database.Delete(&custs, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode("Customer is deleted!!!")

}
