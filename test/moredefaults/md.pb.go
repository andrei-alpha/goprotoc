// Code generated by protoc-gen-dgo.
// source: md.proto
// DO NOT EDIT!

/*
	Package moredefaults is a generated protocol buffer package.

	It is generated from these files:
		md.proto

	It has these top-level messages:
		MoreDefaultsB
		MoreDefaultsA
*/
package moredefaults

import proto "github.com/dropbox/goprotoc/proto"
import math "math"

// discarding unused import gogoproto "github.com/dropbox/goprotoc/gogoproto/gogo.pb"
import test "github.com/dropbox/goprotoc/test/example"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MoreDefaultsB struct {
	Field1           *string `protobuf:"bytes,1,opt" json:"Field1,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MoreDefaultsB) Reset()         { *m = MoreDefaultsB{} }
func (m *MoreDefaultsB) String() string { return proto.CompactTextString(m) }
func (*MoreDefaultsB) ProtoMessage()    {}

func (m *MoreDefaultsB) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

type MoreDefaultsA struct {
	Field1           *int64         `protobuf:"varint,1,opt,def=1234" json:"Field1,omitempty"`
	Field2           int64          `protobuf:"varint,2,opt" json:"Field2"`
	B1               *MoreDefaultsB `protobuf:"bytes,3,opt" json:"B1,omitempty"`
	B2               MoreDefaultsB  `protobuf:"bytes,4,opt" json:"B2"`
	A1               *test.A        `protobuf:"bytes,5,opt" json:"A1,omitempty"`
	A2               test.A         `protobuf:"bytes,6,opt" json:"A2"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *MoreDefaultsA) Reset()         { *m = MoreDefaultsA{} }
func (m *MoreDefaultsA) String() string { return proto.CompactTextString(m) }
func (*MoreDefaultsA) ProtoMessage()    {}

const Default_MoreDefaultsA_Field1 int64 = 1234

func (m *MoreDefaultsA) GetField1() int64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return Default_MoreDefaultsA_Field1
}

func (m *MoreDefaultsA) GetField2() int64 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *MoreDefaultsA) GetB1() *MoreDefaultsB {
	if m != nil {
		return m.B1
	}
	return nil
}

func (m *MoreDefaultsA) GetB2() MoreDefaultsB {
	if m != nil {
		return m.B2
	}
	return MoreDefaultsB{}
}

func (m *MoreDefaultsA) GetA1() *test.A {
	if m != nil {
		return m.A1
	}
	return nil
}

func (m *MoreDefaultsA) GetA2() test.A {
	if m != nil {
		return m.A2
	}
	return test.A{}
}

func init() {
}
