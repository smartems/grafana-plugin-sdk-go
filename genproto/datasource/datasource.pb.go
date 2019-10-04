// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datasource.proto

package datasource

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RowValue_Kind int32

const (
	// Field type null.
	RowValue_TYPE_NULL RowValue_Kind = 0
	// Field type double.
	RowValue_TYPE_DOUBLE RowValue_Kind = 1
	// Field type int64.
	RowValue_TYPE_INT64 RowValue_Kind = 2
	// Field type bool.
	RowValue_TYPE_BOOL RowValue_Kind = 3
	// Field type string.
	RowValue_TYPE_STRING RowValue_Kind = 4
	// Field type bytes.
	RowValue_TYPE_BYTES RowValue_Kind = 5
)

var RowValue_Kind_name = map[int32]string{
	0: "TYPE_NULL",
	1: "TYPE_DOUBLE",
	2: "TYPE_INT64",
	3: "TYPE_BOOL",
	4: "TYPE_STRING",
	5: "TYPE_BYTES",
}

var RowValue_Kind_value = map[string]int32{
	"TYPE_NULL":   0,
	"TYPE_DOUBLE": 1,
	"TYPE_INT64":  2,
	"TYPE_BOOL":   3,
	"TYPE_STRING": 4,
	"TYPE_BYTES":  5,
}

func (x RowValue_Kind) String() string {
	return proto.EnumName(RowValue_Kind_name, int32(x))
}

func (RowValue_Kind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{8, 0}
}

type DatasourceRequest struct {
	TimeRange            *TimeRange      `protobuf:"bytes,1,opt,name=timeRange,proto3" json:"timeRange,omitempty"`
	Datasource           *DatasourceInfo `protobuf:"bytes,2,opt,name=datasource,proto3" json:"datasource,omitempty"`
	Queries              []*Query        `protobuf:"bytes,3,rep,name=queries,proto3" json:"queries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DatasourceRequest) Reset()         { *m = DatasourceRequest{} }
func (m *DatasourceRequest) String() string { return proto.CompactTextString(m) }
func (*DatasourceRequest) ProtoMessage()    {}
func (*DatasourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{0}
}

func (m *DatasourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasourceRequest.Unmarshal(m, b)
}
func (m *DatasourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasourceRequest.Marshal(b, m, deterministic)
}
func (m *DatasourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasourceRequest.Merge(m, src)
}
func (m *DatasourceRequest) XXX_Size() int {
	return xxx_messageInfo_DatasourceRequest.Size(m)
}
func (m *DatasourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatasourceRequest proto.InternalMessageInfo

func (m *DatasourceRequest) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *DatasourceRequest) GetDatasource() *DatasourceInfo {
	if m != nil {
		return m.Datasource
	}
	return nil
}

func (m *DatasourceRequest) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Query struct {
	RefId                string   `protobuf:"bytes,1,opt,name=refId,proto3" json:"refId,omitempty"`
	MaxDataPoints        int64    `protobuf:"varint,2,opt,name=maxDataPoints,proto3" json:"maxDataPoints,omitempty"`
	IntervalMs           int64    `protobuf:"varint,3,opt,name=intervalMs,proto3" json:"intervalMs,omitempty"`
	ModelJson            string   `protobuf:"bytes,4,opt,name=modelJson,proto3" json:"modelJson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{1}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *Query) GetMaxDataPoints() int64 {
	if m != nil {
		return m.MaxDataPoints
	}
	return 0
}

func (m *Query) GetIntervalMs() int64 {
	if m != nil {
		return m.IntervalMs
	}
	return 0
}

func (m *Query) GetModelJson() string {
	if m != nil {
		return m.ModelJson
	}
	return ""
}

