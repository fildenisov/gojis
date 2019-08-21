// Code generated by capnpc-go. DO NOT EDIT.

package snapshot

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Snapshot struct{ capnp.Struct }

// Snapshot_TypeID is the unique identifier for the type Snapshot.
const Snapshot_TypeID = 0xc499a1c127bebec9

func NewSnapshot(s *capnp.Segment) (Snapshot, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Snapshot{st}, err
}

func NewRootSnapshot(s *capnp.Segment) (Snapshot, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Snapshot{st}, err
}

func ReadRootSnapshot(msg *capnp.Message) (Snapshot, error) {
	root, err := msg.RootPtr()
	return Snapshot{root.Struct()}, err
}

func (s Snapshot) String() string {
	str, _ := text.Marshal(0xc499a1c127bebec9, s.Struct)
	return str
}

func (s Snapshot) Uintptrsize() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Snapshot) SetUintptrsize(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Snapshot) Signature() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Snapshot) HasSignature() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Snapshot) SetSignature(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s Snapshot) Nested() (Nested, error) {
	p, err := s.Struct.Ptr(1)
	return Nested{Struct: p.Struct()}, err
}

func (s Snapshot) HasNested() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Snapshot) SetNested(v Nested) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewNested sets the nested field to a newly
// allocated Nested struct, preferring placement in s's segment.
func (s Snapshot) NewNested() (Nested, error) {
	ss, err := NewNested(s.Struct.Segment())
	if err != nil {
		return Nested{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// Snapshot_List is a list of Snapshot.
type Snapshot_List struct{ capnp.List }

// NewSnapshot creates a new list of Snapshot.
func NewSnapshot_List(s *capnp.Segment, sz int32) (Snapshot_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Snapshot_List{l}, err
}

func (s Snapshot_List) At(i int) Snapshot { return Snapshot{s.List.Struct(i)} }

func (s Snapshot_List) Set(i int, v Snapshot) error { return s.List.SetStruct(i, v.Struct) }

func (s Snapshot_List) String() string {
	str, _ := text.MarshalList(0xc499a1c127bebec9, s.List)
	return str
}

// Snapshot_Promise is a wrapper for a Snapshot promised by a client call.
type Snapshot_Promise struct{ *capnp.Pipeline }

func (p Snapshot_Promise) Struct() (Snapshot, error) {
	s, err := p.Pipeline.Struct()
	return Snapshot{s}, err
}

func (p Snapshot_Promise) Nested() Nested_Promise {
	return Nested_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Nested struct{ capnp.Struct }

// Nested_TypeID is the unique identifier for the type Nested.
const Nested_TypeID = 0x9a2e520f421c93df

func NewNested(s *capnp.Segment) (Nested, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Nested{st}, err
}

func NewRootNested(s *capnp.Segment) (Nested, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Nested{st}, err
}

func ReadRootNested(msg *capnp.Message) (Nested, error) {
	root, err := msg.RootPtr()
	return Nested{root.Struct()}, err
}

func (s Nested) String() string {
	str, _ := text.Marshal(0x9a2e520f421c93df, s.Struct)
	return str
}

func (s Nested) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Nested) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Nested) SetData(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s Nested) Pointers() (Pointer_List, error) {
	p, err := s.Struct.Ptr(1)
	return Pointer_List{List: p.List()}, err
}

func (s Nested) HasPointers() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Nested) SetPointers(v Pointer_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewPointers sets the pointers field to a newly
// allocated Pointer_List, preferring placement in s's segment.
func (s Nested) NewPointers(n int32) (Pointer_List, error) {
	l, err := NewPointer_List(s.Struct.Segment(), n)
	if err != nil {
		return Pointer_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Nested_List is a list of Nested.
type Nested_List struct{ capnp.List }

// NewNested creates a new list of Nested.
func NewNested_List(s *capnp.Segment, sz int32) (Nested_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Nested_List{l}, err
}

func (s Nested_List) At(i int) Nested { return Nested{s.List.Struct(i)} }

func (s Nested_List) Set(i int, v Nested) error { return s.List.SetStruct(i, v.Struct) }

func (s Nested_List) String() string {
	str, _ := text.MarshalList(0x9a2e520f421c93df, s.List)
	return str
}

// Nested_Promise is a wrapper for a Nested promised by a client call.
type Nested_Promise struct{ *capnp.Pipeline }

func (p Nested_Promise) Struct() (Nested, error) {
	s, err := p.Pipeline.Struct()
	return Nested{s}, err
}

type Pointer struct{ capnp.Struct }

// Pointer_TypeID is the unique identifier for the type Pointer.
const Pointer_TypeID = 0xea00f2587b64b8bb

func NewPointer(s *capnp.Segment) (Pointer, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Pointer{st}, err
}

func NewRootPointer(s *capnp.Segment) (Pointer, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Pointer{st}, err
}

func ReadRootPointer(msg *capnp.Message) (Pointer, error) {
	root, err := msg.RootPtr()
	return Pointer{root.Struct()}, err
}

func (s Pointer) String() string {
	str, _ := text.Marshal(0xea00f2587b64b8bb, s.Struct)
	return str
}

func (s Pointer) Offset() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Pointer) SetOffset(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Pointer) Target() (Nested, error) {
	p, err := s.Struct.Ptr(0)
	return Nested{Struct: p.Struct()}, err
}

func (s Pointer) HasTarget() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Pointer) SetTarget(v Nested) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTarget sets the target field to a newly
// allocated Nested struct, preferring placement in s's segment.
func (s Pointer) NewTarget() (Nested, error) {
	ss, err := NewNested(s.Struct.Segment())
	if err != nil {
		return Nested{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Pointer_List is a list of Pointer.
type Pointer_List struct{ capnp.List }

// NewPointer creates a new list of Pointer.
func NewPointer_List(s *capnp.Segment, sz int32) (Pointer_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Pointer_List{l}, err
}

func (s Pointer_List) At(i int) Pointer { return Pointer{s.List.Struct(i)} }

func (s Pointer_List) Set(i int, v Pointer) error { return s.List.SetStruct(i, v.Struct) }

func (s Pointer_List) String() string {
	str, _ := text.MarshalList(0xea00f2587b64b8bb, s.List)
	return str
}

// Pointer_Promise is a wrapper for a Pointer promised by a client call.
type Pointer_Promise struct{ *capnp.Pipeline }

func (p Pointer_Promise) Struct() (Pointer, error) {
	s, err := p.Pipeline.Struct()
	return Pointer{s}, err
}

func (p Pointer_Promise) Target() Nested_Promise {
	return Nested_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_bf39c8350769d85a = "x\xdal\x91\xbf\x8b\x13A\x1c\xc5\xdf\x9b\xc9\xba\x16\x89" +
	"\xbbCR\x0a6\x82?\xd0\xa0\xa0\x85iVDAD" +
	"e'\"\x04\xb1Y\xcd&n\xe1f\xdd\x9d hm" +
	"aae\xa5\x9d`/X\x89B\x82\xc7\x15w\xff\xc1" +
	"u\xc7\x15W\xdc\x95\xf7\x0f\xec1\xb9l\x02\xc9U\xdf" +
	"\x99\xc7\x97\xf9\xbc\xf7\xc6\xffyW\xdct<\x02\xba\xee" +
	"\x9c)w\xbf\x9e\xbf\xe7u\xdb\xdf\xa1<\x96/v\x12" +
	"\xf7\xf6\xd6\x9d)\x1c\xe1\x02\xea\xedo5\xb6\xf3\xdd{" +
	"\xb0\xdc\x9eL.\xfd\xff\xf1m\x13\xda\xe3\xda\xe6\xde\x86" +
	":\xb0s\xff\x178\xfd\xf7\xa7\xff\xb1wt\xb8\xb2G" +
	"\x17h~\xe2\xdf\xe6\x97\xd9\xe93\xed\xa3E\x1ae\xc5" +
	"\x9b\x91a\xfbu\x94\xa5Y\xe7\xe9\x85\xb80q?$" +
	"\xf5YY\x03j\x04\xd4\x95\xab\x80\xbe(\xa9o\x08*" +
	"\xb2E+^\x7f\x04\xe8k\x92\xfa\xa1\xa0\xd7\x8fL\xc4" +
	"\x06\x04\x1b`\x99\x8d\x92\xd4\xc4y\x01\x80\xe7\xc0P\x92" +
	"~\xe5\x09\xb4\xd2\x02+\xe6\xd8g\xf3;,\xb8\xbe\x00" +
	"?x\x05\xe8\xfb\x92:\x14\xac\xb8O\xba\x80~,\xa9" +
	"{\x82J\xb0E\x01\xa8\xe7\x1d@\x87\x92\xfa\xa5`9" +
	"NR\x93\x99\xbc\x80\x9b|\x88\xe9@\xd0\xb1\xc8d\x98" +
	"Ff\x9c\x83q\xe54Hga\xe9/?\x01\xa4\x7f" +
	"J-ap\x92i\xa5\x97\xce\xb2\x97E-V\xbb," +
	"\xa9o\x09\x06\xa3\xc1\xa0\x88Me!0Q>\x8c\xcd" +
	":\xee8\x00\x00\xff\xffw\xcd}\xa9"

func init() {
	schemas.Register(schema_bf39c8350769d85a,
		0x9a2e520f421c93df,
		0xc499a1c127bebec9,
		0xea00f2587b64b8bb)
}
