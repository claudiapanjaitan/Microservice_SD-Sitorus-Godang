package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ikyy11321036/TambahAkun/pkg/models"
	"github.com/Ikyy11321036/TambahAkun/pkg/utils"
	"github.com/gorilla/mux"
)

var NewTambahAkun models.TambahAkun

func GetTambahAkun(w http.ResponseWriter, r *http.Request) {
	newTambahAkuns := models.GetAllTambahAkuns()
	res, _ := json.Marshal(newTambahAkuns)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetTamabahAkunById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tambahakunId := vars["tambahakunId"]
	ID, err := strconv.ParseInt(tambahakunId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	tambahakunDetails, _ := models.GetTambahAkunbyId(ID)
	res, _ := json.Marshal(tambahakunDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateTambahAkun(w http.ResponseWriter, r *http.Request) {
	CreateTambahAkun := &models.TambahAkun{}
	utils.ParseBody(r, CreateTambahAkun)
	b := CreateTambahAkun.CreateTambahAkun()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteTambahAkun(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tambahakunId := vars["tambahakunId"]
	ID, err := strconv.ParseInt(tambahakunId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	tambahakun := models.DeleteTambahAkun(ID)
	res, _ := json.Marshal(tambahakun)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateTambahAkun(w http.ResponseWriter, r *http.Request) {
	var UpdateTambahAkun = &models.TambahAkun{}
	utils.ParseBody(r, UpdateTambahAkun)
	vars := mux.Vars(r)
	tambahakunId := vars["tambahakunId"]
	ID, err := strconv.ParseInt(tambahakunId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	tambahakunDetails, db := models.GetTambahAkunbyId(ID)
	if UpdateTambahAkun.Nisn != "" {
		tambahakunDetails.Nisn = UpdateTambahAkun.Nisn
	}
	if UpdateTambahAkun.Nip != "" {
		tambahakunDetails.Nip = UpdateTambahAkun.Nip
	}
	if UpdateTambahAkun.Username != "" {
		tambahakunDetails.Username = UpdateTambahAkun.Username
	}
	if UpdateTambahAkun.Kelas != "" {
		tambahakunDetails.Kelas = UpdateTambahAkun.Kelas
	}
	if UpdateTambahAkun.Jenis_kelamin != "" {
		tambahakunDetails.Jenis_kelamin = UpdateTambahAkun.Jenis_kelamin
	}
	if UpdateTambahAkun.Tanggal_Lahir != "" {
		tambahakunDetails.Tanggal_Lahir = UpdateTambahAkun.Tanggal_Lahir
	}
	if UpdateTambahAkun.Tempat_Lahir != "" {
		tambahakunDetails.Tempat_Lahir = UpdateTambahAkun.Tempat_Lahir
	}
	if UpdateTambahAkun.Jabatan != "" {
		tambahakunDetails.Jabatan = UpdateTambahAkun.Jabatan
	}
	if UpdateTambahAkun.Agama != "" {
		tambahakunDetails.Agama = UpdateTambahAkun.Agama
	}
	if UpdateTambahAkun.Alamat != "" {
		tambahakunDetails.Alamat = UpdateTambahAkun.Alamat
	}
	if UpdateTambahAkun.No_Hp != "" {
		tambahakunDetails.No_Hp = UpdateTambahAkun.No_Hp
	}
	if UpdateTambahAkun.Password != "" {
		tambahakunDetails.Password = UpdateTambahAkun.Password
	}
	if UpdateTambahAkun.Role != "" {
		tambahakunDetails.Role = UpdateTambahAkun.Role
	}
	if UpdateTambahAkun.Foto != "" {
		tambahakunDetails.Foto = UpdateTambahAkun.Foto
	}

	db.Save(&tambahakunDetails)
	res, _ := json.Marshal(tambahakunDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
