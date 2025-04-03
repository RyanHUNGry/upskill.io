package api

import (
	context "context"
	"fmt"
	"interview/src/db"
	"interview/src/llm"
	"interview/src/utils"
	"io"

	"github.com/gocql/gocql"
	"github.com/openai/openai-go/responses"
)

type InterviewServiceServerImpl struct {
	UnimplementedInterviewServiceServer
	Database *db.Database
	Model    *llm.Model
}

func (service InterviewServiceServerImpl) CreateInterviewTemplateCall(ctx context.Context, in *CreateInterviewTemplate) (*InterviewTemplate, error) {
	interviewTemplateId, err := service.Database.CreateInterviewTemplate(in.Company, in.Role, in.Skills, in.Description, in.Questions, in.UserId)
	if err != nil {
		return nil, err
	}

	err = service.Database.InsertUserIdAndInterviewTemplateId(in.UserId, interviewTemplateId)
	if err != nil {
		return nil, err
	}

	err = service.Database.InsertRoleAndCompanyInterviewTemplateId(in.Role, in.Company, interviewTemplateId)
	if err != nil {
		return nil, err
	}

	interviewTemplate, _, err := service.Database.FindInterviewTemplateById(interviewTemplateId)
	if err != nil {
		return nil, err
	}

	return ConvertInterviewTemplateToProto(interviewTemplate), nil
}

func (service InterviewServiceServerImpl) CreateConductedInterviewCall(ctx context.Context, in *CreateConductedInterview) (*ConductedInterview, error) {
	respType := db.ResponseType{Feedback: in.Responses.Feedback, Answers: in.Responses.Answers, Questions: in.Responses.Questions}

	conductedInterviewId, err := service.Database.CreateConductedInterview(in.InterviewTemplateId, in.UserId, in.Score, in.Rating, in.Role, in.Company, respType)
	if err != nil {
		return nil, err
	}

	averageScore, averageRating, amountConducted, err := service.Database.GetInterviewTemplateStats(in.InterviewTemplateId)
	if err != nil {
		return nil, err
	}

	newScore, newRating, newAmountConducted, err := service.Database.UpdateInterviewTemplateStats(in.InterviewTemplateId, averageScore, averageRating, amountConducted, in.Score, in.Rating)
	if err != nil {
		return nil, err
	}

	err = service.Database.UpdateRoleAndCompanyInterviewTemplateId(in.Role, in.Company, gocql.UUID(in.InterviewTemplateId), newScore, newRating, newAmountConducted)

	err = service.Database.InsertUserIdAndConductedInterviewId(in.UserId, conductedInterviewId)
	if err != nil {
		return nil, err
	}

	conductedInterview, _, err := service.Database.FindConductedInterviewById(conductedInterviewId)
	if err != nil {
		return nil, err
	}

	return ConvertConductedInterviewToProto(conductedInterview), nil
}

func (service InterviewServiceServerImpl) GetConductedInterviewsByUserCall(ctx context.Context, in *GetConductedInterviewsByUser) (*ConductedInterviews, error) {
	conductedInterviewIds, err := service.Database.FindConductedInterviewIdsByUserId(in.UserId)
	if err != nil {
		return nil, err
	}

	_, conductedInterviews, err := service.Database.FindConductedInterviewById(conductedInterviewIds)

	if err != nil {
		return nil, err
	}

	return &ConductedInterviews{ConductedInterviews: utils.FunctionMap(conductedInterviews, ConvertConductedInterviewToProto)}, nil
}

func (service InterviewServiceServerImpl) GetInterviewTemplatesByUserCall(ctx context.Context, in *GetInterviewTemplatesByUser) (*InterviewTemplates, error) {
	interviewTemplateIds, err := service.Database.FindInterviewTemplateIdsByUserId(in.UserId)
	if err != nil {
		return nil, err
	}

	_, interviewTemplates, err := service.Database.FindInterviewTemplateById(interviewTemplateIds)

	if err != nil {
		return nil, err
	}

	return &InterviewTemplates{InterviewTemplates: utils.FunctionMap(interviewTemplates, ConvertInterviewTemplateToProto)}, nil
}

