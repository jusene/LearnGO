// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extension_user/extension_user.proto

package extension_user

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	extension_base "github.com/gogo/protobuf/protoc-gen-gogo/testdata/extension_base"
	extension_extra "github.com/gogo/protobuf/protoc-gen-gogo/testdata/extension_extra"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UserMessage struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Rank                 *string  `protobuf:"bytes,2,opt,name=rank" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}
func (*UserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{0}
}
func (m *UserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMessage.Unmarshal(m, b)
}
func (m *UserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMessage.Marshal(b, m, deterministic)
}
func (m *UserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMessage.Merge(m, src)
}
func (m *UserMessage) XXX_Size() int {
	return xxx_messageInfo_UserMessage.Size(m)
}
func (m *UserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMessage proto.InternalMessageInfo

func (m *UserMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UserMessage) GetRank() string {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return ""
}

// Extend inside the scope of another type
type LoudMessage struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *LoudMessage) Reset()         { *m = LoudMessage{} }
func (m *LoudMessage) String() string { return proto.CompactTextString(m) }
func (*LoudMessage) ProtoMessage()    {}
func (*LoudMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{1}
}

var extRange_LoudMessage = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*LoudMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_LoudMessage
}

func (m *LoudMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoudMessage.Unmarshal(m, b)
}
func (m *LoudMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoudMessage.Marshal(b, m, deterministic)
}
func (m *LoudMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoudMessage.Merge(m, src)
}
func (m *LoudMessage) XXX_Size() int {
	return xxx_messageInfo_LoudMessage.Size(m)
}
func (m *LoudMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LoudMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LoudMessage proto.InternalMessageInfo

var E_LoudMessage_Volume = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         8,
	Name:          "extension_user.LoudMessage.volume",
	Tag:           "varint,8,opt,name=volume",
	Filename:      "extension_user/extension_user.proto",
}

// Extend inside the scope of another type, using a message.
type LoginMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginMessage) Reset()         { *m = LoginMessage{} }
func (m *LoginMessage) String() string { return proto.CompactTextString(m) }
func (*LoginMessage) ProtoMessage()    {}
func (*LoginMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{2}
}
func (m *LoginMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginMessage.Unmarshal(m, b)
}
func (m *LoginMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginMessage.Marshal(b, m, deterministic)
}
func (m *LoginMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginMessage.Merge(m, src)
}
func (m *LoginMessage) XXX_Size() int {
	return xxx_messageInfo_LoginMessage.Size(m)
}
func (m *LoginMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LoginMessage proto.InternalMessageInfo

var E_LoginMessage_UserMessage = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*UserMessage)(nil),
	Field:         16,
	Name:          "extension_user.LoginMessage.user_message",
	Tag:           "bytes,16,opt,name=user_message",
	Filename:      "extension_user/extension_user.proto",
}

type Detail struct {
	Color                *string  `protobuf:"bytes,1,opt,name=color" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Detail) Reset()         { *m = Detail{} }
func (m *Detail) String() string { return proto.CompactTextString(m) }
func (*Detail) ProtoMessage()    {}
func (*Detail) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{3}
}
func (m *Detail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Detail.Unmarshal(m, b)
}
func (m *Detail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Detail.Marshal(b, m, deterministic)
}
func (m *Detail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Detail.Merge(m, src)
}
func (m *Detail) XXX_Size() int {
	return xxx_messageInfo_Detail.Size(m)
}
func (m *Detail) XXX_DiscardUnknown() {
	xxx_messageInfo_Detail.DiscardUnknown(m)
}

var xxx_messageInfo_Detail proto.InternalMessageInfo

func (m *Detail) GetColor() string {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return ""
}

// An extension of an extension
type Announcement struct {
	Words                *string  `protobuf:"bytes,1,opt,name=words" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Announcement) Reset()         { *m = Announcement{} }
func (m *Announcement) String() string { return proto.CompactTextString(m) }
func (*Announcement) ProtoMessage()    {}
func (*Announcement) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{4}
}
func (m *Announcement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Announcement.Unmarshal(m, b)
}
func (m *Announcement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Announcement.Marshal(b, m, deterministic)
}
func (m *Announcement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announcement.Merge(m, src)
}
func (m *Announcement) XXX_Size() int {
	return xxx_messageInfo_Announcement.Size(m)
}
func (m *Announcement) XXX_DiscardUnknown() {
	xxx_messageInfo_Announcement.DiscardUnknown(m)
}

var xxx_messageInfo_Announcement proto.InternalMessageInfo

func (m *Announcement) GetWords() string {
	if m != nil && m.Words != nil {
		return *m.Words
	}
	return ""
}

var E_Announcement_LoudExt = &proto.ExtensionDesc{
	ExtendedType:  (*LoudMessage)(nil),
	ExtensionType: (*Announcement)(nil),
	Field:         100,
	Name:          "extension_user.Announcement.loud_ext",
	Tag:           "bytes,100,opt,name=loud_ext",
	Filename:      "extension_user/extension_user.proto",
}

// Something that can be put in a message set.
type OldStyleParcel struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Height               *int32   `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OldStyleParcel) Reset()         { *m = OldStyleParcel{} }
func (m *OldStyleParcel) String() string { return proto.CompactTextString(m) }
func (*OldStyleParcel) ProtoMessage()    {}
func (*OldStyleParcel) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{5}
}
func (m *OldStyleParcel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OldStyleParcel.Unmarshal(m, b)
}
func (m *OldStyleParcel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OldStyleParcel.Marshal(b, m, deterministic)
}
func (m *OldStyleParcel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldStyleParcel.Merge(m, src)
}
func (m *OldStyleParcel) XXX_Size() int {
	return xxx_messageInfo_OldStyleParcel.Size(m)
}
func (m *OldStyleParcel) XXX_DiscardUnknown() {
	xxx_messageInfo_OldStyleParcel.DiscardUnknown(m)
}

