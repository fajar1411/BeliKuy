package handler

import (
	"test/fitur/user"
)

type CostumerRequest struct {
	Name     string `json:"nama" form:"nama"`
	Dob_date string `json:"tanggal" form:"tanggal"`
	Phonenum string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Images   string `json:"images" form:"images"`
}
type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type AdminRequest struct {
	Name     string `json:"nama" form:"nama"`
	Dob_date string `json:"tanggal" form:"tanggal"`
	Phonenum string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Images   string `json:"images" form:"images"`
}

func CostumerRequestToUserCore(data CostumerRequest) user.CustomerEntites {
	return user.CustomerEntites{
		Name:     data.Name,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
		Email:    data.Email,
		Password: data.Password,
	}
}
func AdminRequestToUserCore(data CostumerRequest) user.AdminEntites {
	return user.AdminEntites{
		Name:     data.Name,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
		Email:    data.Email,
		Password: data.Password,
	}
}
