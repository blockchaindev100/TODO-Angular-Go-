package service

import (
	"errors"

	"github.com/blockchaindev100/todo/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func GetUsers() ([]models.Users, error) {
	var users []models.Users

	err := DB.Raw("select * from users").Scan(&users).Error

	return users, err
}

func GetUser(id string) (models.Users, error) {
	var user models.Users

	err := DB.Raw("select * from users where id=?", id).Scan(&user).Error

	return user, err
}

func CreateUsers(data *models.Users) error {
	uuidVal := GenerateUUID()
	var createdUser models.Users
	err := DB.Raw("insert into users (id,name,email,roleid,verification_code,is_verified) values (?,?,?,?,?,?)", uuidVal, data.Name, data.Email, data.Roleid, data.Verification_code, data.Is_verified).Scan(&createdUser).Error

	return err
}

func DeleteUser(id string) error {
	err := DB.Exec("delete from users where id=?", id)
	if err.RowsAffected == 0 {
		return errors.New("no_rows_affected")
	}

	return err.Error
}
