package main

import (
	db2 "go_gorm/db"
	"gorm.io/gorm"
)

/**
CREATE TABLE `dogs` (
  `name` varchar(256) DEFAULT NULL,
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `gril_god_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dogs_gril_god` (`gril_god_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `gril_gods` (
  `name` varchar(256) DEFAULT NULL,
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

*/

type Dog struct {
	Name string
	ID   uint `gorm:"primarykey"`
	GrilGodID uint
	GrilGod GrilGod
}

type GrilGod struct {
	Name string
	ID   uint `gorm:"primarykey"`
}

//属于关系
func belongsTo(db *gorm.DB)  {
	g := GrilGod {
		Name:"琪琪",
	}
	d := Dog {
		Name:      "dog",
		GrilGod:   g,
	}
	//创建一对一
	db.Create(&d)
}

//建立模型关联关系
func one2noe(db *gorm.DB)  {
	d := Dog {
		ID: 1,
	}

	g := GrilGod{
		ID: 1,
	}
	//g2 := GrilGod{
	//	ID: 2,
	//}
	//创建一个舔狗模型去舔女神
	//db.Model(&d).Association("GrilGod").Append(&g)
	//清除某个女神关联关系
	db.Model(&d).Association("GrilGod").Delete(&g)
	//替换关联关系
	//db.Model(&d).Association("GrilGod").Replace(&g,&g2)
	//清除所有关联关系
	//db.Model(&d).Association("GrilGod").Clear()
}

func main()  {

	db,err:= db2.Initdb()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Dog{})

	//判断是否存在表
	//if table := m.HasTable(&Dog{});table == false {
	//	db.AutoMigrate(&Dog{})
	//}

	//belongsTo(db)
	//1.查询单个舔狗数据
	//var god Dog

	//db.First(&god,1)
	//1.查询单个舔狗并且把女神查出来
	//db.Preload("GrilGod").First(&god,1)
	//fmt.Println(god)
	one2noe(db)
}


