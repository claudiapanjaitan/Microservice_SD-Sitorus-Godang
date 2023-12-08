package models

import (
	"github.com/Ikyy11321036/TambahAkun/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type TambahAkun struct {
	gorm.Model
	ID            string `gorm:""json:"id"`
	Nisn          string `json:"nisn"`
	Nip           string `json:"nip"`
	Username      string `json:"username"`
	Kelas         string `json:"kelas"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Tanggal_Lahir string `json:"tgl_lahir"`
	Tempat_Lahir  string `json:"tpt_lahir"`
	Jabatan       string `json:"jabatan"`
	Agama         string `json:"agama"`
	Alamat        string `json:"alamat"`
	No_Hp         string `json:"no_telpon"`
	Password      string `json:"password"`
	Role          string `json:"role"`
	Foto          string `json:"foto"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&TambahAkun{})
}
func (b *TambahAkun) CreateTambahAkun() *TambahAkun {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllTambahAkuns() []TambahAkun {
	var TambahAkuns []TambahAkun
	db.Find(&TambahAkuns)
	return TambahAkuns
}
func GetTambahAkunbyId(id int64) (*TambahAkun, *gorm.DB) {
	var GetTambahAkun TambahAkun
	db := db.Where("ID=?", id).Find(&GetTambahAkun)
	return &GetTambahAkun, db
}
func DeleteTambahAkun(id int64) TambahAkun {
	var tambahakun TambahAkun
	db.Where("ID=?", id).Delete(tambahakun)
	return tambahakun
}