type TimeRange struct {
	FromRaw              string   `protobuf:"bytes,1,opt,name=fromRaw,proto3" json:"fromRaw,omitempty"`
	ToRaw                string   `protobuf:"bytes,2,opt,name=toRaw,proto3" json:"toRaw,omitempty"`
	FromEpochMs          int64    `protobuf:"varint,3,opt,name=fromEpochMs,proto3" json:"fromEpochMs,omitempty"`
	ToEpochMs            int64    `protobuf:"varint,4,opt,name=toEpochMs,proto3" json:"toEpochMs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{2}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetFromRaw() string {
	if m != nil {
		return m.FromRaw
	}
	return ""
}

func (m *TimeRange) GetToRaw() string {
	if m != nil {
		return m.ToRaw
	}
	return ""
}

func (m *TimeRange) GetFromEpochMs() int64 {
	if m != nil {
		return m.FromEpochMs
	}
	return 0
}

func (m *TimeRange) GetToEpochMs() int64 {
	if m != nil {
		return m.ToEpochMs
	}
	return 0
}

type DatasourceResponse struct {
	Results              []*QueryResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DatasourceResponse) Reset()         { *m = DatasourceResponse{} }
func (m *DatasourceResponse) String() string { return proto.CompactTextString(m) }
func (*DatasourceResponse) ProtoMessage()    {}
func (*DatasourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{3}
}

func (m *DatasourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasourceResponse.Unmarshal(m, b)
}
func (m *DatasourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasourceResponse.Marshal(b, m, deterministic)
}
func (m *DatasourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasourceResponse.Merge(m, src)
}
func (m *DatasourceResponse) XXX_Size() int {
	return xxx_messageInfo_DatasourceResponse.Size(m)
}
func (m *DatasourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatasourceResponse proto.InternalMessageInfo

func (m *DatasourceResponse) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type QueryResult struct {
	Error                string        `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	RefId                string        `protobuf:"bytes,2,opt,name=refId,proto3" json:"refId,omitempty"`
	MetaJson             string        `protobuf:"bytes,3,opt,name=metaJson,proto3" json:"metaJson,omitempty"`
	Series               []*TimeSeries `protobuf:"bytes,4,rep,name=series,proto3" json:"series,omitempty"`
	Tables               []*Table      `protobuf:"bytes,5,rep,name=tables,proto3" json:"tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QueryResult) Reset()         { *m = QueryResult{} }
func (m *QueryResult) String() string { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()    {}
func (*QueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{4}
}

func (m *QueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResult.Unmarshal(m, b)
}
func (m *QueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResult.Marshal(b, m, deterministic)
}
func (m *QueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResult.Merge(m, src)
}
func (m *QueryResult) XXX_Size() int {
	return xxx_messageInfo_QueryResult.Size(m)
}
func (m *QueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResult proto.InternalMessageInfo

func (m *QueryResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *QueryResult) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *QueryResult) GetMetaJson() string {
	if m != nil {
		return m.MetaJson
	}
	return ""
}

func (m *QueryResult) GetSeries() []*TimeSeries {
	if m != nil {
		return m.Series
	}
	return nil
}

func (m *QueryResult) GetTables() []*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

type Table struct {
	Columns              []*TableColumn `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows                 []*TableRow    `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{5}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetColumns() []*TableColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Table) GetRows() []*TableRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type TableColumn struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableColumn) Reset()         { *m = TableColumn{} }
func (m *TableColumn) String() string { return proto.CompactTextString(m) }
func (*TableColumn) ProtoMessage()    {}
func (*TableColumn) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{6}
}

func (m *TableColumn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableColumn.Unmarshal(m, b)
}
func (m *TableColumn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableColumn.Marshal(b, m, deterministic)
}
func (m *TableColumn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableColumn.Merge(m, src)
}
func (m *TableColumn) XXX_Size() int {
	return xxx_messageInfo_TableColumn.Size(m)
}
func (m *TableColumn) XXX_DiscardUnknown() {
	xxx_messageInfo_TableColumn.DiscardUnknown(m)
}

var xxx_messageInfo_TableColumn proto.InternalMessageInfo

func (m *TableColumn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TableRow struct {
	Values               []*RowValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TableRow) Reset()         { *m = TableRow{} }
func (m *TableRow) String() string { return proto.CompactTextString(m) }
func (*TableRow) ProtoMessage()    {}
func (*TableRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{7}
}

func (m *TableRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableRow.Unmarshal(m, b)
}
func (m *TableRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableRow.Marshal(b, m, deterministic)
}
func (m *TableRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableRow.Merge(m, src)
}
func (m *TableRow) XXX_Size() int {
	return xxx_messageInfo_TableRow.Size(m)
}
func (m *TableRow) XXX_DiscardUnknown() {
	xxx_messageInfo_TableRow.DiscardUnknown(m)
}

var xxx_messageInfo_TableRow proto.InternalMessageInfo

func (m *TableRow) GetValues() []*RowValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type RowValue struct {
	Kind                 RowValue_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=models.RowValue_Kind" json:"kind,omitempty"`
	DoubleValue          float64       `protobuf:"fixed64,2,opt,name=doubleValue,proto3" json:"doubleValue,omitempty"`
	Int64Value           int64         `protobuf:"varint,3,opt,name=int64Value,proto3" json:"int64Value,omitempty"`
	BoolValue            bool          `protobuf:"varint,4,opt,name=boolValue,proto3" json:"boolValue,omitempty"`
	StringValue          string        `protobuf:"bytes,5,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
	BytesValue           []byte        `protobuf:"bytes,6,opt,name=bytesValue,proto3" json:"bytesValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RowValue) Reset()         { *m = RowValue{} }
func (m *RowValue) String() string { return proto.CompactTextString(m) }
func (*RowValue) ProtoMessage()    {}
func (*RowValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{8}
}

func (m *RowValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RowValue.Unmarshal(m, b)
}
func (m *RowValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RowValue.Marshal(b, m, deterministic)
}
func (m *RowValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RowValue.Merge(m, src)
}
func (m *RowValue) XXX_Size() int {
	return xxx_messageInfo_RowValue.Size(m)
}
func (m *RowValue) XXX_DiscardUnknown() {
	xxx_messageInfo_RowValue.DiscardUnknown(m)
}

var xxx_messageInfo_RowValue proto.InternalMessageInfo

func (m *RowValue) GetKind() RowValue_Kind {
	if m != nil {
		return m.Kind
	}
	return RowValue_TYPE_NULL
}

func (m *RowValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *RowValue) GetInt64Value() int64 {
	if m != nil {
		return m.Int64Value
	}
	return 0
}

func (m *RowValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *RowValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *RowValue) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

type DatasourceInfo struct {
	Id                      int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrgId                   int64             `protobuf:"varint,2,opt,name=orgId,proto3" json:"orgId,omitempty"`
	Name                    string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                    string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Url                     string            `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	JsonData                string            `protobuf:"bytes,6,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
	DecryptedSecureJsonData map[string]string `protobuf:"bytes,7,rep,name=decryptedSecureJsonData,proto3" json:"decryptedSecureJsonData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *DatasourceInfo) Reset()         { *m = DatasourceInfo{} }
func (m *DatasourceInfo) String() string { return proto.CompactTextString(m) }
func (*DatasourceInfo) ProtoMessage()    {}
func (*DatasourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{9}
}

func (m *DatasourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasourceInfo.Unmarshal(m, b)
}
func (m *DatasourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasourceInfo.Marshal(b, m, deterministic)
}
func (m *DatasourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasourceInfo.Merge(m, src)
}
func (m *DatasourceInfo) XXX_Size() int {
	return xxx_messageInfo_DatasourceInfo.Size(m)
}
func (m *DatasourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DatasourceInfo proto.InternalMessageInfo

func (m *DatasourceInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DatasourceInfo) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *DatasourceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatasourceInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DatasourceInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DatasourceInfo) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *DatasourceInfo) GetDecryptedSecureJsonData() map[string]string {
	if m != nil {
		return m.DecryptedSecureJsonData
	}
	return nil
}

type TimeSeries struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tags                 map[string]string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Points               []*Point          `protobuf:"bytes,3,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TimeSeries) Reset()         { *m = TimeSeries{} }
func (m *TimeSeries) String() string { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()    {}
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{10}
}

func (m *TimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeries.Unmarshal(m, b)
}
func (m *TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeries.Marshal(b, m, deterministic)
}
func (m *TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeries.Merge(m, src)
}
func (m *TimeSeries) XXX_Size() int {
	return xxx_messageInfo_TimeSeries.Size(m)
}
func (m *TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeries proto.InternalMessageInfo

func (m *TimeSeries) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeSeries) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type Point struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb096a9d85d590d2, []int{11}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Point) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("models.RowValue_Kind", RowValue_Kind_name, RowValue_Kind_value)
	proto.RegisterType((*DatasourceRequest)(nil), "models.DatasourceRequest")
	proto.RegisterType((*Query)(nil), "models.Query")
	proto.RegisterType((*TimeRange)(nil), "models.TimeRange")
	proto.RegisterType((*DatasourceResponse)(nil), "models.DatasourceResponse")
	proto.RegisterType((*QueryResult)(nil), "models.QueryResult")
	proto.RegisterType((*Table)(nil), "models.Table")
	proto.RegisterType((*TableColumn)(nil), "models.TableColumn")
	proto.RegisterType((*TableRow)(nil), "models.TableRow")
	proto.RegisterType((*RowValue)(nil), "models.RowValue")
	proto.RegisterType((*DatasourceInfo)(nil), "models.DatasourceInfo")
	proto.RegisterMapType((map[string]string)(nil), "models.DatasourceInfo.DecryptedSecureJsonDataEntry")
	proto.RegisterType((*TimeSeries)(nil), "models.TimeSeries")
	proto.RegisterMapType((map[string]string)(nil), "models.TimeSeries.TagsEntry")
	proto.RegisterType((*Point)(nil), "models.Point")
}

