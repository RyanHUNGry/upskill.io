// Driver for the interview service process.
package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"ryhung.upskill.io/internal/cassandra"
	api "ryhung.upskill.io/internal/grpc"
)

// this is from the grpc-ecosystem package, and basically acts as a wrapper for third party loggers to integrate with grpc interceptors
func InterceptorLogger(l log.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		largs := append([]any{"msg", msg}, fields...)
		switch lvl {
		case logging.LevelDebug:
			_ = level.Debug(l).Log(largs...)
		case logging.LevelInfo:
			_ = level.Info(l).Log(largs...)
		case logging.LevelWarn:
			_ = level.Warn(l).Log(largs...)
		case logging.LevelError:
			_ = level.Error(l).Log(largs...)
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
			logger := log.NewLogfmtLogger(os.Stdout)

			opts := []logging.Option{
				logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
				// Add any other option (check functions starting with logging.With).
			}

			fmt.Println("Initializing gRPC API server ⌛")
			lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9999))
			if err != nil {
				panic(err)
			}

			var serverOpts []grpc.ServerOption = []grpc.ServerOption{
				grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(InterceptorLogger(logger), opts...)),
				grpc.ChainStreamInterceptor(logging.StreamServerInterceptor(InterceptorLogger(logger), opts...)),
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
