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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Snapshot{st}, err
}

func NewRootSnapshot(s *capnp.Segment) (Snapshot, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
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

func (s Snapshot) Nested() (Nested, error) {
	p, err := s.Struct.Ptr(0)
	return Nested{Struct: p.Struct()}, err
}

func (s Snapshot) HasNested() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Snapshot) SetNested(v Nested) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewNested sets the nested field to a newly
// allocated Nested struct, preferring placement in s's segment.
func (s Snapshot) NewNested() (Nested, error) {
	ss, err := NewNested(s.Struct.Segment())
	if err != nil {
		return Nested{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Snapshot_List is a list of Snapshot.
type Snapshot_List struct{ capnp.List }

// NewSnapshot creates a new list of Snapshot.
func NewSnapshot_List(s *capnp.Segment, sz int32) (Snapshot_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
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
	return Nested_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Nested struct{ capnp.Struct }

// Nested_TypeID is the unique identifier for the type Nested.
const Nested_TypeID = 0x9a2e520f421c93df

func NewNested(s *capnp.Segment) (Nested, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Nested{st}, err
}

func NewRootNested(s *capnp.Segment) (Nested, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
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

func (s Nested) Len() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Nested) SetLen(v int64) {
	s.Struct.SetUint64(0, uint64(v))
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
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
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

const schema_bf39c8350769d85a = "x\xdal\x90\xbfJ#Q\x18\xc5\xcf\xb97\xb3\x93\"" +
	"\xd9\x99!)wYv\xd9ew\x05\x83\x82\x16\xa6\x99" +
	"(\xa6\x88D\x99\xab\x8d\xd8\x0dfb\x02:\x193#" +
	"\x16>F*\xed|\x02\xc1J\x14\x12\x14\x0b}\x03;" +
	"{K_`d\xc6\xfc\x11\xb5\xba\xdf=|\xfc\xee\xef" +
	"\\\xb3W\x11\xb3\x9aA@\xe5\xb4/\xf1c\xef\xdb\x92" +
	"\xb1^:\x812\xc8x\xeb\xa1\xad\xcf\xdf-\x0c\xa0\x09" +
	"\x1d\xb0\xf6\xce\xad\x83\xe4\xdc?\x03\xe3\xfb~\xff\xef\xf5" +
	"\xe9\xf1-,\xe3\xed\"u\xa0\xf0\x9d7\x85?\xe9\xf4" +
	"\x9368\xb8\xbah\x1cm>?\xbdc\xa6\x0b\x8b\xbc" +
	",\xd4\xd2\xa9\xcaC0\x0e}7\x08[\x9d\x88\xa5m" +
	"7\xf0\x83\xf2\xda\x0f/\x8c\xbc\x86C\xaa\x9c\xcc\x00\x19" +
	"\x02Vu\x0aP\x15IU\x17$\x8bL\xb2\xda/@" +
	"-K*G\xd0\x12,R\x00\xd6\xea\x0a\xa0\xea\x92\xaa" +
	"%h4\xdc\xc8e\x1e\x82yP\xdf\xf5|j\x10\xd4" +
	"\xc08\xe8\xb4\xfd\xc8\xeb\x86\x00\xf8\x15t$i\x8e\x9c" +
	"\xc1$\x1ak\x89\xa1\xd6\xc6\xf0\x8eD,3\x16\xcb\x97" +
	"\x01\x95\x95TEA\xdbO\xcdiN>\x15\xa4\xf9I" +
	"G\xc7~\x15HX\xd91\xeb\x7f\xc2\xfa-\xa9f&" +
	"%\xa7\x93\xec\x9f\xa4\x9a\x13\xb4;\xcdf\xe8E\xa3\x16" +
	"v\xe4vw\xbc\xe8\xe3s/\x01\x00\x00\xff\xff\xf4O" +
	"nD"

func init() {
	schemas.Register(schema_bf39c8350769d85a,
		0x9a2e520f421c93df,
		0xc499a1c127bebec9,
		0xea00f2587b64b8bb)
}