func (service InterviewServiceServerImpl) ConductInterviewCall(stream InterviewService_ConductInterviewCallServer) error {
	finalConductInterviewResponse := &ConductInterviewResponse{}
	var interviewTemplate *db.InterviewTemplate
	var numQuestionsToGenerate *int32
	var userId int32
	var rating int32

	questionIdx := 0
	numQuestionsGenerated := 0
	questions := make([]string, 0)
	answers := make([]string, 0)
	feedback := make([]string, 0)
	var prevResponse *responses.Response

	for {
		conductInterviewRequest, err := stream.Recv()
		// The io.EOF error is returned when the client closes the stream, which programmatically happens when FinalRequest field is populated
		if err == io.EOF {
			fmt.Println(userId, rating)
			if err := stream.Send(finalConductInterviewResponse); err != nil {
				return err
			}
			return nil
		}

		/*
			1. User sends an initial request to provide contextual information
				- Generates an initial question
			2. User sends an answer, which answers the previous question. This is then scored
				- A second question, based on the feedback and the answer, is then generated
			3. Once user is done answering all questions, they send a final request to signal the end of the interview
				- Followed by a cancellation signal, which closes the stream

			Any predefined questions must be manually embedded in the context window of the model
			Otherwise, the previous model response ID is tracked and used for the context window
		*/
		if conductInterviewRequest.FinalRequest != nil {
			rating = conductInterviewRequest.FinalRequest.Rating
		} else if conductInterviewRequest.InitialRequest != nil {
			numQuestionsToGenerate = conductInterviewRequest.InitialRequest.NumQuestionsToGenerate
			userId = conductInterviewRequest.InitialRequest.UserId
			interviewTemplate, _, err = service.Database.FindInterviewTemplateById(conductInterviewRequest.InitialRequest.InterviewTemplateId)
			if err != nil {
				return err
			}

			// Provide the initial question, and expect an initial answer in the following client request
			var question string
			if questionIdx < len(interviewTemplate.Questions) {
				question = interviewTemplate.Questions[questionIdx]
				questionIdx += 1
			} else if numQuestionsGenerated < int(*numQuestionsToGenerate) {
				res, err := service.Model.GenerateQuestion(interviewTemplate.Role, interviewTemplate.Company, interviewTemplate.Description, interviewTemplate.Skills, nil)
				if err != nil {
					return err
				}

				prevResponse = res
				question = res.Output[0].Content[0].Text
				numQuestionsGenerated += 1
			}

			questions = append(questions, question)
			if err := stream.Send(&ConductInterviewResponse{Question: question}); err != nil {
				return err
			}
		} else {
			answer := *conductInterviewRequest.Answer
			answers = append(answers, answer)

			// After answering, generate feedback for the previous question
			prevQuestion := questions[len(questions)-1]
			res, err := service.Model.GenerateFeedback(interviewTemplate.Role, interviewTemplate.Company, interviewTemplate.Description, interviewTemplate.Skills, prevQuestion, answer, prevResponse)
			if err != nil {
				return err
			}

			fb := res.Output[0].Content[0].Text
			feedback = append(feedback, fb)
			fmt.Println(fb)

			// Generate the new question based on the feedback and the answer
			var question string
			if questionIdx < len(interviewTemplate.Questions) {
				question = interviewTemplate.Questions[questionIdx]
				questionIdx += 1
			} else if numQuestionsGenerated < int(*numQuestionsToGenerate) {
				res, err := service.Model.GenerateQuestion(interviewTemplate.Role, interviewTemplate.Company, interviewTemplate.Description, interviewTemplate.Skills, prevResponse)
				if err != nil {
					return err
				}

				question = res.Output[0].Content[0].Text
				numQuestionsGenerated += 1
				questions = append(questions, question)
			}

			if err := stream.Send(&ConductInterviewResponse{Question: question}); err != nil {
				return err
			}
		}
	}
}
