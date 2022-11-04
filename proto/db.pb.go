// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: db.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Create a record in the database. Optionally include an "id" field otherwise it's set automatically.
type UpsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table    *Table `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	DataType int32  `protobuf:"varint,3,opt,name=dataType,proto3" json:"dataType,omitempty"`
	Data     []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpsertRequest) Reset() {
	*x = UpsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertRequest) ProtoMessage() {}

func (x *UpsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertRequest.ProtoReflect.Descriptor instead.
func (*UpsertRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertRequest) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *UpsertRequest) GetDataType() int32 {
	if x != nil {
		return x.DataType
	}
	return 0
}

func (x *UpsertRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpsertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of records upserted
	UpsertedCount int64 `protobuf:"varint,1,opt,name=upsertedCount,proto3" json:"upsertedCount,omitempty"`
	// Number of records matched
	MatchedCount int64 `protobuf:"varint,2,opt,name=matchedCount,proto3" json:"matchedCount,omitempty"`
}

func (x *UpsertResponse) Reset() {
	*x = UpsertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertResponse) ProtoMessage() {}

func (x *UpsertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertResponse.ProtoReflect.Descriptor instead.
func (*UpsertResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertResponse) GetUpsertedCount() int64 {
	if x != nil {
		return x.UpsertedCount
	}
	return 0
}

func (x *UpsertResponse) GetMatchedCount() int64 {
	if x != nil {
		return x.MatchedCount
	}
	return 0
}

// UpsertBinaryRequest is a request to upsert a binary record into storage.
type UpsertBinaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// table is the name of the table or collection to upsert data into.
	Table *Table `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	// binaryColumn is the name of the column to upsert binary data into.
	BinaryColumn string `protobuf:"bytes,2,opt,name=binaryColumn,proto3" json:"binaryColumn,omitempty"`
	// data  is the binary data to upsert.
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// primaryKeyMap is a map of JSON HTTP response column names to their storage analogues.
	PrimaryKeyMap map[string]string `protobuf:"bytes,4,rep,name=primaryKeyMap,proto3" json:"primaryKeyMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpsertBinaryRequest) Reset() {
	*x = UpsertBinaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertBinaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertBinaryRequest) ProtoMessage() {}

func (x *UpsertBinaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertBinaryRequest.ProtoReflect.Descriptor instead.
func (*UpsertBinaryRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertBinaryRequest) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *UpsertBinaryRequest) GetBinaryColumn() string {
	if x != nil {
		return x.BinaryColumn
	}
	return ""
}

func (x *UpsertBinaryRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpsertBinaryRequest) GetPrimaryKeyMap() map[string]string {
	if x != nil {
		return x.PrimaryKeyMap
	}
	return nil
}

// UpsertBinaryReponse is the resupose for upserting binary data into storage.
type UpsertBinaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpsertBinaryResponse) Reset() {
	*x = UpsertBinaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertBinaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertBinaryResponse) ProtoMessage() {}

func (x *UpsertBinaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertBinaryResponse.ProtoReflect.Descriptor instead.
func (*UpsertBinaryResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{3}
}

type Columns struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Columns) Reset() {
	*x = Columns{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Columns) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Columns) ProtoMessage() {}

func (x *Columns) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Columns.ProtoReflect.Descriptor instead.
func (*Columns) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{4}
}

func (x *Columns) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

type ListColumnsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColSet map[string]*Columns `protobuf:"bytes,1,rep,name=colSet,proto3" json:"colSet,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListColumnsResponse) Reset() {
	*x = ListColumnsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListColumnsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListColumnsResponse) ProtoMessage() {}

func (x *ListColumnsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListColumnsResponse.ProtoReflect.Descriptor instead.
func (*ListColumnsResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{5}
}

func (x *ListColumnsResponse) GetColSet() map[string]*Columns {
	if x != nil {
		return x.ColSet
	}
	return nil
}

type PrimaryKeys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PrimaryKeys) Reset() {
	*x = PrimaryKeys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryKeys) ProtoMessage() {}

func (x *PrimaryKeys) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryKeys.ProtoReflect.Descriptor instead.
func (*PrimaryKeys) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{6}
}

