// Package windy contains user related CRUD functionality.
package windy

// Wind ...
type Wind struct {
	TotalWind float64
	UnitCode  string
}

// AggregatedWind ..
type AggregatedWind struct {
	N  Wind
	S  Wind
	W  Wind
	E  Wind
	NE Wind
	NW Wind
	SE Wind
	SW Wind
}

// Report ...
type Report struct {
	NOAAElevation  NOAAElevation
	AggregatedWind AggregatedWind
}

// NOAAGridResponse ..
type NOAAGridResponse struct {
	Properties NOAAProperty `json:"properties"`
}

// NOAAProperty ..
type NOAAProperty struct {
	Elevation NOAAElevation `json:"elevation"`
	Periods   []NOAAPeriod  `json:"periods"`
}

// NOAAElevation ..
type NOAAElevation struct {
	Value    float32 `json:"value"`
	UnitCode string  `json:"unitCode"`
}

// NOAAPeriod ..
type NOAAPeriod struct {
	StartTime     string  `json:"startTime"`
	EndTime       string  `json:"endTime"`
	Temperature   float32 `json:"temperature"`
	WindSpeed     string  `json:"windSpeed"`
	WindDirection string  `json:"windDirection"`
}

// NOAAPointResponse ...
type NOAAPointResponse struct {
	Properties NOAAPointProperty `json:"properties"`
}

// NOAAPointProperty ...
type NOAAPointProperty struct {
	ForecastOffice string `json:"forecastOffice"`
	ForecastHourly string `json:"forecastHourly"`
}
