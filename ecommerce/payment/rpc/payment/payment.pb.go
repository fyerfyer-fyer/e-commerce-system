// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.4
// source: payment.proto

package payment

import (
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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

type PaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int64                  `protobuf:"varint,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Amount        float32                `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	PaymentMethod string                 `protobuf:"bytes,4,opt,name=paymentMethod,proto3" json:"paymentMethod,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	mi := &file_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *PaymentRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PaymentRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

type PaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     int64                  `protobuf:"varint,1,opt,name=paymentId,proto3" json:"paymentId,omitempty"`
	TransactionId string                 `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	PaymentStatus string                 `protobuf:"bytes,3,opt,name=paymentStatus,proto3" json:"paymentStatus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	mi := &file_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentResponse) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *PaymentResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PaymentResponse) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

type GetPaymentStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     int64                  `protobuf:"varint,1,opt,name=paymentId,proto3" json:"paymentId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentStatusRequest) Reset() {
	*x = GetPaymentStatusRequest{}
	mi := &file_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentStatusRequest) ProtoMessage() {}

func (x *GetPaymentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentStatusRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentStatusRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaymentStatusRequest) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

type GetPaymentStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentStatus string                 `protobuf:"bytes,1,opt,name=paymentStatus,proto3" json:"paymentStatus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentStatusResponse) Reset() {
	*x = GetPaymentStatusResponse{}
	mi := &file_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentStatusResponse) ProtoMessage() {}

func (x *GetPaymentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentStatusResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentStatusResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{4}
}

func (x *GetPaymentStatusResponse) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

type RefundRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentId     int64                  `protobuf:"varint,1,opt,name=paymentId,proto3" json:"paymentId,omitempty"`
	RefundAmount  float32                `protobuf:"fixed32,2,opt,name=refundAmount,proto3" json:"refundAmount,omitempty"`
	RefundReason  string                 `protobuf:"bytes,3,opt,name=refundReason,proto3" json:"refundReason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefundRequest) Reset() {
	*x = RefundRequest{}
	mi := &file_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundRequest) ProtoMessage() {}

func (x *RefundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundRequest.ProtoReflect.Descriptor instead.
func (*RefundRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{5}
}

func (x *RefundRequest) GetPaymentId() int64 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *RefundRequest) GetRefundAmount() float32 {
	if x != nil {
		return x.RefundAmount
	}
	return 0
}

func (x *RefundRequest) GetRefundReason() string {
	if x != nil {
		return x.RefundReason
	}
	return ""
}

type RefundResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefundStatus  string                 `protobuf:"bytes,1,opt,name=refundStatus,proto3" json:"refundStatus,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefundResponse) Reset() {
	*x = RefundResponse{}
	mi := &file_payment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefundResponse) ProtoMessage() {}

func (x *RefundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefundResponse.ProtoReflect.Descriptor instead.
func (*RefundResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{6}
}

func (x *RefundResponse) GetRefundStatus() string {
	if x != nil {
		return x.RefundStatus
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x22, 0x7b, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x37, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x75, 0x0a, 0x0d,
	0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72,
	0x65, 0x66, 0x75, 0x6e, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xdf, 0x01, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x66,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_payment_proto_goTypes = []any{
	(*Empty)(nil),                    // 0: payment.Empty
	(*PaymentRequest)(nil),           // 1: payment.PaymentRequest
	(*PaymentResponse)(nil),          // 2: payment.PaymentResponse
	(*GetPaymentStatusRequest)(nil),  // 3: payment.GetPaymentStatusRequest
	(*GetPaymentStatusResponse)(nil), // 4: payment.GetPaymentStatusResponse
	(*RefundRequest)(nil),            // 5: payment.RefundRequest
	(*RefundResponse)(nil),           // 6: payment.RefundResponse
}
var file_payment_proto_depIdxs = []int32{
	1, // 0: payment.Payment.MakePayment:input_type -> payment.PaymentRequest
	3, // 1: payment.Payment.GetPaymentStatus:input_type -> payment.GetPaymentStatusRequest
	5, // 2: payment.Payment.Refund:input_type -> payment.RefundRequest
	2, // 3: payment.Payment.MakePayment:output_type -> payment.PaymentResponse
	4, // 4: payment.Payment.GetPaymentStatus:output_type -> payment.GetPaymentStatusResponse
	6, // 5: payment.Payment.Refund:output_type -> payment.RefundResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}