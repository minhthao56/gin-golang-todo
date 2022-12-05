package controllers

import (
	"gin-golang/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type DataBaseModelStructMock struct {
	mock.Mock
}

func (db *DataBaseModelStructMock) ReadJSON() ([]database.Data, error) {
	var data []database.Data
	args := db.Called()
	return data, args.Error(1)

}

func (db *DataBaseModelStructMock) WriterTodo(data database.Data) (database.Data, error) {
	args := db.Called()
	return data, args.Error(1)

}

func (db *DataBaseModelStructMock) FindToDoById(id int64) (database.Data, error) {
	args := db.Called()
	var data database.Data
	return data, args.Error(1)

}

func (db *DataBaseModelStructMock) UpdateTodo(data database.Data) (database.Data, error) {
	args := db.Called()
	return data, args.Error(1)

}

func (db *DataBaseModelStructMock) DeleteTodoById(id int64) (int64, error) {
	args := db.Called()
	return int64(args.Int(0)), args.Error(1)

}

func (db *DataBaseModelStructMock) FindIndexTodoById(todos []database.Data, id int64) int {
	args := db.Called()
	return args.Int(0)
}

func TestGetManyController(t *testing.T) {
	var data []database.Data
	mockBD := new(DataBaseModelStructMock)

	mockBD.On("ReadJSON").Return(data, nil)

	controller := ControllersStruct{
		DataBaseModelInterface: mockBD,
	}

	r := SetUpRouter()
	r.GET("/", controller.GetManyTodo)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	mockBD.AssertNumberOfCalls(t, "ReadJSON", 1)

}