func init() { proto.RegisterFile("datasource.proto", fileDescriptor_bb096a9d85d590d2) }

var fileDescriptor_bb096a9d85d590d2 = []byte{
	// 841 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdb, 0x8e, 0x23, 0x35,
	0x10, 0xa5, 0x2f, 0xc9, 0xa4, 0x2b, 0x3b, 0xa1, 0xd7, 0xdc, 0x42, 0x34, 0x42, 0xa1, 0xb5, 0x88,
	0x80, 0x44, 0x40, 0xd9, 0xd1, 0x80, 0x40, 0xe2, 0x21, 0x3b, 0x11, 0xca, 0x10, 0x66, 0x06, 0x27,
	0x8b, 0xb4, 0x08, 0x09, 0x75, 0xd2, 0x4e, 0x68, 0xb6, 0xbb, 0x9d, 0xb5, 0xdd, 0x13, 0x22, 0x9e,
	0xf8, 0x18, 0x24, 0x9e, 0xf9, 0x01, 0x1e, 0xf9, 0x2d, 0xe4, 0xea, 0x6b, 0x76, 0x02, 0x12, 0x6f,
	0xae, 0x73, 0x4e, 0x95, 0xcb, 0x55, 0x65, 0x1b, 0xdc, 0xc0, 0x57, 0xbe, 0xe4, 0xa9, 0x58, 0xb1,
	0xe1, 0x56, 0x70, 0xc5, 0x49, 0x33, 0xe6, 0x01, 0x8b, 0xa4, 0xf7, 0xbb, 0x01, 0x0f, 0x2f, 0x4b,
	0x92, 0xb2, 0x17, 0x29, 0x93, 0x8a, 0x7c, 0x0c, 0x8e, 0x0a, 0x63, 0x46, 0xfd, 0x64, 0xc3, 0xba,
	0x46, 0xdf, 0x18, 0xb4, 0x47, 0x0f, 0x87, 0x99, 0xc7, 0x70, 0x51, 0x10, 0xb4, 0xd2, 0x90, 0x0b,
	0x80, 0x6a, 0x8b, 0xae, 0x89, 0x1e, 0x6f, 0x16, 0x1e, 0x55, 0xfc, 0x69, 0xb2, 0xe6, 0xb4, 0xa6,
	0x24, 0xef, 0xc3, 0xc9, 0x8b, 0x94, 0x89, 0x90, 0xc9, 0xae, 0xd5, 0xb7, 0x06, 0xed, 0xd1, 0x69,
	0xe1, 0xf4, 0x6d, 0xca, 0xc4, 0x9e, 0x16, 0xac, 0xf7, 0x9b, 0x01, 0x0d, 0x84, 0xc8, 0xeb, 0xd0,
	0x10, 0x6c, 0x3d, 0x0d, 0x30, 0x2f, 0x87, 0x66, 0x06, 0x79, 0x04, 0xa7, 0xb1, 0xff, 0x8b, 0xde,
	0xe9, 0x96, 0x87, 0x89, 0x92, 0x98, 0x83, 0x45, 0x0f, 0x41, 0xf2, 0x0e, 0x40, 0x98, 0x28, 0x26,
	0xee, 0xfc, 0xe8, 0x1b, 0xbd, 0xa3, 0x96, 0xd4, 0x10, 0x72, 0x06, 0x0e, 0x6e, 0x7f, 0x25, 0x79,
	0xd2, 0xb5, 0x31, 0x7e, 0x05, 0x78, 0xbf, 0x82, 0x53, 0x1e, 0x9e, 0x74, 0xe1, 0x64, 0x2d, 0x78,
	0x4c, 0xfd, 0x5d, 0x9e, 0x48, 0x61, 0xea, 0x04, 0x15, 0xd7, 0xb8, 0x99, 0x25, 0x88, 0x06, 0xe9,
	0x43, 0x5b, 0x0b, 0x26, 0x5b, 0xbe, 0xfa, 0xa9, 0xdc, 0xbb, 0x0e, 0xe9, 0xcd, 0x15, 0x2f, 0x78,
	0x1b, 0xf9, 0x0a, 0xf0, 0x9e, 0x00, 0xa9, 0xf7, 0x49, 0x6e, 0x79, 0x22, 0x19, 0xf9, 0x08, 0x4e,
	0x04, 0x93, 0x69, 0xa4, 0x64, 0xd7, 0xc0, 0xfa, 0xbd, 0x76, 0x58, 0x3f, 0xe4, 0x68, 0xa1, 0xf1,
	0xfe, 0x30, 0xa0, 0x5d, 0x23, 0x74, 0xaa, 0x4c, 0x08, 0x2e, 0x8a, 0x5a, 0xa2, 0x51, 0x55, 0xd8,
	0xac, 0x57, 0xb8, 0x07, 0xad, 0x98, 0x29, 0x1f, 0x4b, 0x63, 0x21, 0x51, 0xda, 0xe4, 0x43, 0x68,
	0xca, 0xac, 0x8b, 0x36, 0x66, 0x41, 0xea, 0xc3, 0x32, 0x47, 0x86, 0xe6, 0x0a, 0xf2, 0x1e, 0x34,
	0x95, 0xbf, 0x8c, 0x98, 0xec, 0x36, 0x0e, 0x3b, 0xbe, 0xd0, 0x28, 0xcd, 0x49, 0xef, 0x07, 0x68,
	0x20, 0xa0, 0x8f, 0xb8, 0xe2, 0x51, 0x1a, 0x27, 0xf7, 0x8e, 0x88, 0xfc, 0x13, 0xe4, 0x68, 0xa1,
	0x21, 0x8f, 0xc0, 0x16, 0x7c, 0xa7, 0xfb, 0xaf, 0xb5, 0xee, 0x61, 0x70, 0xbe, 0xa3, 0xc8, 0x7a,
	0xef, 0x42, 0xbb, 0xe6, 0x4d, 0x08, 0xd8, 0x89, 0x1f, 0xb3, 0xbc, 0x0c, 0xb8, 0xf6, 0xce, 0xa1,
	0x55, 0x38, 0x91, 0x01, 0x34, 0xef, 0xfc, 0x28, 0x65, 0x45, 0x0a, 0x65, 0x58, 0xca, 0x77, 0xdf,
	0x69, 0x82, 0xe6, 0xbc, 0xf7, 0x97, 0x09, 0xad, 0x02, 0x24, 0x1f, 0x80, 0xfd, 0x3c, 0x4c, 0xb2,
	0x49, 0xed, 0x8c, 0xde, 0x78, 0xd9, 0x69, 0xf8, 0x75, 0x98, 0x04, 0x14, 0x25, 0x7a, 0x3c, 0x02,
	0x9e, 0x2e, 0x23, 0x86, 0x0c, 0x56, 0xde, 0xa0, 0x75, 0x28, 0x9f, 0xdd, 0x8b, 0xf3, 0x4c, 0x50,
	0xcd, 0x6e, 0x8e, 0xe8, 0xf1, 0x59, 0x72, 0x1e, 0x65, 0xb4, 0x1e, 0x9f, 0x16, 0xad, 0x00, 0x1d,
	0x5f, 0x2a, 0x11, 0x26, 0x9b, 0x8c, 0x6f, 0xe0, 0x41, 0xeb, 0x90, 0x8e, 0xbf, 0xdc, 0x2b, 0x26,
	0x33, 0x41, 0xb3, 0x6f, 0x0c, 0x1e, 0xd0, 0x1a, 0xe2, 0xad, 0xc1, 0xd6, 0xf9, 0x92, 0x53, 0x70,
	0x16, 0xcf, 0x6e, 0x27, 0x3f, 0x5e, 0x3f, 0x9d, 0xcd, 0xdc, 0x57, 0xc8, 0xab, 0xd0, 0x46, 0xf3,
	0xf2, 0xe6, 0xe9, 0x78, 0x36, 0x71, 0x0d, 0xd2, 0x01, 0x40, 0x60, 0x7a, 0xbd, 0xb8, 0x38, 0x77,
	0xcd, 0x52, 0x3f, 0xbe, 0xb9, 0x99, 0xb9, 0x56, 0xa9, 0x9f, 0x2f, 0xe8, 0xf4, 0xfa, 0x2b, 0xd7,
	0x2e, 0xf5, 0xe3, 0x67, 0x8b, 0xc9, 0xdc, 0x6d, 0x78, 0x7f, 0x9b, 0xd0, 0x39, 0x7c, 0x31, 0x48,
	0x07, 0xcc, 0x30, 0xab, 0xa2, 0x45, 0xcd, 0x30, 0xd0, 0x03, 0xca, 0xc5, 0x26, 0x1f, 0x50, 0x8b,
	0x66, 0x46, 0xd9, 0x44, 0xab, 0x6a, 0xa2, 0xc6, 0xd4, 0x7e, 0xcb, 0xf2, 0xbb, 0x8c, 0x6b, 0xe2,
	0x82, 0x95, 0x8a, 0x28, 0x2f, 0x81, 0x5e, 0xea, 0xd1, 0xfe, 0x59, 0xf2, 0x44, 0xef, 0x8a, 0x07,
	0x77, 0x68, 0x69, 0x93, 0x18, 0xde, 0x0a, 0xd8, 0x4a, 0xec, 0xb7, 0x8a, 0x05, 0x73, 0xb6, 0x4a,
	0x05, 0xbb, 0x2a, 0xa4, 0x27, 0x38, 0x0b, 0x8f, 0x8f, 0x3f, 0x73, 0xc3, 0xcb, 0xe3, 0x5e, 0x93,
	0x44, 0x89, 0x3d, 0xfd, 0xb7, 0x98, 0xbd, 0x2b, 0x38, 0xfb, 0x2f, 0x47, 0x9d, 0xfc, 0x73, 0xb6,
	0xcf, 0x07, 0x55, 0x2f, 0x75, 0x31, 0xee, 0xca, 0x99, 0x71, 0x68, 0x66, 0x7c, 0x6e, 0x7e, 0x66,
	0x78, 0x7f, 0x1a, 0x00, 0xd5, 0x05, 0x3c, 0x36, 0xe4, 0xe4, 0x13, 0xb0, 0x95, 0xbf, 0x29, 0x6e,
	0xcb, 0xd9, 0xfd, 0x6b, 0x3b, 0x5c, 0xf8, 0x1b, 0x99, 0xe5, 0x8c, 0x4a, 0x7d, 0x7d, 0xb7, 0xd9,
	0x0b, 0xfb, 0xd2, 0x83, 0x8d, 0x4f, 0x2c, 0xcd, 0xc9, 0xde, 0xa7, 0xe0, 0x94, 0x9e, 0xff, 0x2b,
	0xe9, 0x2f, 0xa0, 0x81, 0x91, 0xf0, 0x39, 0x0c, 0x63, 0x26, 0x95, 0x1f, 0x6f, 0xf3, 0xde, 0x57,
	0xc0, 0x61, 0x00, 0x23, 0x0f, 0x30, 0xa2, 0xe0, 0x56, 0x5d, 0xb8, 0x8d, 0xd2, 0x4d, 0x98, 0x90,
	0x2f, 0x8b, 0x8f, 0xe3, 0xed, 0xfb, 0x8d, 0xca, 0xff, 0xbb, 0x5e, 0xef, 0x18, 0x95, 0x3d, 0xb1,
	0xe3, 0x07, 0xdf, 0xd7, 0x3e, 0xac, 0x65, 0x13, 0xbf, 0xcf, 0xc7, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xab, 0xf8, 0xda, 0x77, 0x52, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatasourcePluginClient is the client API for DatasourcePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatasourcePluginClient interface {
	Query(ctx context.Context, in *DatasourceRequest, opts ...grpc.CallOption) (*DatasourceResponse, error)
}

type datasourcePluginClient struct {
	cc *grpc.ClientConn
}

func NewDatasourcePluginClient(cc *grpc.ClientConn) DatasourcePluginClient {
	return &datasourcePluginClient{cc}
}

func (c *datasourcePluginClient) Query(ctx context.Context, in *DatasourceRequest, opts ...grpc.CallOption) (*DatasourceResponse, error) {
	out := new(DatasourceResponse)
	err := c.cc.Invoke(ctx, "/models.DatasourcePlugin/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasourcePluginServer is the server API for DatasourcePlugin service.
type DatasourcePluginServer interface {
	Query(context.Context, *DatasourceRequest) (*DatasourceResponse, error)
}

// UnimplementedDatasourcePluginServer can be embedded to have forward compatible implementations.
type UnimplementedDatasourcePluginServer struct {
}

func (*UnimplementedDatasourcePluginServer) Query(ctx context.Context, req *DatasourceRequest) (*DatasourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterDatasourcePluginServer(s *grpc.Server, srv DatasourcePluginServer) {
	s.RegisterService(&_DatasourcePlugin_serviceDesc, srv)
}

func _DatasourcePlugin_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourcePluginServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.DatasourcePlugin/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourcePluginServer).Query(ctx, req.(*DatasourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatasourcePlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.DatasourcePlugin",
	HandlerType: (*DatasourcePluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DatasourcePlugin_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datasource.proto",
}