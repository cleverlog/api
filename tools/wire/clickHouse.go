package wire

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/cleverlog/api/cmd/migrations"
)

func NewClickHouse() (*sql.DB, func(), error) {
	conf := viper.New()

	conf.AutomaticEnv()

	host := conf.GetString("CLICKHOUSE_HOST")
	port := conf.GetString("CLICKHOUSE_PORT")
	user := conf.GetString("CLICKHOUSE_USER")
	pswd := conf.GetString("CLICKHOUSE_PSWD")
	db := conf.GetString("CLICKHOUSE_DB")

	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", host, port)},
		Auth: clickhouse.Auth{
			Database: db,
			Username: user,
			Password: pswd,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Debug: false,
	})

	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)

	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}))
	if err := conn.PingContext(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, nil, err
	}

	cleanup := func() {
		if errClose := conn.Close(); errClose != nil {
			log.Error(errClose)
		}

		log.Info("ClickHouse cleanup")
	}

	if err := migrations.Migrate(conn); err != nil {
		cleanup()
		panic(err)
	}

	return conn, cleanup, nil
}
