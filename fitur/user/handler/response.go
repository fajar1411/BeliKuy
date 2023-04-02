package handler

import (
	"test/fitur/user"
)

type FormResponse struct {
	Name     string `json:"nama"`
	Dob_date string `json:"tanggal lahir"`
	Phonenum string `json:"phone"`
	Email    string `json:"email"`
}
type ProfileResponse struct {
	Name     string `json:"nama"`
	Dob_date string `json:"tanggal lahir"`
	Phonenum string `json:"phone"`
	Email    string `json:"email"`
	Images   string `json:"images"`
}
type LoginResponse struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToFormResponse(data user.CustomerEntites) FormResponse {
	return FormResponse{
		Name:     data.Name,
		Email:    data.Email,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
	}
}
func ToFormResponses(data user.AdminEntites) FormResponse {
	return FormResponse{
		Name:     data.Name,
		Email:    data.Email,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
	}
}
func ToProfileResponse(data user.CustomerEntites) ProfileResponse {
	return ProfileResponse{
		Name:     data.Name,
		Email:    data.Email,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
		Images:   data.Images,
	}
}
func ToLoginRespon(data user.CustomerEntites, token string) LoginResponse {
	return LoginResponse{

		Nama:  data.Name,
		Email: data.Email,
		Token: token,
	}
}
