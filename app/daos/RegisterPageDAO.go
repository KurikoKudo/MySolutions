package daos

import (
	"MySolutions/app/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegisterNewPage(newPage models.Page_Body) uint {
	db := GormConnect()

	db.Create(&newPage)
	db.First(&newPage)

	return newPage.Page_Id
}