var xxx_messageInfo_OldStyleParcel proto.InternalMessageInfo

func (m *OldStyleParcel) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *OldStyleParcel) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

var E_OldStyleParcel_MessageSetExtension = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.OldStyleMessage)(nil),
	ExtensionType: (*OldStyleParcel)(nil),
	Field:         2001,
	Name:          "extension_user.OldStyleParcel",
	Tag:           "bytes,2001,opt,name=message_set_extension",
	Filename:      "extension_user/extension_user.proto",
}

var E_UserMessage = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*UserMessage)(nil),
	Field:         5,
	Name:          "extension_user.user_message",
	Tag:           "bytes,5,opt,name=user_message",
	Filename:      "extension_user/extension_user.proto",
}

var E_ExtraMessage = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*extension_extra.ExtraMessage)(nil),
	Field:         9,
	Name:          "extension_user.extra_message",
	Tag:           "bytes,9,opt,name=extra_message",
	Filename:      "extension_user/extension_user.proto",
}

var E_Width = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         6,
	Name:          "extension_user.width",
	Tag:           "varint,6,opt,name=width",
	Filename:      "extension_user/extension_user.proto",
}

var E_Area = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         7,
	Name:          "extension_user.area",
	Tag:           "varint,7,opt,name=area",
	Filename:      "extension_user/extension_user.proto",
}

var E_Detail = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: ([]*Detail)(nil),
	Field:         17,
	Name:          "extension_user.detail",
	Tag:           "bytes,17,rep,name=detail",
	Filename:      "extension_user/extension_user.proto",
}

func init() {
	proto.RegisterType((*UserMessage)(nil), "extension_user.UserMessage")
	proto.RegisterExtension(E_LoudMessage_Volume)
	proto.RegisterType((*LoudMessage)(nil), "extension_user.LoudMessage")
	proto.RegisterExtension(E_LoginMessage_UserMessage)
	proto.RegisterType((*LoginMessage)(nil), "extension_user.LoginMessage")
	proto.RegisterType((*Detail)(nil), "extension_user.Detail")
	proto.RegisterExtension(E_Announcement_LoudExt)
	proto.RegisterType((*Announcement)(nil), "extension_user.Announcement")
	proto.RegisterExtension(E_OldStyleParcel_MessageSetExtension)
	proto.RegisterType((*OldStyleParcel)(nil), "extension_user.OldStyleParcel")
	proto.RegisterExtension(E_UserMessage)
	proto.RegisterExtension(E_ExtraMessage)
	proto.RegisterExtension(E_Width)
	proto.RegisterExtension(E_Area)
	proto.RegisterExtension(E_Detail)
}

func init() {
	proto.RegisterFile("extension_user/extension_user.proto", fileDescriptor_359ba8abf543ca10)
}

