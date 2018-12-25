// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto_files/HelloWorld.proto

package HelloWorld

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Person_Country int32

const (
	Person_INDIA    Person_Country = 0
	Person_PAKISTAN Person_Country = 1
	Person_USA      Person_Country = 2
)

var Person_Country_name = map[int32]string{
	0: "INDIA",
	1: "PAKISTAN",
	2: "USA",
}

var Person_Country_value = map[string]int32{
	"INDIA":    0,
	"PAKISTAN": 1,
	"USA":      2,
}

func (x Person_Country) String() string {
	return proto.EnumName(Person_Country_name, int32(x))
}

func (Person_Country) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cc0f9d2259375ef, []int{0, 0}
}

type Person struct {
	Id                   int32          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Nationality          Person_Country `protobuf:"varint,3,opt,name=nationality,proto3,enum=HelloWorld.Person_Country" json:"nationality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cc0f9d2259375ef, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetNationality() Person_Country {
	if m != nil {
		return m.Nationality
	}
	return Person_INDIA
}

func init() {
	proto.RegisterEnum("HelloWorld.Person_Country", Person_Country_name, Person_Country_value)
	proto.RegisterType((*Person)(nil), "HelloWorld.Person")
}

func init() { proto.RegisterFile("proto_files/HelloWorld.proto", fileDescriptor_4cc0f9d2259375ef) }

var fileDescriptor_4cc0f9d2259375ef = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x8f, 0x4f, 0xcb, 0xcc, 0x49, 0x2d, 0xd6, 0xf7, 0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0xcf, 0x2f,
	0xca, 0x49, 0xd1, 0x03, 0x0b, 0x0b, 0x71, 0x21, 0x44, 0x94, 0xa6, 0x33, 0x72, 0xb1, 0x05, 0xa4,
	0x16, 0x15, 0xe7, 0xe7, 0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0,
	0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30,
	0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x36, 0x5c, 0xdc, 0x79, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x89,
	0x39, 0x99, 0x25, 0x95, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x52, 0x7a, 0x48, 0x56, 0x40,
	0x0c, 0xd3, 0x73, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x42, 0x56, 0xae, 0xa4, 0xcd, 0xc5,
	0x0e, 0x15, 0x17, 0xe2, 0xe4, 0x62, 0xf5, 0xf4, 0x73, 0xf1, 0x74, 0x14, 0x60, 0x10, 0xe2, 0xe1,
	0xe2, 0x08, 0x70, 0xf4, 0xf6, 0x0c, 0x0e, 0x71, 0xf4, 0x13, 0x60, 0x14, 0x62, 0xe7, 0x62, 0x0e,
	0x0d, 0x76, 0x14, 0x60, 0x4a, 0x62, 0x03, 0x3b, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0x6f, 0x83, 0xad, 0xcc, 0x00, 0x00, 0x00,
}
