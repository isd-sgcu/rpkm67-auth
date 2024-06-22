package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/isd-sgcu/rpkm67-auth/config"
	"github.com/isd-sgcu/rpkm67-auth/database"
	user "github.com/isd-sgcu/rpkm67-auth/internal/user"
	"github.com/isd-sgcu/rpkm67-auth/logger"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	logger := logger.New(conf)

	db, err := database.InitDatabase(&conf.Db, conf.App.IsDevelopment())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	userRepo := user.NewRepository(db)
	userUtils := user.NewUserUtils()
	userSvc := user.NewService(userRepo, userUtils, logger)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.App.Port))
	if err != nil {
		panic(fmt.Sprintf("Failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	userProto.RegisterUserServiceServer(grpcServer, userSvc)

	reflection.Register(grpcServer)
	go func() {
		logger.Sugar().Infof("RPKM67 Auth starting at port %v", conf.App.Port)

		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatal("Failed to start RPKM67 Auth service", zap.Error(err))
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, logger, map[string]operation{
		"server": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"database": func(ctx context.Context) error {
			sqlDB, err := db.DB()
			if err != nil {
				return nil
			}
			return sqlDB.Close()
		},
		// "cache": func(ctx context.Context) error {
		// 	return cacheDb.Close()
		// },
	})

	<-wait

	grpcServer.GracefulStop()
	logger.Info("Closing the listener")
	listener.Close()
	logger.Info("RPKM67 Auth service has been shutdown gracefully")
}
