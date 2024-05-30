package services

import (
	"gofr.dev/pkg/gofr"

	"github.com/tejasmd22/exoplanet-service/models"
)

type Exoplanet interface {
	Create(ctx *gofr.Context, exoplanetCreateRequest *models.ExoplanetCreateRequest) (*models.Exoplanet, error)
}
