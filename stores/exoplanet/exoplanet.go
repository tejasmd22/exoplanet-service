package stores

import (
	"database/sql"
	"strconv"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource"
	"gofr.dev/pkg/gofr/http"

	"github.com/tejasmd22/exoplanet-service/models"
)

type exoplanet struct{}

func New() *exoplanet {
	return &exoplanet{}
}

func (e *exoplanet) Create(ctx *gofr.Context, exoplanetCreateRequest *models.ExoplanetCreateRequest) (*models.Exoplanet, error) {
	query, values := e.generateInsertQuery(exoplanetCreateRequest)

	result, err := ctx.SQL.ExecContext(ctx, query, values...)
	if err != nil {
		return nil, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	exoplanetID, err := result.LastInsertId()
	if err != nil {
		return nil, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	return e.GetByID(ctx, int(exoplanetID))
}

func (e *exoplanet) GetByID(ctx *gofr.Context, id int) (*models.Exoplanet, error) {
	query := selectQuery + "WHERE id = ? "

	var exoplanet models.Exoplanet

	err := ctx.SQL.QueryRowContext(ctx, query, id).Scan(&exoplanet.ID, &exoplanet.Name,
		&exoplanet.Description, &exoplanet.Distance, &exoplanet.Radius, &exoplanet.Mass, &exoplanet.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.ErrorEntityNotFound{
				Name:  "exoplanet",
				Value: strconv.Itoa(id),
			}
		}

		return nil, datasource.ErrorDB{Err: err, Message: "error from db"}
	}

	return &exoplanet, nil
}
