package types

type Weather struct {
	Icon          string        `json:"icon"`
	Summary       string        `json:"summary"`
	Temperature   float32       `json:"temperature"`
	Cloud_cover   float32       `json:"cloud_cover"`
	Wind          Wind          `json:"wind"`
	Precipitation Precipitation `json:"precipitation"`
}
