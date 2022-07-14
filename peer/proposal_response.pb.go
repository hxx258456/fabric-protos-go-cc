// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/proposal_response.proto

package peer

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	common "github.com/hxx258456/fabric-protos-go-cc/common"
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

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement *Endorsement `protobuf:"bytes,6,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
	// The chaincode interest derived from simulating the proposal.
	Interest             *ChaincodeInterest `protobuf:"bytes,7,opt,name=interest,proto3" json:"interest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ProposalResponse) Reset()         { *m = ProposalResponse{} }
func (m *ProposalResponse) String() string { return proto.CompactTextString(m) }
func (*ProposalResponse) ProtoMessage()    {}
func (*ProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed51030656d961a, []int{0}
}

func (m *ProposalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalResponse.Unmarshal(m, b)
}
func (m *ProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalResponse.Marshal(b, m, deterministic)
}
func (m *ProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResponse.Merge(m, src)
}
func (m *ProposalResponse) XXX_Size() int {
	return xxx_messageInfo_ProposalResponse.Size(m)
}
func (m *ProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResponse proto.InternalMessageInfo

func (m *ProposalResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ProposalResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ProposalResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

func (m *ProposalResponse) GetInterest() *ChaincodeInterest {
	if m != nil {
		return m.Interest
	}
	return nil
}

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed51030656d961a, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte).
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposal_hash,json=proposalHash,proto3" json:"proposal_hash,omitempty"`
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension            []byte   `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProposalResponsePayload) Reset()         { *m = ProposalResponsePayload{} }
func (m *ProposalResponsePayload) String() string { return proto.CompactTextString(m) }
func (*ProposalResponsePayload) ProtoMessage()    {}
func (*ProposalResponsePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed51030656d961a, []int{2}
}

func (m *ProposalResponsePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalResponsePayload.Unmarshal(m, b)
}
func (m *ProposalResponsePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalResponsePayload.Marshal(b, m, deterministic)
}
func (m *ProposalResponsePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResponsePayload.Merge(m, src)
}
func (m *ProposalResponsePayload) XXX_Size() int {
	return xxx_messageInfo_ProposalResponsePayload.Size(m)
}
func (m *ProposalResponsePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResponsePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResponsePayload proto.InternalMessageInfo

func (m *ProposalResponsePayload) GetProposalHash() []byte {
	if m != nil {
		return m.ProposalHash
	}
	return nil
}

func (m *ProposalResponsePayload) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	// Identity of the endorser (e.g. its certificate)
	Endorser []byte `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endorsement) Reset()         { *m = Endorsement{} }
func (m *Endorsement) String() string { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()    {}
func (*Endorsement) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed51030656d961a, []int{3}
}

func (m *Endorsement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endorsement.Unmarshal(m, b)
}
func (m *Endorsement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endorsement.Marshal(b, m, deterministic)
}
func (m *Endorsement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endorsement.Merge(m, src)
}
func (m *Endorsement) XXX_Size() int {
	return xxx_messageInfo_Endorsement.Size(m)
}
func (m *Endorsement) XXX_DiscardUnknown() {
	xxx_messageInfo_Endorsement.DiscardUnknown(m)
}

var xxx_messageInfo_Endorsement proto.InternalMessageInfo

func (m *Endorsement) GetEndorser() []byte {
	if m != nil {
		return m.Endorser
	}
	return nil
}

