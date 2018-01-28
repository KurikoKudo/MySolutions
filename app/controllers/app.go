package controllers

import (
	"github.com/revel/revel"
	"strings"
	"fmt"
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



/*POST*/
func (c App) Search() revel.Result {

var all []string
var ptitle []string
var tags []string
var ecode string


if(c.Params.Form.Get("ecode") != ""){
	ecode = c.Params.Form.Get("ecode")
}


	all = strings.Split((c.Params.Form.Get("all")), " ")

	ptitle = strings.Split((c.Params.Form.Get("ptitle")), " ")

	tags = strings.Split((c.Params.Form.Get("tags")), " ")
	fmt.Println(len(all))




	return c.Render(all,ptitle,tags,ecode)
}
