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

func GetManyTodo(c *gin.Context) {
	data, err := database.ReadJSON()

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, data)

}

func AddTodo(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err", err)
	}

	var data database.Data
	err = json.Unmarshal(jsonData, &data)

	result, err := database.WriterTodo(data)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func EditTodo(c *gin.Context) {

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

	result, err := database.FindToDoById(int64(idInt))

	result.Title = data.Title
	result.Description = data.Description
	result.IsCompleted = data.IsCompleted

	if err != nil {
		c.Error(err)
		return
	}

	updatedData, err := database.UpdateTodo(result)

	c.JSON(http.StatusOK, updatedData)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Error(err)
		return
	}

	idReturn, err := database.DeleteTodoById(int64(idInt))

	if err != nil {

		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"id": idReturn})

}
