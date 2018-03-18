package controllers

import (
	"MySolutions/app/daos"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

func (c App) SearchManager(ptitle []string, tags []string, evaluation uint, condition bool) revel.Result {
	var searchingEvaluation uint
	var searchingTags []string
	var searchingTitles []string

	if c.Params.Form.Get("ptitle") != "" {
		searchingTitles = strings.Split((c.Params.Form.Get("ptitle")), ",")
	}

	if c.Params.Form.Get("tags") != "" {
		searchingTags = strings.Split((c.Params.Form.Get("tags")), ",")
	}

	if c.Params.Form.Get("evaluation") != "" {
		evaluationInput, err := strconv.ParseUint(c.Params.Form.Get("evaluation"), 10, 0)
		if err != nil {
			return c.Redirect(App.Home)
		}
		searchingEvaluation = uint(evaluationInput)
	}

	//nil or false 判別のためstring取得
	conditionInputStr := c.Params.Form.Get("condition")

	searchedPageList := daos.SearchPageDAO(conditionInputStr, searchingEvaluation, searchingTags, searchingTitles)
	listlen := len(searchedPageList)
	return c.Render(searchedPageList, listlen)

}
