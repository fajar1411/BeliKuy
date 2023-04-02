package handler

import (
	"test/fitur/owner"
)

type OwnerResponse struct {
	Nama_Toko string `json:"nama toko"`
	Alamat    string `json:"alamat"`
	Ktp       string `json:"pemilik ktp"`
}

func ToFormResponse(data owner.OwnerEntities) OwnerResponse {
	return OwnerResponse{
		Nama_Toko: data.Nama_Toko,
		Ktp:       data.Ktp,
		Alamat:    data.Alamat,
	}
}

// func ListEntitiesToPostsRespon(dataCore []family.FamilyEntities) []FamilyResponse {
// 	var ResponData []FamilyResponse

// 	for _, value := range dataCore {
// 		ResponData = append(ResponData, ToFormResponse(value))
// 	}
// 	return ResponData
// }
