package daos

import (
	"MySolutions/app/models"
	"github.com/jinzhu/gorm"
)

func DeleteFromPageBodiesDAO(pageId int,tx *gorm.DB) error {

	res := tx.Where("page_id = ?",pageId).Delete(&models.Page_Body{})

	return res.Error

}
