package api

import (
	"context"
	"fmt"
	"interview/src/db"
	"interview/src/db/table"
	"log"
	"net"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var client InterviewServiceClient
var grpcCloser func()
var ctx context.Context = context.Background()
var database *db.Database

func init() {
	client, grpcCloser, database = initTestServer(context.Background())
}

func initTestServer(ctx context.Context) (InterviewServiceClient, func(), *db.Database) {
	listenerBufferSize := 10 * 1024 * 1024 // 10 MB
	lis := bufconn.Listen(listenerBufferSize)

	grpcServerChan := make(chan *grpc.Server)
	databaseSessionChan := make(chan *db.Database)

	go func() {
		var serverOpts []grpc.ServerOption = []grpc.ServerOption{}
		grpcServer := grpc.NewServer(serverOpts...)

		dbSession, err := db.Connect("localhost", "9042", ctx)
		table.InitializeTables(dbSession.Session, ctx)

		if err != nil {
			log.Printf("error connecting to database: %v", err)
		}

		interviewServiceServerImpl := &InterviewServiceServerImpl{Database: dbSession}

		RegisterInterviewServiceServer(grpcServer, interviewServiceServerImpl)

		// Client will block until server is listening
		go func() {
			if err := grpcServer.Serve(lis); err != nil {
				log.Printf("error serving server: %v", err)
			}
		}()

		grpcServerChan <- grpcServer
		databaseSessionChan <- dbSession
	}()

	var dialOpts []grpc.DialOption
	dialOpts = append(dialOpts,
		grpc.WithContextDialer(
			func(ctx context.Context, _ string) (net.Conn, error) {
				// Override the default TCP net.conn used for remote dialing with an in-memory full-duplex connection
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

	// Prefix the address to specify passthrough resolver, which passes the address as-is to the dialer without DNS resolution so that dialer can use the in-memory connection
	passthroughAddress := "passthrough:///ARBITRARYADDRESS:21"
	conn, err := grpc.NewClient(passthroughAddress, dialOpts...)

	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	grpcServer := <-grpcServerChan
	databaseSession := <-databaseSessionChan
	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		grpcServer.Stop()
	}

	client := NewInterviewServiceClient(conn)
	return client, closer, databaseSession
}

func TestCreateInterviewTemplateCall(t *testing.T) {
	createInterviewTemplate := &CreateInterviewTemplate{
		Company:     "Google",
		Role:        "Systems Engineer",
		Skills:      []string{"Go", "Docker", "Kubernetes"},
		Description: "Systems Engineer at Google requiring 6+ YOE in embedded systems using C/C++",
		Questions:   []string{"What is the difference between a mutex and a semaphore?", "What is the difference between a process and a thread?"},
		UserId:      32,
	}

	resp, err := client.CreateInterviewTemplateCall(ctx, createInterviewTemplate)

	if err != nil {
		t.Errorf("error creating interview template: %v", err)
	}

	if resp.Company != createInterviewTemplate.Company {
		t.Errorf("expected company %s, got %s", createInterviewTemplate.Company, resp.Company)
	}

	if resp.Role != createInterviewTemplate.Role {
		t.Errorf("expected role %s, got %s", createInterviewTemplate.Role, resp.Role)
	}

	if resp.Description != createInterviewTemplate.Description {
		t.Errorf("expected description %s, got %s", createInterviewTemplate.Description, resp.Description)
	}

	if resp.UserId != createInterviewTemplate.UserId {
		t.Errorf("expected userId %d, got %d", createInterviewTemplate.UserId, resp.UserId)
	}

	if resp.AverageRating != -1 || resp.AverageScore != -1 || resp.AmountConducted != 0 {
		t.Errorf("expected average rating, score to be -1 and amount conducted to be 0, got %d, %d, %d, respectively", resp.AverageRating, resp.AverageScore, resp.AmountConducted)
	}
}

func TestCreateConductedInterviewCall(t *testing.T) {
	createInterviewTemplate := &CreateInterviewTemplate{
		Company:     "Palantir",
		Role:        "Forward Deployed SWE",
		Skills:      []string{"Go", "Terraform", "spiceDB", "Apache Kafka"},
		Description: "Forward Deployed SWE at Palantir requiring 4+ YOE in distributed systems using Go",
		Questions:   []string{"What are distributed transactions?", "What is the CAP theorem?"},
		UserId:      19,
	}

	interviewTemplate, err := client.CreateInterviewTemplateCall(ctx, createInterviewTemplate)
	interviewTemplateId := interviewTemplate.InterviewTemplateId

	if err != nil {
		t.Errorf("error creating interview template: %v", err)
	}

	createConductedInterview := &CreateConductedInterview{
		InterviewTemplateId: interviewTemplateId,
		UserId:              19,
		Score:               86,
		Rating:              4,
		Role:                interviewTemplate.Role,
		Responses: &ResponseType{
			Answers:   []string{"Distributed transactions are a way to ensure consistency across distributed systems.", "CAP theorem states that a distributed system can only guarantee two of the following three properties: Consistency, Availability, and Partition Tolerance."},
			Feedback:  []string{"The answer is correct but could use more detail.", "Palantir requires a deeper understanding of CAP theorem, so please explain why only two properties can be guaranteed."},
			Questions: interviewTemplate.Questions,
		},
	}

	conductedInterview, err := client.CreateConductedInterviewCall(ctx, createConductedInterview)

	if err != nil {
		t.Errorf("error creating conducted interview: %v", err)
	}

	for i := range conductedInterview.InterviewTemplateId {
		if conductedInterview.InterviewTemplateId[i] != interviewTemplateId[i] {
			t.Errorf("expected interview template id %v, got %v", interviewTemplateId, conductedInterview.InterviewTemplateId)
		}
	}

	if conductedInterview.UserId != createConductedInterview.UserId {
		t.Errorf("expected userId %d, got %d", createConductedInterview.UserId, conductedInterview.UserId)
	}

	if conductedInterview.Score != createConductedInterview.Score {
		t.Errorf("expected score %d, got %d", createConductedInterview.Score, conductedInterview.Score)
	}

	if conductedInterview.Rating != createConductedInterview.Rating {
		t.Errorf("expected rating %d, got %d", createConductedInterview.Rating, conductedInterview.Rating)
	}

	if conductedInterview.Role != createConductedInterview.Role {
		t.Errorf("expected role %s, got %s", createConductedInterview.Role, conductedInterview.Role)
	}

	if len(conductedInterview.Responses.Feedback) != len(createConductedInterview.Responses.Feedback) {
		t.Errorf("expected feedback length %d, got %d", len(createConductedInterview.Responses.Feedback), len(conductedInterview.Responses.Feedback))
	}

	if len(conductedInterview.Responses.Answers) != len(createConductedInterview.Responses.Answers) {
		t.Errorf("expected responses length %d, got %d", len(createConductedInterview.Responses.Answers), len(conductedInterview.Responses.Answers))
	}

	if len(conductedInterview.Responses.Questions) != len(createConductedInterview.Responses.Questions) {
		t.Errorf("expected questions length %d, got %d", len(createConductedInterview.Responses.Questions), len(conductedInterview.Responses.Questions))
	}

	for j, feedback := range createConductedInterview.Responses.Feedback {
		if conductedInterview.Responses.Feedback[j] != feedback {
			t.Errorf("expected feedback %s, got %s at index %d", feedback, conductedInterview.Responses.Feedback[j], j)
		}
	}

	for j, answer := range createConductedInterview.Responses.Answers {
		if conductedInterview.Responses.Answers[j] != answer {
			t.Errorf("expected answer %s, got %s at index %d", answer, conductedInterview.Responses.Answers[j], j)
		}
	}

	for j, question := range createConductedInterview.Responses.Questions {
		if conductedInterview.Responses.Questions[j] != question {
			t.Errorf("expected question %s, got %s at index %d", question, conductedInterview.Responses.Questions[j], j)
		}
	}
}

func TestMain(m *testing.M) {
	run := func() int {
		defer grpcCloser()
		defer database.Session.Close()
		defer table.DropAllTables(database.Session, context.Background()) // Note defer statements are LIFO order
		fmt.Println("Running tests through TestMain() entrypoint")
		statusCode := m.Run()
		return statusCode
	}

	os.Exit(run())
}
