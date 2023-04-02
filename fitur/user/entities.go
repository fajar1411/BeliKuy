package user

import "mime/multipart"

type CustomerEntites struct {
	ID       uint
	Name     string `validate:"required,min=5,required"`
	Dob_date string
	Phonenum string `validate:"required,min=5,required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=5,required"`
	Images   string
	Role     string
}
type AdminEntites struct {
	ID       uint
	Name     string `validate:"required,min=5,required"`
	Dob_date string
	Phonenum string `validate:"required,min=5,required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=5,required"`
	Images   string
	Role     string
}
type CostumerService interface {
	FormData(newUser CustomerEntites) (CustomerEntites, error)
	FormAdmin(newAdmin AdminEntites) (AdminEntites, error)
	Login(email, password string) (string, CustomerEntites, error)
	Profile(id int) (CustomerEntites, error)
	Update(id int, file *multipart.FileHeader, Updata CustomerEntites) (CustomerEntites, error)
	Delete(id int) error
}

type CostumerData interface {
	FormData(newUser CustomerEntites) (CustomerEntites, error)
	FormAdmin(newAdmin AdminEntites) (AdminEntites, error)
	Login(email string) (CustomerEntites, error)
	Profile(id int) (CustomerEntites, error)
	Update(id int, Updata CustomerEntites) (CustomerEntites, error)
	Delete(id int) error
}
