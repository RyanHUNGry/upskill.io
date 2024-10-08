// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: enums.proto

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

type Criteria int32

const (
	Criteria_PROFESSIONALISM      Criteria = 0
	Criteria_SKILL_EXPRESSION     Criteria = 1
	Criteria_TAILORING_TO_COMPANY Criteria = 2
)

// Enum value maps for Criteria.
var (
	Criteria_name = map[int32]string{
		0: "PROFESSIONALISM",
		1: "SKILL_EXPRESSION",
		2: "TAILORING_TO_COMPANY",
	}
	Criteria_value = map[string]int32{
		"PROFESSIONALISM":      0,
		"SKILL_EXPRESSION":     1,
		"TAILORING_TO_COMPANY": 2,
	}
)

func (x Criteria) Enum() *Criteria {
	p := new(Criteria)
	*p = x
	return p
}

func (x Criteria) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Criteria) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[0].Descriptor()
}

func (Criteria) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[0]
}

func (x Criteria) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Criteria.Descriptor instead.
func (Criteria) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

var File_enums_proto protoreflect.FileDescriptor

var file_enums_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x2a, 0x4f, 0x0a, 0x08, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x13,
	0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x46, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x49, 0x53,
	0x4d, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x45, 0x58, 0x50,
	0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x41, 0x49,
	0x4c, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e,
	0x59, 0x10, 0x02, 0x42, 0x20, 0x5a, 0x1e, 0x72, 0x79, 0x68, 0x75, 0x6e, 0x67, 0x2e, 0x75, 0x70,
	0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_proto_rawDescOnce sync.Once
	file_enums_proto_rawDescData = file_enums_proto_rawDesc
)

func file_enums_proto_rawDescGZIP() []byte {
	file_enums_proto_rawDescOnce.Do(func() {
		file_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_proto_rawDescData)
	})
	return file_enums_proto_rawDescData
}

var file_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_proto_goTypes = []any{
	(Criteria)(0), // 0: api.Criteria
}
var file_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_proto_init() }
func file_enums_proto_init() {
	if File_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_proto_goTypes,
		DependencyIndexes: file_enums_proto_depIdxs,
		EnumInfos:         file_enums_proto_enumTypes,
	}.Build()
	File_enums_proto = out.File
	file_enums_proto_rawDesc = nil
	file_enums_proto_goTypes = nil
	file_enums_proto_depIdxs = nil
}
