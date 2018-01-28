package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	text := "Hello World!"
	return c.Render(text)
}

func (c App) Home() revel.Result {
	text := "- welcome to YourSolutions -"
	return c.Render(text)
}
