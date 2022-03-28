package logs

import (
	"context"

	"github.com/cleverlog/api/internal/logs/domain"
)

type UseCase interface {
	SendLog(ctx context.Context, log *domain.Log) error
}
