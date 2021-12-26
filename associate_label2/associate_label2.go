package main

import db2 "go_gorm/db"
/**
	foreignKey  指定当前模型的列作为连接表的外键
	references  指定引用表的列名，其将被映射为连接表外键
	joinForeignKey 指定连接表的外键列名，其将被映射到当前表
	joinReferences 指定连接表的外键列名，其将被映射到引用表

*/
type Jiazi struct {
	ID uint `gorm:"primarykey"`
	Name string  `gorm:"index"`
	Xiaofengche []Xiaofengche `gorm:"many2many:jiazi_fengche;foreignKey:Name;references:FcName"`
}

type Xiaofengche struct {
	ID uint `gorm:"primarykey"`
	FcName string   `gorm:"index"`
}

func main()  {
	db,err:= db2.Initdb()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Jiazi{},&Xiaofengche{})

	db.Create(&Jiazi{Name: "夹子",Xiaofengche: []Xiaofengche{
		{FcName: "小风车"},
		{FcName: "大风车"},
	}})

}
