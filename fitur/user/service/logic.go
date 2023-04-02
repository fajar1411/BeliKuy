package service

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"

	"test/fitur/user"
	"test/helper"
	"test/middlewares"
	"test/scripts"
	"test/validasi"

	"github.com/go-playground/validator/v10"
)

type costumerCase struct {
	qry user.CostumerData
	vld *validator.Validate
	ups helper.Uploads
}

func NewService(cd user.CostumerData, vld *validator.Validate, ups helper.Uploads) user.CostumerService {
	return &costumerCase{
		qry: cd,
		vld: vld,
		ups: ups,
	}
}

// FormData implements customer.CostumerService
func (cc *costumerCase) FormData(newUser user.CustomerEntites) (user.CustomerEntites, error) {
	valerr := cc.vld.Struct(&newUser)
	if valerr != nil {
		log.Println("validation error", valerr)
		msg := validasi.ValidationErrorHandle(valerr)
		return user.CustomerEntites{}, errors.New(msg)
	}
	hash := scripts.Bcript(newUser.Password)
	newUser.Password = string(hash)

	res, err := cc.qry.FormData(newUser)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg2 = "email sudah terdaftar"
		} else if strings.Contains(err.Error(), "empty") {
			msg2 = "username not allowed empty"
		} else {
			msg2 = "server error"
		}
		return user.CustomerEntites{}, errors.New(msg2)
	}

	return res, nil
}

// Login implements customer.CostumerService
func (cc *costumerCase) Login(email string, password string) (string, user.CustomerEntites, error) {
	errEmail := cc.vld.Var(email, "required,email")
	if errEmail != nil {
		log.Println("validation error", errEmail)
		msg := validasi.ValidationErrorHandle(errEmail)
		return "", user.CustomerEntites{}, errors.New(msg)
	}
	res, err := cc.qry.Login(email)
	if err != nil {
		log.Println("query login error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "no rows") {
			msg = "email belum terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.CustomerEntites{}, errors.New(msg)
	}
	errPw := cc.vld.Var(password, "required,min=5,required")
	if errPw != nil {
		log.Println("validation error", errPw)
		msg := validasi.ValidationErrorHandle(errPw)
		return "", user.CustomerEntites{}, errors.New(msg)
	}
	if err := scripts.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.CustomerEntites{}, errors.New("password tidak sesuai" + res.Password)
	}

	//Token expires after 1 hour
	token, _ := middlewares.CreateToken(int(res.ID), res.Role)

	return token, res, nil

}

// Profile implements customer.CostumerService
func (cc *costumerCase) Profile(id int) (user.CustomerEntites, error) {
	if id <= 0 {
		log.Println("User belum terdaftar")
	}
	res, err := cc.qry.Profile(id)
	if err != nil {
		log.Println(err)
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user tidak ditemukan harap login lagi"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.CustomerEntites{}, errors.New(msg)
	}
	return res, nil
}

// Delete implements customer.CostumerService
func (cc *costumerCase) Delete(id int) error {
	if id <= 0 {
		log.Println("User belum terdaftar")
	}
	err := cc.qry.Delete(id)

	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}

	return nil
}

// Update implements customer.CostumerService
func (cc *costumerCase) Update(id int, file *multipart.FileHeader, Updata user.CustomerEntites) (user.CustomerEntites, error) {
	if id <= 0 {
		log.Println("User belum terdaftar")
	}
	dobdate := Updata.Dob_date
	if dobdate != "" {
		errdob := cc.vld.Var(dobdate, "required")
		if errdob != nil {
			log.Println("validation error", errdob)
			msg := validasi.ValidationErrorHandle(errdob)
			return user.CustomerEntites{}, errors.New(msg)
		}
	}
	email := Updata.Email
	if email != "" {
		errEmail := cc.vld.Var(email, "required,email")
		if errEmail != nil {
			log.Println("validation error", errEmail)
			msg := validasi.ValidationErrorHandle(errEmail)
			return user.CustomerEntites{}, errors.New(msg)
		}
	}
	name := Updata.Name
	if name != "" {
		errName := cc.vld.Var(name, "required,min=5,required")
		if errName != nil {
			log.Println("validation error", errName)
			msg := validasi.ValidationErrorHandle(errName)
			return user.CustomerEntites{}, errors.New(msg)
		}
	}
	phone := Updata.Phonenum
	if phone != "" {
		errphone := cc.vld.Var(phone, "required,min=5,required")
		if errphone != nil {
			log.Println("validation error", errphone)
			msg := validasi.ValidationErrorHandle(errphone)
			return user.CustomerEntites{}, errors.New(msg)
		}
	}
	pw := Updata.Password
	if pw != "" {
		errPw := cc.vld.Var(pw, "required,min=5,required")
		if errPw != nil {
			log.Println("validation error", errPw)
			msg := validasi.ValidationErrorHandle(errPw)
			return user.CustomerEntites{}, errors.New(msg)
		} else {
			hash := scripts.Bcript(Updata.Password)
			Updata.Password = string(hash)
		}

	}
	if file != nil {
		secureURL, err2 := cc.ups.Upload(file)
		if err2 != nil {
			log.Println(err2)
			// fmt.Print(err2)
			var msg string
			if strings.Contains(err2.Error(), "bad request") {
				msg = err2.Error()
			} else {
				msg = "failed to upload image, server error"
			}
			return user.CustomerEntites{}, errors.New(msg)
		}
		Updata.Images = secureURL
		// fmt.Print("update data image", Updata.Image_url)
	}
	res, err := cc.qry.Update(id, Updata)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg2 = "email sudah terdaftar"
		} else if strings.Contains(err.Error(), "empty") {
			msg2 = "username not allowed empty"
		} else {
			msg2 = "server error"
		}
		return user.CustomerEntites{}, errors.New(msg2)
	}

	return res, nil
}

// FormAdmin implements user.CostumerService
func (cc *costumerCase) FormAdmin(newAdmin user.AdminEntites) (user.AdminEntites, error) {
	valerr := cc.vld.Struct(&newAdmin)
	if valerr != nil {
		log.Println("validation error", valerr)
		msg := validasi.ValidationErrorHandle(valerr)
		return user.AdminEntites{}, errors.New(msg)
	}
	hash := scripts.Bcript(newAdmin.Password)
	newAdmin.Password = string(hash)

	res, err := cc.qry.FormAdmin(newAdmin)
	if err != nil {
		msg2 := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg2 = "email sudah terdaftar"
		} else if strings.Contains(err.Error(), "empty") {
			msg2 = "username not allowed empty"
		} else {
			msg2 = "server error"
		}
		return user.AdminEntites{}, errors.New(msg2)
	}

	return res, nil
}
