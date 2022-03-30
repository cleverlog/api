package delivery

import (
	"context"

	"github.com/cleverlog/api/internal/logs"
	"github.com/cleverlog/api/internal/logs/domain"
	"github.com/cleverlog/api/internal/tools"
	proto "github.com/cleverlog/api/proto/log"
)

//go:generate protoc --proto_path=proto --go_out=../../../proto/log --go_opt=paths=source_relative --go-grpc_out=../../../proto/log --go-grpc_opt=paths=source_relative log.proto

type SourcesDelivery struct {
	proto.UnimplementedLogServiceServer

	logsUseCase logs.UseCase
}

func New(logUC logs.UseCase) *SourcesDelivery {
	return &SourcesDelivery{
		logsUseCase: logUC,
	}
}

func (d *SourcesDelivery) SendLog(ctx context.Context, log *proto.Log) (*proto.LogStatus, error) {
	if err := d.logsUseCase.Create(ctx, &domain.Log{
		ServiceName: log.Service,
		Level:       log.Level.String(),
		SpanID:      tools.StringToUUID(log.SpanId),
		Timestamp:   log.Timestamp.AsTime(),
		Source:      log.Source,
		Message:     log.Message,
	}); err != nil {
		return nil, err
	}

	return &proto.LogStatus{Success: true}, nil
}

func (d *SourcesDelivery) SendLogs(ctx context.Context, logs *proto.Logs) (*proto.LogStatus, error) {
	var newLogs []*domain.Log

	for _, log := range logs.Logs {
		newLogs = append(newLogs, &domain.Log{
			ServiceName: log.Service,
			Level:       log.Level.String(),
			SpanID:      tools.StringToUUID(log.SpanId),
			Timestamp:   log.Timestamp.AsTime(),
			Source:      log.Source,
			Message:     log.Message,
		})
	}

	if err := d.logsUseCase.Create(ctx, newLogs...); err != nil {
		return nil, err
	}

	return &proto.LogStatus{Success: true}, nil
}
