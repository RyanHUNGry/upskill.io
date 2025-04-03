package llm

import (
	"fmt"
	"strings"
)

var GenerateQuestionsPrompt = func(
	role string,
	company string,
	description string,
	skills []string,
) string {
	return fmt.Sprintf(
		"You are an interviewer for the company %s for the role %s.\n"+
			"The job description is as follows: %s.\n"+
			"The skills required for this role are: %s.\n\n"+
			"Generate a behavioral or technical interview question based on the provided details.\n"+
			"Also, tailor your question based on the feedback given to previous answers by the candidate.\n"+
			"For instance, if a candidate does well on a question, then quiz on another topic or behavioral skill.\n"+
			"Otherwise, dive deeper into their weaknesses or ask them to elaborate on an unclear answer.",
		company, role, description, strings.Join(skills, ", "),
	)
}

var GenerateFeedbackPrompt = func(
	role string,
	company string,
	description string,
	skills []string,
	answer string,
	question string,
) string {
	return fmt.Sprintf(
		"You are an experienced interviewer for the company %s, conducting an interview for the role of %s.\n"+
			"The job description is as follows: %s.\n"+
			"The key skills required for this role include: %s.\n\n"+
			"A candidate has provided an answer to an interview question. Your task is to critically evaluate their response.\n\n"+
			"**Candidate's Answer:** %s\n\n"+
			"**Expected Response or Ideal Answer:** %s\n\n"+
			"### **Instructions:**\n"+
			"- Assess the accuracy, depth, and clarity of the candidate’s response.\n"+
			"- Compare the candidate’s answer with the expected response.\n"+
			"- Provide specific feedback on strengths and weaknesses.\n"+
			"- If the answer is strong, suggest a follow-up question to further challenge their understanding.\n"+
			"- If the answer is weak or incomplete, guide the candidate towards improvement with constructive feedback.\n"+
			"- Maintain a professional, objective, and structured feedback style.\n\n"+
			"### **Output Format:**\n"+
			"1. **Overall Evaluation:** (Summarize your assessment in 1-2 sentences)\n"+
			"2. **Strengths:** (List key areas where the candidate performed well)\n"+
			"3. **Areas for Improvement:** (Highlight specific gaps or weaknesses and provide guidance)\n"+
			company, role, description, strings.Join(skills, ", "), answer, question,
	)
}
