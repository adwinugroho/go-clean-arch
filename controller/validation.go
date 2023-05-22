package controller

import (
	"fmt"
	"go-clean-arch/models"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func (route *OrderRoute) Validate(i interface{}) error {
	err := route.Validator.Struct(i)
	if err != nil {
		log.Printf("Error when validate:%+v", err)
		var errMessage string
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required_without_all":
					errMessage = fmt.Sprintf("%s is required", err.Field())
				case "required":
					errMessage = fmt.Sprintf("%s is required", err.Field())
				default:
					errMessage = "Invalid Data"
				}
			}
		}
		// log.Println("error message:", errMessage)
		return models.NewError(http.StatusBadRequest, errMessage)
	}
	return nil
}
