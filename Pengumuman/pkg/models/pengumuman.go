package models

import (
	"github.com/Ikyy11321036/Pengumuman/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Pengumuman struct {
	gorm.Model
	ID           string `gorm:""json:"id"`
	Hari_tanggal string `json:"hari_tanggal"`
	Judul        string `json:"judul"`
	Deskripsi    string `json:"deskripsi"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Pengumuman{})
}
func (b *Pengumuman) CreatePengumuman() *Pengumuman {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllPengumumans() []Pengumuman {
	var Pengumumans []Pengumuman
	db.Find(&Pengumumans)
	return Pengumumans
}
func GetPengumumanbyId(id int64) (*Pengumuman, *gorm.DB) {
	var GetPemgumuman Pengumuman
	db := db.Where("ID=?", id).Find(&GetPemgumuman)
	return &GetPemgumuman, db
}
func DeletePengumuman(id int64) Pengumuman {
	var pengumuman Pengumuman
	db.Where("ID=?", id).Delete(pengumuman)
	return pengumuman
}
