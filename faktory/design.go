package faktory

import (
	"test/config"
	ownerdata "test/fitur/owner/data"
	ownerservice "test/fitur/owner/service"
	customerdata "test/fitur/user/data"
	customerservice "test/fitur/user/service"
	"test/helper"
	customerhandler "test/routes"
	ownerhandler "test/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	v := validator.New()
	cfg := config.GetConfig()
	cld := helper.NewCloud(cfg)
	userRepofaktory := customerdata.NewCustomer(db)
	userServiceFaktory := customerservice.NewService(userRepofaktory, v, cld)
	customerhandler.NewHandlerCostumer(userServiceFaktory, e)

	ownerRepofaktory := ownerdata.NewOwner(db)
	ownerServiceFaktory := ownerservice.NewService(ownerRepofaktory, v, cld)
	ownerhandler.NewHandlerOwner(ownerServiceFaktory, e)

}
