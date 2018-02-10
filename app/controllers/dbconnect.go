package controllers

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func gormConnect() *gorm.DB {
  DBMS     = "mysql"
  USER     = "mysolutions"
  PASS     = "MySystem2017!"
  PROTOCOL = "tcp(localhost:3306)"
  DBNAME   = "mysolutions"

  CONNECT = USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
  db,err := gorm.Open(DBMS, CONNECT)

  if err != nil {
    panic(err.Error())
  }
  return db
}


func DBSearch(body []string,ptitle []string,tags []string, ecode string) ([]models.Page) {

  // DB接続
  db := gormConnect()
  defer db.Close()

  // sql文先頭
  sqlsentence := "SELECT pages.page_id,pages.title,pages.error_code,pages.body,pages.importance,tags.tag_name,tags.page_id FROM pages INNER JOIN tags ON pages.page_id = tags.page_id WHERE 1"

  //　body　空白文字でない時のみsql追加
  if(0 < len(body)){
    bodysentence := ""
    for i:=0; i < len(body); i++{
      if(body[i] != ""){
        bodysentence += " AND pages.body LIKE '%" + body[i] + "%'"
      }
    }
    sqlsentence += bodysentence
  }

  //　title　空白文字でない時のみsql追加
  if(0 < len(ptitle)){
    ptitlesentence := ""
    for i:=0; i < len(ptitle); i++{
      if(ptitle[i] != ""){
        ptitlesentence += " AND pages.title LIKE '%" + ptitle[i] + "%'"
      }
    }
    sqlsentence += ptitlesentence
  }

  // tags 空白文字でない時のみsql追加
  if(0 < len(tags)){
    tagssentence := ""
    for i:=0; i<len(tags); i++{
      if(tags[i] != ""){
        tagssentence += " AND tags.tag_name LIKE '%" + tags[i] + "%'"
      }
    }
    sqlsentence += tagssentence
  }

  // error_code 空白文字でない時のみsql追加
  if(ecode != ""){
    ecodesentence := " AND pages.error_code LIKE '%" + ecode + "%'"
    sqlsentence += ecodesentence
  }

  // sql文実行
  rows, err := db.Query(sqlsentence)
  if err != nil {
    panic(err.Error())
  }

  //検索結果　sliceの初期化
  var list []models.Page

  for rows.Next() {
    var pages models.Page
    var tmpString sql.NullString
    var tmpInt sql.NullInt64
    //　必要のない結果はpagesに代入せず、とりあえず退避（ここstructの見直しで綺麗になるはず）
    err = rows.Scan(&pages.PageId,&pages.PageTitle,&tmpString,&tmpString,&tmpInt,&tmpString,&tmpInt)
    if err != nil {
      panic(err.Error())
    }

    //タグが２つ以上マッチするとその分page_idも追加されるので、同idがきたらリストに追加しない
    if (contains(list,pages.PageId) >= 0){
      // list[contains(list,pages.PageId)].TagName = append(list[contains(list,pages.PageId)].TagName,pages.TagName[0])
    } else {
      list = append(list, pages)
    }

  }

  return list;
}
