package model

type RentListResponse struct {
	Status string                `json:"status"`
	Values map[string]RentDetail `json:"values"`
}

type RentDetail struct {
	ID      string `json:"id"`
	Phone   string `json:"phone"`
	EndDate string `json:"endDate"`
}
