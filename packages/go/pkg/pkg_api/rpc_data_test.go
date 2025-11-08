package pkgapi_test

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderStatus int32

const (
	OrderStatus_ORDER_STATUS_UNSPECIFIED OrderStatus = 0
	OrderStatus_ORDER_STATUS_PENDING     OrderStatus = 1
	OrderStatus_ORDER_STATUS_COMPLETED   OrderStatus = 2
	OrderStatus_ORDER_STATUS_CANCELLED   OrderStatus = 3
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "ORDER_STATUS_UNSPECIFIED",
		1: "ORDER_STATUS_PENDING",
		2: "ORDER_STATUS_COMPLETED",
		3: "ORDER_STATUS_CANCELLED",
	}
	OrderStatus_value = map[string]int32{
		"ORDER_STATUS_UNSPECIFIED": 0,
		"ORDER_STATUS_PENDING":     1,
		"ORDER_STATUS_COMPLETED":   2,
		"ORDER_STATUS_CANCELLED":   3,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_saas_v1_entity_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_saas_v1_entity_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_saas_v1_entity_proto_rawDescGZIP(), []int{0}
}

type CustomerProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerProto) Reset() {
	*x = CustomerProto{}
	mi := &file_saas_v1_entity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerProto) ProtoMessage() {}

func (x *CustomerProto) ProtoReflect() protoreflect.Message {
	mi := &file_saas_v1_entity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerProto.ProtoReflect.Descriptor instead.
func (*CustomerProto) Descriptor() ([]byte, []int) {
	return file_saas_v1_entity_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerProto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CustomerProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerProto) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LineItemProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LineItemProto) Reset() {
	*x = LineItemProto{}
	mi := &file_saas_v1_entity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LineItemProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItemProto) ProtoMessage() {}

func (x *LineItemProto) ProtoReflect() protoreflect.Message {
	mi := &file_saas_v1_entity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItemProto.ProtoReflect.Descriptor instead.
func (*LineItemProto) Descriptor() ([]byte, []int) {
	return file_saas_v1_entity_proto_rawDescGZIP(), []int{1}
}

func (x *LineItemProto) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *LineItemProto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LineItemProto) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *LineItemProto) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type OrderProto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Customer      *CustomerProto         `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	LineItems     []*LineItemProto       `protobuf:"bytes,3,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
	TotalAmount   float64                `protobuf:"fixed64,4,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Status        OrderStatus            `protobuf:"varint,6,opt,name=status,proto3,enum=saas.v1.OrderStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderProto) Reset() {
	*x = OrderProto{}
	mi := &file_saas_v1_entity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProto) ProtoMessage() {}

func (x *OrderProto) ProtoReflect() protoreflect.Message {
	mi := &file_saas_v1_entity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProto.ProtoReflect.Descriptor instead.
func (*OrderProto) Descriptor() ([]byte, []int) {
	return file_saas_v1_entity_proto_rawDescGZIP(), []int{2}
}

func (x *OrderProto) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderProto) GetCustomer() *CustomerProto {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *OrderProto) GetLineItems() []*LineItemProto {
	if x != nil {
		return x.LineItems
	}
	return nil
}

func (x *OrderProto) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *OrderProto) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OrderProto) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_ORDER_STATUS_UNSPECIFIED
}

var File_saas_v1_entity_proto protoreflect.FileDescriptor

const file_saas_v1_entity_proto_rawDesc = "" +
	"\n" +
	"\x13saas/v1/entity.proto\x12\x06saas.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"D\n" +
	"\bCustomer\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\"o\n" +
	"\bLineItem\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bquantity\x18\x03 \x01(\x05R\bquantity\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\"\x8c\x02\n" +
	"\x05Order\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12,\n" +
	"\bcustomer\x18\x02 \x01(\v2\x10.saas.v1.CustomerR\bcustomer\x12/\n" +
	"\n" +
	"line_items\x18\x03 \x03(\v2\x10.saas.v1.LineItemR\tlineItems\x12!\n" +
	"\ftotal_amount\x18\x04 \x01(\x01R\vtotalAmount\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12+\n" +
	"\x06status\x18\x06 \x01(\x0e2\x13.saas.v1.OrderStatusR\x06status*}\n" +
	"\vOrderStatus\x12\x1c\n" +
	"\x18ORDER_STATUS_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14ORDER_STATUS_PENDING\x10\x01\x12\x1a\n" +
	"\x16ORDER_STATUS_COMPLETED\x10\x02\x12\x1a\n" +
	"\x16ORDER_STATUS_CANCELLED\x10\x03B\x8e\x01\n" +
	"\n" +
	"com.saas.v1B\vEntityProtoP\x01Z:github.com/vtuanjs/saas_management/packages/go/buf/saas/v1\xa2\x02\x03IXX\xaa\x02\x06Iam.V1\xca\x02\x06Iam\\V1\xe2\x02\x12Iam\\V1\\GPBMetadata\xea\x02\aIam::V1b\x06proto3"

var (
	file_saas_v1_entity_proto_rawDescOnce sync.Once
	file_saas_v1_entity_proto_rawDescData []byte
)

func file_saas_v1_entity_proto_rawDescGZIP() []byte {
	file_saas_v1_entity_proto_rawDescOnce.Do(func() {
		file_saas_v1_entity_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_saas_v1_entity_proto_rawDesc), len(file_saas_v1_entity_proto_rawDesc)))
	})
	return file_saas_v1_entity_proto_rawDescData
}

var file_saas_v1_entity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_saas_v1_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_saas_v1_entity_proto_goTypes = []any{
	(OrderStatus)(0),              // 0: saas.v1.OrderStatus
	(*CustomerProto)(nil),         // 1: saas.v1.CustomerProto
	(*LineItemProto)(nil),         // 2: saas.v1.LineItemProto
	(*OrderProto)(nil),            // 3: saas.v1.OrderProto
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_saas_v1_entity_proto_depIdxs = []int32{
	1, // 0: saas.v1.OrderProto.customer:type_name -> saas.v1.CustomerProto
	2, // 1: saas.v1.OrderProto.line_items:type_name -> saas.v1.LineItemProto
	4, // 2: saas.v1.OrderProto.created_at:type_name -> google.protobuf.Timestamp
	0, // 3: saas.v1.OrderProto.status:type_name -> saas.v1.OrderStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_saas_v1_entity_proto_init() }
func file_saas_v1_entity_proto_init() {
	if File_saas_v1_entity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_saas_v1_entity_proto_rawDesc), len(file_saas_v1_entity_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_saas_v1_entity_proto_goTypes,
		DependencyIndexes: file_saas_v1_entity_proto_depIdxs,
		EnumInfos:         file_saas_v1_entity_proto_enumTypes,
		MessageInfos:      file_saas_v1_entity_proto_msgTypes,
	}.Build()
	File_saas_v1_entity_proto = out.File
	file_saas_v1_entity_proto_goTypes = nil
	file_saas_v1_entity_proto_depIdxs = nil
}
