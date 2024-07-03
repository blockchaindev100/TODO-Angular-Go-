package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Roles struct {
	Id        int
	Role_type string
}

type Todo struct {
	Id        int
	Task_name string
	Status    string
	Active    bool
	User_id   int
}

type Users struct {
	Id                int
	Name              string
	Email             string
	Roleid            int
	Verification_code string
	Is_verified       bool
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=TODO port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error\n", err)
	}
	// var users Users
	// var roles Roles
	// var todo Todo

	err1 := db.Raw("select * from  users")
	// err2 := db.Raw("select * from roles").Scan(&roles)
	// err3 := db.Raw("Select * from todo").Scan(&todo)

	if err1 != nil {
		// fmt.Println("Error", err1, err2, err3)
		log.Fatalf("%#v", err1)
		// fmt.Println(string(err1))
	}

	// fmt.Printf("%#v\n", users)
	// fmt.Printf("%s\n", roles)
	// fmt.Printf("%s\n", todo)
}
