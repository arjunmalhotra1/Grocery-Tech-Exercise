package store

import (
	"fmt"
	"strconv"
	"strings"

	structTagValidator "github.com/go-playground/validator/v10"
)

const (
	errInvalidProduceCode = "invalid produce code"
)

var validate *structTagValidator.Validate

func isValidItem(items []Item) error {
	validate = structTagValidator.New()
	for _, item := range items {
		// Converting produce name to lower case
		item.Name = strings.ToLower(item.Name)
		stringPrice := fmt.Sprintf("%.2f", item.Price)
		// TODO: handle this error for Internal Server Error.
		item.Price, _ = strconv.ParseFloat(stringPrice, 64)
		err := isValidProduceCode(item.ProduceCode)
		if err != nil {
			return err
		}
		err = validate.Struct(item)
		if err != nil {
			validationErrs, ok := err.(structTagValidator.ValidationErrors)
			if ok {
				for _, validationErr := range validationErrs {
					if validationErr.Tag() == "required" {
						return fmt.Errorf("field missing %s expected type %s", validationErr.Field(), validationErr.Type())
					}
				}
			}
		}
	}
	return nil
}

func isValidProduceCode(ProduceCode string) error {
	splitCodes := strings.Split(ProduceCode, "-")
	if len(splitCodes) != 4 {
		return fmt.Errorf(errInvalidProduceCode)
	}
	for _, v := range splitCodes {
		if len(v) != 4 {
			return fmt.Errorf(errInvalidProduceCode)
		}
	}

	return nil
}
