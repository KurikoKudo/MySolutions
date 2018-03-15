package daos

import (
	"MySolutions/app/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SearchPage(newPage models.Page_Body) []models.Page_Display {
  var page models.Page_Display

  db := gormConnect()


	//db.joins("JOIN page_bodies ON page_bodies.page_id = tags.page_id ").Where("page_title in (?) AND evaluation = ? AND ", ptitle).Find(&page)

  return nil
}
