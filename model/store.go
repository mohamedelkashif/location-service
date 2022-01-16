package model

type Store struct {
	StoreId     string    `json:"store_id"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	CountryCode string    `json:"country_code"`
	Location    Locationn `json:"location"`
	SlowService bool      `json:"slow_service"`
}
