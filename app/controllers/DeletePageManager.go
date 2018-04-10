package controllers

import "MySolutions/app/daos"

func DeletePageManager(pageId int) bool {

	db := daos.gormConnect()

	tx := db.Begin()

	daos.DeleteFromPageBodiesDAO(pageId,tx)


	daos.DeleteFromPageRelationsDAO(pageId,tx)


	return false
}
