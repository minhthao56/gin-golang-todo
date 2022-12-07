package controllers

import (
	"encoding/json"
	"fmt"
	"gin-golang/database"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DataBaseInterface interface {
	Read() ([]database.Data, error)
	Write(data database.Data) (database.Data, error)
	Delete(id int64) (int64, error)
	Update(data database.Data) (database.Data, error)
	ReadOne(id int64) (database.Data, error)
}
type ControllersStruct struct {
	DataBaseInterface DataBaseInterface
}

func (controller *ControllersStruct) GetManyTodo(c *gin.Context) {
	fmt.Println("------GetManyTodo-----", c)

	data, err := controller.DataBaseInterface.Read()

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, data)

}

func (controller *ControllersStruct) AddTodo(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err", err)
	}

	var data database.Data

	err = json.Unmarshal(jsonData, &data)

	result, err := controller.DataBaseInterface.Write(data)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (controller *ControllersStruct) EditTodo(c *gin.Context) {

	id := c.Param("id")

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err", err)
	}

	var data database.Data
	err = json.Unmarshal(jsonData, &data)

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Error(err)
		return
	}

	result, err := controller.DataBaseInterface.ReadOne(int64(idInt))

	result.Title = data.Title
	result.Description = data.Description
	result.IsCompleted = data.IsCompleted

	if err != nil {
		c.Error(err)
		return
	}

	updatedData, err := controller.DataBaseInterface.Update(result)

	c.JSON(http.StatusOK, updatedData)
}

func (controller *ControllersStruct) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Error(err)
		return
	}

	idReturn, err := controller.DataBaseInterface.Delete(int64(idInt))

	if err != nil {

		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"id": idReturn})

}
