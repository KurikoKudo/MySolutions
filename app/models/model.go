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
	Page_Id    uint   `gorm:"primary_key"`
	Page_Title string `gorm:"not null;unique"`
	Page_Body  string `gorm:"not null"`
	Evaluation uint   `gorm:"not null"`
	Condition  bool   `gorm:"not null"`
}

type Tag struct {
	Page_Id  uint   `gorm:"not null"`
	Tag_Name string `gorm:"not null"`
}

type Page_Relation struct {
	Page_id0 uint `gorm:"not null"`
	Page_id1 uint `gorm:"not null"`
}

type Summary_Body struct {
	Summary_Id    uint `gorm:"primary_key"`
	Summary_Title uint `gorm:"not null"`
	Page_Total    uint `gorm:"not null"`
}

type Summary_Page struct {
	Summary_Id  uint `gorm:"not null"`
	Page_Id     uint `gorm:"not null"`
	Page_Number uint `gorm:"not null"`
}

type Summary_Display struct {
	Summary_Body Summary_Body
	Summary_Page []Summary_Page
}
