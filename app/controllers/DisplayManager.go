package controllers

import (
	"MySolutions/app/daos"
	"MySolutions/app/helpers"
	"MySolutions/app/models"
)

func DisplayManager(pageId int) models.Page_Display {

	page := daos.DisplayPage(pageId)

	//Tags -> Tag_List
	page.Tag_List = helpers.TagConverter(page.Page_Body.Tags)

	var relations []uint
	for _, v := range page.Page_Relations {
		if v.Page_id0 == uint(pageId) {
			relations = append(relations, v.Page_id1)
		} else {
			relations = append(relations, v.Page_id0)
		}
	}

	//md->html
	//page.Page_Body.Page_Body = string(blackfriday.MarkdownBasic([]byte(page.Page_Body.Page_Body)))

	//relations -> []relationPages
	page.Page_Relation_Links = daos.DisplayRelation(relations)

	return page

}
