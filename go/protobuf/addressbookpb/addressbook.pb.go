// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

package addressbookpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

// [START messages]
type Person struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0}
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

func (m *Person) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor_1eb1a68c9dd6d429) }

var fileDescriptor_1eb1a68c9dd6d429 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x9a, 0x86, 0x76, 0xa2, 0x35, 0x2e, 0x22, 0x21, 0x08, 0x96, 0x9e, 0x02, 0xc2,
	0x16, 0xaa, 0xe0, 0xc9, 0x83, 0x85, 0x82, 0xa2, 0xb5, 0x25, 0xb4, 0x08, 0x5e, 0x64, 0x43, 0xd6,
	0x1a, 0x9a, 0x64, 0x96, 0xec, 0xe6, 0xd0, 0xdf, 0xe6, 0x9f, 0x93, 0xee, 0x26, 0x52, 0xc4, 0xdb,
	0xcc, 0x9b, 0xc7, 0xec, 0xf7, 0x66, 0xe1, 0x8c, 0xa5, 0x69, 0xc5, 0xa5, 0x4c, 0x10, 0xb7, 0x54,
	0x54, 0xa8, 0x90, 0xf4, 0x54, 0xad, 0xb0, 0xca, 0x58, 0x1e, 0x5e, 0x6d, 0x10, 0x37, 0x39, 0x1f,
	0x6b, 0x3d, 0xa9, 0x3f, 0xc7, 0x2a, 0x2b, 0xb8, 0x54, 0xac, 0x10, 0xc6, 0x3a, 0xfa, 0xb6, 0xc1,
	0x5d, 0xf2, 0x4a, 0x62, 0x49, 0x06, 0x60, 0x67, 0x69, 0x60, 0x0d, 0xad, 0xa8, 0x1f, 0xdb, 0x59,
	0x4a, 0x08, 0x38, 0x25, 0x2b, 0x78, 0x60, 0x6b, 0x45, 0xd7, 0xe4, 0x1c, 0xba, 0xbc, 0x60, 0x59,
	0x1e, 0x74, 0xb4, 0x68, 0x1a, 0x72, 0x0b, 0xae, 0xf8, 0xc2, 0x92, 0xcb, 0xc0, 0x19, 0x76, 0x22,
	0x6f, 0x72, 0x49, 0x5b, 0x00, 0x6a, 0x76, 0xd3, 0xe5, 0x7e, 0xfc, 0x5a, 0x17, 0x09, 0xaf, 0xe2,
	0xc6, 0x4b, 0xee, 0xe1, 0x38, 0x67, 0x52, 0x7d, 0xd4, 0x22, 0x65, 0x8a, 0xa7, 0x41, 0x77, 0x68,
	0x45, 0xde, 0x24, 0xa4, 0x06, 0x99, 0xb6, 0xc8, 0x74, 0xd5, 0x22, 0xc7, 0xde, 0xde, 0xbf, 0x36,
	0xf6, 0x70, 0x0d, 0xde, 0xc1, 0x56, 0x72, 0x01, 0x6e, 0xa9, 0xab, 0x26, 0x41, 0xd3, 0x11, 0x0a,
	0x8e, 0xda, 0x09, 0x93, 0x62, 0x30, 0x09, 0xff, 0x27, 0x5b, 0xed, 0x04, 0x8f, 0xb5, 0x6f, 0x74,
	0x0d, 0xfd, 0x5f, 0x89, 0x00, 0xb8, 0xf3, 0xc5, 0xf4, 0xe9, 0x65, 0xe6, 0x1f, 0x91, 0x1e, 0x38,
	0x8f, 0x8b, 0xf9, 0xcc, 0xb7, 0xf6, 0xd5, 0xdb, 0x22, 0x7e, 0xf6, 0xed, 0xd1, 0x1d, 0x78, 0x0f,
	0xe6, 0xfa, 0x53, 0xc4, 0x2d, 0x89, 0xc0, 0x15, 0x1c, 0x45, 0xce, 0x03, 0x4b, 0xdf, 0xc1, 0xff,
	0xfb, 0x5a, 0xdc, 0xcc, 0xa7, 0xa7, 0xef, 0x27, 0x07, 0xdf, 0x26, 0x92, 0xc4, 0xd5, 0x71, 0x6f,
	0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0x94, 0xb4, 0x24, 0xce, 0x01, 0x00, 0x00,
}
