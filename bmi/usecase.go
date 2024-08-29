package bmi

type Usecase interface {
	CalculateBMI(height, weight float64) (float64, error)
}
