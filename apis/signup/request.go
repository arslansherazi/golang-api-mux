package signup_api

import (
	"find_competitor/common"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	logger, err := common.GetLogger("signup_api")
	if err != nil {
		common.Generate500ErrorResponse(r.URL.Path, w)
	} else {
		user, profileImage, err, isValidationError := processRequestParams(logger, r)
		if err != nil {
			if isValidationError {
				validationMessage := common.ParseValidationError(err)
				common.Generate422Response(r.URL.Path, validationMessage, w)
			} else {
				logger.Println(err)
				common.Generate500ErrorResponse(r.URL.Path, w)
			}
		} else {
			user.ProfileImageUrl, err = generateProfileImageUrl(profileImage)
			if err != nil {
				common.Generate500ErrorResponse(r.URL.Path, w)
			} else {
				user.Password, err = createHashOfPassword(user.Password)
				if err != nil {
					common.Generate500ErrorResponse(r.URL.Path, w)
				} else {
					err = insertUserIntoDB(user)
					if err != nil {
						common.Generate500ErrorResponse(r.URL.Path, w)
					} else {
						generateSignupSuccessResponse(r.URL.Path, w)
					}
				}
			}
		}
	}
}
