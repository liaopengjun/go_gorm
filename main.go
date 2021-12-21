package main

import (
	db2 "go_gorm/db"
	"gorm.io/gorm"
)

type Dog struct {
	Name string
	gorm.Model
	GrilGodID uint
	GrilGod GrilGod
}

type GrilGod struct {
	Name string
	gorm.Model
}

type DogTwo struct {
	Name string
	GrilGodTowID uint
	gorm.Model
}

type GrilGodTow struct {
	Name string
	DogTwo DogTwo
	gorm.Model
}

//属于关系
func belongsTo(db *gorm.DB)  {
	g := GrilGod {
		Name:"琪琪",
		Model:gorm.Model{
			ID:        1,
		},
	}
	d := Dog {
		Name:      "dog",
		Model:     gorm.Model{},
		GrilGod:   g,
	}
	//创建一对一
	db.Create(&d)
}

//包含关系
func hasOne(db *gorm.DB)  {
	d := DogTwo{
		Name:         "dog",

	}
	g := GrilGodTow{
		Name:   "qqiq",
		DogTwo: d,
	}
	//创建一对一
	db.Create(&g)
}

//一对多关系
func hanMany() {

}

func main()  {
	db,err:= db2.Initdb()
	if err != nil {
		panic(err)
	}
	m := db.Migrator()
	//判断是否存在表
	//if table := m.HasTable(&Dog{});table == false{
	//	db.AutoMigrate(&Dog{})
	//}
	//belongsTo(db)

	if table := m.HasTable(&GrilGodTow{});table == false{
		db.AutoMigrate(&GrilGodTow{},&DogTwo{})
	}
	hasOne(db)
}

