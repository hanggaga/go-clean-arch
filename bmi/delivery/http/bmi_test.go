package http_test

import (
	"github.com/bxcodec/go-clean-arch/models"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	bmiHttp "github.com/bxcodec/go-clean-arch/bmi/delivery/http"
	"github.com/bxcodec/go-clean-arch/bmi/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateBMISuccess(t *testing.T) {
	mockBMIUsecase := new(mocks.Usecase)
	handler := bmiHttp.BMIHandler{
		BMIUsecase: mockBMIUsecase,
	}

	height := 180.0 // cm
	weight := 75.0  // kg
	expectedBMI := 23.15

	// Setting up mock expectations
	mockBMIUsecase.On("CalculateBMI", height, weight).Return(expectedBMI, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/bmi?height="+strconv.FormatFloat(height, 'f', -1, 64)+"&weight="+strconv.FormatFloat(weight, 'f', -1, 64), strings.NewReader(""))
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = handler.CalculateBMI(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), `"bmi":23.15`)

	mockBMIUsecase.AssertExpectations(t)
}

func TestCalculateBMIInvalidHeight(t *testing.T) {
	mockBMIUsecase := new(mocks.Usecase)
	handler := bmiHttp.BMIHandler{
		BMIUsecase: mockBMIUsecase,
	}

	height := "invalid" // invalid height
	weight := 75.0      // kg

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/bmi?height="+height+"&weight="+strconv.FormatFloat(weight, 'f', -1, 64), strings.NewReader(""))
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = handler.CalculateBMI(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid height parameter")

	mockBMIUsecase.AssertNotCalled(t, "CalculateBMI")
}

func TestCalculateBMIInvalidWeight(t *testing.T) {
	mockBMIUsecase := new(mocks.Usecase)
	handler := bmiHttp.BMIHandler{
		BMIUsecase: mockBMIUsecase,
	}

	height := 180.0     // cm
	weight := "invalid" // invalid weight

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/bmi?height="+strconv.FormatFloat(height, 'f', -1, 64)+"&weight="+weight, strings.NewReader(""))
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = handler.CalculateBMI(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid weight parameter")

	mockBMIUsecase.AssertNotCalled(t, "CalculateBMI")
}

func TestCalculateBMIUsecaseError(t *testing.T) {
	mockBMIUsecase := new(mocks.Usecase)
	handler := bmiHttp.BMIHandler{
		BMIUsecase: mockBMIUsecase,
	}

	height := 180.0 // cm
	weight := 75.0  // kg

	// Setting up mock expectations
	mockBMIUsecase.On("CalculateBMI", height, weight).Return(0.0, models.ErrBadParamInput)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/bmi?height="+strconv.FormatFloat(height, 'f', -1, 64)+"&weight="+strconv.FormatFloat(weight, 'f', -1, 64), strings.NewReader(""))
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = handler.CalculateBMI(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Given Param is not valid")

	mockBMIUsecase.AssertExpectations(t)
}
