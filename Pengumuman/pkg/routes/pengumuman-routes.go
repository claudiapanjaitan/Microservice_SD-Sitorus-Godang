package routes

import (
	"github.com/Ikyy11321036/Pengumuman/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPengumumansRoutes = func(router *mux.Router) {
	router.HandleFunc("/pengumuman/",
		controllers.CreatePengumuman).Methods("POST")
	router.HandleFunc("/pengumuman/", controllers.GetPengumumanById).Methods("GET")
	router.HandleFunc("/pengumuman/{pengumumanId}",
		controllers.GetPengumumanById).Methods("GET")
	router.HandleFunc("/pengumuman/{pengumumanId}",
		controllers.UpdatePengumuman).Methods("PUT")
	router.HandleFunc("/pengumuman/{pengumumanId}",
		controllers.DeletePengumuman).Methods("DELETE")
}
