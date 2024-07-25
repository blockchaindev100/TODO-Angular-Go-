package repository

import (
	"errors"

	"github.com/blockchaindev100/todo/models"
	"github.com/blockchaindev100/todo/service"
)

func GetUsers() ([]models.Users, error) {
	var users []models.Users

	result := DB.Where("active=?", "true").Find(&users)

	return users, result.Error
}

func GetUser(id string) (models.Users, error) {
	var user models.Users

	result := DB.Where("active=? AND id=?", "true", id).First(&user)

	return user, result.Error
}

func CreateUsers(data *models.Users) error {
	uuidVal := service.GenerateUUID()
	data.Id = uuidVal
	data.Active = true
	result := DB.Create(data)

	return result.Error
}

func DeleteUser(id string) error {
	var user models.Users
	DB.Where("active=? AND id=?", true, id).First(&user)
	if !user.Active {
		return errors.New("user is doesn't exist")
	}
	user.Active = false
	result := DB.Save(&user)
	return result.Error
}
