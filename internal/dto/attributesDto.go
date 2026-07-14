package dto

type SpendAttributePointsRequest struct {
	Attribute string `json:"attribute" validate:"required,oneof=hp attack defense critical_chance"`
}
