package main

import (
	"learn-crud/rest"
	"net/http"
)

func main() {
	http.HandleFunc("/api/students", rest.GetAll)
	http.HandleFunc("/api/student", rest.SingleHandler)
	http.ListenAndServe(":9000", nil)
}
