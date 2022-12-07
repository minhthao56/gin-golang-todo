package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"time"
)

type Data struct {
	Id          int64
	Title       string
	Description string
	IsCompleted bool
}

type DataBaseModelStruct struct {
	IoReader func(filename string) ([]byte, error)
	IoWriter func(filename string, data []byte, perm fs.FileMode) error
}

func (db *DataBaseModelStruct) Read() ([]Data, error) {

	byteValue, err := db.IoReader("database/db.json")
	var dataTodo []Data
	err = json.Unmarshal(byteValue, &dataTodo)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return dataTodo, err
}

func (db *DataBaseModelStruct) Write(data Data) (Data, error) {

	todo := Data{}
	todo.Id = time.Now().Unix()
	todo.Description = data.Description
	todo.Title = data.Title
	todo.IsCompleted = data.IsCompleted || false

	listTodo, err := db.Read()

	listTodo = append(listTodo, todo)

	content, err := json.Marshal(listTodo)

	if err != nil {
		fmt.Println(err)
	}
	err = db.IoWriter("database/db.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return todo, err
}

func (db *DataBaseModelStruct) Update(data Data) (Data, error) {

	listTodo, err := db.Read()

	for i, v := range listTodo {

		if v.Id == data.Id {
			listTodo[i] = data
		}
	}
	content, err := json.Marshal(listTodo)

	if err != nil {
		fmt.Println(err)
	}
	err = db.IoWriter("database/db.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return data, err

}

func (db *DataBaseModelStruct) Delete(id int64) (int64, error) {

	listTodo, err := db.Read()

	index := db.ReadOneIndex(listTodo, id)
	if index != -1 {
		listTodo = append(listTodo[:index], listTodo[index+1:]...)

		content, err := json.Marshal(listTodo)

		if err != nil {
			fmt.Println(err)
		}
		err = db.IoWriter("database/db.json", content, 0644)
		if err != nil {
			log.Fatal(err)
		}
		return int64(index), err

	} else {
		err = errors.New("Not found todo")
		return int64(index), err
	}

}

func (db *DataBaseModelStruct) ReadOneIndex(todos []Data, id int64) int {

	for i, v := range todos {
		if id == v.Id {
			return i
		}
	}
	return -1
}

func (db *DataBaseModelStruct) ReadOne(id int64) (Data, error) {

	listTodo, err := db.Read()
	var todo Data

	for _, v := range listTodo {
		if v.Id == id {
			todo = v
		}
	}
	return todo, err

}
