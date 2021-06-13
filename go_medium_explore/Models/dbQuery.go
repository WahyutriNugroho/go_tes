//Models/User.go
package Models

import (
	"fmt"
	"go_medium/Config"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllData(predict *[]Predict) (err error) {
	if err = Config.DB.Find(predict).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateData(predict *Predict) (err error) {
	if err = Config.DB.Create(predict).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetDataByID(user *Predict, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateData(predict *Predict, id string) (err error) {
	fmt.Println(predict)
	Config.DB.Save(predict)
	return nil
}

//DeleteUser ... Delete user
func DeleteData(predict *Predict, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(predict)
	return nil
}
