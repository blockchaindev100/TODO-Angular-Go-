package repository

import "github.com/blockchaindev100/todo/models"

func GetTodo(userid string) ([]models.Todo, error) {
	var data []models.Todo
	err := DB.Exec("select * from todo where user_id=?", userid).Scan(&data).Error

	return data, err
}

func GetTodoByID(userid string, id string) (models.Todo, error) {
	var data models.Todo
	err := DB.Exec("select * from todo where user_id=? and id=?", userid, id).Scan(&data).Error

	return data, err
}

func CreateTodo(data *models.Todo) error {
	err := DB.Exec("insert into todo (id,task_name,status,active,user_id) values (?,?,?,?,?)", data.Id, data.Task_name, data.Status, true, data.User_id).Error

	return err
}

func DeleteTodo(id string) error {
	err := DB.Exec("update todo set active=? where id=?", false, id).Error

	return err
}

func UpdateTodo(data *models.Todo) error {
	err := DB.Exec("update todo set task_name=?,status=? where id=?", data.Task_name, data.Status, data.User_id).Error
	return err
}
