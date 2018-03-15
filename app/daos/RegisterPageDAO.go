package daos

import (
	"MySolutions/app/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegisterNewPage(newPage models.Page_Body) uint {
	db := gormConnect()

	db.Create(&newPage)
	db.First(&newPage)

	return newPage.Page_Id
}
