package daos

import (
	"github.com/jinzhu/gorm"
	"MySolutions/app/models"
)

func DeleteFromPageRelationsDAO(pageId int,tx *gorm.DB) error {

	res := tx.Where("page_id0 = ?",pageId).Or("page_id1 = ?",pageId).Delete(&models.Page_Relation{})

	return res.Error


}
