package daos

import (
	"MySolutions/app/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DisplayPage(pageId int) models.Page_Display {
	var displayPage models.Page_Display

	db := gormConnect()

	db.First(&displayPage.Page_Body, pageId)
	db.Where("page_id0 = ? OR page_id1 = ?", pageId, pageId).Find(&displayPage.Page_Relations)

	return displayPage
}

func DisplayRelation(relations []uint) []models.Page_Body {
	db := gormConnect()

	var relationPages []models.Page_Body

	db.Select("page_id,page_title").Where("page_id in (?)", relations).Find(&relationPages)

	return relationPages
}
