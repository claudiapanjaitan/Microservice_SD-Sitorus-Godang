package routes

import (
	"github.com/Ikyy11321036/TambahAkun/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTambahAKunsRoutes = func(router *mux.Router) {
	router.HandleFunc("/tambahakun/",
		controllers.CreateTambahAkun).Methods("POST")
	router.HandleFunc("/tambahakun/", controllers.GetTambahAkunById).Methods("GET")
	router.HandleFunc("/tambahakun/{tambahakunId}",
		controllers.GetTambahAkunsById).Methods("GET")
	router.HandleFunc("/tambahakun/{tambahakunId}",
		controllers.UpdateTambahAkun).Methods("PUT")
	router.HandleFunc("/tambahakun/{tambahakunId}",
		controllers.DeleteTambahAkun).Methods("DELETE")
}
