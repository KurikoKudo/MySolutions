package controllers

import (
	"github.com/revel/revel"
	"strings"
	"strconv"
	"fmt"
	"database/sql"
	 "time"
	  _ "github.com/go-sql-driver/mysql"
	 "MySolutions/app/models"
 )

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	text := "It Works!"
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
	//fmt.Println("長さ：",listlen)

	return c.Render(list,listlen)
}

func (c App) PageDisplay(pageId int) revel.Result {

	start := time.Now();
	var page models.Page
	display := make(chan models.Page)

	go func() {
		db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close() // 関数がリターンする直前に呼び出される

		rows, err := db.Query("SELECT pages.page_id,pages.title,pages.error_code,pages.importance,pages.summary_id,pages.summary_page,pages.complete, body.error_abst,body.error_detail,body.solutions FROM pages INNER JOIN body ON pages.page_id = body.page_id WHERE pages.page_id = ?",pageId) //
		if err != nil {
			panic(err.Error())
		}

		var page models.Page

		for rows.Next() {

			var ecode sql.NullString
			var summaryId sql.NullInt64
			var summaryPage sql.NullInt64
			var eabst sql.NullString
			var edetail sql.NullString

			err = rows.Scan(&page.PageId,&page.PageTitle,&ecode,&page.Importance,&summaryId,&summaryPage,&page.Complete,&eabst,&edetail,&page.Solutions)
			if err != nil {
				panic(err.Error())
			}

			if(ecode.Valid){
				page.ErrorCode = ecode.String
			}
			if(summaryId.Valid){
				page.SummaryId = summaryId.Int64
			}
			if(summaryPage.Valid){
				page.SummaryPage = summaryPage.Int64
			}
			if(eabst.Valid){
				page.ErrorAbst = eabst.String
			}
			if(edetail.Valid){
				page.ErrorDetail = edetail.String
			}

		}

		display <- page

	}()
	page = <- display
	end := time.Now();
	fmt.Printf("DBDisplay")
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	start = time.Now();
	tags := make(chan []string)

	go func() {
		db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
		if err != nil {
			panic(err.Error())
		}

		rows, err := db.Query("SELECT * FROM tags WHERE page_id = ?",pageId) //
		if err != nil {
			panic(err.Error())
		}

		var tagList []string

		for rows.Next() {
			var tagName string
			var pageId int
			err = rows.Scan(&pageId,&tagName)
			if err != nil {
				panic(err.Error())
			}
			tagList = append(tagList,tagName)
		}

		db.Close()
		tags <- tagList

	}()


	page.TagName = <- tags
	end = time.Now();
	fmt.Printf("DBTags")
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())


	start = time.Now();
	var relations []int
	relations = DBRelations(pageId)
	page.Relation = relations
	end = time.Now();
	fmt.Printf("DBRelations")
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	var relationPage []models.Title
	relationPage = DBTitlelist(relations)
	page.RelationPage = relationPage


	start = time.Now();
	var references []models.Reference
	references = DBReferences(pageId)
	page.Reference = references
	end = time.Now();
	fmt.Printf("DBReferences")
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	return c.Render(page)
}

func (c App) Regist() revel.Result {

	pageId := DBPage()

	return c.Render(pageId)
}

  func DBSearch(body []string,ptitle []string,tags []string, ecode string) ([]models.Page) {
    db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close() // 関数がリターンする直前に呼び出される


		sqlsentence := "SELECT pages.page_id,pages.title,pages.error_code,pages.body,pages.importance,tags.tag_name,tags.page_id FROM pages INNER JOIN tags ON pages.page_id = tags.page_id WHERE 1"
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

    rows, err := db.Query(sqlsentence) //
    if err != nil {
      panic(err.Error())
    }

		//sliceの初期化
		var list []models.Page

    for rows.Next() {
			var pages models.Page
			var tmpString sql.NullString
			var tmpInt sql.NullInt64
      err = rows.Scan(&pages.PageId,&pages.PageTitle,&tmpString,&tmpString,&tmpInt,&tmpString,&tmpInt)
      if err != nil {
        panic(err.Error())
      }

			if (contains(list,pages.PageId) >= 0){ //タグが２つ以上マッチ
				// list[contains(list,pages.PageId)].TagName = append(list[contains(list,pages.PageId)].TagName,pages.TagName[0])
			} else {
				list = append(list, pages)
			}

    }

		return list;
  }

	func contains(s []models.Page, e int) int {
		i := 0
		for _, v := range s {
			if e == v.PageId {
				return i
			}
			i += 1
		}
		return -1
}

func DBRelations(pageId int) []int {
	db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
	if err != nil {
		panic(err.Error())
	}

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
		for i,col := range values {
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
			}
			fmt.Println(columns[i], ": ", value)
		}



		//fmt.Println("-----------------------------------")
	}
	db.Close()
	return pageList

}

func DBReferences(pageId int) []models.Reference {
	db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
	if err != nil {
		panic(err.Error())
	}


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
	db.Close()
	return referenceList

}

func DBTitlelist(list []int) []models.Title {
	db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
	if err != nil {
		panic(err.Error())
	}


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
	db.Close()
	return pageList

}

func DBPage() int {

	db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
	if err != nil {
		panic(err.Error())
	}


	sqlsentence := "SELECT * FROM pages"

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

	pageLength := 0

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		pageLength += 1

	}
	db.Close()
	return pageLength+1

}
