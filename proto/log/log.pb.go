// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: log.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Logs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*Log `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *Logs) Reset() {
	*x = Logs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logs) ProtoMessage() {}

func (x *Logs) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logs.ProtoReflect.Descriptor instead.
func (*Logs) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{0}
}

func (x *Logs) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service   string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	SpanId    string                 `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Source    string                 `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Message   string                 `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Log) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *Log) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Log) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Log) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type LogStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LogStatus) Reset() {
	*x = LogStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogStatus) ProtoMessage() {}

func (x *LogStatus) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogStatus.ProtoReflect.Descriptor instead.
func (*LogStatus) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{2}
}

func (x *LogStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_log_proto protoreflect.FileDescriptor

var file_log_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x04,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0xa4,
	0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x48, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x53, 0x65,
	0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x0a, 0x2e, 0x4c, 0x6f,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4c,
	0x6f, 0x67, 0x73, 0x12, 0x05, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x1a, 0x0a, 0x2e, 0x4c, 0x6f, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x79, 0x63, 0x6e, 0x69, 0x63, 0x6b, 0x2f, 0x63, 0x6c, 0x65,
	0x76, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_proto_rawDescOnce sync.Once
	file_log_proto_rawDescData = file_log_proto_rawDesc
)

func file_log_proto_rawDescGZIP() []byte {
	file_log_proto_rawDescOnce.Do(func() {
		file_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_proto_rawDescData)
	})
	return file_log_proto_rawDescData
}

var file_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_log_proto_goTypes = []interface{}{
	(*Logs)(nil),                  // 0: Logs
	(*Log)(nil),                   // 1: Log
	(*LogStatus)(nil),             // 2: LogStatus
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_log_proto_depIdxs = []int32{
	1, // 0: Logs.logs:type_name -> Log
	3, // 1: Log.timestamp:type_name -> google.protobuf.Timestamp
	1, // 2: LogService.SendLog:input_type -> Log
	0, // 3: LogService.SendLogs:input_type -> Logs
	2, // 4: LogService.SendLog:output_type -> LogStatus
	2, // 5: LogService.SendLogs:output_type -> LogStatus
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_log_proto_init() }
func file_log_proto_init() {
	if File_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logs); i {
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
		file_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogStatus); i {
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
			RawDescriptor: file_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_proto_goTypes,
		DependencyIndexes: file_log_proto_depIdxs,
		MessageInfos:      file_log_proto_msgTypes,
	}.Build()
	File_log_proto = out.File
	file_log_proto_rawDesc = nil
	file_log_proto_goTypes = nil
	file_log_proto_depIdxs = nil
}
