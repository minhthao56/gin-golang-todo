package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-golang/database"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type DataBaseModelStructMock struct {
	mock.Mock
}

func creatMockDataRespJSON() ([]database.Data, database.Data) {
	var data []database.Data
	var mockData database.Data
	mockData.Id = 1
	mockData.Description = "Description"
	mockData.Title = "Title"
	mockData.IsCompleted = false
	data = append(data, mockData)
	return data, mockData
}

func (db *DataBaseModelStructMock) Read() ([]database.Data, error) {
	data, _ := creatMockDataRespJSON()

	args := db.Called()
	return data, args.Error(1)

}

func (db *DataBaseModelStructMock) Write(data database.Data) (database.Data, error) {
	args := db.Called(data)
	return data, args.Error(1)

}

func (db *DataBaseModelStructMock) Delete(id int64) (int64, error) {
	args := db.Called(id)
	return id, args.Error(1)
}

func (db *DataBaseModelStructMock) Update(data database.Data) (database.Data, error) {
	args := db.Called(data)
	return data, args.Error(1)
}

func (db *DataBaseModelStructMock) ReadOne(id int64) (database.Data, error) {
	_, mockDatOne := creatMockDataRespJSON()

	args := db.Called(id)
	return mockDatOne, args.Error(1)
}

func TestGetManyController(t *testing.T) {
	data, _ := creatMockDataRespJSON()
	mockBD := new(DataBaseModelStructMock)

	mockBD.On("Read").Return(data, nil)

	controller := ControllersStruct{
		DataBaseInterface: mockBD,
	}

	r := SetUpRouter()
	r.GET("/", controller.GetManyTodo)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	res, _ := json.Marshal(data)

	mockBD.AssertNumberOfCalls(t, "Read", 1)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(res), w.Body.String())

}
func TestAddTodoController(t *testing.T) {
	_, mockData := creatMockDataRespJSON()
	mockBD := new(DataBaseModelStructMock)
	mockData.Id = 0
	mockBD.On("Write", mockData).Return(mockData, nil)

	controller := ControllersStruct{
		DataBaseInterface: mockBD,
	}

	company := database.Data{
		Description: mockData.Description,
		Title:       mockData.Title,
		IsCompleted: mockData.IsCompleted,
	}
	jsonValue, _ := json.Marshal(company)

	r := SetUpRouter()
	r.POST("/add", controller.AddTodo)
	req, _ := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	res, _ := json.Marshal(mockData)

	mockBD.AssertNumberOfCalls(t, "Write", 1)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(res), w.Body.String())

}

func TestEditTodoController(t *testing.T) {
	_, mockData := creatMockDataRespJSON()
	mockBD := new(DataBaseModelStructMock)

	mockBD.On("Update", mockData).Return(mockData, nil)
	mockBD.On("ReadOne", mockData.Id).Return(mockData, nil)

	controller := ControllersStruct{
		DataBaseInterface: mockBD,
	}

	todo := database.Data{
		Description: mockData.Description,
		Title:       mockData.Title,
		IsCompleted: mockData.IsCompleted,
	}
	jsonValue, _ := json.Marshal(todo)

	r := SetUpRouter()
	r.PATCH("/:id", controller.EditTodo)
	req, _ := http.NewRequest(http.MethodPatch, "/"+strconv.Itoa(int(mockData.Id)), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	res, _ := json.Marshal(mockData)

	mockBD.AssertNumberOfCalls(t, "Update", 1)
	mockBD.AssertNumberOfCalls(t, "ReadOne", 1)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(res), w.Body.String())

}

func TestDeleteTodoController(t *testing.T) {
	_, mockData := creatMockDataRespJSON()
	mockBD := new(DataBaseModelStructMock)

	mockBD.On("Delete", mockData.Id).Return(mockData.Id, nil)

	controller := ControllersStruct{
		DataBaseInterface: mockBD,
	}

	r := SetUpRouter()
	r.DELETE("/:id", controller.DeleteTodo)
	req, _ := http.NewRequest(http.MethodDelete, "/"+strconv.Itoa(int(mockData.Id)), nil)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	res := []byte(`{"id":1}`)

	fmt.Println(string(res))
	mockBD.AssertNumberOfCalls(t, "Delete", 1)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(res), w.Body.String())

}
