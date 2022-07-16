package validate_phone_number_api

type Validator struct {
	PhoneNumber string `validate:"required,min=4,max=20"`
}