func (m *Endorsement) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// ChaincodeInterest defines an interest about an endorsement
// for a specific single chaincode invocation.
// Multiple chaincodes indicate chaincode to chaincode invocations.
type ChaincodeInterest struct {
	Chaincodes           []*ChaincodeCall `protobuf:"bytes,1,rep,name=chaincodes,proto3" json:"chaincodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ChaincodeInterest) Reset()         { *m = ChaincodeInterest{} }
func (m *ChaincodeInterest) String() string { return proto.CompactTextString(m) }
func (*ChaincodeInterest) ProtoMessage()    {}
func (*ChaincodeInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed51030656d961a, []int{4}
}

func (m *ChaincodeInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeInterest.Unmarshal(m, b)
}
func (m *ChaincodeInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeInterest.Marshal(b, m, deterministic)
}
func (m *ChaincodeInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeInterest.Merge(m, src)
}
func (m *ChaincodeInterest) XXX_Size() int {
	return xxx_messageInfo_ChaincodeInterest.Size(m)
}
func (m *ChaincodeInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeInterest.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeInterest proto.InternalMessageInfo

func (m *ChaincodeInterest) GetChaincodes() []*ChaincodeCall {
	if m != nil {
		return m.Chaincodes
	}
	return nil
}

// ChaincodeCall defines a call to a chaincode.
// It may have collections that are related to the chaincode
type ChaincodeCall struct {
	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CollectionNames []string `protobuf:"bytes,2,rep,name=collection_names,json=collectionNames,proto3" json:"collection_names,omitempty"`
	NoPrivateReads  bool     `protobuf:"varint,3,opt,name=no_private_reads,json=noPrivateReads,proto3" json:"no_private_reads,omitempty"`
	NoPublicWrites  bool     `protobuf:"varint,4,opt,name=no_public_writes,json=noPublicWrites,proto3" json:"no_public_writes,omitempty"`
	// The set of signature policies associated with states in the write-set
	// that have state-based endorsement policies.
	KeyPolicies []*common.SignaturePolicyEnvelope `protobuf:"bytes,5,rep,name=key_policies,json=keyPolicies,proto3" json:"key_policies,omitempty"`
	// Indicates we wish to ignore the namespace endorsement policy
	DisregardNamespacePolicy bool     `protobuf:"varint,6,opt,name=disregard_namespace_policy,json=disregardNamespacePolicy,proto3" json:"disregard_namespace_policy,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *ChaincodeCall) Reset()         { *m = ChaincodeCall{} }
func (m *ChaincodeCall) String() string { return proto.CompactTextString(m) }
func (*ChaincodeCall) ProtoMessage()    {}
func (*ChaincodeCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ed51030656d961a, []int{5}
}

func (m *ChaincodeCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeCall.Unmarshal(m, b)
}
func (m *ChaincodeCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeCall.Marshal(b, m, deterministic)
}
func (m *ChaincodeCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeCall.Merge(m, src)
}
func (m *ChaincodeCall) XXX_Size() int {
	return xxx_messageInfo_ChaincodeCall.Size(m)
}
func (m *ChaincodeCall) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeCall.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeCall proto.InternalMessageInfo

func (m *ChaincodeCall) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeCall) GetCollectionNames() []string {
	if m != nil {
		return m.CollectionNames
	}
	return nil
}

func (m *ChaincodeCall) GetNoPrivateReads() bool {
	if m != nil {
		return m.NoPrivateReads
	}
	return false
}

func (m *ChaincodeCall) GetNoPublicWrites() bool {
	if m != nil {
		return m.NoPublicWrites
	}
	return false
}

func (m *ChaincodeCall) GetKeyPolicies() []*common.SignaturePolicyEnvelope {
	if m != nil {
		return m.KeyPolicies
	}
	return nil
}

func (m *ChaincodeCall) GetDisregardNamespacePolicy() bool {
	if m != nil {
		return m.DisregardNamespacePolicy
	}
	return false
}

func init() {
	proto.RegisterType((*ProposalResponse)(nil), "protos.ProposalResponse")
	proto.RegisterType((*Response)(nil), "protos.Response")
	proto.RegisterType((*ProposalResponsePayload)(nil), "protos.ProposalResponsePayload")
	proto.RegisterType((*Endorsement)(nil), "protos.Endorsement")
	proto.RegisterType((*ChaincodeInterest)(nil), "protos.ChaincodeInterest")
	proto.RegisterType((*ChaincodeCall)(nil), "protos.ChaincodeCall")
}

func init() { proto.RegisterFile("peer/proposal_response.proto", fileDescriptor_2ed51030656d961a) }

