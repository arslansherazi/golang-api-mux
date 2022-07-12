package login_api

type Validator struct {
	PhoneNumber string `validate:"required,min=4,max=20"`
	Password    string `validate:"required,min=8,max=16"`
}
