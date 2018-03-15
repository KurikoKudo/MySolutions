package controllers

import (
	"MySolutions/app/daos"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	text := "It Works!"
	daos.Migration()
	//fmt.Println("Index")
	return c.Render(text)
}

func (c App) Home() revel.Result {
	text := "- welcome to YourSolutions -"
	return c.Render(text)

}

func (c App) RegistForm() revel.Result {
	/*将来的にユーザが複数いる場合ここでユーザ情報の保持とかする?*/
	return c.Render()

}

func (c App) Search() revel.Result {

	if c.Params.Form.Get("ptitle") != "" {
		ptitle = strings.Split((c.Params.Form.Get("ptitle")), " ")
	} else {

	}

	if c.Params.Form.Get("tags") != "" {
		tags = strings.Split((c.Params.Form.Get("tags")), " ")
	}

	if c.Params.Form.Get("evaluation") != "" {
		evaluationInput, err := strconv.ParseUint(c.Params.Form.Get("evaluation"), 10, 0)
		if err != nil {
			return err
		}
		evaluation = uint(evaluationInput)
	}

	if c.Params.Form.Get("ptitle") != "" {
		condition, err := strconv.ParseBool(c.Params.Form.Get("condition"))
		if err != nil {
			return err
		}
	}

	list := SearchController(ptitle, tags, evaluation, condition)

	listlen := len(list)
	//fmt.Println("長さ：",listlen)

	return c.Render(list, listlen)
}

/*
func (c App) PageDisplay(pageId int) revel.Result {

	start := time.Now();
	var page models.Page
	display := make(chan models.Page)
	tags := make(chan []string)

	go func() {

		fmt.Println("display start")
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
		fmt.Println("display end")
		display <- page

	}()

	go func() {
		fmt.Println("tag start")
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

		fmt.Println("tag end")
		db.Close()
		tags <- tagList

	}()
	page = <- display
	close(display)
	end := time.Now();
	fmt.Printf("DBDisplay")
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	page.TagName = <- tags
	close(tags)
	end = time.Now();
	fmt.Printf("DBTags")
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	var relations []int
	relations = DBRelations(pageId)
	page.Relation = relations
	end = time.Now();
	fmt.Printf("DBRelations")
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	var relationPage []models.Title
	relationPage = DBTitlelist(relations)
	page.RelationPage = relationPage

	var references []models.Reference
	references = DBReferences(pageId)
	page.Reference = references
	end = time.Now();
	fmt.Printf("DBReferences")
	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	return c.Render(page)
}

func (c App) Regist() revel.Result {

	//fmt.Println("hoge")
	return c.Render()
}

func (c App) Insert() revel.Result {

	//var ptitle,ecode,eabst,edatail,solutions string

	var tags []string
	var relList []int
	var refList []models.Reference

		ptitle := c.Params.Form.Get("ptitle")
		ecode := c.Params.Form.Get("ecode")
		eabst := c.Params.Form.Get("eabst")
		edetail := c.Params.Form.Get("edetail")
		solutions := c.Params.Form.Get("solutions")
		tags = strings.Split((c.Params.Form.Get("tags")), " ")
		importance,err := strconv.Atoi(c.Params.Form.Get("importance"))
		if err != nil {
			panic(err.Error())
		}

		for i:=1; i <= 5; i++{
			var ref models.Reference
			getT := "ltitle" + strconv.Itoa(i)
			getL := "link" + strconv.Itoa(i)
			if(c.Params.Form.Get(getT) != ""){
				ref.LinkTitle = c.Params.Form.Get(getT)
				ref.Link = c.Params.Form.Get(getL)
				refList = append(refList,ref)
			}
		}
		condition := c.Params.Form.Get("condition")



	for j:=1; j<=5; j++{
		getR := "relation"+strconv.Itoa(j)
		if(c.Params.Form.Get(getR) != ""){
			str := strings.Split((c.Params.Form.Get(getR)), "=")
			fmt.Println(str[0])
			fmt.Println(str[1])
			num,err := strconv.Atoi(str[1])
			if err != nil {
				panic(err.Error())
			}
			relList = append(relList,num)
		}
	}

	page := make(chan int)
	go func () {

		db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
		if err != nil {
			panic(err.Error())
		}


		sqlsentence := "SELECT * FROM pages"

		//fmt.Println(sqlsentence)
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

		page <- pageLength+1

	}()

	pageId := <- page

*/

