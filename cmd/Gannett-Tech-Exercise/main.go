package main

import (
	"net/http"

	"github.com/arjunmalhotra1/Grocery-Tech-Exercise/handler"
)

func main() {
	h := handler.New()
	http.ListenAndServe(":8086", h.Router)

}
