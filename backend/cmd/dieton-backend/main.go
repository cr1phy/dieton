package main

import (
	"log"
	"os"

	"github.com/cr1phy/dieton-backend/pkg/models"
	"github.com/cr1phy/dieton-backend/pkg/router"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var v *validator.Validate

func InitValidation() {
	v = validator.New()
	v.RegisterStructValidation(func(sl validator.StructLevel) {
		user := sl.Current().Interface().(models.User)

		if len(user.FirstName) == 0 && len(user.LastName) == 0 {
			sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
			sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
		}
	}, models.User{})
}

func main() {
	InitValidation()

	dsn := os.Getenv("POSTGRES_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Something went wrong with database: %s", err)
	}

	dbRepo := models.NewDBRepo(db)
	handler := router.NewRouterHandler(dbRepo)

	r := router.InitRouter(handler)
	r.Run("0.0.0.0:8080")
}
