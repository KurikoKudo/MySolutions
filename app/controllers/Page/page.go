package controllers

import (
	"MySolutions/app/models"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

type Page struct {
	*revel.Controller
}

func (h Page) Regist() revel.Result {
	var newPage models.Page_Body

	newPage.Page_Title = h.Params.Form.Get("ptitle")
	tagList := strings.Split((h.Params.Form.Get("tags")), " ")
	evaluation, err := strconv.ParseUint(h.Params.Form.Get("evaluation"), 10, 0)
	if err != nil {
		panic(err.Error())
	}
	condition, err := strconv.ParseBool(h.Params.Form.Get("condition"))
	if err != nil {
		panic(err.Error())
	}
	newPage.Evaluation = uint(evaluation)
	newPage.Condition = condition

	return h.Render(newPage, tagList)

}
