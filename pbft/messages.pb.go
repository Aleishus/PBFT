// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package pbft is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Message
	Request
	PrePrepare
	Prepare
	Commit
	Checkpoint
	ViewChange
	NewView
	RequestBatch
	BatchMessage
*/
package pbft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	// Types that are valid to be assigned to Payload:
	//	*Message_RequestBatch
	//	*Message_PrePrepare
	//	*Message_Prepare
	//	*Message_Commit
	//	*Message_Checkpoint
	//	*Message_ReturnRequestBatch
	Payload isMessage_Payload `protobuf_oneof:"payload"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_RequestBatch struct {
	RequestBatch *RequestBatch `protobuf:"bytes,1,opt,name=request_batch,json=requestBatch,oneof"`
}
type Message_PrePrepare struct {
	PrePrepare *PrePrepare `protobuf:"bytes,2,opt,name=pre_prepare,json=prePrepare,oneof"`
}
type Message_Prepare struct {
	Prepare *Prepare `protobuf:"bytes,3,opt,name=prepare,oneof"`
}
type Message_Commit struct {
	Commit *Commit `protobuf:"bytes,4,opt,name=commit,oneof"`
}
type Message_Checkpoint struct {
	Checkpoint *Checkpoint `protobuf:"bytes,5,opt,name=checkpoint,oneof"`
}
type Message_ReturnRequestBatch struct {
	ReturnRequestBatch *RequestBatch `protobuf:"bytes,8,opt,name=return_request_batch,json=returnRequestBatch,oneof"`
}

func (*Message_RequestBatch) isMessage_Payload()       {}
func (*Message_PrePrepare) isMessage_Payload()         {}
func (*Message_Prepare) isMessage_Payload()            {}
func (*Message_Commit) isMessage_Payload()             {}
func (*Message_Checkpoint) isMessage_Payload()         {}
func (*Message_ReturnRequestBatch) isMessage_Payload() {}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRequestBatch() *RequestBatch {
	if x, ok := m.GetPayload().(*Message_RequestBatch); ok {
		return x.RequestBatch
	}
	return nil
}

func (m *Message) GetPrePrepare() *PrePrepare {
	if x, ok := m.GetPayload().(*Message_PrePrepare); ok {
		return x.PrePrepare
	}
	return nil
}

func (m *Message) GetPrepare() *Prepare {
	if x, ok := m.GetPayload().(*Message_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Message) GetCommit() *Commit {
	if x, ok := m.GetPayload().(*Message_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Message) GetCheckpoint() *Checkpoint {
	if x, ok := m.GetPayload().(*Message_Checkpoint); ok {
		return x.Checkpoint
	}
	return nil
}

func (m *Message) GetReturnRequestBatch() *RequestBatch {
	if x, ok := m.GetPayload().(*Message_ReturnRequestBatch); ok {
		return x.ReturnRequestBatch
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_RequestBatch)(nil),
		(*Message_PrePrepare)(nil),
		(*Message_Prepare)(nil),
		(*Message_Commit)(nil),
		(*Message_Checkpoint)(nil),
		(*Message_ReturnRequestBatch)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_RequestBatch:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestBatch); err != nil {
			return err
		}
	case *Message_PrePrepare:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PrePrepare); err != nil {
			return err
		}
	case *Message_Prepare:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *Message_Commit:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *Message_Checkpoint:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Checkpoint); err != nil {
			return err
		}
	case *Message_ReturnRequestBatch:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReturnRequestBatch); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Payload has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 1: // payload.request_batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestBatch)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_RequestBatch{msg}
		return true, err
	case 2: // payload.pre_prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrePrepare)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_PrePrepare{msg}
		return true, err
	case 3: // payload.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Prepare)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Prepare{msg}
		return true, err
	case 4: // payload.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Commit)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Commit{msg}
		return true, err
	case 5: // payload.checkpoint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Checkpoint)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Checkpoint{msg}
		return true, err
	case 8: // payload.return_request_batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestBatch)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_ReturnRequestBatch{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_RequestBatch:
		s := proto.Size(x.RequestBatch)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PrePrepare:
		s := proto.Size(x.PrePrepare)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Prepare:
		s := proto.Size(x.Prepare)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Commit:
		s := proto.Size(x.Commit)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Checkpoint:
		s := proto.Size(x.Checkpoint)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_ReturnRequestBatch:
		s := proto.Size(x.ReturnRequestBatch)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Request struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	ReplicaId uint64                     `protobuf:"varint,3,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	Signature []byte                     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Request) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Request) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *Request) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type PrePrepare struct {
	View           uint64        `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64        `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	BatchDigest    string        `protobuf:"bytes,3,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	RequestBatch   *RequestBatch `protobuf:"bytes,4,opt,name=request_batch,json=requestBatch" json:"request_batch,omitempty"`
	ReplicaId      uint64        `protobuf:"varint,5,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
}

