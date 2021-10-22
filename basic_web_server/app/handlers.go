package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct{
	Name string     `json:"name"     xml:"name"`
	City string     `json:"city"     xml:"city"`
	Zipcode string  `json:"Zipcode"  xml:"zipcode"`
}

func Greet(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello-world")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request){
	Customers := []Customer{
		{"Sundar","Theni","625531",},
		{"Raji","Madurai","500300",},
		{"Muni","Madras","674321",}}

	if r.Header.Get("Content-Type") == "application/json"{
		w.Header().Add("Content-Type","application/json")
		json.NewEncoder(w).Encode(Customers)
	}else{
		w.Header().Add("Content-Type","application/xml")
		xml.NewEncoder(w).Encode(Customers)
	}

}