var fileDescriptor_2ed51030656d961a = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x6b, 0xdb, 0x3e,
	0x14, 0x25, 0xe9, 0x57, 0x72, 0x93, 0xfe, 0x7e, 0x99, 0x46, 0x37, 0x2f, 0x14, 0x1a, 0xbc, 0x97,
	0x0c, 0x5a, 0x1b, 0x3a, 0x0a, 0x7b, 0xd8, 0x53, 0x4b, 0xd9, 0xc7, 0x43, 0x09, 0xda, 0xd8, 0x60,
	0x0c, 0x82, 0x62, 0xdf, 0x3a, 0x22, 0xb6, 0x64, 0x24, 0x25, 0x9b, 0xff, 0x97, 0x3d, 0xee, 0x0f,
	0x1d, 0x96, 0x2c, 0x27, 0x5d, 0xf7, 0x64, 0xdf, 0xa3, 0x73, 0xcf, 0xbd, 0xf7, 0xe8, 0x03, 0x4e,
	0x4b, 0x44, 0x15, 0x97, 0x4a, 0x96, 0x52, 0xb3, 0x7c, 0xae, 0x50, 0x97, 0x52, 0x68, 0x8c, 0x4a,
	0x25, 0x8d, 0x24, 0x87, 0xf6, 0xa3, 0xc7, 0x67, 0x99, 0x94, 0x59, 0x8e, 0xb1, 0x0d, 0x17, 0xeb,
	0xfb, 0xd8, 0xf0, 0x02, 0xb5, 0x61, 0x45, 0xe9, 0x88, 0xe3, 0x93, 0x44, 0x16, 0x85, 0x14, 0x71,
	0x29, 0x73, 0x9e, 0x70, 0xd4, 0x0e, 0x0e, 0x7f, 0x75, 0x61, 0x34, 0x6b, 0xb4, 0x69, 0x23, 0x4d,
	0x02, 0x38, 0xda, 0xa0, 0xd2, 0x5c, 0x8a, 0xa0, 0x33, 0xe9, 0x4c, 0x0f, 0xa8, 0x0f, 0xc9, 0x1b,
	0xe8, 0xb7, 0xc2, 0x41, 0x77, 0xd2, 0x99, 0x0e, 0x2e, 0xc7, 0x91, 0x2b, 0x1d, 0xf9, 0xd2, 0xd1,
	0x67, 0xcf, 0xa0, 0x5b, 0x32, 0x39, 0x87, 0x9e, 0x6f, 0x3d, 0xd8, 0xb7, 0x89, 0x23, 0x97, 0xa1,
	0x23, 0x5f, 0x97, 0xb6, 0x8c, 0xba, 0x83, 0x92, 0x55, 0xb9, 0x64, 0x69, 0x70, 0x30, 0xe9, 0x4c,
	0x87, 0xd4, 0x87, 0xe4, 0x0a, 0x06, 0x28, 0x52, 0xa9, 0x34, 0x16, 0x28, 0x4c, 0x70, 0x68, 0xa5,
	0x9e, 0x7a, 0xa9, 0xdb, 0xed, 0x12, 0xdd, 0xe5, 0x91, 0x2b, 0xe8, 0x71, 0x61, 0x50, 0xa1, 0x36,
	0xc1, 0x91, 0xcd, 0x79, 0xe1, 0x73, 0x6e, 0x96, 0x8c, 0x8b, 0x44, 0xa6, 0xf8, 0xa1, 0x21, 0xd0,
	0x96, 0x1a, 0x7e, 0x81, 0x5e, 0xeb, 0xca, 0x33, 0x38, 0xd4, 0x86, 0x99, 0xb5, 0x6e, 0x4c, 0x69,
	0xa2, 0xba, 0xd7, 0x02, 0xb5, 0x66, 0x19, 0x5a, 0x47, 0xfa, 0xd4, 0x87, 0xbb, 0x53, 0xec, 0x3d,
	0x98, 0x22, 0xfc, 0x0e, 0xcf, 0xff, 0x76, 0x7d, 0xd6, 0x0c, 0xf8, 0x12, 0x8e, 0xdb, 0xcd, 0x5e,
	0x32, 0xbd, 0xb4, 0xd5, 0x86, 0x74, 0xe8, 0xc1, 0xf7, 0x4c, 0x2f, 0xc9, 0x29, 0xf4, 0xf1, 0xa7,
	0x41, 0x61, 0xf7, 0xa8, 0x6b, 0x09, 0x5b, 0x20, 0x7c, 0x07, 0x83, 0x1d, 0x23, 0xc8, 0x18, 0x7a,
	0x8d, 0x15, 0xaa, 0x11, 0x6b, 0xe3, 0x5a, 0x48, 0xf3, 0x4c, 0x30, 0xb3, 0x56, 0xe8, 0x85, 0x5a,
	0x20, 0xfc, 0x08, 0x4f, 0x1e, 0xb9, 0x43, 0xae, 0x00, 0x12, 0x0f, 0xd6, 0x5e, 0xec, 0x4d, 0x07,
	0x97, 0x27, 0x8f, 0xcc, 0xbc, 0x61, 0x79, 0x4e, 0x77, 0x88, 0xe1, 0xef, 0x2e, 0x1c, 0x3f, 0x58,
	0x25, 0x04, 0xf6, 0x05, 0x2b, 0xd0, 0xf6, 0xd4, 0xa7, 0xf6, 0x9f, 0xbc, 0x82, 0x51, 0x22, 0xf3,
	0x1c, 0x13, 0xc3, 0xa5, 0x98, 0xd7, 0x90, 0x0e, 0xba, 0x93, 0xbd, 0x69, 0x9f, 0xfe, 0xbf, 0xc5,
	0xef, 0x6a, 0x98, 0x4c, 0x61, 0x24, 0xe4, 0xbc, 0x54, 0x7c, 0xc3, 0x0c, 0xce, 0x15, 0xb2, 0x54,
	0x5b, 0x9b, 0x7b, 0xf4, 0x3f, 0x21, 0x67, 0x0e, 0xa6, 0x35, 0xea, 0x99, 0xeb, 0x45, 0xce, 0x93,
	0xf9, 0x0f, 0xc5, 0x0d, 0x6a, 0x7b, 0x06, 0x1d, 0xd3, 0xc2, 0x5f, 0x2d, 0x4a, 0xae, 0x61, 0xb8,
	0xc2, 0x6a, 0xee, 0x2f, 0x49, 0x70, 0x60, 0xa7, 0x3b, 0x8b, 0xdc, 0xe5, 0x89, 0x3e, 0x79, 0x67,
	0x66, 0x35, 0xa1, 0xba, 0x15, 0x1b, 0xcc, 0x65, 0x89, 0x74, 0xb0, 0xc2, 0x6a, 0xd6, 0xe4, 0x90,
	0xb7, 0x30, 0x4e, 0xb9, 0x56, 0x98, 0x31, 0x95, 0xba, 0x09, 0x4a, 0x96, 0xa0, 0xd3, 0xac, 0xec,
	0x81, 0xed, 0xd1, 0xa0, 0x65, 0xdc, 0x79, 0x82, 0x93, 0xbc, 0x5e, 0x41, 0x28, 0x55, 0x16, 0x2d,
	0xab, 0x12, 0x55, 0x8e, 0x69, 0x86, 0x2a, 0xba, 0x67, 0x0b, 0xc5, 0x13, 0xef, 0x70, 0xfd, 0x1c,
	0x5c, 0xff, 0xe3, 0xf4, 0x24, 0x2b, 0x96, 0xe1, 0xb7, 0xf3, 0x8c, 0x9b, 0xe5, 0x7a, 0x51, 0x37,
	0x1c, 0xef, 0x68, 0xc4, 0x4e, 0xe3, 0xc2, 0x69, 0x5c, 0x64, 0x32, 0xae, 0x65, 0x16, 0xee, 0xf5,
	0x78, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x2e, 0xf7, 0xb3, 0x64, 0x04, 0x00, 0x00,
}
