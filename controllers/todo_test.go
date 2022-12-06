package controllers

import (
	"encoding/json"
	"gin-golang/database"
	"net/http"
	"net/http/httptest"
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

	args := db.Called()
	return data, args.Error(1)

}

func (db *DataBaseModelStructMock) Delete(id int64) (int64, error) {
	args := db.Called()
	return id, args.Error(1)
}

func (db *DataBaseModelStructMock) Update(data database.Data) (database.Data, error) {
	args := db.Called()
	return data, args.Error(1)
}

func (db *DataBaseModelStructMock) ReadOne(id int64) (database.Data, error) {
	_, mockDatOne := creatMockDataRespJSON()

	args := db.Called()
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
