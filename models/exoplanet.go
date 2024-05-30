package models

import (
	"fmt"
	"net/http"
	"strings"
)

type Exoplanet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    float64 `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass,omitempty"`
	Type        string  `json:"type"`
}

type ExoplanetType string

const (
	GasGiant    = "GASGIANT"
	Terrestrial = "TERRESTRIAL"
)

type ExoplanetCreateRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    float64 `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass"`
	Type        string  `json:"type"`
}

type ExoplanetResponse struct {
	Exoplanet *Exoplanet `json:"exoplanet"`
}

type ExoplanetDetail struct {
	Data   []*Exoplanet `json:"exoplanet"`
	Count  int          `json:"count"`
	Offset int          `json:"offset"`
	Limit  int          `json:"limit"`
}

func (e *ExoplanetCreateRequest) Validate() error {
	if e.Mass > 0 && !strings.EqualFold(string(e.Type), string(Terrestrial)) {
		err := customError{error: "mass will be provided only in case of Terrestrial type of planet"}
		return err
	}

	if e.Type != "" {
		if !strings.EqualFold(string(e.Type), string(Terrestrial)) && !strings.EqualFold(string(e.Type), string(GasGiant)) {
			err := customError{error: "only GasGiant or Terrestrial types are allowed"}
			return err
		}

		e.Type = strings.ToUpper(e.Type)
	}

	return nil
}

type ExoplanetUpdateRequest struct {
	ID          int     `json:"-"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    float64 `json:"distance"`
	Radius      float64 `json:"radius"`
	Mass        float64 `json:"mass"`
	Type        string  `json:"type"`
}

func (e *ExoplanetUpdateRequest) Validate() error {
	if e.Mass > 0 && !strings.EqualFold(string(e.Type), string(Terrestrial)) {
		err := customError{error: "mass will be provided only in case of Terrestrial type of planet"}
		return err
	}

	if e.Type != "" {
		if !strings.EqualFold(string(e.Type), string(Terrestrial)) && !strings.EqualFold(string(e.Type), string(GasGiant)) {
			err := customError{error: "only GasGiant or Terrestrial types are allowed"}
			return err
		}

		e.Type = strings.ToUpper(e.Type)
	}

	return nil
}

type customError struct {
	error string
}

func (c customError) Error() string {
	return fmt.Sprintf("error: %s", c.error)
}

func (c customError) StatusCode() int {
	return http.StatusBadRequest
}
