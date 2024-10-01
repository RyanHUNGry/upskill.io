// Driver for the interview service process.
package main

import (
	"context"
	"fmt"
	"net"
	"os"

	kitlog "github.com/go-kit/log"
	kitlogLevel "github.com/go-kit/log/level"
	loggingWrapper "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"ryhung.upskill.io/internal/cassandra"
	api "ryhung.upskill.io/internal/grpc"
)

// InterceptorLogger adapts go-kit logger to interceptor logger.
// this is also where kafka logs are produced
func InterceptorLogger(l kitlog.Logger) loggingWrapper.Logger {
	return loggingWrapper.LoggerFunc(func(_ context.Context, lvl loggingWrapper.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case loggingWrapper.LevelDebug:
			_ = kitlogLevel.Debug(l).Log(largs...)
		case loggingWrapper.LevelInfo:
			_ = kitlogLevel.Info(l).Log(largs...)
		case loggingWrapper.LevelWarn:
			_ = kitlogLevel.Warn(l).Log(largs...)
		case loggingWrapper.LevelError:
			_ = kitlogLevel.Error(l).Log(largs...)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	runlevel := os.Getenv("DEV_RUNLEVEL")

	// TODO: initialize Kafka queue

	// initialize Cassandra database
	dbChannel := make(chan *cassandra.InterviewServiceDatabase)
	if runlevel == "5" {
		go func(ctx context.Context) {
			fmt.Println("Initializing Cassandra database ⌛")
			db := cassandra.InitializeInterviewServiceDatabase(ctx)
			fmt.Println("Cassandra database initialized")
			dbChannel <- db
			defer db.CloseInterviewServiceDatabase()
		}(ctx)
	}

	// initialize gRPC API server
	if runlevel == "5" || runlevel == "1" {
		go func() {
			logger := kitlog.NewSyncLogger(kitlog.NewJSONLogger(os.Stdout)) // because each gRPC request will be handled in a separate goroutine, logger must be thread-safe so ordering is preserved
			logger = kitlog.With(logger, "service", "InterviewService")

			opts := []loggingWrapper.Option{
				loggingWrapper.WithLogOnEvents(loggingWrapper.StartCall, loggingWrapper.FinishCall),
				loggingWrapper.WithLevels(func(code codes.Code) loggingWrapper.Level {
					switch code {
					case codes.OK:
						return loggingWrapper.LevelInfo
					default:
						return loggingWrapper.LevelError
					}
				}),
				// Add any other option (check functions starting with loggingWrapper.With).
			}

			fmt.Println("Initializing gRPC API server ⌛")
			lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9999))
			if err != nil {
				panic(err)
			}

			var serverOpts []grpc.ServerOption = []grpc.ServerOption{
				grpc.ChainUnaryInterceptor(loggingWrapper.UnaryServerInterceptor(InterceptorLogger(logger), opts...)),
				grpc.ChainStreamInterceptor(loggingWrapper.StreamServerInterceptor(InterceptorLogger(logger), opts...)),
			}

			grpcServer := grpc.NewServer(serverOpts...)
			var db *cassandra.InterviewServiceDatabase // nil if not set
			if runlevel == "5" {
				db = <-dbChannel
			}
			api.RegisterInterviewServiceServer(grpcServer, api.CreateInterviewServiceServer(ctx, db))
			fmt.Println("gRPC API server initialized ✅")
			grpcServer.Serve(lis) // each gRPC request will be handled in a separate goroutine
			defer grpcServer.GracefulStop()
		}()
	}

	// don't stop process
	select {}
}
