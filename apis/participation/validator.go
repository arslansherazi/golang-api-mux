package participation_api

type Validator struct {
	CompetitionID uint64 `vaidator:"required"`
	UserID        uint64 `vaidator:"required"`
}
