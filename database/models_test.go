package database

import (
	"encoding/json"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createMockFunc() ([]byte, func(filename string) ([]byte, error), func(filename string, data []byte, perm fs.FileMode) error) {
	dataMock := []byte(`[{
		"Id":0,
		"Title":"Title",
		"Description":"Description",
		"IsCompleted": false
		}]`,
	)
	mockRender := func(filename string) ([]byte, error) {
		return dataMock, nil
	}

	mockWrite := func(filename string, data []byte, perm fs.FileMode) error {

		return nil
	}
	return dataMock, mockRender, mockWrite
}

func TestReadFunc(t *testing.T) {

	dataMock, mockRender, mockWrite := createMockFunc()

	r := DataBaseModelStruct{
		IoReader: mockRender,
		IoWriter: mockWrite,
	}
	result, err := r.Read()

	var dataTodo []Data
	json.Unmarshal(dataMock, &dataTodo)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, dataTodo, result)
}

func TestWriteFunc(t *testing.T) {
	_, mockRender, mockWrite := createMockFunc()

	var dataTodoMock Data
	dataTodoMock.Description = "Description"
	dataTodoMock.Title = "Title"
	dataTodoMock.IsCompleted = false
	dataTodoMock.Id = 0

	r := DataBaseModelStruct{
		IoReader: mockRender,
		IoWriter: mockWrite,
	}
	result, err := r.Write(dataTodoMock)
	result.Id = 0

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, dataTodoMock, result)
}

func TestUpdateFunc(t *testing.T) {
	_, mockRender, mockWrite := createMockFunc()

	var dataTodoMock Data
	dataTodoMock.Description = "Description + edited"
	dataTodoMock.Title = "Title"
	dataTodoMock.IsCompleted = false
	dataTodoMock.Id = 0

	r := DataBaseModelStruct{
		IoReader: mockRender,
		IoWriter: mockWrite,
	}
	result, err := r.Update(dataTodoMock)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, dataTodoMock, result)
}

func TestDeleteFunc(t *testing.T) {
	_, mockRender, mockWrite := createMockFunc()

	r := DataBaseModelStruct{
		IoReader: mockRender,
		IoWriter: mockWrite,
	}
	result, err := r.Delete(0)

	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, int64(0), result)
}
