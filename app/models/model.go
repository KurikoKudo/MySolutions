package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Page_Display struct { //not DB
	Page_Body     Page_Body
	Tag           []Tag
	Page_Relation []Page_Relation
}

type Page_Body struct {
	Page_Id    uint   `json:page_id gorm:"primary_key"`
	Page_Title string `json:page_title gorm:"not null"`
	Page_Body  string `json:page_body gorm:"not null"`
	Evaluation uint   `json:evalution gorm:"not null"`
	Condition  bool   `json:condition gorm:"not null"`
}

type Tag struct {
	Page_Id  uint   `json:page_id gorm:"not null"`
	Tag_Name string `json:tag_name gorm:"not null"`
}

type Page_Relation struct {
	Page_id0 uint `json:page_id gorm:"not null"`
	Page_id1 uint `json:page_id gorm:"not null"`
}

type Summary_Body struct {
	Summary_Id    uint `json:summary_id gorm:"primary_key"`
	Summary_Title uint `json:summary_title gorm:"not null"`
	Page_Total    uint `json:page_total gorm:"not null"`
}

type Summary_Page struct {
	Summary_Id  uint `json:summary_id gorm:"not null"`
	Page_Id     uint `json:page_id gorm:"not null"`
	Page_Number uint `json:page_number gorm:"not null"`
}

type Summary_Display struct {
	Summary_Body Summary_Body
	Summary_Page []Summary_Page
}
