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

func UpdateTodo(data *models.Todo) {
	// err:=DB.Exec("update todo set task_name=?, ")
}
