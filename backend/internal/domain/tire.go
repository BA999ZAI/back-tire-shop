package domain

type Tire struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Price       float64 `json:"price"`
	Brand       string  `json:"brand"`
	Model       string  `json:"model"`
	Season      string  `json:"season"`
	Width       int     `json:"width"`
	Length      int     `json:"length"`
	Radius      int     `json:"radius"`
	Amount      int     `json:"amount"`
	SpeedLimit  int     `json:"speed_limit"`
	LoadLimit   int     `json:"load_limit"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	Spikes      int     `json:"spikes"`
}
