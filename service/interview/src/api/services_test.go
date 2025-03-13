package api

import (
	"context"
	"interview/src/db"
	"interview/src/db/table"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var client InterviewServiceClient
var closer func()
var ctx context.Context = context.Background()

func init() {
	client, closer = initTestServer(context.Background())
}

func initTestServer(ctx context.Context) (InterviewServiceClient, func()) {
	buffer := 10 * 1024 * 1024
	lis := bufconn.Listen(buffer) // no need for port, since bufconn controls in-memory IPC

	grpcServerChan := make(chan *grpc.Server)
	go func(grpcServerChan chan *grpc.Server) {
		var serverOpts []grpc.ServerOption
		grpcServer := grpc.NewServer(serverOpts...)

		// initialize database connection with test server
		dbSession, err := db.Connect("localhost", "9042", ctx)
		table.InitializeTables(dbSession.Session, ctx)

		if err != nil {
			log.Printf("error connecting to database: %v", err)
		}

		interviewServiceServerImpl := &InterviewServiceServerImpl{session: dbSession}

		RegisterInterviewServiceServer(grpcServer, interviewServiceServerImpl)

		// client will block until server is listening
		go func() {
			if err := grpcServer.Serve(lis); err != nil {
				log.Printf("error serving server: %v", err)
			}
		}()

		grpcServerChan <- grpcServer
	}(grpcServerChan)

	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts,
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (net.Conn, error) {
				// override the default TCP net.conn used for remote dialing with the bufconn connection
				conn, err := lis.DialContext(ctx)
				if err != nil {
					log.Printf("error dialing: %v", err)
					return nil, err
				}
				return conn, nil
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	// using passthrough resolver: https://github.com/grpc/grpc-go/blob/v1.64.0/internal/resolver/passthrough/passthrough.go
	// does not map host to an address, just returns the address as-is, but still isn't used by *bufconn.Listener
	passthroughAddress := "passthrough:///ARBITRARYADDRESS:21"
	conn, err := grpc.NewClient(passthroughAddress, dialOpts...)

	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	grpcServer := <-grpcServerChan
	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		grpcServer.Stop()
	}

	client := NewInterviewServiceClient(conn)
	return client, closer
}

func TestCreateInterviewTemplateCall(t *testing.T) {
	defer closer()

	createInterviewTemplate := &CreateInterviewTemplate{
		Company:     "Google",
		Role:        "Systems Engineer",
		Skills:      []string{"Go", "Docker", "Kubernetes"},
		Description: "Systems Engineer at Google requiring 6+ YOE in embedded systems using C/C++",
		Questions:   []string{"What is the difference between a mutex and a semaphore?", "What is the difference between a process and a thread?"},
		UserId:      32,
	}

	resp, err := client.CreateInterviewTemplateCall(ctx, createInterviewTemplate)

	t.Cleanup(func() {

	})
}
