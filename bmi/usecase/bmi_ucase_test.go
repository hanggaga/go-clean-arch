package usecase_test

import (
	"testing"

	"github.com/bxcodec/go-clean-arch/bmi/usecase"
	"github.com/stretchr/testify/assert"
)

func TestCalculateBMI(t *testing.T) {
	u := usecase.NewBMIUsecase()

	t.Run("success", func(t *testing.T) {
		height := 180.0      // height in cm
		weight := 75.0       // weight in kg
		expectedBMI := 23.15 // Expected BMI value

		bmiValue, err := u.CalculateBMI(height, weight)
		assert.NoError(t, err)
		assert.InEpsilon(t, expectedBMI, bmiValue, 0.01) // Allowing a small delta for float comparison
	})

	t.Run("error with zero height", func(t *testing.T) {
		height := 0.0  // height in cm
		weight := 75.0 // weight in kg

		bmiValue, err := u.CalculateBMI(height, weight)
		assert.Error(t, err)
		assert.Equal(t, 0.0, bmiValue)
	})

	t.Run("error with zero weight", func(t *testing.T) {
		height := 180.0 // height in cm
		weight := 0.0   // weight in kg

		bmiValue, err := u.CalculateBMI(height, weight)
		assert.Error(t, err)
		assert.Equal(t, 0.0, bmiValue)
	})

	t.Run("error with negative height", func(t *testing.T) {
		height := -180.0 // negative height in cm
		weight := 75.0   // weight in kg

		bmiValue, err := u.CalculateBMI(height, weight)
		assert.Error(t, err)
		assert.Equal(t, 0.0, bmiValue)
	})

	t.Run("error with negative weight", func(t *testing.T) {
		height := 180.0 // height in cm
		weight := -75.0 // negative weight in kg

		bmiValue, err := u.CalculateBMI(height, weight)
		assert.Error(t, err)
		assert.Equal(t, 0.0, bmiValue)
	})
}