func (m *PrePrepare) Reset()                    { *m = PrePrepare{} }
func (m *PrePrepare) String() string            { return proto.CompactTextString(m) }
func (*PrePrepare) ProtoMessage()               {}
func (*PrePrepare) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PrePrepare) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *PrePrepare) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *PrePrepare) GetBatchDigest() string {
	if m != nil {
		return m.BatchDigest
	}
	return ""
}

func (m *PrePrepare) GetRequestBatch() *RequestBatch {
	if m != nil {
		return m.RequestBatch
	}
	return nil
}

func (m *PrePrepare) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

type Prepare struct {
	View           uint64 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	BatchDigest    string `protobuf:"bytes,3,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,4,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
}

func (m *Prepare) Reset()                    { *m = Prepare{} }
func (m *Prepare) String() string            { return proto.CompactTextString(m) }
func (*Prepare) ProtoMessage()               {}
func (*Prepare) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Prepare) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *Prepare) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *Prepare) GetBatchDigest() string {
	if m != nil {
		return m.BatchDigest
	}
	return ""
}

func (m *Prepare) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

type Commit struct {
	View           uint64 `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	BatchDigest    string `protobuf:"bytes,3,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,4,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Commit) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *Commit) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *Commit) GetBatchDigest() string {
	if m != nil {
		return m.BatchDigest
	}
	return ""
}

func (m *Commit) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

type Checkpoint struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	ReplicaId      uint64 `protobuf:"varint,2,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	Id             string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *Checkpoint) Reset()                    { *m = Checkpoint{} }
func (m *Checkpoint) String() string            { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()               {}
func (*Checkpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Checkpoint) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *Checkpoint) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *Checkpoint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ViewChange struct {
	View      uint64           `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	H         uint64           `protobuf:"varint,2,opt,name=h" json:"h,omitempty"`
	Cset      []*ViewChange_C  `protobuf:"bytes,3,rep,name=cset" json:"cset,omitempty"`
	Pset      []*ViewChange_PQ `protobuf:"bytes,4,rep,name=pset" json:"pset,omitempty"`
	Qset      []*ViewChange_PQ `protobuf:"bytes,5,rep,name=qset" json:"qset,omitempty"`
	ReplicaId uint64           `protobuf:"varint,6,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
	Signature []byte           `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ViewChange) Reset()                    { *m = ViewChange{} }
func (m *ViewChange) String() string            { return proto.CompactTextString(m) }
func (*ViewChange) ProtoMessage()               {}
func (*ViewChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ViewChange) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *ViewChange) GetH() uint64 {
	if m != nil {
		return m.H
	}
	return 0
}

func (m *ViewChange) GetCset() []*ViewChange_C {
	if m != nil {
		return m.Cset
	}
	return nil
}

func (m *ViewChange) GetPset() []*ViewChange_PQ {
	if m != nil {
		return m.Pset
	}
	return nil
}

func (m *ViewChange) GetQset() []*ViewChange_PQ {
	if m != nil {
		return m.Qset
	}
	return nil
}

func (m *ViewChange) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

func (m *ViewChange) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type ViewChange_C struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	Id             string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *ViewChange_C) Reset()                    { *m = ViewChange_C{} }
func (m *ViewChange_C) String() string            { return proto.CompactTextString(m) }
func (*ViewChange_C) ProtoMessage()               {}
func (*ViewChange_C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *ViewChange_C) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *ViewChange_C) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ViewChange_PQ struct {
	SequenceNumber uint64 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	BatchDigest    string `protobuf:"bytes,2,opt,name=batch_digest,json=batchDigest" json:"batch_digest,omitempty"`
	View           uint64 `protobuf:"varint,3,opt,name=view" json:"view,omitempty"`
}

func (m *ViewChange_PQ) Reset()                    { *m = ViewChange_PQ{} }
func (m *ViewChange_PQ) String() string            { return proto.CompactTextString(m) }
func (*ViewChange_PQ) ProtoMessage()               {}
func (*ViewChange_PQ) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 1} }

func (m *ViewChange_PQ) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *ViewChange_PQ) GetBatchDigest() string {
	if m != nil {
		return m.BatchDigest
	}
	return ""
}

func (m *ViewChange_PQ) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

type NewView struct {
	View      uint64            `protobuf:"varint,1,opt,name=view" json:"view,omitempty"`
	Vset      []*ViewChange     `protobuf:"bytes,2,rep,name=vset" json:"vset,omitempty"`
	Xset      map[uint64]string `protobuf:"bytes,3,rep,name=xset" json:"xset,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ReplicaId uint64            `protobuf:"varint,4,opt,name=replica_id,json=replicaId" json:"replica_id,omitempty"`
}

func (m *NewView) Reset()                    { *m = NewView{} }
func (m *NewView) String() string            { return proto.CompactTextString(m) }
func (*NewView) ProtoMessage()               {}
func (*NewView) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NewView) GetView() uint64 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *NewView) GetVset() []*ViewChange {
	if m != nil {
		return m.Vset
	}
	return nil
}