func (x *PrimaryKeys) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

type ListPrimaryKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PKSet map[string]*PrimaryKeys `protobuf:"bytes,3,rep,name=PKSet,proto3" json:"PKSet,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListPrimaryKeysResponse) Reset() {
	*x = ListPrimaryKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPrimaryKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPrimaryKeysResponse) ProtoMessage() {}

func (x *ListPrimaryKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPrimaryKeysResponse.ProtoReflect.Descriptor instead.
func (*ListPrimaryKeysResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{7}
}

func (x *ListPrimaryKeysResponse) GetPKSet() map[string]*PrimaryKeys {
	if x != nil {
		return x.PKSet
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size     int64  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Database string `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{8}
}

func (x *Table) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

type ListTablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableSet map[string]*Table `protobuf:"bytes,1,rep,name=tableSet,proto3" json:"tableSet,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListTablesResponse) Reset() {
	*x = ListTablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTablesResponse) ProtoMessage() {}

func (x *ListTablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTablesResponse.ProtoReflect.Descriptor instead.
func (*ListTablesResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{9}
}

func (x *ListTablesResponse) GetTableSet() map[string]*Table {
	if x != nil {
		return x.TableSet
	}
	return nil
}

// Read data from a table. Lookup can be by ID or via querying any field in the record.
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional table name. Defaults to 'default'
	ReaderBuilder []byte           `protobuf:"bytes,1,opt,name=readerBuilder,proto3" json:"readerBuilder,omitempty"`
	Required      *structpb.Struct `protobuf:"bytes,2,opt,name=required,proto3" json:"required,omitempty"`
	Options       *structpb.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	Table         *Table           `protobuf:"bytes,4,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{10}
}

func (x *ReadRequest) GetReaderBuilder() []byte {
	if x != nil {
		return x.ReaderBuilder
	}
	return nil
}

func (x *ReadRequest) GetRequired() *structpb.Struct {
	if x != nil {
		return x.Required
	}
	return nil
}

func (x *ReadRequest) GetOptions() *structpb.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ReadRequest) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON encoded records
	Records []*structpb.Struct `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{11}
}

func (x *ReadResponse) GetRecords() []*structpb.Struct {
	if x != nil {
		return x.Records
	}
	return nil
}

type TruncateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tables is a slice of tables to truncate.
	Tables []*Table `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *TruncateRequest) Reset() {
	*x = TruncateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TruncateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TruncateRequest) ProtoMessage() {}

func (x *TruncateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TruncateRequest.ProtoReflect.Descriptor instead.
func (*TruncateRequest) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{12}
}

func (x *TruncateRequest) GetTables() []*Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

type TruncateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DeletedCount of records deleted
	DeletedCount int32 `protobuf:"varint,1,opt,name=deletedCount,proto3" json:"deletedCount,omitempty"`
}

func (x *TruncateResponse) Reset() {
	*x = TruncateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TruncateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TruncateResponse) ProtoMessage() {}

func (x *TruncateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_db_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TruncateResponse.ProtoReflect.Descriptor instead.
func (*TruncateResponse) Descriptor() ([]byte, []int) {
	return file_db_proto_rawDescGZIP(), []int{13}
}

func (x *TruncateResponse) GetDeletedCount() int32 {
	if x != nil {
		return x.DeletedCount
	}
	return 0
}

var File_db_proto protoreflect.FileDescriptor

var file_db_proto_rawDesc = []byte{
	0x0a, 0x08, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x63, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x75,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x88, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x53, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x4d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x1a, 0x40, 0x0a, 0x12, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x16, 0x0a, 0x14, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x63, 0x6f,
	0x6c, 0x53, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x53, 0x65, 0x74, 0x1a, 0x49, 0x0a, 0x0b, 0x43, 0x6f,
	0x6c, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x21, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x50, 0x4b, 0x53, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x50, 0x4b, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x50, 0x4b, 0x53, 0x65, 0x74, 0x1a, 0x4c, 0x0a, 0x0a, 0x50, 0x4b, 0x53, 0x65, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x4b, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x22, 0xa4, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x1a, 0x49, 0x0a, 0x0d,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbf, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a,
	0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x41, 0x0a, 0x0c, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x37, 0x0a, 0x0f,
	0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x06, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x10, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_proto_rawDescOnce sync.Once
	file_db_proto_rawDescData = file_db_proto_rawDesc
)

func file_db_proto_rawDescGZIP() []byte {
	file_db_proto_rawDescOnce.Do(func() {
		file_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_proto_rawDescData)
	})
	return file_db_proto_rawDescData
}

