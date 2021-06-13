//Controllers/User.go
package Controllers

import (
	"fmt"
	"go_medium/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get ... Get all from table predict
func GetData(c *gin.Context) {
	var predict []Models.Predict
	err := Models.GetAllData(&predict)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, predict)
	}
}

//Create ... Create data
func CreateData(c *gin.Context) {
	var predict Models.Predict
	c.BindJSON(&predict)
	err := Models.CreateData(&predict)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, predict)
	}
}

//GetUserByID ... Get the user by id
func GetDataByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var predict Models.Predict
	err := Models.GetDataByID(&predict, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, predict)
	}
}

//UpdateUser ... Update the user information
func UpdateData(c *gin.Context) {
	var predict Models.Predict
	id := c.Params.ByName("id")
	err := Models.GetDataByID(&predict, id)
	if err != nil {
		c.JSON(http.StatusNotFound, predict)
	}
	c.BindJSON(&predict)
	err = Models.UpdateData(&predict, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, predict)
	}
}

//DeleteUser ... Delete the user
func DeleteData(c *gin.Context) {
	var predict Models.Predict
	id := c.Params.ByName("id")
	err := Models.DeleteData(&predict, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
