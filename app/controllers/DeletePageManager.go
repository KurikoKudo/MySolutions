package controllers

import "MySolutions/app/daos"

func DeletePageManager(pageId int) bool {

	db := daos.GormConnect()

	tx := db.Begin()


	pageBodyErr := daos.DeleteFromPageBodiesDAO(pageId,tx)

	if pageBodyErr != nil {
		tx.Rollback()
		return false
	}

	relationsErr := daos.DeleteFromPageRelationsDAO(pageId,tx)

	if relationsErr != nil {
		tx.Rollback()
		return false
	}

	return true
}
