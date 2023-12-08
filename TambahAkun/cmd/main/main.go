package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ikyy11321036/TambahAkun/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTambahAkunsRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9020")
	log.Fatal(http.ListenAndServe("localhost:9020", r))
}
