package logs

import (
	"context"

	"github.com/cleverlog/api/internal/logs/domain"
)

type Repository interface {
	Create(ctx context.Context, logs ...*domain.Log) error
}
