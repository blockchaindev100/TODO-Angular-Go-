package repository

import "github.com/blockchaindev100/todo/models"

func GetTodo(userid string) ([]models.Todo, error) {
	var data []models.Todo
	result := DB.Where("user_id=? AND active=?", userid, true).Find(&data)
	return data, result.Error
}

func GetTodoByID(userid string, id string) (models.Todo, error) {
	var data models.Todo
	result := DB.Where("user_id=? AND active=?", userid, true).First(&data, id)
	return data, result.Error
}

func CreateTodo(data *models.Todo) error {
	data.Active = true
	result := DB.Create(data)
	return result.Error
}

func DeleteTodo(id string) error {
	var data models.Todo
	if err := DB.First(&data, id).Error; err != nil {
		return err
	}
	data.Active = false
	result := DB.Save(&data)
	return result.Error
}

func UpdateTodo(data *models.Todo) error {
	result := DB.Save(data)
	return result.Error
}
