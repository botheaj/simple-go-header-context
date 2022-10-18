package main

import (
	"fmt"
	"github.com/botheaj/goauth/auth"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	err := http.ListenAndServe(":8080", auth.Middleware(router))
	if err != nil {
		fmt.Println("oops error")
	}
}
