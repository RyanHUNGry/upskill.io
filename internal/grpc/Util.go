package api

import "errors"

func answerAggregator(answerResponsePairs []AnswerResponsePair) (*GetAnswerScores, error) {
	if len(answerResponsePairs) == 0 {
		return nil, errors.New("No answer response pairs in slice")
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

		var averageScore float64;
		for key, value := range answerResponse {
			averageScore += value["score"]
		}
		averageScore /= float64(len(answerResponse))

		getAnswerScores_AnswerScore := &GetAnswerScores_AnswerScore{
			QuestionIdx: answerRequest.QuestionIdx,
			Question:    answerRequest.Question,
			Answer:      answerRequest.Answer,
			AnswerIdx: i,

		}}

		answerResponse[]
	}

	return nil, nil
}
