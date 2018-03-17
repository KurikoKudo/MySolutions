package daos

import (
	"MySolutions/app/models"
	"fmt"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SearchPageDAO(searchingConditionStr string, searchingEvaluation uint, searchingTags []string, searchingTitles []string) []models.Page_Body {
	var pageList []models.Page_Body

	db := gormConnect()

	sqlQuery := "SELECT * FROM page_bodies WHERE 1 "

	if searchingConditionStr != "" {
		sqlQuery += "AND condition = " + searchingConditionStr + " "
	}
	if searchingEvaluation != 0 {
		str := strconv.Itoa(int(searchingEvaluation))
		sqlQuery += "AND evaluation = " + str + " "
	}
	for _, v := range searchingTitles {
		v = "%" + v + "%"
		sqlQuery += "AND page_title LIKE " + v + " "
	}
	for _, v := range searchingTags {
		v = "%" + v + "%"
		sqlQuery += "AND tags LIKE " + v + " "
	}
	fmt.Println(sqlQuery)
	db.Raw(sqlQuery).Find(&pageList)

	fmt.Println(pageList)

	return pageList
}
