package dto

type AddExperienceRequest struct {
	Experience int `json:"experience" validate:"required,min=1"`
}
