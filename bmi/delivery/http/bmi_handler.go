package http

import (
	"github.com/bxcodec/go-clean-arch/bmi"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type BMIHandler struct {
	BMIUsecase bmi.Usecase
}

func NewBMIHandler(e *echo.Echo, usecase bmi.Usecase) {
	handler := &BMIHandler{
		BMIUsecase: usecase,
	}
	e.GET("/bmi", handler.CalculateBMI)
}

func (h *BMIHandler) CalculateBMI(c echo.Context) error {
	heightStr := c.QueryParam("height")
	weightStr := c.QueryParam("weight")

	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid height parameter")
	}

	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid weight parameter")
	}

	bmiValue, err := h.BMIUsecase.CalculateBMI(height, weight)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]float64{"bmi": bmiValue})
}
