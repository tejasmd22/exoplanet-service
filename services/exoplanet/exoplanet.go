package services

import (
	"gofr.dev/pkg/gofr"

	"github.com/tejasmd22/exoplanet-service/models"
	"github.com/tejasmd22/exoplanet-service/services"
	"github.com/tejasmd22/exoplanet-service/stores"
)

type exoplanet struct {
	exoplanetStores stores.Exoplanet
}

func New(exoplanetStores stores.Exoplanet) services.Exoplanet {
	return &exoplanet{exoplanetStores: exoplanetStores}
}

func (e *exoplanet) Create(ctx *gofr.Context, exoplanetCreateRequest *models.ExoplanetCreateRequest) (*models.Exoplanet, error) {
	if err := exoplanetCreateRequest.Validate(); err != nil {
		return nil, err
	}

	exoplanet, err := e.exoplanetStores.Create(ctx, exoplanetCreateRequest)
	if err != nil {
		return nil, err
	}

	return exoplanet, nil
}
