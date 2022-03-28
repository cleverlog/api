package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"

	"github.com/cleverlog/api/internal/logs/domain"
)

type LogsRepository struct {
	clickHouseConn *sql.DB
	sqlGen         squirrel.StatementBuilderType
}

func New(clickConn *sql.DB) *LogsRepository {
	return &LogsRepository{
		clickHouseConn: clickConn,
		sqlGen:         squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *LogsRepository) Create(ctx context.Context, log *domain.Log) error {
	sqlQuery, args, err := r.sqlGen.Insert("logs").
		SetMap(map[string]interface{}{
			"service_name": log.ServiceName,
			"span_id":      log.SpanID,
			"tstamp":       log.Timestamp,
			"src":          log.Source,
			"message":      log.Message,
		}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.clickHouseConn.Exec(sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
