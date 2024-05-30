package http

import (
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"

	"github.com/tejasmd22/exoplanet-service/models"
	"github.com/tejasmd22/exoplanet-service/services"
)

type handler struct {
	exoplanetServices services.Exoplanet
}

func New(exoplanetServices services.Exoplanet) *handler {
	return &handler{exoplanetServices: exoplanetServices}
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var exoplanetCreateRequest models.ExoplanetCreateRequest

	err := ctx.Bind(exoplanetCreateRequest)
	if err != nil {
		return nil, http.ErrorInvalidParam{Params: []string{"json request body"}}
	}

	exoplanet, err := h.exoplanetServices.Create(ctx, &exoplanetCreateRequest)
	if err != nil {
		return nil, err
	}

	return models.ExoplanetResponse{
		Exoplanet: *exoplanet,
	}, nil
}
