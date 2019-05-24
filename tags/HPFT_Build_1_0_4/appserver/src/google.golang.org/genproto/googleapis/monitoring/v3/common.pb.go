// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/common.proto

package monitoring

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	distribution "google.golang.org/genproto/googleapis/api/distribution"
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

// Specifies an ordering relationship on two arguments, here called left and
// right.
type ComparisonType int32

const (
	// No ordering relationship is specified.
	ComparisonType_COMPARISON_UNSPECIFIED ComparisonType = 0
	// The left argument is greater than the right argument.
	ComparisonType_COMPARISON_GT ComparisonType = 1
	// The left argument is greater than or equal to the right argument.
	ComparisonType_COMPARISON_GE ComparisonType = 2
	// The left argument is less than the right argument.
	ComparisonType_COMPARISON_LT ComparisonType = 3
	// The left argument is less than or equal to the right argument.
	ComparisonType_COMPARISON_LE ComparisonType = 4
	// The left argument is equal to the right argument.
	ComparisonType_COMPARISON_EQ ComparisonType = 5
	// The left argument is not equal to the right argument.
	ComparisonType_COMPARISON_NE ComparisonType = 6
)

var ComparisonType_name = map[int32]string{
	0: "COMPARISON_UNSPECIFIED",
	1: "COMPARISON_GT",
	2: "COMPARISON_GE",
	3: "COMPARISON_LT",
	4: "COMPARISON_LE",
	5: "COMPARISON_EQ",
	6: "COMPARISON_NE",
}

var ComparisonType_value = map[string]int32{
	"COMPARISON_UNSPECIFIED": 0,
	"COMPARISON_GT":          1,
	"COMPARISON_GE":          2,
	"COMPARISON_LT":          3,
	"COMPARISON_LE":          4,
	"COMPARISON_EQ":          5,
	"COMPARISON_NE":          6,
}

func (x ComparisonType) String() string {
	return proto.EnumName(ComparisonType_name, int32(x))
}

func (ComparisonType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_013c57c1dcbb8d65, []int{0}
}

// The tier of service for a Workspace. Please see the
// [service tiers
// documentation](https://cloud.google.com/monitoring/workspaces/tiers) for more
// details.
type ServiceTier int32 // Deprecated: Do not use.
const (
	// An invalid sentinel value, used to indicate that a tier has not
	// been provided explicitly.
	ServiceTier_SERVICE_TIER_UNSPECIFIED ServiceTier = 0
	// The Stackdriver Basic tier, a free tier of service that provides basic
	// features, a moderate allotment of logs, and access to built-in metrics.
	// A number of features are not available in this tier. For more details,
	// see [the service tiers
	// documentation](https://cloud.google.com/monitoring/workspaces/tiers).
	ServiceTier_SERVICE_TIER_BASIC ServiceTier = 1
	// The Stackdriver Premium tier, a higher, more expensive tier of service
	// that provides access to all Stackdriver features, lets you use Stackdriver
	// with AWS accounts, and has a larger allotments for logs and metrics. For
	// more details, see [the service tiers
	// documentation](https://cloud.google.com/monitoring/workspaces/tiers).
	ServiceTier_SERVICE_TIER_PREMIUM ServiceTier = 2
)

var ServiceTier_name = map[int32]string{
	0: "SERVICE_TIER_UNSPECIFIED",
	1: "SERVICE_TIER_BASIC",
	2: "SERVICE_TIER_PREMIUM",
}

var ServiceTier_value = map[string]int32{
	"SERVICE_TIER_UNSPECIFIED": 0,
	"SERVICE_TIER_BASIC":       1,
	"SERVICE_TIER_PREMIUM":     2,
}

func (x ServiceTier) String() string {
	return proto.EnumName(ServiceTier_name, int32(x))
}

func (ServiceTier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_013c57c1dcbb8d65, []int{1}
}

// The Aligner describes how to bring the data points in a single
// time series into temporal alignment.
type Aggregation_Aligner int32

