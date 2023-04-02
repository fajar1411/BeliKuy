package handler

import (
	"test/fitur/owner"
)

type OwnerRequest struct {
	Nama_toko string `json:"namatoko" form:"namatoko"`
	Alamat    string `json:"alamat" form:"alamat"`
	Ktp       string `json:"ktp" form:"ktp"`
}

func OwnerRequestToEnitities(data OwnerRequest) owner.OwnerEntities {
	return owner.OwnerEntities{
		Nama_Toko: data.Nama_toko,
		Ktp:       data.Ktp,
		Alamat:    data.Alamat,
	}
}
