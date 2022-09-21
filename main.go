package main

import (
	"golang-blogs/controller"
	"golang-blogs/database"
	"net/http"
)

func main() {
	database.Connect()

	http.HandleFunc("/create", controller.Create)

	http.ListenAndServe(":3000", nil)
}
