// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pkg/logchain/logchain.proto

package logchain

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

type Loglevel int32

const (
	Loglevel_PANIC    Loglevel = 0
	Loglevel_ALERT    Loglevel = 1
	Loglevel_CRITICAL Loglevel = 2
	Loglevel_ERROR    Loglevel = 3
	Loglevel_WARNING  Loglevel = 4
	Loglevel_NOTICE   Loglevel = 5
	Loglevel_INFO     Loglevel = 6
	Loglevel_DEBUG    Loglevel = 7
	Loglevel_FATAL    Loglevel = 8
	Loglevel_TRACE    Loglevel = 9
)

// Enum value maps for Loglevel.
var (
	Loglevel_name = map[int32]string{
		0: "PANIC",
		1: "ALERT",
		2: "CRITICAL",
		3: "ERROR",
		4: "WARNING",
		5: "NOTICE",
		6: "INFO",
		7: "DEBUG",
		8: "FATAL",
		9: "TRACE",
	}
	Loglevel_value = map[string]int32{
		"PANIC":    0,
		"ALERT":    1,
		"CRITICAL": 2,
		"ERROR":    3,
		"WARNING":  4,
		"NOTICE":   5,
		"INFO":     6,
		"DEBUG":    7,
		"FATAL":    8,
		"TRACE":    9,
	}
)

func (x Loglevel) Enum() *Loglevel {
	p := new(Loglevel)
	*p = x
	return p
}

func (x Loglevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Loglevel) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_logchain_logchain_proto_enumTypes[0].Descriptor()
}

func (Loglevel) Type() protoreflect.EnumType {
	return &file_pkg_logchain_logchain_proto_enumTypes[0]
}

func (x Loglevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Loglevel.Descriptor instead.
func (Loglevel) EnumDescriptor() ([]byte, []int) {
	return file_pkg_logchain_logchain_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName   string    `protobuf:"bytes,1,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	LogLine   []byte    `protobuf:"bytes,2,opt,name=log_line,json=logLine,proto3" json:"log_line,omitempty"`
	Labels    []*Labels `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	LogLevel  string    `protobuf:"bytes,4,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	Timestamp int64     `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_logchain_logchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_logchain_logchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_pkg_logchain_logchain_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *LogRequest) GetLogLine() []byte {
	if x != nil {
		return x.LogLine
	}
	return nil
}

func (x *LogRequest) GetLabels() []*Labels {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *LogRequest) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *LogRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Labels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Labels) Reset() {
	*x = Labels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_logchain_logchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labels) ProtoMessage() {}

func (x *Labels) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_logchain_logchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labels.ProtoReflect.Descriptor instead.
func (*Labels) Descriptor() ([]byte, []int) {
	return file_pkg_logchain_logchain_proto_rawDescGZIP(), []int{1}
}

func (x *Labels) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Labels) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MetricRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricGroup string    `protobuf:"bytes,1,opt,name=metric_group,json=metricGroup,proto3" json:"metric_group,omitempty"`
	MetricName  string    `protobuf:"bytes,2,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	MetricValue float64   `protobuf:"fixed64,3,opt,name=metric_value,json=metricValue,proto3" json:"metric_value,omitempty"`
	Timestamp   int64     `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Labels      []*Labels `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *MetricRequest) Reset() {
	*x = MetricRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_logchain_logchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricRequest) ProtoMessage() {}

func (x *MetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_logchain_logchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricRequest.ProtoReflect.Descriptor instead.
func (*MetricRequest) Descriptor() ([]byte, []int) {
	return file_pkg_logchain_logchain_proto_rawDescGZIP(), []int{2}
}

func (x *MetricRequest) GetMetricGroup() string {
	if x != nil {
		return x.MetricGroup
	}
	return ""
}

func (x *MetricRequest) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *MetricRequest) GetMetricValue() float64 {
	if x != nil {
		return x.MetricValue
	}
	return 0
}

func (x *MetricRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MetricRequest) GetLabels() []*Labels {
	if x != nil {
		return x.Labels
	}
	return nil
}

// The response message containing the greetings
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_logchain_logchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_logchain_logchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pkg_logchain_logchain_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_pkg_logchain_logchain_proto protoreflect.FileDescriptor

var file_pkg_logchain_logchain_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6c,
	0x6f, 0x67, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01,
	0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x4c, 0x69,
	0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x32,
	0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x49, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x7d, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x4c, 0x45, 0x52, 0x54, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49,
	0x43, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0a, 0x0a,
	0x06, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46,
	0x4f, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x07, 0x12, 0x09,
	0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41,
	0x43, 0x45, 0x10, 0x09, 0x32, 0x52, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x1f, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x0b, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x25, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x0e, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_logchain_logchain_proto_rawDescOnce sync.Once
	file_pkg_logchain_logchain_proto_rawDescData = file_pkg_logchain_logchain_proto_rawDesc
)

func file_pkg_logchain_logchain_proto_rawDescGZIP() []byte {
	file_pkg_logchain_logchain_proto_rawDescOnce.Do(func() {
		file_pkg_logchain_logchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_logchain_logchain_proto_rawDescData)
	})
	return file_pkg_logchain_logchain_proto_rawDescData
}

var file_pkg_logchain_logchain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_logchain_logchain_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_logchain_logchain_proto_goTypes = []interface{}{
	(Loglevel)(0),         // 0: Loglevel
	(*LogRequest)(nil),    // 1: LogRequest
	(*Labels)(nil),        // 2: Labels
	(*MetricRequest)(nil), // 3: MetricRequest
	(*Response)(nil),      // 4: Response
}
var file_pkg_logchain_logchain_proto_depIdxs = []int32{
	2, // 0: LogRequest.labels:type_name -> Labels
	2, // 1: MetricRequest.labels:type_name -> Labels
	1, // 2: LogChain.Log:input_type -> LogRequest
	3, // 3: LogChain.Metric:input_type -> MetricRequest
	4, // 4: LogChain.Log:output_type -> Response
	4, // 5: LogChain.Metric:output_type -> Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_logchain_logchain_proto_init() }
func file_pkg_logchain_logchain_proto_init() {
	if File_pkg_logchain_logchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_logchain_logchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_pkg_logchain_logchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Labels); i {
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
		file_pkg_logchain_logchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricRequest); i {
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
		file_pkg_logchain_logchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_pkg_logchain_logchain_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_logchain_logchain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_logchain_logchain_proto_goTypes,
		DependencyIndexes: file_pkg_logchain_logchain_proto_depIdxs,
		EnumInfos:         file_pkg_logchain_logchain_proto_enumTypes,
		MessageInfos:      file_pkg_logchain_logchain_proto_msgTypes,
	}.Build()
	File_pkg_logchain_logchain_proto = out.File
	file_pkg_logchain_logchain_proto_rawDesc = nil
	file_pkg_logchain_logchain_proto_goTypes = nil
	file_pkg_logchain_logchain_proto_depIdxs = nil
}
