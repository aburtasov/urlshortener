package url

import (
	"context"

	echo "github.com/labstack/echo/v4"

	"github.com/aburtasov/urlshortener/internal/models"
)

type EchoDelivery interface {
	Create(ectx echo.Context) error
	List(ectx echo.Context) error
	Update(ectx echo.Context) error
	Delete(ectx echo.Context) error
}

type Usecase interface {
	Set(ctx context.Context, url *models.Url) (string, error)
	Get(ctx context.Context, url *models.Url) (string, error)
	Delete(ctx context.Context, url *models.Url) (string, error)
	Update(ctx context.Context, url *models.Url) (string, error)
}

type Repository interface {
	Set(ctx context.Context, url *models.Url) (string, error)
	Get(ctx context.Context, url *models.Url) (string, error)
	Delete(ctx context.Context, url *models.Url) (string, error)
	Update(ctx context.Context, url *models.Url) (string, error)
}
