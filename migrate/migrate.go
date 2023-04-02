package migrate

import (
	owner "test/fitur/owner/data"
	user "test/fitur/user/data"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&owner.Owner{})

}
