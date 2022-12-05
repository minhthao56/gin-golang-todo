package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gin-golang/database"

	"github.com/gin-gonic/gin"
)

type ControllersInterface interface {
	GetManyTodo(c *gin.Context)
	AddTodo(c *gin.Context)
	EditTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type ControllersStruct struct {
	DataBaseModelInterface database.DataBaseModelInterface
}

func (controller *ControllersStruct) GetManyTodo(c *gin.Context) {
	data, err := controller.DataBaseModelInterface.ReadJSON()

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

	result, err := controller.DataBaseModelInterface.WriterTodo(data)

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

	result, err := controller.DataBaseModelInterface.FindToDoById(int64(idInt))

	result.Title = data.Title
	result.Description = data.Description
	result.IsCompleted = data.IsCompleted

	if err != nil {
		c.Error(err)
		return
	}

	updatedData, err := controller.DataBaseModelInterface.UpdateTodo(result)

	c.JSON(http.StatusOK, updatedData)
}

func (controller *ControllersStruct) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Error(err)
		return
	}

	idReturn, err := controller.DataBaseModelInterface.DeleteTodoById(int64(idInt))

	if err != nil {

		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"id": idReturn})

}
