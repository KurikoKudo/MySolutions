package controllers

import (
	"github.com/revel/revel"
	"strings"
	"strconv"
	"fmt"
	"database/sql"
	  _ "github.com/go-sql-driver/mysql"
	 "MySolutions/app/models"
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

var body []string
var ptitle []string
var tags []string
var ecode string = ""


if(c.Params.Form.Get("ecode") != ""){
	ecode = c.Params.Form.Get("ecode")
}


	body = strings.Split((c.Params.Form.Get("body")), " ")

	ptitle = strings.Split((c.Params.Form.Get("ptitle")), " ")

	tags = strings.Split((c.Params.Form.Get("tags")), " ")

	var list []models.Page
	list = DBSearch(body,ptitle,tags,ecode)

	listlen := len(list)




	return c.Render(list,listlen)
}


  func DBSearch(body []string,ptitle []string,tags []string, ecode string) ([]models.Page) {
    db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close() // 関数がリターンする直前に呼び出される


		sqlsentence := "SELECT * FROM pages INNER JOIN tags ON pages.page_id = tags.page_id WHERE 1"
		if(0 < len(body)){
			bodysentence := ""
			for i:=0; i < len(body); i++{
				bodysentence += " AND pages.body LIKE '%" + body[i] + "%'"
			}
			sqlsentence += bodysentence
		}

		if(0 < len(ptitle)){
			ptitlesentence := ""
			for i:=0; i<len(ptitle); i++{
				ptitlesentence += " AND pages.title LIKE '%" + ptitle[i] + "%'"
			}
			sqlsentence += ptitlesentence
		}

		if(0 < len(tags)){
			tagssentence := ""
			for i:=0; i<len(tags); i++{
				tagssentence += " AND tags.tag_name LIKE '%" + tags[i] + "%'"
			}
			sqlsentence += tagssentence
		}

		if(ecode != ""){
			ecodesentence := " AND pages.error_code '%" + ecode + "%'"
			sqlsentence += ecodesentence
		}


		//fmt.Println(sqlsentence)



    rows, err := db.Query(sqlsentence + ";") //
    if err != nil {
      panic(err.Error())
    }

    columns, err := rows.Columns() // カラム名を取得
    if err != nil {
      panic(err.Error())
    }

    values := make([]sql.RawBytes, len(columns))

    //  rows.Scan は引数に `[]interface{}`が必要.

    scanArgs := make([]interface{}, len(values))
    for i := range values {
      scanArgs[i] = &values[i]
    }

		//sliceの初期化
		var list []models.Page

    for rows.Next() {
      err = rows.Scan(scanArgs...)
      if err != nil {
        panic(err.Error())
      }

			pages := models.Page{}

      var value string
      for i, col := range values {
        // Here we can check if the value is nil (NULL value)



        if col == nil {
          value = "NULL"
        } else {
          value = string(col)
        }

				switch columns[i] {
				case "page_id":
					pages.PageId,err = strconv.Atoi(value)
				case "title":
					pages.PageTitle = value
				case "error_code":
					pages.ErrorCode = value
				case "body":
					pages.Body = value
				case "summary_id":
					pages.SummaryId,err = strconv.Atoi(value)
				case "summary_page":
					pages.SummaryPage,err = strconv.Atoi(value)
				case "importance":
					pages.Importance,err = strconv.Atoi(value)
				case "complete":
					pages.Complete,err = strconv.ParseBool(value)
				case "tag_name":
					 pages.TagName = value
				}

        fmt.Println(columns[i], ": ", value)
      }

			list = append(list, pages)

      fmt.Println("-----------------------------------")
    }

		return list;
  }
