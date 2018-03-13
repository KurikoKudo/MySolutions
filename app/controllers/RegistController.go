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

func (c App) RegistForm() revel.Result {
	/*将来的にユーザが複数いる場合ここでユーザ情報の保持とかする?*/
	return c.Render()

}

func (c App) RegistPage() revel.Result {
	var newPage models.Page_Body

	/*c.Validation.Required(c.Params.Form.Get("ptitle")).Message("タイトルを入力してください")
	c.Validation.Required(c.Params.Form.Get("tags")).Message("タグを１つ以上入力してください")
	c.Validation.Required(c.Params.Form.Get("condition")).Message("conditionを入力してください")
	c.Validation.Required(c.Params.Form.Get("evaluation")).Message("evaluationを入力してください")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.RegistForm)
	}*/

	newPage.Page_Title = c.Params.Form.Get("ptitle")
	tagList := strings.Split((c.Params.Form.Get("tags")), " ")
	evaluation, err := strconv.ParseUint(c.Params.Form.Get("evaluation"), 10, 0)
	if err != nil {
		panic(err.Error())
	}
	condition, err := strconv.ParseBool(c.Params.Form.Get("condition"))
	if err != nil {
		panic(err.Error())
	}
	newPage.Evaluation = uint(evaluation)
	newPage.Condition = condition

	fmt.Println(newPage, tagList)

	registedPageId := daos.RegisterNewPage(newPage)
	daos.RegisterNewTags(registedPageId, tagList)

	return c.Render(registedPageId)

}
