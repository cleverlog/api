package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/cleverlog/api/cmd/builder"
	logsDelivery "github.com/cleverlog/api/internal/logs/delivery"
	"github.com/cleverlog/api/internal/logs/domain"
	logsRepository "github.com/cleverlog/api/internal/logs/repository"
	proto "github.com/cleverlog/api/proto/log"
)

func main() {
	service, cancel, err := builder.Initialize()
	if err != nil {
		cancel()
		panic(err)
	}

	delivery := logsDelivery.New()
	repository := logsRepository.New(service.ClickHouse)

	grpcServer := grpc.NewServer(grpc.EmptyServerOption{})

	proto.RegisterLogServiceServer(grpcServer, delivery)

	if err := repository.Create(context.Background(), &domain.Log{
		ServiceName: "TestService",
		SpanID:      uuid.New(),
		Timestamp:   time.Now(),
		Source:      "main.go",
		Message:     "hello world",
	}); err != nil {
		log.Error(err)
	}
}
