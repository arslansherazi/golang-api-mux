package participation_api

type Validator struct {
	CompetitionID int64 `vaidator:"required"`
	UserID        int64 `vaidator:"required"`
}
