package controller

import (
	"AccubitsMachineTest/accubits.machine.test/logger"
	models "AccubitsMachineTest/accubits.machine.test/model"
	test "AccubitsMachineTest/accubits.machine.test/service"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type TestController struct {
	Usecase test.TestUsecase
}

/***********************************************Save Data**************************************/
// SaveDatafromCoursera
//@Tags coursera service
// @Summary save data from coursera
// @Description save data from coursera
// @Accept  json
// @Produce  json
// @Param call body models.Req true "save data from coursera"
// @Success 200 {object} interface{}
// @Failure 404  "Not Found"
// @Router /api/v1/coursera/saveData [post]
func (r *TestController) SaveData(c echo.Context) error {
	var req models.Req
	c.Bind(&req)
	if req.Query == "" || req.Page == 0 {
		logger.Logger.Error("Query or page no. is missing.")
		return c.JSON(400, &models.Response{IsSuccess: false, StatusCode: http.StatusBadRequest, Message: "Please enter query and page no. both."})
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Logger.Info("Request from controller part of save data is ", req)
	authResponse, _ := r.Usecase.SaveData(ctx, req)
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}

/**********************************************Fetch Call logs**************************************/
// FetchData
//@Tags coursera service
// @Summary Fetch data
// @Description fetch data
// @Accept  json
// @Produce  json
// @Param query query string true "query"
// @Param limit query int false "limit"
// @Param page query int false "page"
// @Success 200 {object} interface{}
// @Failure 404  "Not Found"
// @Router /api/v1/coursera/fetchData [get]
func (r *TestController) FetchData(c echo.Context) error {
	source := c.QueryParam("query")
	query := c.Request().URL.Query()
	if source == "" {
		logger.Logger.Error("Query is missing.")
		return c.JSON(400, &models.Response{IsSuccess: false, StatusCode: http.StatusBadRequest, Message: "Please enter query."})
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	logger.Logger.Info("Request from controller part of fetch data is ", logrus.Fields{"query": source})
	authResponse, _ := r.Usecase.FetchData(ctx, source, query)
	if authResponse == nil {
		return c.JSON(http.StatusUnauthorized, authResponse)
	}
	return c.JSON(http.StatusOK, authResponse)
}
