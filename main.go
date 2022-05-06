package main

import (
	"fmt"
	"gorm.io/gorm"
	"sync"
)

// 学生
//type Student struct {
//	gorm.Model
//	StudentName string
//	ClassID uint
//	IdCard IdCard
//	Teachers []Teacher `gorm:"many2many:student_teacher;"`
//}
//
//// 身份证
//type IdCard struct {
//	gorm.Model
//	Code int64
//	StudentID uint
//}
//
//// 班级
//type Class struct {
//	gorm.Model
//	ClassName string
//	Students []Student
//}
//
//// 老师
//type Teacher struct {
//	gorm.Model
//	TeacherName string
//	Students []Student `gorm:"many2many:student_teacher;"`
//}
type CreditCard struct {
	gorm.Model
	Number   string
	UserID   uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}



func main()  {
	var wg = sync.WaitGroup{}

	var lock sync.Mutex
	wg.Add(10000)
	for i:=0;i<10000;i++{
		lock.Lock() //加锁
		go func(i int) {
			fmt.Println("hello groutine",i)
			lock.Unlock() //加锁
			wg.Done()
		}(i)

	}
	wg.Wait()

	//dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//
	//db.AutoMigrate(&User{},&CreditCard{})
	//user := &User{
	//	Name: "jinzhu",
	//	CreditCard: CreditCard{Number: "411111111111"},
	//}
	//db.Omit("CreditCard").Create(&user)


	//i := IdCard{
	//	Code: 3624291000,
	//}
	//
	//s := Student{
	//	StudentName: "jun",
	//	IdCard: i,
	//}
	//
	//t := Teacher{
	//	TeacherName:"qm",
	//	Students: []Student{s},
	//}
	//
	//c := Class{
	//	ClassName: "五年级二班",
	//	Students: []Student{s},
	//}
	//
	//_ = db.Create(&c).Error
	//_ = db.Create(&t).Error

	//fmt.Println(err)

}

