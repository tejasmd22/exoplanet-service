package models

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
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type ExoplanetCreateRequest struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    float64       `json:"distance"`
	Radius      float64       `json:"radius"`
	Mass        float64       `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type"`
}

type ExoplanetResponse struct {
	Exoplanet Exoplanet `json:"exoplanet"`
}

func (e *ExoplanetCreateRequest) Validate() error {
	return nil
}