const (
	// No alignment. Raw data is returned. Not valid if cross-time
	// series reduction is requested. The value type of the result is
	// the same as the value type of the input.
	Aggregation_ALIGN_NONE Aggregation_Aligner = 0
	// Align and convert to delta metric type. This alignment is valid
	// for cumulative metrics and delta metrics. Aligning an existing
	// delta metric to a delta metric requires that the alignment
	// period be increased. The value type of the result is the same
	// as the value type of the input.
	//
	// One can think of this aligner as a rate but without time units; that
	// is, the output is conceptually (second_point - first_point).
	Aggregation_ALIGN_DELTA Aggregation_Aligner = 1
	// Align and convert to a rate. This alignment is valid for
	// cumulative metrics and delta metrics with numeric values. The output is a
	// gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	//
	// One can think of this aligner as conceptually providing the slope of
	// the line that passes through the value at the start and end of the
	// window. In other words, this is conceptually ((y1 - y0)/(t1 - t0)),
	// and the output unit is one that has a "/time" dimension.
	//
	// If, by rate, you are looking for percentage change, see the
	// `ALIGN_PERCENT_CHANGE` aligner option.
	Aggregation_ALIGN_RATE Aggregation_Aligner = 2
	// Align by interpolating between adjacent points around the
	// period boundary. This alignment is valid for gauge
	// metrics with numeric values. The value type of the result is the same
	// as the value type of the input.
	Aggregation_ALIGN_INTERPOLATE Aggregation_Aligner = 3
	// Align by shifting the oldest data point before the period
	// boundary to the boundary. This alignment is valid for gauge
	// metrics. The value type of the result is the same as the
	// value type of the input.
	Aggregation_ALIGN_NEXT_OLDER Aggregation_Aligner = 4
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the minimum of all data points in the
	// period. This alignment is valid for gauge and delta metrics with numeric
	// values. The value type of the result is the same as the value
	// type of the input.
	Aggregation_ALIGN_MIN Aggregation_Aligner = 10
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the maximum of all data points in the
	// period. This alignment is valid for gauge and delta metrics with numeric
	// values. The value type of the result is the same as the value
	// type of the input.
	Aggregation_ALIGN_MAX Aggregation_Aligner = 11
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the average or arithmetic mean of all
	// data points in the period. This alignment is valid for gauge and delta
	// metrics with numeric values. The value type of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_MEAN Aggregation_Aligner = 12
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the count of all data points in the
	// period. This alignment is valid for gauge and delta metrics with numeric
	// or Boolean values. The value type of the output is
	// [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_ALIGN_COUNT Aggregation_Aligner = 13
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the sum of all data points in the
	// period. This alignment is valid for gauge and delta metrics with numeric
	// and distribution values. The value type of the output is the
	// same as the value type of the input.
	Aggregation_ALIGN_SUM Aggregation_Aligner = 14
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the standard deviation of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with numeric values. The value type of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_STDDEV Aggregation_Aligner = 15
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the count of True-valued data points in the
	// period. This alignment is valid for gauge metrics with
	// Boolean values. The value type of the output is
	// [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_ALIGN_COUNT_TRUE Aggregation_Aligner = 16
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the count of False-valued data points in the
	// period. This alignment is valid for gauge metrics with
	// Boolean values. The value type of the output is
	// [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_ALIGN_COUNT_FALSE Aggregation_Aligner = 24
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the fraction of True-valued data points in the
	// period. This alignment is valid for gauge metrics with Boolean values.
	// The output value is in the range [0, 1] and has value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_FRACTION_TRUE Aggregation_Aligner = 17
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the 99th percentile of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with distribution values. The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENTILE_99 Aggregation_Aligner = 18
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the 95th percentile of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with distribution values. The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENTILE_95 Aggregation_Aligner = 19
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the 50th percentile of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with distribution values. The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENTILE_50 Aggregation_Aligner = 20
	// Align time series via aggregation. The resulting data point in
	// the alignment period is the 5th percentile of all data
	// points in the period. This alignment is valid for gauge and delta metrics
	// with distribution values. The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENTILE_05 Aggregation_Aligner = 21
	// Align and convert to a percentage change. This alignment is valid for
	// gauge and delta metrics with numeric values. This alignment conceptually
	// computes the equivalent of "((current - previous)/previous)*100"
	// where previous value is determined based on the alignmentPeriod.
	// In the event that previous is 0 the calculated value is infinity with the
	// exception that if both (current - previous) and previous are 0 the
	// calculated value is 0.
	// A 10 minute moving mean is computed at each point of the time window
	// prior to the above calculation to smooth the metric and prevent false
	// positives from very short lived spikes.
	// Only applicable for data that is >= 0. Any values < 0 are treated as
	// no data. While delta metrics are accepted by this alignment special care
	// should be taken that the values for the metric will always be positive.
	// The output is a gauge metric with value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_ALIGN_PERCENT_CHANGE Aggregation_Aligner = 23
)

var Aggregation_Aligner_name = map[int32]string{
	0:  "ALIGN_NONE",
	1:  "ALIGN_DELTA",
	2:  "ALIGN_RATE",
	3:  "ALIGN_INTERPOLATE",
	4:  "ALIGN_NEXT_OLDER",
	10: "ALIGN_MIN",
	11: "ALIGN_MAX",
	12: "ALIGN_MEAN",
	13: "ALIGN_COUNT",
	14: "ALIGN_SUM",
	15: "ALIGN_STDDEV",
	16: "ALIGN_COUNT_TRUE",
	24: "ALIGN_COUNT_FALSE",
	17: "ALIGN_FRACTION_TRUE",
	18: "ALIGN_PERCENTILE_99",
	19: "ALIGN_PERCENTILE_95",
	20: "ALIGN_PERCENTILE_50",
	21: "ALIGN_PERCENTILE_05",
	23: "ALIGN_PERCENT_CHANGE",
}

