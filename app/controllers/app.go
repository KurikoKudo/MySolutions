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
	fmt.Println("長さ：",listlen)

	return c.Render(list,listlen)
}

func (c App) PageDisplay(pageId int) revel.Result {

	var page models.Page
	page = DBDisplay(pageId)

	var tagList []string
	tagList = DBTags(pageId)
	page.TagName = tagList

	var relations []int
	relations = DBRelations(pageId)
	page.Relation = relations

	var relationPage []models.Title
	relationPage = DBTitlelist(relations)
	page.RelationPage = relationPage

	var references []models.Reference
	references = DBReferences(pageId)
	page.Reference = references

	return c.Render(page)
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


		fmt.Println(sqlsentence)



    rows, err := db.Query(sqlsentence) //
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
					 pages.TagName = append(pages.TagName,value)
				}

        //fmt.Println(columns[i], ": ", value)
      }

			if (contains(list,pages.PageId) >= 0){ //タグが２つ以上マッチ
				 list[contains(list,pages.PageId)].TagName = append(list[contains(list,pages.PageId)].TagName,pages.TagName[0])
			} else {
				list = append(list, pages)
			}

      //fmt.Println("-----------------------------------")
    }

		return list;
  }

	func contains(s []models.Page, e int) int{
		i := 0
	for _, v := range s {
		if e == v.PageId {
			return i
		}
		i += 1
	}
	return -1
}

	func DBDisplay(pageId int) models.Page {
		db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close() // 関数がリターンする直前に呼び出される

		rows, err := db.Query("SELECT * FROM pages INNER JOIN body ON pages.page_id = body.page_id WHERE pages.page_id = ?",pageId) //
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

		var page models.Page

		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				panic(err.Error())
			}

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
					page.PageId,err = strconv.Atoi(value)
				case "title":
					page.PageTitle = value
				case "error_code":
					page.ErrorCode = value
				case "body":
					page.Body = value
				case "summary_id":
					page.SummaryId,err = strconv.Atoi(value)
				case "summary_page":
					page.SummaryPage,err = strconv.Atoi(value)
				case "importance":
					page.Importance,err = strconv.Atoi(value)
				case "complete":
					page.Complete,err = strconv.ParseBool(value)
				case "error_abst":
					page.ErrorAbst = value
				case "error_detail":
					page.ErrorDetail = value
				case "solutions":
					page.Solutions = value

				}

				//fmt.Println(columns[i], ": ", value)
			}



			//fmt.Println("-----------------------------------")
		}

		return page;

	}

func DBTags(pageId int) []string {
	db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM tags WHERE page_id = ?",pageId) //
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

	var tagList []string

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			if(columns[i] == "tag_name"){
				tagList = append(tagList,value)
			}

			//fmt.Println(columns[i], ": ", value)
		}



		//fmt.Println("-----------------------------------")
	}

	return tagList

}

func DBRelations(pageId int) []int {
	db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM relation WHERE page_id1 = ? OR page_id2 = ?",pageId,pageId) //
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

	var pageList []int

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			var valueId int
			valueId,err = strconv.Atoi(value)
			if(valueId != pageId){
				pageList = append(pageList,valueId)
				//fmt.Println(valueId,"addList")
			}

			fmt.Println(columns[i], ": ", value)
		}



		fmt.Println("-----------------------------------")
	}

	return pageList

}

func DBReferences(pageId int) []models.Reference {
	db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT pages.page_id,reference.page_id,reference.link,reference.link_title FROM reference INNER JOIN pages ON reference.page_id = pages.page_id WHERE reference.page_id = ? ",pageId) //
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

	var referenceList []models.Reference

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var reference models.Reference

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)


			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			if(columns[i] == "link"){
				reference.Link = value
			} else if(columns[i] == "link_title"){
				reference.LinkTitle = value
			}


			//fmt.Println(columns[i], ": ", value)
		}

		referenceList = append(referenceList,reference)

		//fmt.Println("-----------------------------------")
	}

	return referenceList

}

func DBTitlelist(list []int) []models.Title {
	db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	sqlsentence := "SELECT * FROM pages WHERE 0 "

	for i:=0; i < len(list); i++{
		str :=strconv.Itoa(list[i])
		sqlsentence += "OR page_id = " + str
	}

	fmt.Println(sqlsentence)
	rows, err := db.Query(sqlsentence) //
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

	var pageList []models.Title

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var page models.Title

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)


			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			switch columns[i]{
			case "page_id":
				page.PageId,err = strconv.Atoi(value)
			case "title":
				page.PageTitle = value
			}


			//fmt.Println(columns[i], ": ", value)
		}

		pageList = append(pageList,page)

		//fmt.Println("-----------------------------------")
	}

	return pageList

}
