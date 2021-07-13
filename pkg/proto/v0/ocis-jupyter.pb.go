// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.3
// source: ocis-jupyter.proto

package proto

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type JupyterNotebookJSON struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JSONString string `protobuf:"bytes,1,opt,name=JSONString,proto3" json:"JSONString,omitempty"`
}

func (x *JupyterNotebookJSON) Reset() {
	*x = JupyterNotebookJSON{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_jupyter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JupyterNotebookJSON) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JupyterNotebookJSON) ProtoMessage() {}

func (x *JupyterNotebookJSON) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_jupyter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JupyterNotebookJSON.ProtoReflect.Descriptor instead.
func (*JupyterNotebookJSON) Descriptor() ([]byte, []int) {
	return file_ocis_jupyter_proto_rawDescGZIP(), []int{0}
}

func (x *JupyterNotebookJSON) GetJSONString() string {
	if x != nil {
		return x.JSONString
	}
	return ""
}

type JupyterNotebookHTML struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HTMLString string `protobuf:"bytes,1,opt,name=HTMLString,proto3" json:"HTMLString,omitempty"`
	Err        string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *JupyterNotebookHTML) Reset() {
	*x = JupyterNotebookHTML{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocis_jupyter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JupyterNotebookHTML) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JupyterNotebookHTML) ProtoMessage() {}

func (x *JupyterNotebookHTML) ProtoReflect() protoreflect.Message {
	mi := &file_ocis_jupyter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JupyterNotebookHTML.ProtoReflect.Descriptor instead.
func (*JupyterNotebookHTML) Descriptor() ([]byte, []int) {
	return file_ocis_jupyter_proto_rawDescGZIP(), []int{1}
}

func (x *JupyterNotebookHTML) GetHTMLString() string {
	if x != nil {
		return x.HTMLString
	}
	return ""
}

func (x *JupyterNotebookHTML) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_ocis_jupyter_proto protoreflect.FileDescriptor

var file_ocis_jupyter_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x63, 0x69, 0x73, 0x2d, 0x6a, 0x75, 0x70, 0x79, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x13, 0x4a, 0x75, 0x70, 0x79, 0x74,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1e,
	0x0a, 0x0a, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x47,
	0x0a, 0x13, 0x4a, 0x75, 0x70, 0x79, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x48, 0x54, 0x4d, 0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x54, 0x4d, 0x4c, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x48, 0x54, 0x4d, 0x4c, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x7c, 0x0a, 0x16, 0x4a, 0x75, 0x70, 0x79, 0x74,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x62, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x48, 0x54, 0x4d,
	0x4c, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x75, 0x70, 0x79, 0x74, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x4a, 0x53, 0x4f, 0x4e, 0x1a, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x75, 0x70, 0x79, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x74,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x48, 0x54, 0x4d, 0x4c, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x5e, 0x5a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x30, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x92, 0x41, 0x47, 0x12, 0x1d,
	0x0a, 0x16, 0x4a, 0x75, 0x70, 0x79, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65, 0x62, 0x6f, 0x6f,
	0x6b, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01,
	0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ocis_jupyter_proto_rawDescOnce sync.Once
	file_ocis_jupyter_proto_rawDescData = file_ocis_jupyter_proto_rawDesc
)

func file_ocis_jupyter_proto_rawDescGZIP() []byte {
	file_ocis_jupyter_proto_rawDescOnce.Do(func() {
		file_ocis_jupyter_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocis_jupyter_proto_rawDescData)
	})
	return file_ocis_jupyter_proto_rawDescData
}

var file_ocis_jupyter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ocis_jupyter_proto_goTypes = []interface{}{
	(*JupyterNotebookJSON)(nil), // 0: proto.JupyterNotebookJSON
	(*JupyterNotebookHTML)(nil), // 1: proto.JupyterNotebookHTML
}
var file_ocis_jupyter_proto_depIdxs = []int32{
	0, // 0: proto.JupyterNotebookSupport.GenerateHTML:input_type -> proto.JupyterNotebookJSON
	1, // 1: proto.JupyterNotebookSupport.GenerateHTML:output_type -> proto.JupyterNotebookHTML
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ocis_jupyter_proto_init() }
func file_ocis_jupyter_proto_init() {
	if File_ocis_jupyter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocis_jupyter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JupyterNotebookJSON); i {
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
		file_ocis_jupyter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JupyterNotebookHTML); i {
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
			RawDescriptor: file_ocis_jupyter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ocis_jupyter_proto_goTypes,
		DependencyIndexes: file_ocis_jupyter_proto_depIdxs,
		MessageInfos:      file_ocis_jupyter_proto_msgTypes,
	}.Build()
	File_ocis_jupyter_proto = out.File
	file_ocis_jupyter_proto_rawDesc = nil
	file_ocis_jupyter_proto_goTypes = nil
	file_ocis_jupyter_proto_depIdxs = nil
}
