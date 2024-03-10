package validators

import (
	"github.com/go-playground/validator/v10"
)

type FieldValidator struct {
	Validator *validator.Validate
}

func (fv *FieldValidator) Validate(i interface{}) error {
	return fv.Validator.Struct(i)
}

func NewFieldValidator() (*FieldValidator, error) {
	fieldValidator := &FieldValidator{Validator: validator.New()}
	err := fieldValidator.Init()

	return fieldValidator, err
}

func (fv *FieldValidator) Init() error {
	// register custom validations here
	return nil
}
