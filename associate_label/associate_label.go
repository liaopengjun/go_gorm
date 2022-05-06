package main

import db2 "go_gorm/db"

type Jiazi struct {
	ID uint
	Name string
	Xiaofengche Xiaofengche `gorm:"polymorphic:Owner;polymorphicValue:huhu"`
}
type Yujie struct {
	ID uint
	Name string
	Xiaofengche Xiaofengche `gorm:"polymorphic:Owner;polymorphicValue:ababa""`
}
type Xiaofengche struct {
	ID uint
	Name string
	OwnerType string
	OwnerID uint
}
func main()  {
	db,err:= db2.Initdb()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Jiazi{},&Yujie{},&Xiaofengche{})

	db.Create(&Jiazi{Name: "夹子",Xiaofengche:Xiaofengche{Name: "小风车"}})
	db.Create(&Yujie{Name: "御姐",Xiaofengche:Xiaofengche{Name: "大风车"}})
}
