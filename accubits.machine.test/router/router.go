package router

import (
	controller "AccubitsMachineTest/accubits.machine.test/controller"
	test "AccubitsMachineTest/accubits.machine.test/service"

	"github.com/labstack/echo/v4"
)

func NewTestController(e *echo.Echo, testusecase test.TestUsecase) {
	handler := &controller.TestController{
		Usecase: testusecase,
	}
	e.POST("api/v1/coursera/saveData", handler.SaveData)
	e.GET("api/v1/coursera/fetchData", handler.FetchData)
}