var file_db_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_db_proto_goTypes = []interface{}{
	(*UpsertRequest)(nil),           // 0: proto.UpsertRequest
	(*UpsertResponse)(nil),          // 1: proto.UpsertResponse
	(*UpsertBinaryRequest)(nil),     // 2: proto.UpsertBinaryRequest
	(*UpsertBinaryResponse)(nil),    // 3: proto.UpsertBinaryResponse
	(*Columns)(nil),                 // 4: proto.Columns
	(*ListColumnsResponse)(nil),     // 5: proto.ListColumnsResponse
	(*PrimaryKeys)(nil),             // 6: proto.PrimaryKeys
	(*ListPrimaryKeysResponse)(nil), // 7: proto.ListPrimaryKeysResponse
	(*Table)(nil),                   // 8: proto.Table
	(*ListTablesResponse)(nil),      // 9: proto.ListTablesResponse
	(*ReadRequest)(nil),             // 10: proto.ReadRequest
	(*ReadResponse)(nil),            // 11: proto.ReadResponse
	(*TruncateRequest)(nil),         // 12: proto.TruncateRequest
	(*TruncateResponse)(nil),        // 13: proto.TruncateResponse
	nil,                             // 14: proto.UpsertBinaryRequest.PrimaryKeyMapEntry
	nil,                             // 15: proto.ListColumnsResponse.ColSetEntry
	nil,                             // 16: proto.ListPrimaryKeysResponse.PKSetEntry
	nil,                             // 17: proto.ListTablesResponse.TableSetEntry
	(*structpb.Struct)(nil),         // 18: google.protobuf.Struct
}
var file_db_proto_depIdxs = []int32{
	8,  // 0: proto.UpsertRequest.table:type_name -> proto.Table
	8,  // 1: proto.UpsertBinaryRequest.table:type_name -> proto.Table
	14, // 2: proto.UpsertBinaryRequest.primaryKeyMap:type_name -> proto.UpsertBinaryRequest.PrimaryKeyMapEntry
	15, // 3: proto.ListColumnsResponse.colSet:type_name -> proto.ListColumnsResponse.ColSetEntry
	16, // 4: proto.ListPrimaryKeysResponse.PKSet:type_name -> proto.ListPrimaryKeysResponse.PKSetEntry
	17, // 5: proto.ListTablesResponse.tableSet:type_name -> proto.ListTablesResponse.TableSetEntry
	18, // 6: proto.ReadRequest.required:type_name -> google.protobuf.Struct
	18, // 7: proto.ReadRequest.options:type_name -> google.protobuf.Struct
	8,  // 8: proto.ReadRequest.table:type_name -> proto.Table
	18, // 9: proto.ReadResponse.records:type_name -> google.protobuf.Struct
	8,  // 10: proto.TruncateRequest.tables:type_name -> proto.Table
	4,  // 11: proto.ListColumnsResponse.ColSetEntry.value:type_name -> proto.Columns
	6,  // 12: proto.ListPrimaryKeysResponse.PKSetEntry.value:type_name -> proto.PrimaryKeys
	8,  // 13: proto.ListTablesResponse.TableSetEntry.value:type_name -> proto.Table
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_db_proto_init() }
func file_db_proto_init() {
	if File_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertBinaryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertBinaryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Columns); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListColumnsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryKeys); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPrimaryKeysResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTablesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TruncateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TruncateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_proto_goTypes,
		DependencyIndexes: file_db_proto_depIdxs,
		MessageInfos:      file_db_proto_msgTypes,
	}.Build()
	File_db_proto = out.File
	file_db_proto_rawDesc = nil
	file_db_proto_goTypes = nil
	file_db_proto_depIdxs = nil
}
