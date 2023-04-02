package service

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"
	"test/fitur/owner"
	"test/helper"
	"test/validasi"

	"github.com/go-playground/validator/v10"
)

type ownerCase struct {
	qry owner.OwnerData
	vld *validator.Validate
	ups helper.Uploads
}

func NewService(od owner.OwnerData, vld *validator.Validate, ups helper.Uploads) owner.OwnerService {
	return &ownerCase{
		qry: od,
		vld: vld,
		ups: ups,
	}
}

// Add implements family.FamilyService
func (oc *ownerCase) AddOwner(ktpfile *multipart.FileHeader, newOwner owner.OwnerEntities) (owner.OwnerEntities, error) {

	valerr := oc.vld.Struct(&newOwner)
	if valerr != nil {
		log.Println("validation error", valerr)
		msg := validasi.ValidationErrorHandle(valerr)
		return owner.OwnerEntities{}, errors.New(msg)
	}
	if ktpfile != nil {
		secureURL, err2 := oc.ups.Upload(ktpfile)
		if err2 != nil {
			log.Println(err2)
			// fmt.Print(err2)
			var msg string
			if strings.Contains(err2.Error(), "bad request") {
				msg = err2.Error()
			} else {
				msg = "failed to upload image, server error"
			}
			return owner.OwnerEntities{}, errors.New(msg)
		}
		newOwner.Ktp = secureURL
		// fmt.Print("update data image", Updata.Image_url)
	}
	res, err := oc.qry.AddOwner(newOwner)
	if err != nil {
		// fmt.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "content family not found"
		} else {
			msg = "internal server error"
		}
		return owner.OwnerEntities{}, errors.New(msg)
	}

	return res, nil
}
