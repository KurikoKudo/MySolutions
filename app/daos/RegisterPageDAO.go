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

func RegisterNewTags(pageId uint, tagNameList []string) {
	db := gormConnect()

	var tagList []models.Tag
	for _, v := range tagNameList {
		var tag models.Tag
		tag.Page_Id = pageId
		tag.Tag_Name = v

		tagList = append(tagList, tag)
	}

	for _, v := range tagList {
		db.Create(&v)
	}

}
