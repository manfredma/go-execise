package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(127.0.0.1:3306)/coupon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("----", err.Error())
	}

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := db.Create(&user) // pass pointer of data to Create
	fmt.Println("-------", user.ID, result)

	email := "xxx@yyy.zzz"
	fmt.Println("-------", reflect.TypeOf(email))
	user2 := User{Name: "Jinzhu2", Age: 18, Email: &email, Birthday: time.Now()}
	result2 := db.Create(&user2) // pass pointer of data to Create
	fmt.Println("-------", user2.ID, result2)
}
