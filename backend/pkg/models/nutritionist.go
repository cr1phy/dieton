package models

type Nutritionist interface {
	CalculateCalories(proteins float64, fats float64, carbs float64) float64
}
