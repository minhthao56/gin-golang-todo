package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Data struct {
	Id          int64
	Title       string
	Description string
	IsCompleted bool
}

type DataBaseModelInterface interface {
	ReadJSON() ([]Data, error)
	WriterTodo(data Data) (Data, error)
	FindToDoById(id int64) (Data, error)
	UpdateTodo(data Data) (Data, error)
	DeleteTodoById(id int64) (int64, error)
	FindIndexTodoById(todos []Data, id int64) int
}

type DataBaseModelStruct struct{}

func (db *DataBaseModelStruct) ReadJSON() ([]Data, error) {

	byteValue, err := ioutil.ReadFile("database/db.json")
	var dataTodo []Data
	err = json.Unmarshal(byteValue, &dataTodo)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return dataTodo, err
}

func (db *DataBaseModelStruct) WriterTodo(data Data) (Data, error) {

	todo := Data{}
	todo.Id = time.Now().Unix()
	todo.Description = data.Description
	todo.Title = data.Title
	todo.IsCompleted = data.IsCompleted || false

	listTodo, err := db.ReadJSON()

	listTodo = append(listTodo, todo)

	content, err := json.Marshal(listTodo)

	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("database/db.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return todo, err
}

func (db *DataBaseModelStruct) FindToDoById(id int64) (Data, error) {

	listTodo, err := db.ReadJSON()

	var todo Data

	for _, v := range listTodo {
		if v.Id == id {
			todo = v
		}
	}

	return todo, err

}

func (db *DataBaseModelStruct) UpdateTodo(data Data) (Data, error) {

	listTodo, err := db.ReadJSON()

	for i, v := range listTodo {

		if v.Id == data.Id {
			listTodo[i] = data
		}
	}

	content, err := json.Marshal(listTodo)

	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("database/db.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return data, err

}

func (db *DataBaseModelStruct) DeleteTodoById(id int64) (int64, error) {

	listTodo, err := db.ReadJSON()

	index := db.FindIndexTodoById(listTodo, id)
	if index != -1 {
		listTodo = append(listTodo[:index], listTodo[index+1:]...)

		content, err := json.Marshal(listTodo)

		if err != nil {
			fmt.Println(err)
		}
		err = ioutil.WriteFile("database/db.json", content, 0644)
		if err != nil {
			log.Fatal(err)
		}
		return int64(index), err

	} else {
		err = errors.New("Not found todo")
		return int64(index), err
	}

}

func (db *DataBaseModelStruct) FindIndexTodoById(todos []Data, id int64) int {

	for i, v := range todos {
		if id == v.Id {
			return i
		}
	}
	return -1
}
