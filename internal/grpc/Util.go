package api

import (
	"errors"
	"strings"
)

func answerAggregator(answerResponsePairs []AnswerResponsePair) (*GetAnswerScores, error) {
	if len(answerResponsePairs) == 0 {
		return nil, errors.New("no answer response pairs in slice")
	}

	getAnswerScores := new(GetAnswerScores)
	getAnswerScores.InterviewId = answerResponsePairs[0].answerRequest.InterviewId
	getAnswerScores.UserId = answerResponsePairs[0].answerRequest.UserId
	getAnswerScores.SessionId = answerResponsePairs[0].answerRequest.SessionId

	answerScores := make([]*GetAnswerScores_AnswerScore, 0, 10)
	// loop through slice to create *GetAnswerScores_AnswerScore
	for i := 0; i < len(answerResponsePairs); i++ {
		answerResponsePair := answerResponsePairs[i]
		answerResponse := answerResponsePair.answerResponse
		answerRequest := answerResponsePair.answerRequest

		var averageScore float64
		for _, value := range answerResponse {
			if score, ok := value["score"].(float64); ok {
				averageScore += score
			} else {
				return nil, errors.New("score is not a float64")
			}
		}
		averageScore /= float64(len(answerResponse))

		getAnswerScores_AnswerScore := &GetAnswerScores_AnswerScore{
			QuestionIdx:  answerRequest.QuestionIdx,
			Question:     answerRequest.Question,
			Answer:       answerRequest.Answer,
			AnswerIdx:    int32(i),
			AverageScore: averageScore,
		}

		scoreBreakdowns := make([]*GetAnswerScores_AnswerScore_ScoreBreakdown, 0, 10)
		for key, value := range answerResponse {
			criteria, err := stringToCriteria(key)
			if err != nil {
				return nil, err
			}

			score, ok := value["score"].(float64)
			if !ok {
				return nil, errors.New("score is not a float64")
			}

			feedback, ok := value["reason"].(string)
			if !ok {
				return nil, errors.New("reason is not a string")
			}

			scoreBreakdown := &GetAnswerScores_AnswerScore_ScoreBreakdown{
				Criteria: criteria,
				Score:    score,
				Feedback: feedback,
			}
			scoreBreakdowns = append(scoreBreakdowns, scoreBreakdown)
		}

		getAnswerScores_AnswerScore.ScoreBreakdowns = scoreBreakdowns
		answerScores = append(answerScores, getAnswerScores_AnswerScore)
	}

	getAnswerScores.AnswerScores = answerScores

	return getAnswerScores, nil
}

func stringToCriteria(key string) (Criteria, error) {
	key = strings.ToUpper(key)

	switch key {
	case "PROFESSIONALISM":
		return Criteria_PROFESSIONALISM, nil
	case "SKILL_EXPRESSION":
		return Criteria_SKILL_EXPRESSION, nil
	case "TAILORING_TO_COMPANY":
		return Criteria_TAILORING_TO_COMPANY, nil
	default:
		return 0, errors.New("invalid criteria")
	}
}
