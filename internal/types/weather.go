package types

type Weather struct {
	Date          string        `json:"date"`
	Weather       string        `json:"weather"`
	Summary       string        `json:"summary"`
	Temperature   float32       `json:"temperature"`
	Cloud_cover   Cloud_cover   `json:"cloud_cover"`
	Icon          int           `json:"icon"`
	Wind          Wind          `json:"wind"`
	Precipitation Precipitation `json:"precipitation"`
}

func (i Weather) Title() string       { return i.Weather }
func (i Weather) Description() string { return i.Summary }
func (w Weather) FilterValue() string { return "" }
