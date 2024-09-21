// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message for creating an answer in an interview session
type CreateAnswerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer      string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	SessionId   []byte `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	InterviewId []byte `protobuf:"bytes,3,opt,name=interview_id,json=interviewId,proto3" json:"interview_id,omitempty"`
	UserId      []byte `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	QuestionIdx int32  `protobuf:"varint,5,opt,name=question_idx,json=questionIdx,proto3" json:"question_idx,omitempty"`
}

func (x *CreateAnswerRequest) Reset() {
	*x = CreateAnswerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnswerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnswerRequest) ProtoMessage() {}

func (x *CreateAnswerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnswerRequest.ProtoReflect.Descriptor instead.
func (*CreateAnswerRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnswerRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *CreateAnswerRequest) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *CreateAnswerRequest) GetInterviewId() []byte {
	if x != nil {
		return x.InterviewId
	}
	return nil
}

func (x *CreateAnswerRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *CreateAnswerRequest) GetQuestionIdx() int32 {
	if x != nil {
		return x.QuestionIdx
	}
	return 0
}

// Message for retrieving answer scores
type GetAnswerScores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       []byte                         `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	InterviewId  []byte                         `protobuf:"bytes,2,opt,name=interview_id,json=interviewId,proto3" json:"interview_id,omitempty"`
	AnswerScores []*GetAnswerScores_AnswerScore `protobuf:"bytes,3,rep,name=answer_scores,json=answerScores,proto3" json:"answer_scores,omitempty"`
}

func (x *GetAnswerScores) Reset() {
	*x = GetAnswerScores{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnswerScores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnswerScores) ProtoMessage() {}

func (x *GetAnswerScores) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnswerScores.ProtoReflect.Descriptor instead.
func (*GetAnswerScores) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetAnswerScores) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *GetAnswerScores) GetInterviewId() []byte {
	if x != nil {
		return x.InterviewId
	}
	return nil
}

func (x *GetAnswerScores) GetAnswerScores() []*GetAnswerScores_AnswerScore {
	if x != nil {
		return x.AnswerScores
	}
	return nil
}

// Request to create an interview with a company and questions
type CreateInterviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      []byte                             `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CompanyName string                             `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Questions   []*CreateInterviewRequest_Question `protobuf:"bytes,3,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *CreateInterviewRequest) Reset() {
	*x = CreateInterviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInterviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInterviewRequest) ProtoMessage() {}

func (x *CreateInterviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInterviewRequest.ProtoReflect.Descriptor instead.
func (*CreateInterviewRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateInterviewRequest) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *CreateInterviewRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CreateInterviewRequest) GetQuestions() []*CreateInterviewRequest_Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

// Response message for retrieving interview details
type GetInterview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterviewId []byte                             `protobuf:"bytes,1,opt,name=interview_id,json=interviewId,proto3" json:"interview_id,omitempty"`
	CompanyName string                             `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Questions   []*CreateInterviewRequest_Question `protobuf:"bytes,3,rep,name=questions,proto3" json:"questions,omitempty"`
}

func (x *GetInterview) Reset() {
	*x = GetInterview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInterview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInterview) ProtoMessage() {}

func (x *GetInterview) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInterview.ProtoReflect.Descriptor instead.
func (*GetInterview) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetInterview) GetInterviewId() []byte {
	if x != nil {
		return x.InterviewId
	}
	return nil
}

func (x *GetInterview) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *GetInterview) GetQuestions() []*CreateInterviewRequest_Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

// Inner message representing the score of a single answer
type GetAnswerScores_AnswerScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score       float64 `protobuf:"fixed64,1,opt,name=score,proto3" json:"score,omitempty"`
	Answer      string  `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	AnswerIdx   int32   `protobuf:"varint,3,opt,name=answer_idx,json=answerIdx,proto3" json:"answer_idx,omitempty"`
	QuestionIdx int32   `protobuf:"varint,4,opt,name=question_idx,json=questionIdx,proto3" json:"question_idx,omitempty"`
}

func (x *GetAnswerScores_AnswerScore) Reset() {
	*x = GetAnswerScores_AnswerScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnswerScores_AnswerScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnswerScores_AnswerScore) ProtoMessage() {}

func (x *GetAnswerScores_AnswerScore) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnswerScores_AnswerScore.ProtoReflect.Descriptor instead.
func (*GetAnswerScores_AnswerScore) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAnswerScores_AnswerScore) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *GetAnswerScores_AnswerScore) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *GetAnswerScores_AnswerScore) GetAnswerIdx() int32 {
	if x != nil {
		return x.AnswerIdx
	}
	return 0
}

func (x *GetAnswerScores_AnswerScore) GetQuestionIdx() int32 {
	if x != nil {
		return x.QuestionIdx
	}
	return 0
}

// Inner message representing a single question in an interview
type CreateInterviewRequest_Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question    string `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	QuestionIdx int32  `protobuf:"varint,2,opt,name=question_idx,json=questionIdx,proto3" json:"question_idx,omitempty"`
}

func (x *CreateInterviewRequest_Question) Reset() {
	*x = CreateInterviewRequest_Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInterviewRequest_Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInterviewRequest_Question) ProtoMessage() {}

func (x *CreateInterviewRequest_Question) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInterviewRequest_Question.ProtoReflect.Descriptor instead.
func (*CreateInterviewRequest_Question) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CreateInterviewRequest_Question) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *CreateInterviewRequest_Question) GetQuestionIdx() int32 {
	if x != nil {
		return x.QuestionIdx
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0xab, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x78, 0x22, 0x93,
	0x02, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x45,
	0x0a, 0x0d, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x0c, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x1a, 0x7d, 0x0a, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x49, 0x64,
	0x78, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x78, 0x22, 0xe3, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x49, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x78, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x42, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x9b, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x22,
	0x00, 0x28, 0x01, 0x42, 0x20, 0x5a, 0x1e, 0x72, 0x79, 0x68, 0x75, 0x6e, 0x67, 0x2e, 0x75, 0x70,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []any{
	(*CreateAnswerRequest)(nil),             // 0: api.CreateAnswerRequest
	(*GetAnswerScores)(nil),                 // 1: api.GetAnswerScores
	(*CreateInterviewRequest)(nil),          // 2: api.CreateInterviewRequest
	(*GetInterview)(nil),                    // 3: api.GetInterview
	(*GetAnswerScores_AnswerScore)(nil),     // 4: api.GetAnswerScores.AnswerScore
	(*CreateInterviewRequest_Question)(nil), // 5: api.CreateInterviewRequest.Question
}
var file_api_proto_depIdxs = []int32{
	4, // 0: api.GetAnswerScores.answer_scores:type_name -> api.GetAnswerScores.AnswerScore
	5, // 1: api.CreateInterviewRequest.questions:type_name -> api.CreateInterviewRequest.Question
	5, // 2: api.GetInterview.questions:type_name -> api.CreateInterviewRequest.Question
	2, // 3: api.InterviewService.CreateInterview:input_type -> api.CreateInterviewRequest
	0, // 4: api.InterviewService.CreateAnswer:input_type -> api.CreateAnswerRequest
	3, // 5: api.InterviewService.CreateInterview:output_type -> api.GetInterview
	1, // 6: api.InterviewService.CreateAnswer:output_type -> api.GetAnswerScores
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAnswerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAnswerScores); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateInterviewRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetInterview); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAnswerScores_AnswerScore); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateInterviewRequest_Question); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