/*tag := make(chan string[])
ref := make(chan string[])
ref := make(chan string[])*/

/*
	go func(){
		db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()

			sqlsentence := "INSERT INTO pages (page_id,title,body,complete,importance"
			if(ecode != ""){
				sqlsentence += ",error_code)"
			} else {
				sqlsentence += ")"
			}
			sqlsentence += "VALUES("+strconv.Itoa(pageId)+",'"+ptitle+"','"+eabst+edetail+solutions+"',"+condition+","+strconv.Itoa(importance)
			if(ecode != ""){
				sqlsentence += ",'"+ecode+"')"
			} else {
				sqlsentence += ")"
			}
			fmt.Println("page",sqlsentence)
			_, err = db.Exec(sqlsentence)
			if err != nil {
    		panic(err.Error())
			}

	}()

	go func(){
		db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()

			sqlsentence := "INSERT INTO body (page_id,solutions"

			if(eabst != ""){
				sqlsentence += ",error_abst"
				if(edetail != ""){
					sqlsentence += ",error_detail)"
				} else {
					sqlsentence += ")"
				}
			} else {
				if(edetail != ""){
					sqlsentence += ",error_detail)"
				} else {
					sqlsentence += ")"
				}
			}

			sqlsentence += "VALUES("+strconv.Itoa(pageId)+",'"+solutions+"'"

			if(eabst != ""){
				sqlsentence += ",'"+eabst+"'"
				if(edetail != ""){
					sqlsentence += ",'"+edetail+"')"
				} else {
					sqlsentence += ")"
				}
			} else {
				if(edetail != ""){
					sqlsentence += ",'"+edetail+"')"
				} else {
					sqlsentence += ")"
				}
			}
			fmt.Println("body",sqlsentence)
			_, err = db.Exec(sqlsentence)
			if err != nil {
    		panic(err.Error())
			}

	}()

	go func(){
		db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()

		if(tags[0] != ""){
			sqlsentence := "INSERT INTO tags (page_id,tag_name) VALUES"

			for i:=0; i < len(tags); i++ {
				if(i==0){
					sqlsentence += " (" + strconv.Itoa(pageId) + ",'" + tags[i] + "')"
				}else{
					sqlsentence += " ,(" + strconv.Itoa(pageId) + ",'" + tags[i] + "')"
				}
			}
			fmt.Println("tag",sqlsentence)
			_, err = db.Exec(sqlsentence)
			if err != nil {
    		panic(err.Error())
			}
		}
	}()

	func(){
		db, err := sql.Open("mysql", "mysolutions:MySystem2017!@tcp(localhost:3306)/mysolutions")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()

		if(refList[0].LinkTitle != ""){

			sqlsentence := "INSERT INTO reference (page_id,link_title,link) VALUES"

			for i:=0; i < len(refList); i++ {
				if(i==0){
					sqlsentence += " (" + strconv.Itoa(pageId) + ",'" + refList[i].LinkTitle + "','" + refList[i].Link + "')"
				}else{
					sqlsentence += " ,(" + strconv.Itoa(pageId) + ",'" + refList[i].LinkTitle + "','" + refList[i].Link + "')"
				}
			}
			fmt.Println("ref",sqlsentence)
			_, err = db.Exec(sqlsentence)
			if err != nil {
    		panic(err.Error())
			}
		}
	}()

	return c.Render(ptitle,pageId)

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
*/
