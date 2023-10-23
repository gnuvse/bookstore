package main

import (
	"fmt"
	"log"
	"net/http"

	"go-bookstore/pkg/routes"

	"github.com/gorilla/mux"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is starting at :9000...")
	log.Fatal(http.ListenAndServe(":9000", r))
}
