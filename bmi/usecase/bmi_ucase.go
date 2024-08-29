package usecase

import (
	"fmt"
	"github.com/bxcodec/go-clean-arch/bmi"
	"github.com/shopspring/decimal"
)

type bmiUsecase struct{}

func NewBMIUsecase() bmi.Usecase {
	return &bmiUsecase{}
}

func (u *bmiUsecase) CalculateBMI(height, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, fmt.Errorf("height and weight must be greater than zero")
	}

	heightInMeters := decimal.NewFromFloat(height).Div(decimal.NewFromFloat(100))
	weightInKg := decimal.NewFromFloat(weight)

	heightSquared := heightInMeters.Mul(heightInMeters)

	bmiDecimal := weightInKg.Div(heightSquared)
	bmiFloat, _ := bmiDecimal.Float64()

	return bmiFloat, nil
}
