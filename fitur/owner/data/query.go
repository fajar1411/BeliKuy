package data

import (
	"errors"
	"log"
	"strings"

	"test/fitur/owner"
	"test/fitur/user/data"

	"gorm.io/gorm"
)

type ownerData struct {
	db *gorm.DB
}

func NewOwner(db *gorm.DB) owner.OwnerData {
	return &ownerData{
		db: db,
	}
}

// Add implements family.FamilyData
func (od *ownerData) AddOwner(newOwner owner.OwnerEntities) (owner.OwnerEntities, error) {
	// var customer data.Customer
	var user data.User
	data := Todata(newOwner)
	data.Status = "Aktifasi"
	err := od.db.Create(&data)
	if err.Error != nil {
		log.Println("add Owner query error", err.Error.Error())
		msg := ""
		if strings.Contains(err.Error.Error(), "not valid") {
			msg = "wrong input"

		} else {
			msg = "server error"
		}
		return owner.OwnerEntities{}, errors.New(msg)
	}
	erruser := od.db.Model(&user).Where("id", data.UserID).Update("Role", "Owner").Error
	if erruser != nil {
		log.Println("update role query error", erruser.Error())
		return owner.OwnerEntities{}, erruser
	}
	return newOwner, nil
}
