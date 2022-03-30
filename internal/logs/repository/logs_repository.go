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

func (r *LogsRepository) Create(ctx context.Context, logs ...*domain.Log) error {
	builder := r.sqlGen.Insert("logs").
		Columns(
			"service_name",
			"log_level",
			"span_id",
			"tstamp",
			"src",
			"message")

	for _, log := range logs {
		builder = builder.Values(
			log.ServiceName,
			log.Level,
			log.SpanID,
			log.Timestamp,
			log.Source,
			log.Message)
	}

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	if _, err = r.clickHouseConn.Exec(sqlQuery, args...); err != nil {
		return err
	}

	return nil
}
