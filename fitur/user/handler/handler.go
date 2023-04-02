package handler

import (
	"fmt"
	"log"
	"net/http"
	"test/fitur/user"
	"test/helper"
	"test/middlewares"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	CostumerServices user.CostumerService
}

func (cd *CustomerHandler) FormData(c echo.Context) error {

	Inputform := CostumerRequest{}

	errbind := c.Bind(&Inputform)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	dataCore := CostumerRequestToUserCore(Inputform)
	res, err := cd.CostumerServices.FormData(dataCore)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Mengisi Form Berhasil", dataResp))

}
func (cd *CustomerHandler) Login(c echo.Context) error {
	input := LoginRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}

	token, res, err := cd.CostumerServices.Login(input.Email, input.Password)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToLoginRespon(res, token)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Login berhasil", dataResp))
}

func (cd *CustomerHandler) Profile(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	res, err := cd.CostumerServices.Profile(id)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToProfileResponse(res)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Melihat Profile Berhasil", dataResp))
}
func (cd *CustomerHandler) Update(c echo.Context) error {
	input := CostumerRequest{}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	id := middlewares.ExtractTokenUserId(c)
	formHeader, errimg := c.FormFile("images")
	fmt.Print("image handler", formHeader.Filename)
	if errimg != nil {
		log.Println(errimg)
	}
	res, err := cd.CostumerServices.Update(id, formHeader, CostumerRequestToUserCore(input))
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToFormResponse(res)
	return c.JSON(http.StatusOK, helper.PesanDataBerhasilHelper("Update berhasil", dataResp))
}
func (cd *CustomerHandler) Delete(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	err := cd.CostumerServices.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}
	return c.NoContent(204)
}
func (cd *CustomerHandler) FormAdmin(c echo.Context) error {

	Inputform := CostumerRequest{}

	errbind := c.Bind(&Inputform)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, "format inputan salah")
	}
	dataCore := AdminRequestToUserCore(Inputform)
	res, err := cd.CostumerServices.FormAdmin(dataCore)
	if err != nil {
		return c.JSON(helper.PesanGagalHelper(err.Error()))
	}
	dataResp := ToFormResponses(res)
	return c.JSON(http.StatusCreated, helper.PesanDataBerhasilHelper("Mengisi Form Berhasil", dataResp))

}
