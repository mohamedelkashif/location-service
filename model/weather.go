package model

type Weather struct {
	Temp                 int64  `json:"temp"`
	PrecipitationLevel   string `json:"precipitationLevel"`
	PrecipitationType    string `json:"precipitationType"`
	PrecipitationLast24h int64  `json:"precipitationLast24h"`
}
