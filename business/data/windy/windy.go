// Package windy contains user related CRUD functionality.
package windy

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/leogoesger/goservices/foundation/web"
	"github.com/pkg/errors"
)

var (
	// ErrNoPointResponse is used when a http requeat failed to get GRID info from a point
	ErrNoPointResponse = errors.New("cannot get GRID info from lat and long from NOAA")

	// ErrParsePointResponse occurs when POINT response cannot be parsed
	ErrParsePointResponse = errors.New("cannot parse POINT response from NOAA")

	// ErrNoWeatherInfo occurs when a http request failed to get GRID weather in at NOAA
	ErrNoWeatherInfo = errors.New("cannot get weather info from NOAA")

	// ErrParseGridResponse occurs when unmarshalling weather response from grid
	ErrParseGridResponse = errors.New("cannot parse grid response")

	// ErrParseWindSpeed occurs when the wind speed cannot be parsed. Valid wind speed is "12 mph"
	ErrParseWindSpeed = errors.New("cannot parse wind speed")
)

// GetReport returns the aggragated wind report at a given location. It converts lat and long to
// a NOAA grid system. With the GRID number, we can pull weather information directly from NOAA
func GetReport(ctx context.Context, traceID string, lat string, long string, hours int) (Report, error) {

	// Using lat and long calling POINTS api from NOAA to get GRID numbers
	pointRes, err := http.Get(`https://api.weather.gov/points/` + lat + `,` + long)
	if err != nil {
		web.NewRequestError(ErrNoPointResponse, http.StatusBadRequest)
	}
	defer pointRes.Body.Close()
	pr := new(NOAAPointResponse)
	pointBody, err := ioutil.ReadAll(pointRes.Body)
	if err != nil {
		web.NewRequestError(ErrParsePointResponse, http.StatusBadRequest)
	}
	json.Unmarshal(pointBody, &pr)

	// FROM points API, there is a property call forecastHourly. It is the same URL can
	// be used for GRID weather info. We could construct the URL, or use it directly here
	resp, err := http.Get(pr.Properties.ForecastHourly)
	if err != nil {
		web.NewRequestError(ErrNoWeatherInfo, http.StatusBadRequest)
	}
	defer resp.Body.Close()

	r := new(NOAAGridResponse)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		web.NewRequestError(ErrParseGridResponse, http.StatusBadRequest)
	}
	json.Unmarshal(body, r)

	wd := Report{
		AggregatedWind: AggregatedWind{
			N:  Wind{TotalWind: 0, UnitCode: "mph"},
			W:  Wind{TotalWind: 0, UnitCode: "mph"},
			E:  Wind{TotalWind: 0, UnitCode: "mph"},
			S:  Wind{TotalWind: 0, UnitCode: "mph"},
			NE: Wind{TotalWind: 0, UnitCode: "mph"},
			NW: Wind{TotalWind: 0, UnitCode: "mph"},
			SE: Wind{TotalWind: 0, UnitCode: "mph"},
			SW: Wind{TotalWind: 0, UnitCode: "mph"}},
		NOAAElevation: r.Properties.Elevation,
	}

	// Loop through all the wind speed, and break at the # of hours given
	for idx, value := range r.Properties.Periods {
		if idx > hours {
			break
		}

		// Windspeed format is "12 mph", therefore we need to split and parse into float64
		windSpeed, err := strconv.ParseFloat(strings.Split(value.WindSpeed, " ")[0], 32)
		if err != nil {
			web.NewRequestError(ErrParseWindSpeed, http.StatusBadRequest)
		}
		switch value.WindDirection {
		case "N":
			wd.AggregatedWind.N.TotalWind += windSpeed
		case "S":
			wd.AggregatedWind.S.TotalWind += windSpeed
		case "W":
			wd.AggregatedWind.W.TotalWind += windSpeed
		case "E":
			wd.AggregatedWind.E.TotalWind += windSpeed
		case "NW":
			wd.AggregatedWind.NW.TotalWind += windSpeed
		case "NE":
			wd.AggregatedWind.NE.TotalWind += windSpeed
		case "SW":
			wd.AggregatedWind.SW.TotalWind += windSpeed
		case "SE":
			wd.AggregatedWind.SE.TotalWind += windSpeed
		default:
		}
	}

	return wd, nil
}
