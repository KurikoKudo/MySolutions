package controllers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	text := "It Works!"
	//daos.Migration()
	//fmt.Println("Index")
	return c.Render(text)
}

func (c App) Home() revel.Result {
	text := "- Welcome to YourSolutions -"
	return c.Render(text)

}

func (c App) RegistForm() revel.Result {
	//将来的にユーザが複数いる場合ここでユーザ情報の保持とかする?
	return c.Render()

}

func (c App) PageDisplay(pageId int) revel.Result {

	page := DisplayManager(pageId)

	return c.Render(page)
}

func (c App) DeletePage(pageId int) revel.Result {

	if DeletePageManager(pageId) {
		return c.Render(pageId)
	} else {
		return c.RenderText("削除に失敗しました。")
	}



}