var fileDescriptor_359ba8abf543ca10 = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xdb, 0xc6, 0x4d, 0xc7, 0x69, 0x29, 0x06, 0xaa, 0xa8, 0x40, 0xb1, 0x8c, 0x90, 0x2c,
	0xa4, 0xc6, 0xc2, 0x88, 0x8b, 0x4f, 0x10, 0x91, 0x13, 0x41, 0x20, 0x17, 0x2e, 0x70, 0xb0, 0x36,
	0xf6, 0xe0, 0x58, 0xb5, 0x77, 0xd1, 0xee, 0x9a, 0x06, 0x4e, 0xf9, 0x26, 0xfe, 0x84, 0x3f, 0x42,
	0x5e, 0xaf, 0x5b, 0xc7, 0x48, 0x11, 0xbd, 0x44, 0xfb, 0x66, 0xdf, 0xbc, 0x99, 0x7d, 0x33, 0x31,
	0x3c, 0xc5, 0x95, 0x44, 0x2a, 0x72, 0x46, 0xe3, 0x4a, 0x20, 0xf7, 0x37, 0xe1, 0xe4, 0x3b, 0x67,
	0x92, 0xd9, 0x47, 0x9b, 0xd1, 0xd3, 0x4e, 0xd2, 0x82, 0x08, 0xf4, 0x37, 0x61, 0x93, 0x74, 0xfa,
	0xec, 0x26, 0x8a, 0x2b, 0xc9, 0x89, 0xdf, 0xc3, 0x0d, 0xcd, 0x7d, 0x05, 0xd6, 0x67, 0x81, 0xfc,
	0x3d, 0x0a, 0x41, 0x32, 0xb4, 0x6d, 0xd8, 0xa3, 0xa4, 0xc4, 0xb1, 0xe1, 0x18, 0xde, 0x41, 0xa4,
	0xce, 0x75, 0x8c, 0x13, 0x7a, 0x39, 0xde, 0x69, 0x62, 0xf5, 0xd9, 0x9d, 0x83, 0x35, 0x67, 0x55,
	0xaa, 0xd3, 0x9e, 0x0f, 0x87, 0xe9, 0xf1, 0x7a, 0xbd, 0x5e, 0xef, 0x04, 0x2f, 0xc1, 0xfc, 0xc1,
	0x8a, 0xaa, 0x44, 0xfb, 0xe1, 0xa4, 0xd7, 0xd7, 0x94, 0x08, 0xd4, 0x09, 0xe3, 0xa1, 0x63, 0x78,
	0x87, 0x91, 0xa6, 0xba, 0x97, 0x30, 0x9a, 0xb3, 0x2c, 0xa7, 0xfa, 0x36, 0xf8, 0x0a, 0xa3, 0xfa,
	0xa1, 0x71, 0xa9, 0xbb, 0xda, 0x2a, 0x75, 0xec, 0x18, 0x9e, 0x15, 0x74, 0x29, 0xca, 0xba, 0xce,
	0xab, 0x22, 0xab, 0xba, 0x01, 0xee, 0x19, 0x98, 0x6f, 0x51, 0x92, 0xbc, 0xb0, 0xef, 0xc3, 0x20,
	0x61, 0x05, 0xe3, 0xfa, 0xb5, 0x0d, 0x70, 0x7f, 0xc1, 0xe8, 0x0d, 0xa5, 0xac, 0xa2, 0x09, 0x96,
	0x48, 0x65, 0xcd, 0xba, 0x62, 0x3c, 0x15, 0x2d, 0x4b, 0x81, 0xe0, 0x13, 0x0c, 0x0b, 0x56, 0xa5,
	0xb5, 0x97, 0xf6, 0x3f, 0xb5, 0x3b, 0xd6, 0x8c, 0x53, 0xd5, 0xde, 0xa3, 0x3e, 0xa5, 0x5b, 0x22,
	0xda, 0xaf, 0xa5, 0x66, 0x2b, 0xe9, 0xfe, 0x36, 0xe0, 0xe8, 0x43, 0x91, 0x5e, 0xc8, 0x9f, 0x05,
	0x7e, 0x24, 0x3c, 0xc1, 0xa2, 0x33, 0x91, 0x9d, 0xeb, 0x89, 0x9c, 0x80, 0xb9, 0xc4, 0x3c, 0x5b,
	0x4a, 0x35, 0x93, 0x41, 0xa4, 0x51, 0x20, 0xe1, 0x81, 0xb6, 0x2c, 0x16, 0x28, 0xe3, 0xeb, 0x92,
	0xf6, 0x93, 0xbe, 0x81, 0x6d, 0x91, 0xb6, 0xcb, 0x3f, 0x77, 0x54, 0x9b, 0x67, 0xfd, 0x36, 0x37,
	0x9b, 0x89, 0xee, 0x69, 0xf9, 0x0b, 0x94, 0xb3, 0x96, 0x18, 0xde, 0x6a, 0x5a, 0x83, 0xdb, 0x4d,
	0x2b, 0x8c, 0xe1, 0x50, 0xad, 0xeb, 0xff, 0xa9, 0x1f, 0x28, 0xf5, 0xc7, 0x93, 0xfe, 0xae, 0xcf,
	0xea, 0xdf, 0x56, 0x7f, 0x84, 0x1d, 0x14, 0xbe, 0x80, 0xc1, 0x55, 0x9e, 0xca, 0xe5, 0x76, 0x61,
	0x53, 0xf9, 0xdc, 0x30, 0x43, 0x1f, 0xf6, 0x08, 0x47, 0xb2, 0x3d, 0x63, 0xdf, 0x31, 0xbc, 0xdd,
	0x48, 0x11, 0xc3, 0x77, 0x60, 0xa6, 0xcd, 0xca, 0x6d, 0x4d, 0xb9, 0xeb, 0xec, 0x7a, 0x56, 0x70,
	0xd2, 0xf7, 0xa6, 0xd9, 0xd6, 0x48, 0x4b, 0x4c, 0xa7, 0x5f, 0x5e, 0x67, 0xb9, 0x5c, 0x56, 0x8b,
	0x49, 0xc2, 0x4a, 0x3f, 0x63, 0x19, 0xf3, 0xd5, 0x5f, 0x79, 0x51, 0x7d, 0x6b, 0x0e, 0xc9, 0x79,
	0x86, 0xf4, 0x5c, 0x5d, 0x48, 0x14, 0x32, 0x25, 0x92, 0xf4, 0xbe, 0x2b, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xc5, 0x88, 0x70, 0x88, 0x77, 0x04, 0x00, 0x00,
}
