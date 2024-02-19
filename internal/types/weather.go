package types

type Weather struct {
	Date          string        `json:"date"`
	Weather       string        `json:"weather"`
	Summary       string        `json:"summary"`
	Temperature   float32       `json:"temperature"`
	Cloud_cover   float32       `json:"cloud_cover"`
	Icon          int           `json:"icon"`
	Wind          Wind          `json:"wind"`
	Precipitation Precipitation `json:"precipitation"`
}