func (m *NewView) GetXset() map[uint64]string {
	if m != nil {
		return m.Xset
	}
	return nil
}

func (m *NewView) GetReplicaId() uint64 {
	if m != nil {
		return m.ReplicaId
	}
	return 0
}

type RequestBatch struct {
	Batch []*Request `protobuf:"bytes,1,rep,name=batch" json:"batch,omitempty"`
}

func (m *RequestBatch) Reset()                    { *m = RequestBatch{} }
func (m *RequestBatch) String() string            { return proto.CompactTextString(m) }
func (*RequestBatch) ProtoMessage()               {}
func (*RequestBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RequestBatch) GetBatch() []*Request {
	if m != nil {
		return m.Batch
	}
	return nil
}

type BatchMessage struct {
	// Types that are valid to be assigned to Payload:
	//	*BatchMessage_Request
	//	*BatchMessage_RequestBatch
	//	*BatchMessage_PbftMessage
	//	*BatchMessage_Complaint
	Payload isBatchMessage_Payload `protobuf_oneof:"payload"`
}

func (m *BatchMessage) Reset()                    { *m = BatchMessage{} }
func (m *BatchMessage) String() string            { return proto.CompactTextString(m) }
func (*BatchMessage) ProtoMessage()               {}
func (*BatchMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type isBatchMessage_Payload interface {
	isBatchMessage_Payload()
}

type BatchMessage_Request struct {
	Request *Request `protobuf:"bytes,1,opt,name=request,oneof"`
}
type BatchMessage_RequestBatch struct {
	RequestBatch *RequestBatch `protobuf:"bytes,2,opt,name=request_batch,json=requestBatch,oneof"`
}
type BatchMessage_PbftMessage struct {
	PbftMessage []byte `protobuf:"bytes,3,opt,name=pbft_message,json=pbftMessage,proto3,oneof"`
}
type BatchMessage_Complaint struct {
	Complaint *Request `protobuf:"bytes,4,opt,name=complaint,oneof"`
}

func (*BatchMessage_Request) isBatchMessage_Payload()      {}
func (*BatchMessage_RequestBatch) isBatchMessage_Payload() {}
func (*BatchMessage_PbftMessage) isBatchMessage_Payload()  {}
func (*BatchMessage_Complaint) isBatchMessage_Payload()    {}

func (m *BatchMessage) GetPayload() isBatchMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *BatchMessage) GetRequest() *Request {
	if x, ok := m.GetPayload().(*BatchMessage_Request); ok {
		return x.Request
	}
	return nil
}

func (m *BatchMessage) GetRequestBatch() *RequestBatch {
	if x, ok := m.GetPayload().(*BatchMessage_RequestBatch); ok {
		return x.RequestBatch
	}
	return nil
}

