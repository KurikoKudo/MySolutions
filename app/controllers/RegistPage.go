package controllers

import (
	"MySolutions/app/daos"
	"MySolutions/app/models"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

func (c App) RegistPage() revel.Result {
	var newPage models.Page_Body

	c.Validation.Required(c.Params.Form.Get("ptitle")).Message("タイトルを入力してください")
	c.Validation.Required(c.Params.Form.Get("solutions")).Message("タイトルを入力してください")
	c.Validation.Required(c.Params.Form.Get("tags")).Message("タグを１つ以上入力してください")
	c.Validation.Required(c.Params.Form.Get("condition")).Message("conditionを入力してください")
	c.Validation.Required(c.Params.Form.Get("evaluation")).Message("evaluationを入力してください")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.RegistForm)
	}

	newPage.Page_Body = c.Params.Form.Get("solutions")
	newPage.Page_Title = c.Params.Form.Get("ptitle")
	tagList := strings.Split((c.Params.Form.Get("tags")), ",")
	for i, v := range tagList {
		tagList[i] = " " + v + "\n"
	}
	evaluation, err := strconv.ParseUint(c.Params.Form.Get("evaluation"), 10, 0)
	if err != nil {
		return error
	}
	if 1 <= evaluation && evaluation <= 5 {
		fmt.Println("evaluationの入力が不正です")
		return error
	}
	condition, err := strconv.ParseBool(c.Params.Form.Get("condition"))
	if err != nil {
		return error
	}
	newPage.Evaluation = uint(evaluation)
	newPage.Condition = condition
	newPage.Tag = tagList

	fmt.Println(newPage)

	registedPageId := daos.RegisterNewPage(newPage)

	return c.Render(registedPageId)

}
