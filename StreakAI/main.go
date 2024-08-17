package main

import (
	"Task/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()

	fmt.Println("server begin")

	r.Use(middleware.Logger)
	r.Post("/find-pairs", handler.FindPairsHandler)
	if err := http.ListenAndServe(":8086", r); err != nil {
		fmt.Println("error serving on port", err)
	}
	fmt.Println("Server listening to 8086")

}