var Aggregation_Aligner_value = map[string]int32{
	"ALIGN_NONE":           0,
	"ALIGN_DELTA":          1,
	"ALIGN_RATE":           2,
	"ALIGN_INTERPOLATE":    3,
	"ALIGN_NEXT_OLDER":     4,
	"ALIGN_MIN":            10,
	"ALIGN_MAX":            11,
	"ALIGN_MEAN":           12,
	"ALIGN_COUNT":          13,
	"ALIGN_SUM":            14,
	"ALIGN_STDDEV":         15,
	"ALIGN_COUNT_TRUE":     16,
	"ALIGN_COUNT_FALSE":    24,
	"ALIGN_FRACTION_TRUE":  17,
	"ALIGN_PERCENTILE_99":  18,
	"ALIGN_PERCENTILE_95":  19,
	"ALIGN_PERCENTILE_50":  20,
	"ALIGN_PERCENTILE_05":  21,
	"ALIGN_PERCENT_CHANGE": 23,
}

func (x Aggregation_Aligner) String() string {
	return proto.EnumName(Aggregation_Aligner_name, int32(x))
}

func (Aggregation_Aligner) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_013c57c1dcbb8d65, []int{2, 0}
}

// A Reducer describes how to aggregate data points from multiple
// time series into a single time series.
type Aggregation_Reducer int32

const (
	// No cross-time series reduction. The output of the aligner is
	// returned.
	Aggregation_REDUCE_NONE Aggregation_Reducer = 0
	// Reduce by computing the mean across time series for each
	// alignment period. This reducer is valid for delta and
	// gauge metrics with numeric or distribution values. The value type of the
	// output is [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_REDUCE_MEAN Aggregation_Reducer = 1
	// Reduce by computing the minimum across time series for each
	// alignment period. This reducer is valid for delta and
	// gauge metrics with numeric values. The value type of the output
	// is the same as the value type of the input.
	Aggregation_REDUCE_MIN Aggregation_Reducer = 2
	// Reduce by computing the maximum across time series for each
	// alignment period. This reducer is valid for delta and
	// gauge metrics with numeric values. The value type of the output
	// is the same as the value type of the input.
	Aggregation_REDUCE_MAX Aggregation_Reducer = 3
	// Reduce by computing the sum across time series for each
	// alignment period. This reducer is valid for delta and
	// gauge metrics with numeric and distribution values. The value type of
	// the output is the same as the value type of the input.
	Aggregation_REDUCE_SUM Aggregation_Reducer = 4
	// Reduce by computing the standard deviation across time series
	// for each alignment period. This reducer is valid for delta
	// and gauge metrics with numeric or distribution values. The value type of
	// the output is [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_REDUCE_STDDEV Aggregation_Reducer = 5
	// Reduce by computing the count of data points across time series
	// for each alignment period. This reducer is valid for delta
	// and gauge metrics of numeric, Boolean, distribution, and string value
	// type. The value type of the output is
	// [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_REDUCE_COUNT Aggregation_Reducer = 6
	// Reduce by computing the count of True-valued data points across time
	// series for each alignment period. This reducer is valid for delta
	// and gauge metrics of Boolean value type. The value type of
	// the output is [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_REDUCE_COUNT_TRUE Aggregation_Reducer = 7
	// Reduce by computing the count of False-valued data points across time
	// series for each alignment period. This reducer is valid for delta
	// and gauge metrics of Boolean value type. The value type of
	// the output is [INT64][google.api.MetricDescriptor.ValueType.INT64].
	Aggregation_REDUCE_COUNT_FALSE Aggregation_Reducer = 15
	// Reduce by computing the fraction of True-valued data points across time
	// series for each alignment period. This reducer is valid for delta
	// and gauge metrics of Boolean value type. The output value is in the
	// range [0, 1] and has value type
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE].
	Aggregation_REDUCE_FRACTION_TRUE Aggregation_Reducer = 8
	// Reduce by computing 99th percentile of data points across time series
	// for each alignment period. This reducer is valid for gauge and delta
	// metrics of numeric and distribution type. The value of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE]
	Aggregation_REDUCE_PERCENTILE_99 Aggregation_Reducer = 9
	// Reduce by computing 95th percentile of data points across time series
	// for each alignment period. This reducer is valid for gauge and delta
	// metrics of numeric and distribution type. The value of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE]
	Aggregation_REDUCE_PERCENTILE_95 Aggregation_Reducer = 10
	// Reduce by computing 50th percentile of data points across time series
	// for each alignment period. This reducer is valid for gauge and delta
	// metrics of numeric and distribution type. The value of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE]
	Aggregation_REDUCE_PERCENTILE_50 Aggregation_Reducer = 11
	// Reduce by computing 5th percentile of data points across time series
	// for each alignment period. This reducer is valid for gauge and delta
	// metrics of numeric and distribution type. The value of the output is
	// [DOUBLE][google.api.MetricDescriptor.ValueType.DOUBLE]
	Aggregation_REDUCE_PERCENTILE_05 Aggregation_Reducer = 12
)

