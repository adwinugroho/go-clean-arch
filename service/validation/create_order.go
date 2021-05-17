package validation

import (
	"fmt"
	"go-clean-arch/models/request"
	"log"

	"github.com/go-playground/validator"
)

func ValidateCreateOrder(request request.CreateOrderLRequest) (string, error) {
	validate := validator.New()
	err := validate.Struct(request)
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
		return errMessage, err
	}

	return "", nil
}
