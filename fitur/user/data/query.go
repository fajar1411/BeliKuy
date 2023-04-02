package data

import (
	"errors"
	"log"
	"strings"

	"test/fitur/user"

	"gorm.io/gorm"
)

type costumerData struct {
	db *gorm.DB
}

// Profile implements user.UserData

func NewCustomer(db *gorm.DB) user.CostumerData {
	return &costumerData{
		db: db,
	}
}

// FormAdmin implements user.CostumerData
func (cc *costumerData) FormAdmin(newAdmin user.AdminEntites) (user.AdminEntites, error) {
	userGorm := FromEntitiesAdmin(newAdmin)
	userGorm.Role = "Admin"
	tx := cc.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		log.Println("register query error", tx.Error.Error())
		msg := ""
		if strings.Contains(tx.Error.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return user.AdminEntites{}, errors.New(msg)
	}
	return newAdmin, nil
}

// FormData implements customer.CostumerData
func (cd *costumerData) FormData(newUser user.CustomerEntites) (user.CustomerEntites, error) {
	userGorm := FromEntities(newUser)
	userGorm.Role = "Customer"
	tx := cd.db.Create(&userGorm) // proses insert data

	if tx.Error != nil {
		log.Println("register query error", tx.Error.Error())
		msg := ""
		if strings.Contains(tx.Error.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "server error"
		}
		return user.CustomerEntites{}, errors.New(msg)
	}
	return newUser, nil
}

// Login implements customer.CostumerData
func (cd *costumerData) Login(email string) (user.CustomerEntites, error) {
	res := User{}

	if err := cd.db.Where("email = ?", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.CustomerEntites{}, errors.New("data not found")
	}

	return ToCore(res), nil
}

// Profile implements customer.CostumerData
func (cd *costumerData) Profile(id int) (user.CustomerEntites, error) {
	users := UserName{}
	err := cd.db.Raw("SELECT users.dob_date, users.email, users.phonenum, users.name, users.images  FROM users  WHERE users.id = ?", id).Find(&users).Error

	if err != nil {
		return user.CustomerEntites{}, err
	}
	gorms := users.ModelsToCore()
	return gorms, nil

}

// Delete implements customer.CostumerData
func (cd *costumerData) Delete(id int) error {
	users := User{}
	qry := cd.db.Delete(&users, id)

	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete account fail")
	}

	return nil
}

// Update implements customer.CostumerData
func (cd *costumerData) Update(id int, Updata user.CustomerEntites) (user.CustomerEntites, error) {
	users := User{}
	userGorm := FromEntities(Updata)
	qry := cd.db.Model(&users).Where("id = ?", id).Updates(&userGorm)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return user.CustomerEntites{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return user.CustomerEntites{}, err
	}

	return ToCore(userGorm), nil
}
