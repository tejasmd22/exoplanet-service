package stores

import (
	"strings"

	"github.com/tejasmd22/exoplanet-service/models"
)

const (
	insertQuery  = "INSERT INTO exoplanet "
	selectQuery = "SELECT id, name, description, distance, radius, mass, type FROM exoplanet "
)

func (e *exoplanet) generateInsertQuery(exoplanetCreateRequest *models.ExoplanetCreateRequest) (query string, values []interface{}) {
	if exoplanetCreateRequest.Name != "" {
		query += "name, "

		values = append(values, exoplanetCreateRequest.Name)
	}

	if exoplanetCreateRequest.Description != "" {
		query += "description, "

		values = append(values, exoplanetCreateRequest.Description)
	}

	if exoplanetCreateRequest.Distance > 0 {
		query += "distance, "

		values = append(values, exoplanetCreateRequest.Distance)
	}

	if exoplanetCreateRequest.Radius > 0 {
		query += "radius, "

		values = append(values, exoplanetCreateRequest.Radius)
	}

	if exoplanetCreateRequest.Mass > 0 {
		query += "mass, "

		values = append(values, exoplanetCreateRequest.Mass)
	}

	if exoplanetCreateRequest.Type != "" {
		query += "type, "

		values = append(values, exoplanetCreateRequest.Type)
	}

	query = strings.Trim(query, ", ")
	if len(values) > 0 {
		query = "(" + query + ") VALUES"
		q := strings.Repeat("?, ", len(values))
		q = strings.Trim(q, ", ")
		query += "(" + q + ")"
	}

	query = insertQuery + query

	return query, values
}

// func (e *exoplanet) ge