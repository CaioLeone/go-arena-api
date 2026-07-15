package dto

type SpendAttributePointsRequest struct {
	HP             int `json:"hp" validate:"min=0,max=100"`
	Attack         int `json:"attack" validate:"min=0,max=100"`
	Defense        int `json:"defense" validate:"min=0,max=100"`
	CriticalChance int `json:"critical_chance" validate:"min=0,max=100"`
}