func (m *BatchMessage) GetPbftMessage() []byte {
	if x, ok := m.GetPayload().(*BatchMessage_PbftMessage); ok {
		return x.PbftMessage
	}
	return nil
}

func (m *BatchMessage) GetComplaint() *Request {
	if x, ok := m.GetPayload().(*BatchMessage_Complaint); ok {
		return x.Complaint
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BatchMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BatchMessage_OneofMarshaler, _BatchMessage_OneofUnmarshaler, _BatchMessage_OneofSizer, []interface{}{
		(*BatchMessage_Request)(nil),
		(*BatchMessage_RequestBatch)(nil),
		(*BatchMessage_PbftMessage)(nil),
		(*BatchMessage_Complaint)(nil),
	}
}

func _BatchMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BatchMessage)
	// payload
	switch x := m.Payload.(type) {
	case *BatchMessage_Request:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Request); err != nil {
			return err
		}
	case *BatchMessage_RequestBatch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestBatch); err != nil {
			return err
		}
	case *BatchMessage_PbftMessage:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.PbftMessage)
	case *BatchMessage_Complaint:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Complaint); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BatchMessage.Payload has unexpected type %T", x)
	}
	return nil
}

func _BatchMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BatchMessage)
	switch tag {
	case 1: // payload.request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request)
		err := b.DecodeMessage(msg)
		m.Payload = &BatchMessage_Request{msg}
		return true, err
	case 2: // payload.request_batch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestBatch)
		err := b.DecodeMessage(msg)
		m.Payload = &BatchMessage_RequestBatch{msg}
		return true, err
	case 3: // payload.pbft_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &BatchMessage_PbftMessage{x}
		return true, err
	case 4: // payload.complaint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Request)
		err := b.DecodeMessage(msg)
		m.Payload = &BatchMessage_Complaint{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BatchMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BatchMessage)
	// payload
	switch x := m.Payload.(type) {
	case *BatchMessage_Request:
		s := proto.Size(x.Request)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BatchMessage_RequestBatch:
		s := proto.Size(x.RequestBatch)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BatchMessage_PbftMessage:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PbftMessage)))
		n += len(x.PbftMessage)
	case *BatchMessage_Complaint:
		s := proto.Size(x.Complaint)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Message)(nil), "pbft.message")
	proto.RegisterType((*Request)(nil), "pbft.request")
	proto.RegisterType((*PrePrepare)(nil), "pbft.pre_prepare")
	proto.RegisterType((*Prepare)(nil), "pbft.prepare")
	proto.RegisterType((*Commit)(nil), "pbft.commit")
	proto.RegisterType((*Checkpoint)(nil), "pbft.checkpoint")
	proto.RegisterType((*ViewChange)(nil), "pbft.view_change")
	proto.RegisterType((*ViewChange_C)(nil), "pbft.view_change.C")
	proto.RegisterType((*ViewChange_PQ)(nil), "pbft.view_change.PQ")
	proto.RegisterType((*NewView)(nil), "pbft.new_view")
	proto.RegisterType((*RequestBatch)(nil), "pbft.request_batch")
	proto.RegisterType((*BatchMessage)(nil), "pbft.batch_message")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0x3a, 0x9b, 0xa6, 0x9e, 0xb8, 0xa5, 0x5d, 0x7a, 0x88, 0x22, 0x10, 0xe0, 0x0a, 0xda,
	0x4a, 0xe0, 0x4a, 0xa5, 0x12, 0x55, 0xc5, 0xa9, 0x05, 0x11, 0x0e, 0xa0, 0x76, 0xc5, 0x81, 0x9b,
	0xe5, 0x38, 0xdb, 0xc4, 0x6a, 0xfc, 0x53, 0x7b, 0xd3, 0xd2, 0x27, 0x80, 0x0b, 0x2f, 0xc0, 0x03,
	0x71, 0xe1, 0x0c, 0xcf, 0x83, 0x76, 0xbd, 0x6b, 0xc7, 0x4e, 0x28, 0xbd, 0x20, 0x6e, 0xde, 0x99,
	0x6f, 0xc6, 0xdf, 0x7c, 0x33, 0xbb, 0x03, 0xab, 0x21, 0xcb, 0x32, 0x6f, 0xc4, 0x32, 0x27, 0x49,
	0x63, 0x1e, 0x13, 0x9c, 0x0c, 0xce, 0x78, 0xef, 0xc1, 0x28, 0x8e, 0x47, 0x13, 0xb6, 0x2b, 0x6d,
	0x83, 0xe9, 0xd9, 0x2e, 0x0f, 0x42, 0x96, 0x71, 0x2f, 0x4c, 0x72, 0x98, 0xfd, 0xcb, 0x80, 0xb6,
	0x8a, 0x24, 0x87, 0xb0, 0x92, 0xb2, 0x8b, 0x29, 0xcb, 0xb8, 0x3b, 0xf0, 0xb8, 0x3f, 0xee, 0xa2,
	0x87, 0x68, 0xbb, 0xb3, 0x77, 0xd7, 0x11, 0xa9, 0x9c, 0x8a, 0xab, 0xdf, 0xa0, 0x96, 0x32, 0x1c,
	0x89, 0x33, 0xd9, 0x87, 0x4e, 0x92, 0x32, 0x37, 0x49, 0x59, 0xe2, 0xa5, 0xac, 0x6b, 0xc8, 0xc8,
	0xf5, 0x3c, 0x72, 0xc6, 0xd1, 0x6f, 0x50, 0x48, 0x52, 0x76, 0x92, 0x9f, 0xc8, 0x0e, 0xb4, 0x75,
	0x44, 0x53, 0x46, 0xac, 0x14, 0x11, 0x0a, 0xad, 0xfd, 0xe4, 0x09, 0x2c, 0xf9, 0x71, 0x18, 0x06,
	0xbc, 0x8b, 0x25, 0xd2, 0xca, 0x91, 0xb9, 0xad, 0xdf, 0xa0, 0xca, 0x4b, 0xf6, 0x00, 0xfc, 0x31,
	0xf3, 0xcf, 0x93, 0x38, 0x88, 0x78, 0xb7, 0x25, 0xb1, 0x6b, 0x0a, 0x5b, 0xd8, 0x05, 0x8d, 0xf2,
	0x44, 0xde, 0xc0, 0x46, 0xca, 0xf8, 0x34, 0x8d, 0xdc, 0x6a, 0xfd, 0xcb, 0x37, 0xd5, 0x4f, 0xf2,
	0x10, 0x3a, 0xa3, 0xc2, 0x91, 0x09, 0xed, 0xc4, 0xbb, 0x9e, 0xc4, 0xde, 0xd0, 0xfe, 0x86, 0xa0,
	0xad, 0x42, 0xc8, 0x01, 0x98, 0x85, 0xee, 0x4a, 0xd4, 0x9e, 0x93, 0x77, 0xc6, 0xd1, 0x9d, 0x71,
	0x3e, 0x68, 0x04, 0x2d, 0xc1, 0xa4, 0x5b, 0x24, 0x94, 0x92, 0x5a, 0x54, 0x1f, 0xc9, 0x7d, 0x80,
	0x94, 0x25, 0x93, 0xc0, 0xf7, 0xdc, 0x60, 0x28, 0xd5, 0xc3, 0xd4, 0x54, 0x96, 0xb7, 0x43, 0x72,
	0x0f, 0xcc, 0x2c, 0x18, 0x45, 0x1e, 0x9f, 0xa6, 0x4c, 0x2a, 0x66, 0xd1, 0xd2, 0x60, 0x7f, 0x47,
	0x95, 0x76, 0x11, 0x02, 0xf8, 0x32, 0x60, 0x57, 0x92, 0x1b, 0xa6, 0xf2, 0x9b, 0x6c, 0xc1, 0x9d,
	0x4c, 0xf0, 0x8f, 0x7c, 0xe6, 0x46, 0xd3, 0x70, 0xc0, 0x52, 0x49, 0x01, 0xd3, 0x55, 0x6d, 0x7e,
	0x2f, 0xad, 0xe4, 0x11, 0x58, 0x52, 0x13, 0x77, 0x18, 0x8c, 0x58, 0xc6, 0x25, 0x17, 0x93, 0x76,
	0xa4, 0xed, 0x95, 0x34, 0x91, 0x83, 0xfa, 0x64, 0xe1, 0x3f, 0x2a, 0x5b, 0x9b, 0xab, 0x6a, 0x99,
	0xad, 0x5a, 0x99, 0xf6, 0x17, 0x54, 0x4c, 0xd0, 0x3f, 0x2f, 0xa2, 0x4a, 0x05, 0xd7, 0xa9, 0x7c,
	0x46, 0x7a, 0x42, 0xff, 0x37, 0x93, 0xe1, 0xec, 0x15, 0x58, 0xf4, 0x63, 0xb4, 0xf0, 0xc7, 0xd5,
	0xac, 0x46, 0x7d, 0xa2, 0x56, 0xc1, 0x50, 0x83, 0x66, 0x52, 0x23, 0x18, 0xda, 0x5f, 0x9b, 0xd0,
	0x11, 0x95, 0xb9, 0xfe, 0xd8, 0x8b, 0x46, 0x8b, 0xe5, 0xb7, 0x00, 0x8d, 0x55, 0x26, 0x34, 0x26,
	0x5b, 0x80, 0xfd, 0x8c, 0x89, 0x8a, 0x9a, 0x65, 0xf3, 0x67, 0x52, 0x38, 0xc7, 0x54, 0x02, 0xc8,
	0x36, 0xe0, 0x44, 0x00, 0xb1, 0x04, 0x6e, 0xcc, 0x03, 0x4f, 0x4e, 0xa9, 0x44, 0x08, 0xe4, 0x85,
	0x40, 0xb6, 0x6e, 0x42, 0x0a, 0x44, 0xad, 0xba, 0xa5, 0x1b, 0xef, 0x4b, 0xbb, 0x76, 0x5f, 0x7a,
	0x2f, 0x01, 0x1d, 0xdf, 0x5e, 0xc8, 0x9a, 0x52, 0xbd, 0x21, 0x18, 0x27, 0xa7, 0xb7, 0x0f, 0xaf,
	0x0f, 0x80, 0x31, 0x3f, 0x00, 0x5a, 0xeb, 0x66, 0xa9, 0xb5, 0xfd, 0x03, 0xc1, 0x72, 0xc4, 0xae,
	0x5c, 0x29, 0xfc, 0xa2, 0x66, 0x3c, 0x06, 0x7c, 0x29, 0xb4, 0x32, 0xa4, 0x56, 0xeb, 0x73, 0x5a,
	0x51, 0xe9, 0x26, 0x4f, 0x01, 0x7f, 0x2a, 0xbb, 0xd4, 0xcd, 0x61, 0x3a, 0xb1, 0xf3, 0x31, 0x63,
	0xfc, 0x75, 0xc4, 0xd3, 0x6b, 0x2a, 0x51, 0x7f, 0x19, 0xc5, 0xde, 0x0b, 0x30, 0x8b, 0x08, 0xb2,
	0x06, 0xcd, 0x73, 0x76, 0xad, 0x38, 0x89, 0x4f, 0xb2, 0x01, 0xad, 0x4b, 0x6f, 0x32, 0x65, 0xaa,
	0xc6, 0xfc, 0x70, 0x68, 0x1c, 0x20, 0x7b, 0xbf, 0xf6, 0x62, 0x90, 0x4d, 0x68, 0xe9, 0xa5, 0xd4,
	0x2c, 0x17, 0x85, 0xc2, 0xd0, 0xdc, 0x67, 0xff, 0x44, 0xb0, 0x92, 0x6b, 0xa7, 0x77, 0xda, 0x4e,
	0xf1, 0x0a, 0xab, 0x87, 0xb7, 0x1a, 0x28, 0x36, 0x8c, 0x7e, 0xa5, 0xe7, 0xd6, 0x9f, 0x71, 0xfb,
	0xf5, 0xb7, 0x09, 0x96, 0x40, 0xe9, 0xdf, 0xca, 0xc6, 0x58, 0xfd, 0x06, 0xed, 0x08, 0xeb, 0x3b,
	0xc5, 0xe5, 0x19, 0x98, 0x7e, 0x1c, 0x26, 0x13, 0x4f, 0x6c, 0x26, 0xbc, 0x98, 0x4d, 0x89, 0x98,
	0x59, 0x26, 0x83, 0x25, 0xb9, 0x25, 0x9e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xf0, 0xd9,
	0xb4, 0xe5, 0x07, 0x00, 0x00,
}
