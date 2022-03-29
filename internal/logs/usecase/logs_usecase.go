package usecase

import (
	"context"

	"github.com/cleverlog/api/internal/logs"
	"github.com/cleverlog/api/internal/logs/domain"
)

type LogsUseCase struct {
	logsRepository logs.Repository
}

func New(repo logs.Repository) *LogsUseCase {
	return &LogsUseCase{
		logsRepository: repo,
	}
}

func (r *LogsUseCase) Create(ctx context.Context, logs ...*domain.Log) error {
	return r.logsRepository.SendLog(ctx, logs...)
}
