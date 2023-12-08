package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ikyy11321036/Jadwal/pkg/models"
	"github.com/Ikyy11321036/Jadwal/pkg/utils"
	"github.com/gorilla/mux"
)

var NewPengumuman models.Pengumuman

func GetPemgumuman(w http.ResponseWriter, r *http.Request) {
	newPengumumans := models.GetAllPengumumans()
	res, _ := json.Marshal(newPengumumans)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetPengumumanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pengumumanId := vars["pengumumanId"]
	ID, err := strconv.ParseInt(pengumumanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pengumumanDetails, _ := models.GetPengumumanbyId(ID)
	res, _ := json.Marshal(pengumumanDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreatePengumuman(w http.ResponseWriter, r *http.Request) {
	CreatePengumuman := &models.Pengumuman{}
	utils.ParseBody(r, CreatePengumuman)
	b := CreatePengumuman.CreatePengumuman()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeletePengumuman(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pengumumanId := vars["pengumumanId"]
	ID, err := strconv.ParseInt(pengumumanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pengumuman := models.DeletePengumuman(ID)
	res, _ := json.Marshal(pengumuman)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdatePengumuman(w http.ResponseWriter, r *http.Request) {
	var UpdatePengumuman = &models.Pengumuman{}
	utils.ParseBody(r, UpdatePengumuman)
	vars := mux.Vars(r)
	pengumumanId := vars["pengumumanId"]
	ID, err := strconv.ParseInt(pengumumanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pengumumanDetails, db := models.GetPengumumanbyId(ID)
	if UpdatePengumuman.Hari_tanggal != "" {
		pengumumanDetails.Hari_tanggal = UpdatePengumuman.Hari_tanggal
	}
	if UpdatePengumuman.Judul != "" {
		pengumumanDetails.Judul = UpdatePengumuman.Judul
	}
	if UpdatePengumuman.Deskripsi != "" {
		pengumumanDetails.Deskripsi = UpdatePengumuman.Deskripsi
	}
	db.Save(&pengumumanDetails)
	res, _ := json.Marshal(pengumumanDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
