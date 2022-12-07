package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type ControllersStructMock struct {
	mock.Mock
}

func (ctrl *ControllersStructMock) GetManyTodo(c *gin.Context) {
	ctrl.Called(c)
}
func (ctrl *ControllersStructMock) AddTodo(c *gin.Context) {
	ctrl.Called(c)
}
func (ctrl *ControllersStructMock) EditTodo(c *gin.Context) {
	ctrl.Called(c)
}
func (ctrl *ControllersStructMock) DeleteTodo(c *gin.Context) {
	ctrl.Called(c)
}

func TestTodoRouters(t *testing.T) {

	controlMock := new(ControllersStructMock)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	c.Request = req
	controlMock.On("GetManyTodo", c).Return()

	router := TodoRoutersStruct{
		R:                    r,
		ControllersInterface: controlMock,
	}

	router.TodoRouters()

	router.R.ServeHTTP(w, req)

	controlMock.AssertNumberOfCalls(t, "GetManyTodo", 1)

}
