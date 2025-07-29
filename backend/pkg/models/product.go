package models

type Product struct {
	ID       uint
	Name     string
	Calories float64
	Proteins float64
	Fats     float64
	Carbs    float64
}

func (p *Product) CalculateCalories(proteins float64, fats float64, carbs float64) float64 {
	return proteins*4.0 + fats*9.0 + carbs*4.0
}

func NewProduct(name string, calories float64, proteins float64, fats float64, carbs float64) *Product {
	return &Product{
		ID:       0,
		Name:     name,
		Calories: calories,
		Proteins: proteins,
		Fats:     fats,
		Carbs:    carbs,
	}
}
