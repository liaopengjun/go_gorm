package main

import (
	"fmt"
	db2 "go_gorm/db"
	"gorm.io/gorm"
)

/**
	has many
**/
type Info struct {
	ID uint `gorm:"primarykey"`
	Money int
	DogID uint
}

type Dog struct {
	ID  uint `gorm:"primarykey"`
	Name string
	GirlGodID uint
	Info Info
}

type GirlGod struct {
	ID  uint `gorm:"primarykey"`
	Name string
	Dogs []Dog
}

func onetomany(db *gorm.DB){
	d := Dog{
		ID:        1,
		Name:      "dog1",
	}
	d2 := Dog{
		ID:        2,
		Name:      "dog2",
	}
	g := GirlGod{
		ID:   1,
		Name: "hehe",
		Dogs: []Dog{d,d2},
	}
	db.Create(&g)
}

func main()  {
	db,err:= db2.Initdb()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&GirlGod{},&Dog{})
	db.AutoMigrate(&Info{})

	//创建
	//onetomany(db)
	var girl GirlGod

	//查询所有女神
	//db.First(&girl)

	//查询女神的指定舔狗
	//db.Preload("Dogs.Info").Preload("Dogs","name = ?","dog2").First(&girl)

	//自定义sql where
	//db.Preload("Dogs", func(db *gorm.DB) *gorm.DB{
	//	return db.Where("name = ?","dog2")
	//}).First(&girl)
	//

	//该女神舔狗的钱包
	//db.Preload("Dogs.Info","money > ?","1000").Preload("Dogs").First(&girl)

	// Join 加载
	db.Preload("Dogs", func(db *gorm.DB) *gorm.DB {
		return db.Joins("Info").Where("money > ?","1000")
	}).First(&girl)

	fmt.Println(girl)
}