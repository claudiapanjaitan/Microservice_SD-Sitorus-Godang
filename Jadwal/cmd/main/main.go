package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ikyy11321036/Jadwal/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterPengumumansRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
