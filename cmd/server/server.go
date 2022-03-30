package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/cleverlog/api/cmd/builder"
	logsDelivery "github.com/cleverlog/api/internal/logs/delivery"
	logsRepository "github.com/cleverlog/api/internal/logs/repository"
	logsUseCase "github.com/cleverlog/api/internal/logs/useCase"
	proto "github.com/cleverlog/api/proto/log"
)

func main() {
	service, cancel, err := builder.Initialize()
	if err != nil {
		cancel()
		panic(err)
	}
	repository := logsRepository.New(service.ClickHouse)
	useCase := logsUseCase.New(repository)
	delivery := logsDelivery.New(useCase)

	grpcServer := grpc.NewServer(grpc.EmptyServerOption{})

	proto.RegisterLogServiceServer(grpcServer, delivery)

	viper.New()

	viper.AutomaticEnv()

	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "5555")

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s",
		viper.GetString("SERVER_HOST"),
		viper.GetString("SERVER_PORT")))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	if err = grpcServer.Serve(lis); err != nil {
		logrus.Fatal(err)
	}
}
