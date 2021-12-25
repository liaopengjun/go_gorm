package main
//hasone
import (
	db2 "go_gorm/db"
	"gorm.io/gorm"
)
/**
CREATE TABLE `dogs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `gril_god_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_gril_gods_dog` (`gril_god_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `gril_gods` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
 */
type Dog struct {
	ID   uint `gorm:"primarykey"`
	Name string
	GrilGodID uint
}

type GrilGod struct {
	ID   uint `gorm:"primarykey"`
	Name string
	Dog Dog
}

//包含关系
func hasOne(db *gorm.DB) {
	d := Dog{
		Name: "dog2",
	}
	g := GrilGod{
		Name:   "qq2",
		Dog: d,
	}
	//创建一对一
	db.Create(&g)
}

//关联关系
func one2one(db *gorm.DB){
	d := Dog{
		ID :1,
	}
	d2 := Dog{
		ID :2,
	}
	g := GrilGod{
		ID: 3,
	}

	//给女神推送舔狗:创建女神模型建立跟舔狗的关系追加指定舔狗关系
	db.Model(&g).Association("Dog").Append(&d)
	db.Model(&g).Association("Dog").Delete(&d)
	//替换关联关系
	db.Model(&g).Association("Dog").Replace(&d,&d2)
	//清除所有关联关系
	db.Model(&g).Association("Dog").Clear()

}

func main()  {
	db,err:= db2.Initdb()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&GrilGod{},&Dog{})
	hasOne(db)

	//1.查询单个女神数据
	//var girl GrilGod
	//db.First(&girl,1)

	//1.查询单个女神并且把舔狗查出来
	//db.Preload("Dog").First(&girl,1)
	//fmt.Println(girl)

	//one2one(db)
}