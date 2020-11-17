package view

// PromoCompany ...
type Promo struct {
	PromoID     int    `json:"promo_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
