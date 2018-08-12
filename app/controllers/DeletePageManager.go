package controllers

import (
	"MySolutions/app/daos"
	"fmt"
)

func DeletePageManager(pageId int) bool {

	db := daos.GormConnect()

	tx := db.Begin()

	fmt.Println(pageId)

	pageBodyErr := daos.DeleteFromPageBodiesDAO(pageId,tx)

	if pageBodyErr != nil {
		tx.Rollback()
		fmt.Println("bodiesのエラー")
		return false
	}

	relationsErr := daos.DeleteFromPageRelationsDAO(pageId,tx)

	if relationsErr != nil {
		tx.Rollback()
		fmt.Println("relationのエラー")
		return false
	}

	tx.Commit()
	return true
}
