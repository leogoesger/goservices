// Package windy contains user related CRUD functionality.
package windy

// UserInput contains information needed to create a new User.
type UserInput struct {
	X     float32 `json:"x" validate:"required"`
	Y     float32 `json:"y" validate:"required,email"`
	Hours int     `json:"hours" validate:"required"`
}

// Wind ...
type Wind struct {
	TotalWind float32
}

// Info ..
type Info struct {
	N  Wind
	S  Wind
	W  Wind
	E  Wind
	NE Wind
	NW Wind
	SE Wind
	SW Wind
}

// NOAAResponse ..
type NOAAResponse struct {
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
