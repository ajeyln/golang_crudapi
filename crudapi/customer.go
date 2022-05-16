package main

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustID   string `json:"custid"`
	CustName string `json:"name"`
	ConInfo  int    `json:"contact"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}
