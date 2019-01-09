// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bq_table.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BigQueryMessageOptions struct {
	// Specifies a name of table in BigQuery for the message.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// If true, BigQuery field names will default to a field's JSON name,
	// not its original/proto field name.
	UseJsonNames         bool     `protobuf:"varint,2,opt,name=use_json_names,json=useJsonNames,proto3" json:"use_json_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BigQueryMessageOptions) Reset()         { *m = BigQueryMessageOptions{} }
func (m *BigQueryMessageOptions) String() string { return proto.CompactTextString(m) }
func (*BigQueryMessageOptions) ProtoMessage()    {}
func (*BigQueryMessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_bq_table_b947be292cbdac0c, []int{0}
}
func (m *BigQueryMessageOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigQueryMessageOptions.Unmarshal(m, b)
}
func (m *BigQueryMessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigQueryMessageOptions.Marshal(b, m, deterministic)
}
func (dst *BigQueryMessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigQueryMessageOptions.Merge(dst, src)
}
func (m *BigQueryMessageOptions) XXX_Size() int {
	return xxx_messageInfo_BigQueryMessageOptions.Size(m)
}
func (m *BigQueryMessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_BigQueryMessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_BigQueryMessageOptions proto.InternalMessageInfo

func (m *BigQueryMessageOptions) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *BigQueryMessageOptions) GetUseJsonNames() bool {
	if m != nil {
		return m.UseJsonNames
	}
	return false
}

var E_BigqueryOpts = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*BigQueryMessageOptions)(nil),
	Field:         1021,
	Name:          "gen_bq_schema.bigquery_opts",
	Tag:           "bytes,1021,opt,name=bigquery_opts,json=bigqueryOpts",
	Filename:      "bq_table.proto",
}

func init() {
	proto.RegisterType((*BigQueryMessageOptions)(nil), "gen_bq_schema.BigQueryMessageOptions")
	proto.RegisterExtension(E_BigqueryOpts)
}

func init() { proto.RegisterFile("bq_table.proto", fileDescriptor_bq_table_b947be292cbdac0c) }

var fileDescriptor_bq_table_b947be292cbdac0c = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4b, 0x43, 0x31,
	0x10, 0x87, 0x89, 0x83, 0xb6, 0xb1, 0xed, 0x90, 0x41, 0x1e, 0x82, 0x18, 0x44, 0xe1, 0x4d, 0x29,
	0xe8, 0xe6, 0xd8, 0x51, 0xd0, 0xe2, 0x1b, 0x05, 0x09, 0x49, 0x3d, 0x63, 0xe4, 0xbd, 0x5c, 0x5e,
	0x2e, 0x19, 0xfc, 0xe3, 0x05, 0x69, 0x1e, 0x1d, 0x0a, 0x4e, 0x07, 0x77, 0xbf, 0xfb, 0xee, 0x3b,
	0xbe, 0xb2, 0xa3, 0xce, 0xc6, 0xf6, 0xa0, 0x62, 0xc2, 0x8c, 0x62, 0xe9, 0x20, 0x68, 0x3b, 0x6a,
	0xda, 0x7d, 0xc1, 0x60, 0x2e, 0xa5, 0x43, 0x74, 0x3d, 0xac, 0xeb, 0xd0, 0x96, 0xcf, 0xf5, 0x07,
	0xd0, 0x2e, 0xf9, 0x98, 0x31, 0x4d, 0x0b, 0x37, 0xef, 0xfc, 0x62, 0xe3, 0xdd, 0x6b, 0x81, 0xf4,
	0xf3, 0x0c, 0x44, 0xc6, 0xc1, 0x36, 0x66, 0x8f, 0x81, 0xc4, 0x15, 0xe7, 0x95, 0xac, 0x83, 0x19,
	0xa0, 0x61, 0x92, 0xb5, 0xf3, 0x6e, 0x5e, 0x3b, 0x2f, 0x66, 0x00, 0x71, 0xcb, 0x57, 0x85, 0x40,
	0x7f, 0x13, 0x86, 0x9a, 0xa0, 0xe6, 0x44, 0xb2, 0x76, 0xd6, 0x2d, 0x0a, 0xc1, 0x13, 0x61, 0xd8,
	0x87, 0xe8, 0xb1, 0xe7, 0x4b, 0xeb, 0xdd, 0xb8, 0xc7, 0x6b, 0x8c, 0x99, 0xc4, 0xb5, 0x9a, 0x94,
	0xd4, 0x41, 0x49, 0x1d, 0x9f, 0x6d, 0x7e, 0xcf, 0x24, 0x6b, 0xcf, 0xef, 0xef, 0xd4, 0xd1, 0x27,
	0xea, 0x7f, 0xc9, 0x6e, 0x71, 0xa0, 0x6f, 0x63, 0xa6, 0xcd, 0xec, 0xed, 0xb4, 0x62, 0xc9, 0x4e,
	0xf5, 0xe1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x91, 0xcc, 0xd5, 0xaf, 0x20, 0x01, 0x00, 0x00,
}