var Aggregation_Reducer_name = map[int32]string{
	0:  "REDUCE_NONE",
	1:  "REDUCE_MEAN",
	2:  "REDUCE_MIN",
	3:  "REDUCE_MAX",
	4:  "REDUCE_SUM",
	5:  "REDUCE_STDDEV",
	6:  "REDUCE_COUNT",
	7:  "REDUCE_COUNT_TRUE",
	15: "REDUCE_COUNT_FALSE",
	8:  "REDUCE_FRACTION_TRUE",
	9:  "REDUCE_PERCENTILE_99",
	10: "REDUCE_PERCENTILE_95",
	11: "REDUCE_PERCENTILE_50",
	12: "REDUCE_PERCENTILE_05",
}

var Aggregation_Reducer_value = map[string]int32{
	"REDUCE_NONE":          0,
	"REDUCE_MEAN":          1,
	"REDUCE_MIN":           2,
	"REDUCE_MAX":           3,
	"REDUCE_SUM":           4,
	"REDUCE_STDDEV":        5,
	"REDUCE_COUNT":         6,
	"REDUCE_COUNT_TRUE":    7,
	"REDUCE_COUNT_FALSE":   15,
	"REDUCE_FRACTION_TRUE": 8,
	"REDUCE_PERCENTILE_99": 9,
	"REDUCE_PERCENTILE_95": 10,
	"REDUCE_PERCENTILE_50": 11,
	"REDUCE_PERCENTILE_05": 12,
}

func (x Aggregation_Reducer) String() string {
	return proto.EnumName(Aggregation_Reducer_name, int32(x))
}

func (Aggregation_Reducer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_013c57c1dcbb8d65, []int{2, 1}
}

// A single strongly-typed value.
type TypedValue struct {
	// The typed value field.
	//
	// Types that are valid to be assigned to Value:
	//	*TypedValue_BoolValue
	//	*TypedValue_Int64Value
	//	*TypedValue_DoubleValue
	//	*TypedValue_StringValue
	//	*TypedValue_DistributionValue
	Value                isTypedValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TypedValue) Reset()         { *m = TypedValue{} }
func (m *TypedValue) String() string { return proto.CompactTextString(m) }
func (*TypedValue) ProtoMessage()    {}
func (*TypedValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_013c57c1dcbb8d65, []int{0}
}

func (m *TypedValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedValue.Unmarshal(m, b)
}
func (m *TypedValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedValue.Marshal(b, m, deterministic)
}
func (m *TypedValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedValue.Merge(m, src)
}
func (m *TypedValue) XXX_Size() int {
	return xxx_messageInfo_TypedValue.Size(m)
}
func (m *TypedValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedValue.DiscardUnknown(m)
}

var xxx_messageInfo_TypedValue proto.InternalMessageInfo

type isTypedValue_Value interface {
	isTypedValue_Value()
}

type TypedValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,1,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type TypedValue_Int64Value struct {
	Int64Value int64 `protobuf:"varint,2,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type TypedValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type TypedValue_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type TypedValue_DistributionValue struct {
	DistributionValue *distribution.Distribution `protobuf:"bytes,5,opt,name=distribution_value,json=distributionValue,proto3,oneof"`
}

func (*TypedValue_BoolValue) isTypedValue_Value() {}

func (*TypedValue_Int64Value) isTypedValue_Value() {}

func (*TypedValue_DoubleValue) isTypedValue_Value() {}

func (*TypedValue_StringValue) isTypedValue_Value() {}

func (*TypedValue_DistributionValue) isTypedValue_Value() {}

func (m *TypedValue) GetValue() isTypedValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TypedValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*TypedValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *TypedValue) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*TypedValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *TypedValue) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*TypedValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *TypedValue) GetStringValue() string {
	if x, ok := m.GetValue().(*TypedValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *TypedValue) GetDistributionValue() *distribution.Distribution {
	if x, ok := m.GetValue().(*TypedValue_DistributionValue); ok {
		return x.DistributionValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TypedValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TypedValue_BoolValue)(nil),
		(*TypedValue_Int64Value)(nil),
		(*TypedValue_DoubleValue)(nil),
		(*TypedValue_StringValue)(nil),
		(*TypedValue_DistributionValue)(nil),
	}
}

// A time interval extending just after a start time through an end time.
// If the start time is the same as the end time, then the interval
// represents a single point in time.
type TimeInterval struct {
	// Required. The end of the time interval.
	EndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Optional. The beginning of the time interval.  The default value
	// for the start time is the end time. The start time must not be
	// later than the end time.
	StartTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeInterval) Reset()         { *m = TimeInterval{} }
func (m *TimeInterval) String() string { return proto.CompactTextString(m) }
func (*TimeInterval) ProtoMessage()    {}
func (*TimeInterval) Descriptor() ([]byte, []int) {
	return fileDescriptor_013c57c1dcbb8d65, []int{1}
}

func (m *TimeInterval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeInterval.Unmarshal(m, b)
}
func (m *TimeInterval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeInterval.Marshal(b, m, deterministic)
}
func (m *TimeInterval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeInterval.Merge(m, src)
}
func (m *TimeInterval) XXX_Size() int {
	return xxx_messageInfo_TimeInterval.Size(m)
}
func (m *TimeInterval) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeInterval.DiscardUnknown(m)
}

var xxx_messageInfo_TimeInterval proto.InternalMessageInfo

func (m *TimeInterval) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *TimeInterval) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

// Describes how to combine multiple time series to provide different views of
// the data.  Aggregation consists of an alignment step on individual time
// series (`alignment_period` and `per_series_aligner`) followed by an optional
// reduction step of the data across the aligned time series
// (`cross_series_reducer` and `group_by_fields`).  For more details, see
// [Aggregation](/monitoring/api/learn_more#aggregation).
type Aggregation struct {
	// The alignment period for per-[time series][google.monitoring.v3.TimeSeries]
	// alignment. If present, `alignmentPeriod` must be at least 60
	// seconds.  After per-time series alignment, each time series will
	// contain data points only on the period boundaries. If
	// `perSeriesAligner` is not specified or equals `ALIGN_NONE`, then
	// this field is ignored. If `perSeriesAligner` is specified and
	// does not equal `ALIGN_NONE`, then this field must be defined;
	// otherwise an error is returned.
	AlignmentPeriod *duration.Duration `protobuf:"bytes,1,opt,name=alignment_period,json=alignmentPeriod,proto3" json:"alignment_period,omitempty"`
	// The approach to be used to align individual time series. Not all
	// alignment functions may be applied to all time series, depending
	// on the metric type and value type of the original time
	// series. Alignment may change the metric type or the value type of
	// the time series.
	//
	// Time series data must be aligned in order to perform cross-time
	// series reduction. If `crossSeriesReducer` is specified, then
	// `perSeriesAligner` must be specified and not equal `ALIGN_NONE`
	// and `alignmentPeriod` must be specified; otherwise, an error is
	// returned.
	PerSeriesAligner Aggregation_Aligner `protobuf:"varint,2,opt,name=per_series_aligner,json=perSeriesAligner,proto3,enum=google.monitoring.v3.Aggregation_Aligner" json:"per_series_aligner,omitempty"`
	// The approach to be used to combine time series. Not all reducer
	// functions may be applied to all time series, depending on the
	// metric type and the value type of the original time
	// series. Reduction may change the metric type of value type of the
	// time series.
	//
	// Time series data must be aligned in order to perform cross-time
	// series reduction. If `crossSeriesReducer` is specified, then
	// `perSeriesAligner` must be specified and not equal `ALIGN_NONE`
	// and `alignmentPeriod` must be specified; otherwise, an error is
	// returned.
	CrossSeriesReducer Aggregation_Reducer `protobuf:"varint,4,opt,name=cross_series_reducer,json=crossSeriesReducer,proto3,enum=google.monitoring.v3.Aggregation_Reducer" json:"cross_series_reducer,omitempty"`
	// The set of fields to preserve when `crossSeriesReducer` is
	// specified. The `groupByFields` determine how the time series are
	// partitioned into subsets prior to applying the aggregation
	// function. Each subset contains time series that have the same
	// value for each of the grouping fields. Each individual time
	// series is a member of exactly one subset. The
	// `crossSeriesReducer` is applied to each subset of time series.
	// It is not possible to reduce across different resource types, so
	// this field implicitly contains `resource.type`.  Fields not
	// specified in `groupByFields` are aggregated away.  If
	// `groupByFields` is not specified and all the time series have
	// the same resource type, then the time series are aggregated into
	// a single output time series. If `crossSeriesReducer` is not
	// defined, this field is ignored.
	GroupByFields        []string `protobuf:"bytes,5,rep,name=group_by_fields,json=groupByFields,proto3" json:"group_by_fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Aggregation) Reset()         { *m = Aggregation{} }
func (m *Aggregation) String() string { return proto.CompactTextString(m) }
func (*Aggregation) ProtoMessage()    {}
func (*Aggregation) Descriptor() ([]byte, []int) {
	return fileDescriptor_013c57c1dcbb8d65, []int{2}
}

func (m *Aggregation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Aggregation.Unmarshal(m, b)
}
func (m *Aggregation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Aggregation.Marshal(b, m, deterministic)
}
func (m *Aggregation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Aggregation.Merge(m, src)
}
func (m *Aggregation) XXX_Size() int {
	return xxx_messageInfo_Aggregation.Size(m)
}
func (m *Aggregation) XXX_DiscardUnknown() {
	xxx_messageInfo_Aggregation.DiscardUnknown(m)
}

var xxx_messageInfo_Aggregation proto.InternalMessageInfo

func (m *Aggregation) GetAlignmentPeriod() *duration.Duration {
	if m != nil {
		return m.AlignmentPeriod
	}
	return nil
}

func (m *Aggregation) GetPerSeriesAligner() Aggregation_Aligner {
	if m != nil {
		return m.PerSeriesAligner
	}
	return Aggregation_ALIGN_NONE
}

func (m *Aggregation) GetCrossSeriesReducer() Aggregation_Reducer {
	if m != nil {
		return m.CrossSeriesReducer
	}
	return Aggregation_REDUCE_NONE
}

func (m *Aggregation) GetGroupByFields() []string {
	if m != nil {
		return m.GroupByFields
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.monitoring.v3.ComparisonType", ComparisonType_name, ComparisonType_value)
	proto.RegisterEnum("google.monitoring.v3.ServiceTier", ServiceTier_name, ServiceTier_value)
	proto.RegisterEnum("google.monitoring.v3.Aggregation_Aligner", Aggregation_Aligner_name, Aggregation_Aligner_value)
	proto.RegisterEnum("google.monitoring.v3.Aggregation_Reducer", Aggregation_Reducer_name, Aggregation_Reducer_value)
	proto.RegisterType((*TypedValue)(nil), "google.monitoring.v3.TypedValue")
	proto.RegisterType((*TimeInterval)(nil), "google.monitoring.v3.TimeInterval")
	proto.RegisterType((*Aggregation)(nil), "google.monitoring.v3.Aggregation")
}

func init() { proto.RegisterFile("google/monitoring/v3/common.proto", fileDescriptor_013c57c1dcbb8d65) }

var fileDescriptor_013c57c1dcbb8d65 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xc1, 0x6e, 0xe3, 0x44,
	0x18, 0xc7, 0xe3, 0x64, 0xdb, 0x34, 0x9f, 0xdb, 0x66, 0x3a, 0xdb, 0xed, 0x86, 0x68, 0x61, 0xb3,
	0x45, 0x42, 0x61, 0x0f, 0x4e, 0xd5, 0x12, 0xa4, 0x0a, 0x09, 0xc9, 0x75, 0xa6, 0xad, 0xa5, 0xc4,
	0x09, 0x13, 0xa7, 0x54, 0x50, 0xc9, 0x72, 0x9a, 0x59, 0xcb, 0x52, 0xe2, 0xb1, 0x6c, 0xa7, 0x52,
	0x6f, 0xdc, 0x79, 0x07, 0x2e, 0xdc, 0xb8, 0xf1, 0x1a, 0x3c, 0x0c, 0x17, 0x5e, 0x00, 0x79, 0xc6,
	0x59, 0x3b, 0x21, 0x08, 0x8e, 0xdf, 0xef, 0xff, 0xff, 0xbe, 0x99, 0xf9, 0x8f, 0x35, 0x86, 0x77,
	0x1e, 0xe7, 0xde, 0x9c, 0x75, 0x16, 0x3c, 0xf0, 0x13, 0x1e, 0xf9, 0x81, 0xd7, 0x79, 0xba, 0xe8,
	0x3c, 0xf2, 0xc5, 0x82, 0x07, 0x5a, 0x18, 0xf1, 0x84, 0xe3, 0x63, 0x69, 0xd1, 0x72, 0x8b, 0xf6,
	0x74, 0xd1, 0x7c, 0x93, 0x35, 0xba, 0xa1, 0xdf, 0x71, 0x83, 0x80, 0x27, 0x6e, 0xe2, 0xf3, 0x20,
	0x96, 0x3d, 0xcd, 0x4f, 0x0b, 0xea, 0xcc, 0x8f, 0x93, 0xc8, 0x9f, 0x2e, 0x53, 0x3d, 0x93, 0x3f,
	0xcb, 0x64, 0x51, 0x4d, 0x97, 0x1f, 0x3a, 0xb3, 0x65, 0xe4, 0x16, 0xf4, 0xb7, 0x9b, 0x7a, 0xe2,
	0x2f, 0x58, 0x9c, 0xb8, 0x8b, 0x50, 0x1a, 0x4e, 0xff, 0x54, 0x00, 0xec, 0xe7, 0x90, 0xcd, 0xee,
	0xdc, 0xf9, 0x92, 0xe1, 0xb7, 0x00, 0x53, 0xce, 0xe7, 0xce, 0x53, 0x5a, 0x35, 0x94, 0x96, 0xd2,
	0xde, 0xbb, 0x2d, 0xd1, 0x5a, 0xca, 0xa4, 0xe1, 0x1d, 0xa8, 0x7e, 0x90, 0x7c, 0xfd, 0x55, 0xe6,
	0x28, 0xb7, 0x94, 0x76, 0xe5, 0xb6, 0x44, 0x41, 0x40, 0x69, 0xf9, 0x1c, 0xf6, 0x67, 0x7c, 0x39,
	0x9d, 0xb3, 0xcc, 0x53, 0x69, 0x29, 0x6d, 0xe5, 0xb6, 0x44, 0x55, 0x49, 0x3f, 0x9a, 0xd2, 0xc3,
	0x04, 0x5e, 0x66, 0x7a, 0xd1, 0x52, 0xda, 0xb5, 0xd4, 0x24, 0xa9, 0x34, 0x99, 0x80, 0x8b, 0x67,
	0xce, 0xac, 0x3b, 0x2d, 0xa5, 0xad, 0x9e, 0x37, 0xb4, 0x2c, 0x4d, 0x37, 0xf4, 0xb5, 0x5e, 0xc1,
	0x75, 0x5b, 0xa2, 0x47, 0xc5, 0x2e, 0x31, 0xea, 0xaa, 0x0a, 0x3b, 0xa2, 0xfb, 0xf4, 0x27, 0x05,
	0xf6, 0x6d, 0x7f, 0xc1, 0xcc, 0x20, 0x61, 0xd1, 0x93, 0x3b, 0xc7, 0x5d, 0xd8, 0x63, 0xc1, 0xcc,
	0x49, 0x83, 0x11, 0xc7, 0x51, 0xcf, 0x9b, 0xab, 0xd1, 0xab, 0xd4, 0x34, 0x7b, 0x95, 0x1a, 0xad,
	0xb2, 0x60, 0x96, 0x56, 0xf8, 0x12, 0x20, 0x4e, 0xdc, 0x28, 0x91, 0x8d, 0xca, 0x7f, 0x36, 0xd6,
	0x84, 0x3b, 0xad, 0x4f, 0xff, 0xaa, 0x82, 0xaa, 0x7b, 0x5e, 0xc4, 0x3c, 0x71, 0x55, 0xb8, 0x07,
	0xc8, 0x9d, 0xfb, 0x5e, 0xb0, 0x60, 0x41, 0xe2, 0x84, 0x2c, 0xf2, 0xf9, 0x2c, 0x1b, 0xf8, 0xc9,
	0x3f, 0x06, 0xf6, 0xb2, 0xfb, 0xa5, 0xf5, 0x8f, 0x2d, 0x23, 0xd1, 0x81, 0xbf, 0x07, 0x1c, 0xb2,
	0xc8, 0x89, 0x59, 0xe4, 0xb3, 0xd8, 0x11, 0x2a, 0x8b, 0xc4, 0x89, 0x0e, 0xcf, 0xbf, 0xd4, 0xb6,
	0x7d, 0x7a, 0x5a, 0x61, 0x13, 0x9a, 0x2e, 0x1b, 0x28, 0x0a, 0x59, 0x34, 0x16, 0x33, 0x32, 0x82,
	0x7f, 0x84, 0xe3, 0xc7, 0x88, 0xc7, 0xf1, 0x6a, 0x74, 0xc4, 0x66, 0xcb, 0x47, 0x16, 0x89, 0x2b,
	0xfb, 0x5f, 0xa3, 0xa9, 0x6c, 0xa0, 0x58, 0x8c, 0x91, 0xc3, 0x33, 0x86, 0xbf, 0x80, 0xba, 0x17,
	0xf1, 0x65, 0xe8, 0x4c, 0x9f, 0x9d, 0x0f, 0x3e, 0x9b, 0xcf, 0xe2, 0xc6, 0x4e, 0xab, 0xd2, 0xae,
	0xd1, 0x03, 0x81, 0xaf, 0x9e, 0xaf, 0x05, 0x3c, 0xfd, 0xb9, 0x02, 0xd5, 0xd5, 0x86, 0x0e, 0x01,
	0xf4, 0xbe, 0x79, 0x63, 0x39, 0xd6, 0xd0, 0x22, 0xa8, 0x84, 0xeb, 0xa0, 0xca, 0xba, 0x47, 0xfa,
	0xb6, 0x8e, 0x94, 0xdc, 0x40, 0x75, 0x9b, 0xa0, 0x32, 0x7e, 0x05, 0x47, 0xb2, 0x36, 0x2d, 0x9b,
	0xd0, 0xd1, 0xb0, 0x9f, 0xe2, 0x0a, 0x3e, 0x06, 0x94, 0xcd, 0x21, 0xf7, 0xb6, 0x33, 0xec, 0xf7,
	0x08, 0x45, 0x2f, 0xf0, 0x01, 0xd4, 0x24, 0x1d, 0x98, 0x16, 0x82, 0x42, 0xa9, 0xdf, 0x23, 0x35,
	0x1f, 0x3d, 0x20, 0xba, 0x85, 0xf6, 0xf3, 0xb5, 0x8d, 0xe1, 0xc4, 0xb2, 0xd1, 0x41, 0xee, 0x1f,
	0x4f, 0x06, 0xe8, 0x10, 0x23, 0xd8, 0xcf, 0x4a, 0xbb, 0xd7, 0x23, 0x77, 0xa8, 0x9e, 0xaf, 0x2a,
	0x3a, 0x1c, 0x9b, 0x4e, 0x08, 0x42, 0xf9, 0x16, 0x25, 0xbd, 0xd6, 0xfb, 0x63, 0x82, 0x1a, 0xf8,
	0x35, 0xbc, 0x94, 0xf8, 0x9a, 0xea, 0x86, 0x6d, 0x0e, 0x2d, 0xe9, 0x3f, 0xca, 0x85, 0x11, 0xa1,
	0x06, 0xb1, 0x6c, 0xb3, 0x4f, 0x9c, 0xcb, 0x4b, 0x84, 0xb7, 0x0b, 0x5d, 0xf4, 0x72, 0xab, 0xd0,
	0x3d, 0x43, 0xc7, 0x5b, 0x85, 0xb3, 0x2e, 0x7a, 0x85, 0x1b, 0x70, 0xbc, 0x26, 0x38, 0xc6, 0xad,
	0x6e, 0xdd, 0x10, 0xf4, 0xfa, 0xf4, 0xf7, 0x32, 0x54, 0x57, 0x37, 0x58, 0x07, 0x95, 0x92, 0xde,
	0xc4, 0x20, 0x85, 0xeb, 0xc8, 0x80, 0xc8, 0x48, 0x5c, 0xc7, 0x0a, 0x98, 0x16, 0x2a, 0x17, 0x6b,
	0xfd, 0x1e, 0x55, 0x0a, 0x75, 0x9a, 0xd9, 0x0b, 0x7c, 0x04, 0x07, 0xab, 0x5a, 0x86, 0xb6, 0x93,
	0xc6, 0x98, 0x21, 0x99, 0xf3, 0x6e, 0x1a, 0x58, 0x91, 0xc8, 0x5c, 0xaa, 0xf8, 0x04, 0xf0, 0x1a,
	0x96, 0x41, 0xd6, 0xd3, 0xb3, 0x64, 0x7c, 0x3d, 0xc9, 0xbd, 0x82, 0xb2, 0x1e, 0x65, 0xed, 0x5f,
	0x94, 0x2e, 0x82, 0xed, 0x4a, 0xf7, 0x0c, 0xa9, 0xdb, 0x95, 0xb3, 0x2e, 0xda, 0x7f, 0xff, 0x8b,
	0x02, 0x87, 0x06, 0x5f, 0x84, 0x6e, 0xe4, 0xc7, 0x3c, 0x48, 0xdf, 0x5c, 0xdc, 0x84, 0x13, 0x63,
	0x38, 0x18, 0xe9, 0xd4, 0x1c, 0x0f, 0x2d, 0x67, 0x62, 0x8d, 0x47, 0xc4, 0x30, 0xaf, 0x4d, 0xd2,
	0x43, 0xa5, 0x34, 0x84, 0x82, 0x76, 0x63, 0x23, 0x65, 0x13, 0xa5, 0x5f, 0xf6, 0x3a, 0xea, 0xdb,
	0xa8, 0xb2, 0x89, 0x88, 0x0c, 0xb4, 0x80, 0xc8, 0x77, 0x68, 0x67, 0x03, 0x59, 0x04, 0xed, 0xbe,
	0x77, 0x41, 0x1d, 0xb3, 0xe8, 0xc9, 0x7f, 0x64, 0xb6, 0xcf, 0x22, 0xfc, 0x06, 0x1a, 0x63, 0x42,
	0xef, 0x4c, 0x83, 0x38, 0xb6, 0x49, 0xe8, 0xc6, 0xf6, 0x4e, 0x00, 0xaf, 0xa9, 0x57, 0xfa, 0xd8,
	0x34, 0x90, 0x92, 0x9e, 0x7f, 0x8d, 0x8f, 0x28, 0x19, 0x98, 0x93, 0x01, 0x2a, 0x37, 0xcb, 0x0d,
	0xe5, 0xea, 0x57, 0x05, 0x1a, 0x8f, 0x7c, 0xb1, 0xf5, 0xc9, 0xb8, 0x52, 0x0d, 0xf1, 0xb3, 0x1c,
	0xa5, 0x4f, 0xdd, 0x48, 0xf9, 0xe1, 0xdb, 0xcc, 0xe4, 0xf1, 0xb9, 0x1b, 0x78, 0x1a, 0x8f, 0xbc,
	0x8e, 0xc7, 0x02, 0xf1, 0x10, 0x76, 0xa4, 0xe4, 0x86, 0x7e, 0xbc, 0xfe, 0xbf, 0xfd, 0x26, 0xaf,
	0x7e, 0x2b, 0x37, 0x6f, 0xe4, 0x00, 0x63, 0xce, 0x97, 0x33, 0x6d, 0x90, 0xaf, 0x75, 0x77, 0xf1,
	0xc7, 0x4a, 0x7c, 0x10, 0xe2, 0x43, 0x2e, 0x3e, 0xdc, 0x5d, 0x4c, 0x77, 0xc5, 0x22, 0x17, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x78, 0xd9, 0x96, 0xd3, 0x07, 0x00, 0x00,
}
