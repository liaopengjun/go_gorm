package main
/**
CREATE TABLE `infos` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `money` bigint(20) DEFAULT NULL,
  `dog_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_dogs_info` (`dog_id`),
  CONSTRAINT `fk_dogs_info` FOREIGN KEY (`dog_id`) REFERENCES `dogs` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `dogs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `girl_god_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `girl_gods` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `dog_girl_god` (
  `girl_god_id` bigint(20) unsigned NOT NULL,
  `dog_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`girl_god_id`,`dog_id`),
  KEY `fk_dog_girl_god_dog` (`dog_id`),
  CONSTRAINT `fk_dog_girl_god_dog` FOREIGN KEY (`dog_id`) REFERENCES `dogs` (`id`),
  CONSTRAINT `fk_dog_girl_god_girl_god` FOREIGN KEY (`girl_god_id`) REFERENCES `girl_gods` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

 */
import (
	db2 "go_gorm/db"
	"gorm.io/gorm"
)

type Info struct {
	ID uint `gorm:"primarykey"`
	Money int
	DogID uint
}

type Dog struct {
	ID  uint `gorm:"primarykey"`
	Name string
	Info Info
	GirlGod []GirlGod `gorm:"many2many:dog_girl_god"`
}

type GirlGod struct {
	ID  uint `gorm:"primarykey"`
	Name string
	Dogs []Dog `gorm:"many2many:dog_girl_god"`
}
//创建多对多关系
func many2many(db *gorm.DB)  {
	i := Info{
		Money: 12000,
	}
	g := GirlGod{
		ID: 1,
		Name: "qq1",
	}
	g2 := GirlGod{
		ID: 2,
		Name: "qq2",
	}
	d := Dog{
		ID: 3,
		Name: "tg3",
		Info: i,
		GirlGod: []GirlGod{g,g2},
	}
	db.Create(&d)
}

//关联关系
func many2many2(db *gorm.DB) {
	g := GirlGod{
		ID: 1,
		Name: "qq1",
	}
	g2 := GirlGod{
		ID: 2,
		Name: "qq2",
	}
	d := Dog{
		ID: 3,
	}
	db.Model(&d).Association("GirlGod").Append(&g,&g2)
	//db.Model(&d).Association("GirlGod").Delete(&g)
	//db.Model(&d).Association("GirlGod").Clear()
	//db.Model(&d).Association("GirlGod").Replace(&g2)
}

func main()  {
	db,err:= db2.Initdb()
	if err != nil {
		panic(err)
	}
	//db.AutoMigrate(&Dog{},&GirlGod{},&Info{})
	//many2many(db)
	many2many2(db)

	//dog := Dog{
	//	ID: 1,
	//}
	//查找舔狗钱包和舔的所有女神
	//db.Find(&dog)
	//db.Preload("GirlGod").Preload("Info").Find(&dog)

	//查找该舔狗女神和舔狗钱包
	//var grils []GirlGod
	//db.Model(&dog).Preload("Dogs.Info").Association("GirlGod").Find(&grils)
	//
	//fmt.Println(grils)
}
