package data

import (
	"test/fitur/owner"
	"test/fitur/user/data"

	"gorm.io/gorm"
)

type Owner struct {
	gorm.Model
	Nama_Toko string `gorm:"type:varchar(50);not null"`
	Ktp       string
	Alamat    string `gorm:"type:varchar(50);not null"`
	Status    string
	UserID    uint
	User      data.User
}
type OwnerUser struct {
	ID        uint
	Nama_Toko string
	Ktp       string
	Alamat    string
	Status    string
}

func Todata(data owner.OwnerEntities) Owner {
	return Owner{
		Nama_Toko: data.Nama_Toko,
		Ktp:       data.Ktp,
		Alamat:    data.Alamat,
		Status:    data.Status,
		UserID:    data.UserID,
	}
}
func (dataModel *OwnerUser) ModelsToCore() owner.OwnerEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return owner.OwnerEntities{
		ID:        dataModel.ID,
		Nama_Toko: dataModel.Nama_Toko,
		Alamat:    dataModel.Alamat,
		Status:    dataModel.Status,
	}
}
func ToCore(data Owner) owner.OwnerEntities {
	return owner.OwnerEntities{
		ID:        data.ID,
		Nama_Toko: data.Nama_Toko,
		Status:    data.Status,
		Alamat:    data.Alamat,
		Ktp:       data.Ktp,
		UserID:    data.UserID,
	}
}
func ListModelTOCore(dataModel []OwnerUser) []owner.OwnerEntities { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	var dataCore []owner.OwnerEntities
	for _, value := range dataModel {
		dataCore = append(dataCore, value.ModelsToCore())
	}
	return dataCore //  untuk menampilkan data ke controller
}
