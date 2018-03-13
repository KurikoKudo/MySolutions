package controllers

import (
	"MySolutions/app/models"

	_ "github.com/go-sql-driver/mysql"
)

func SearchController(ptitle []string, tags []string, evaluation uint, condition bool) []models.Page_Display {
	var searchPage models.Page_Display

	//searchPage.Page_Body.Page_Title = ptitle
	searchPage.Page_Body.Evaluation = evaluation
	searchPage.Page_Body.Condition = condition

	var list []models.Page_Display

	return list

}
