package data

import (
	"test/fitur/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:char(50);not null"`
	Dob_date string `gorm:"not null"`
	Phonenum string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(50);unique;not null"`
	Password string `gorm:"not null"`
	Images   string
	Role     string
	Owners   []Owner `gorm:"foreignkey:CustomerID"`
}

type Owner struct {
	gorm.Model
	Nama_Toko  string `gorm:"type:varchar(50);not null"`
	Ktp        string
	Alamat     string `gorm:"type:varchar(50);not null"`
	Status     string
	CustomerID uint
}
type UserName struct {
	ID               uint
	Dob_date         string
	Phonenum         string
	Email            string
	Name             string
	Password         string
	Images           string
	Role             string
	Nationality_name string
}

// register
func FromEntities(dataCore user.CustomerEntites) User { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	return User{

		Email:    dataCore.Email,
		Dob_date: dataCore.Dob_date,
		Phonenum: dataCore.Phonenum,
		Name:     dataCore.Name,
		Password: dataCore.Password,
		Images:   dataCore.Images,
		Role:     dataCore.Role,
	}

}
func FromEntitiesAdmin(dataCore user.AdminEntites) User { //fungsi yang mengambil data dari entities usercore dan merubah data ke user gorm(model.go)
	return User{

		Email:    dataCore.Email,
		Dob_date: dataCore.Dob_date,
		Phonenum: dataCore.Phonenum,
		Name:     dataCore.Name,
		Password: dataCore.Password,
		Images:   dataCore.Images,
		Role:     dataCore.Role,
	}

}

// update dan login
func ToCore(data User) user.CustomerEntites {
	return user.CustomerEntites{
		ID:       data.ID,
		Email:    data.Email,
		Dob_date: data.Dob_date,
		Phonenum: data.Phonenum,
		Name:     data.Name,
		Images:   data.Images,
		Password: data.Password,
	}
}

// profile user
func (dataModel *UserName) ModelsToCore() user.CustomerEntites { //fungsi yang mengambil data dari  user gorm(model.go)  dan merubah data ke entities usercore
	return user.CustomerEntites{
		Name:     dataModel.Name,
		Email:    dataModel.Email, //mapping data core ke data gorm model
		Password: dataModel.Password,
		Dob_date: dataModel.Dob_date,
		Phonenum: dataModel.Phonenum,
		Images:   dataModel.Images,
	}
}
