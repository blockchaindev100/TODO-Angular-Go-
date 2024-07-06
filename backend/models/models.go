package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Roles struct {
	Id        string
	Role_type string
}

type Todo struct {
	Id        string
	Task_name string
	Status    string
	Active    bool
	User_id   int
}

type Users struct {
	Id                string
	Name              string
	Email             string
	Roleid            int
	Verification_code string
	Is_verified       bool
}

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=TODO port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error\n", err)
	}

	return db, err
}
