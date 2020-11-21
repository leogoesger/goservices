// Package windy contains user related CRUD functionality.
package windy

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrNotFound is used when a specific User is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("attempted action is not allowed")
)

// GetReport ...
func GetReport(ctx context.Context, traceID string) (Info, error) {
	resp, err := http.Get("https://api.weather.gov/gridpoints/TOP/31,80/forecast/hourly")
	wd := Info{
		N:  Wind{TotalWind: 12},
		W:  Wind{TotalWind: 12},
		E:  Wind{TotalWind: 12},
		S:  Wind{TotalWind: 12},
		NE: Wind{TotalWind: 12},
		NW: Wind{TotalWind: 12},
		SE: Wind{TotalWind: 12},
		SW: Wind{TotalWind: 12},
	}
	defer resp.Body.Close()

	r := new(NOAAResponse)
	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, r)
	fmt.Println(r)
	if err != nil {
		panic(err)
	}
	return wd, nil

}
