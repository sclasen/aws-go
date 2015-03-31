// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudwatch provides a client for Amazon CloudWatch.
package cloudwatch

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// DeleteAlarmsRequest generates a request for the DeleteAlarms operation.
func (c *CloudWatch) DeleteAlarmsRequest(input *DeleteAlarmsInput) (req *aws.Request, output *DeleteAlarmsOutput) {
	if opDeleteAlarms == nil {
		opDeleteAlarms = &aws.Operation{
			Name:       "DeleteAlarms",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteAlarms, input, output)
	output = &DeleteAlarmsOutput{}
	req.Data = output
	return
}

// Deletes all specified alarms. In the event of an error, no alarms are deleted.
func (c *CloudWatch) DeleteAlarms(input *DeleteAlarmsInput) (output *DeleteAlarmsOutput, err error) {
	req, out := c.DeleteAlarmsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteAlarms *aws.Operation

// DescribeAlarmHistoryRequest generates a request for the DescribeAlarmHistory operation.
func (c *CloudWatch) DescribeAlarmHistoryRequest(input *DescribeAlarmHistoryInput) (req *aws.Request, output *DescribeAlarmHistoryOutput) {
	if opDescribeAlarmHistory == nil {
		opDescribeAlarmHistory = &aws.Operation{
			Name:       "DescribeAlarmHistory",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeAlarmHistory, input, output)
	output = &DescribeAlarmHistoryOutput{}
	req.Data = output
	return
}

// Retrieves history for the specified alarm. Filter alarms by date range or
// item type. If an alarm name is not specified, Amazon CloudWatch returns histories
// for all of the owner's alarms.
func (c *CloudWatch) DescribeAlarmHistory(input *DescribeAlarmHistoryInput) (output *DescribeAlarmHistoryOutput, err error) {
	req, out := c.DescribeAlarmHistoryRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeAlarmHistory *aws.Operation

// DescribeAlarmsRequest generates a request for the DescribeAlarms operation.
func (c *CloudWatch) DescribeAlarmsRequest(input *DescribeAlarmsInput) (req *aws.Request, output *DescribeAlarmsOutput) {
	if opDescribeAlarms == nil {
		opDescribeAlarms = &aws.Operation{
			Name:       "DescribeAlarms",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeAlarms, input, output)
	output = &DescribeAlarmsOutput{}
	req.Data = output
	return
}

// Retrieves alarms with the specified names. If no name is specified, all alarms
// for the user are returned. Alarms can be retrieved by using only a prefix
// for the alarm name, the alarm state, or a prefix for any action.
func (c *CloudWatch) DescribeAlarms(input *DescribeAlarmsInput) (output *DescribeAlarmsOutput, err error) {
	req, out := c.DescribeAlarmsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeAlarms *aws.Operation

// DescribeAlarmsForMetricRequest generates a request for the DescribeAlarmsForMetric operation.
func (c *CloudWatch) DescribeAlarmsForMetricRequest(input *DescribeAlarmsForMetricInput) (req *aws.Request, output *DescribeAlarmsForMetricOutput) {
	if opDescribeAlarmsForMetric == nil {
		opDescribeAlarmsForMetric = &aws.Operation{
			Name:       "DescribeAlarmsForMetric",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeAlarmsForMetric, input, output)
	output = &DescribeAlarmsForMetricOutput{}
	req.Data = output
	return
}

// Retrieves all alarms for a single metric. Specify a statistic, period, or
// unit to filter the set of alarms further.
func (c *CloudWatch) DescribeAlarmsForMetric(input *DescribeAlarmsForMetricInput) (output *DescribeAlarmsForMetricOutput, err error) {
	req, out := c.DescribeAlarmsForMetricRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeAlarmsForMetric *aws.Operation

// DisableAlarmActionsRequest generates a request for the DisableAlarmActions operation.
func (c *CloudWatch) DisableAlarmActionsRequest(input *DisableAlarmActionsInput) (req *aws.Request, output *DisableAlarmActionsOutput) {
	if opDisableAlarmActions == nil {
		opDisableAlarmActions = &aws.Operation{
			Name:       "DisableAlarmActions",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDisableAlarmActions, input, output)
	output = &DisableAlarmActionsOutput{}
	req.Data = output
	return
}

// Disables actions for the specified alarms. When an alarm's actions are disabled
// the alarm's state may change, but none of the alarm's actions will execute.
func (c *CloudWatch) DisableAlarmActions(input *DisableAlarmActionsInput) (output *DisableAlarmActionsOutput, err error) {
	req, out := c.DisableAlarmActionsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDisableAlarmActions *aws.Operation

// EnableAlarmActionsRequest generates a request for the EnableAlarmActions operation.
func (c *CloudWatch) EnableAlarmActionsRequest(input *EnableAlarmActionsInput) (req *aws.Request, output *EnableAlarmActionsOutput) {
	if opEnableAlarmActions == nil {
		opEnableAlarmActions = &aws.Operation{
			Name:       "EnableAlarmActions",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEnableAlarmActions, input, output)
	output = &EnableAlarmActionsOutput{}
	req.Data = output
	return
}

// Enables actions for the specified alarms.
func (c *CloudWatch) EnableAlarmActions(input *EnableAlarmActionsInput) (output *EnableAlarmActionsOutput, err error) {
	req, out := c.EnableAlarmActionsRequest(input)
	output = out
	err = req.Send()
	return
}

var opEnableAlarmActions *aws.Operation

// GetMetricStatisticsRequest generates a request for the GetMetricStatistics operation.
func (c *CloudWatch) GetMetricStatisticsRequest(input *GetMetricStatisticsInput) (req *aws.Request, output *GetMetricStatisticsOutput) {
	if opGetMetricStatistics == nil {
		opGetMetricStatistics = &aws.Operation{
			Name:       "GetMetricStatistics",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetMetricStatistics, input, output)
	output = &GetMetricStatisticsOutput{}
	req.Data = output
	return
}

// Gets statistics for the specified metric.
//
//  The maximum number of data points returned from a single GetMetricStatistics
// request is 1,440, wereas the maximum number of data points that can be queried
// is 50,850. If you make a request that generates more than 1,440 data points,
// Amazon CloudWatch returns an error. In such a case, you can alter the request
// by narrowing the specified time range or increasing the specified period.
// Alternatively, you can make multiple requests across adjacent time ranges.
//
//  Amazon CloudWatch aggregates data points based on the length of the period
// that you specify. For example, if you request statistics with a one-minute
// granularity, Amazon CloudWatch aggregates data points with time stamps that
// fall within the same one-minute period. In such a case, the data points queried
// can greatly outnumber the data points returned.
//
//  The following examples show various statistics allowed by the data point
// query maximum of 50,850 when you call GetMetricStatistics on Amazon EC2 instances
// with detailed (one-minute) monitoring enabled:
//
//  Statistics for up to 400 instances for a span of one hour Statistics for
// up to 35 instances over a span of 24 hours Statistics for up to 2 instances
// over a span of 2 weeks   For information about the namespace, metric names,
// and dimensions that other Amazon Web Services products use to send metrics
// to Cloudwatch, go to Amazon CloudWatch Metrics, Namespaces, and Dimensions
// Reference (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html)
// in the Amazon CloudWatch Developer Guide.
func (c *CloudWatch) GetMetricStatistics(input *GetMetricStatisticsInput) (output *GetMetricStatisticsOutput, err error) {
	req, out := c.GetMetricStatisticsRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetMetricStatistics *aws.Operation

// ListMetricsRequest generates a request for the ListMetrics operation.
func (c *CloudWatch) ListMetricsRequest(input *ListMetricsInput) (req *aws.Request, output *ListMetricsOutput) {
	if opListMetrics == nil {
		opListMetrics = &aws.Operation{
			Name:       "ListMetrics",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListMetrics, input, output)
	output = &ListMetricsOutput{}
	req.Data = output
	return
}

// Returns a list of valid metrics stored for the AWS account owner. Returned
// metrics can be used with GetMetricStatistics to obtain statistical data for
// a given metric.
func (c *CloudWatch) ListMetrics(input *ListMetricsInput) (output *ListMetricsOutput, err error) {
	req, out := c.ListMetricsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListMetrics *aws.Operation

// PutMetricAlarmRequest generates a request for the PutMetricAlarm operation.
func (c *CloudWatch) PutMetricAlarmRequest(input *PutMetricAlarmInput) (req *aws.Request, output *PutMetricAlarmOutput) {
	if opPutMetricAlarm == nil {
		opPutMetricAlarm = &aws.Operation{
			Name:       "PutMetricAlarm",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutMetricAlarm, input, output)
	output = &PutMetricAlarmOutput{}
	req.Data = output
	return
}

// Creates or updates an alarm and associates it with the specified Amazon CloudWatch
// metric. Optionally, this operation can associate one or more Amazon Simple
// Notification Service resources with the alarm.
//
//  When this operation creates an alarm, the alarm state is immediately set
// to INSUFFICIENT_DATA. The alarm is evaluated and its StateValue is set appropriately.
// Any actions associated with the StateValue is then executed.
func (c *CloudWatch) PutMetricAlarm(input *PutMetricAlarmInput) (output *PutMetricAlarmOutput, err error) {
	req, out := c.PutMetricAlarmRequest(input)
	output = out
	err = req.Send()
	return
}

var opPutMetricAlarm *aws.Operation

// PutMetricDataRequest generates a request for the PutMetricData operation.
func (c *CloudWatch) PutMetricDataRequest(input *PutMetricDataInput) (req *aws.Request, output *PutMetricDataOutput) {
	if opPutMetricData == nil {
		opPutMetricData = &aws.Operation{
			Name:       "PutMetricData",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutMetricData, input, output)
	output = &PutMetricDataOutput{}
	req.Data = output
	return
}

// Publishes metric data points to Amazon CloudWatch. Amazon Cloudwatch associates
// the data points with the specified metric. If the specified metric does not
// exist, Amazon CloudWatch creates the metric. It can take up to fifteen minutes
// for a new metric to appear in calls to the ListMetrics action.
//
//  The size of a PutMetricData request is limited to 8 KB for HTTP GET requests
// and 40 KB for HTTP POST requests.
//
//  Although the Value parameter accepts numbers of type Double, Amazon CloudWatch
// truncates values with very large exponents. Values with base-10 exponents
// greater than 126 (1 x 10^126) are truncated. Likewise, values with base-10
// exponents less than -130 (1 x 10^-130) are also truncated.  Data that is
// timestamped 24 hours or more in the past may take in excess of 48 hours to
// become available from submission time using GetMetricStatistics.
func (c *CloudWatch) PutMetricData(input *PutMetricDataInput) (output *PutMetricDataOutput, err error) {
	req, out := c.PutMetricDataRequest(input)
	output = out
	err = req.Send()
	return
}

var opPutMetricData *aws.Operation

// SetAlarmStateRequest generates a request for the SetAlarmState operation.
func (c *CloudWatch) SetAlarmStateRequest(input *SetAlarmStateInput) (req *aws.Request, output *SetAlarmStateOutput) {
	if opSetAlarmState == nil {
		opSetAlarmState = &aws.Operation{
			Name:       "SetAlarmState",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetAlarmState, input, output)
	output = &SetAlarmStateOutput{}
	req.Data = output
	return
}

// Temporarily sets the state of an alarm. When the updated StateValue differs
// from the previous value, the action configured for the appropriate state
// is invoked. This is not a permanent change. The next periodic alarm check
// (in about a minute) will set the alarm to its actual state.
func (c *CloudWatch) SetAlarmState(input *SetAlarmStateInput) (output *SetAlarmStateOutput, err error) {
	req, out := c.SetAlarmStateRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetAlarmState *aws.Operation

// The AlarmHistoryItem data type contains descriptive information about the
// history of a specific alarm. If you call DescribeAlarmHistory, Amazon CloudWatch
// returns this data type as part of the DescribeAlarmHistoryResult data type.
type AlarmHistoryItem struct {
	// The descriptive name for the alarm.
	AlarmName *string `type:"string"`

	// Machine-readable data about the alarm in JSON format.
	HistoryData *string `type:"string"`

	// The type of alarm history item.
	HistoryItemType *string `type:"string"`

	// A human-readable summary of the alarm history.
	HistorySummary *string `type:"string"`

	// The time stamp for the alarm history item. Amazon CloudWatch uses Coordinated
	// Universal Time (UTC) when returning time stamps, which do not accommodate
	// seasonal adjustments such as daylight savings time. For more information,
	// see Time stamps (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp)
	// in the Amazon CloudWatch Developer Guide.
	Timestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	metadataAlarmHistoryItem `json:"-", xml:"-"`
}

type metadataAlarmHistoryItem struct {
	SDKShapeTraits bool `type:"structure"`
}

const (
	ComparisonOperatorGreaterThanOrEqualToThreshold = "GreaterThanOrEqualToThreshold"
	ComparisonOperatorGreaterThanThreshold          = "GreaterThanThreshold"
	ComparisonOperatorLessThanOrEqualToThreshold    = "LessThanOrEqualToThreshold"
	ComparisonOperatorLessThanThreshold             = "LessThanThreshold"
)

// The Datapoint data type encapsulates the statistical data that Amazon CloudWatch
// computes from metric data.
type Datapoint struct {
	// The average of metric values that correspond to the datapoint.
	Average *float64 `type:"double"`

	// The maximum of the metric value used for the datapoint.
	Maximum *float64 `type:"double"`

	// The minimum metric value used for the datapoint.
	Minimum *float64 `type:"double"`

	// The number of metric values that contributed to the aggregate value of this
	// datapoint.
	SampleCount *float64 `type:"double"`

	// The sum of metric values used for the datapoint.
	Sum *float64 `type:"double"`

	// The time stamp used for the datapoint. Amazon CloudWatch uses Coordinated
	// Universal Time (UTC) when returning time stamps, which do not accommodate
	// seasonal adjustments such as daylight savings time. For more information,
	// see Time stamps (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp)
	// in the Amazon CloudWatch Developer Guide.
	Timestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The standard unit used for the datapoint.
	Unit *string `type:"string"`

	metadataDatapoint `json:"-", xml:"-"`
}

type metadataDatapoint struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteAlarmsInput struct {
	// A list of alarms to be deleted.
	AlarmNames []*string `type:"list" required:"true"`

	metadataDeleteAlarmsInput `json:"-", xml:"-"`
}

type metadataDeleteAlarmsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteAlarmsOutput struct {
	metadataDeleteAlarmsOutput `json:"-", xml:"-"`
}

type metadataDeleteAlarmsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeAlarmHistoryInput struct {
	// The name of the alarm.
	AlarmName *string `type:"string"`

	// The ending date to retrieve alarm history.
	EndDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The type of alarm histories to retrieve.
	HistoryItemType *string `type:"string"`

	// The maximum number of alarm history records to retrieve.
	MaxRecords *int64 `type:"integer"`

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string `type:"string"`

	// The starting date to retrieve alarm history.
	StartDate *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	metadataDescribeAlarmHistoryInput `json:"-", xml:"-"`
}

type metadataDescribeAlarmHistoryInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the DescribeAlarmHistory action.
type DescribeAlarmHistoryOutput struct {
	// A list of alarm histories in JSON format.
	AlarmHistoryItems []*AlarmHistoryItem `type:"list"`

	// A string that marks the start of the next batch of returned results.
	NextToken *string `type:"string"`

	metadataDescribeAlarmHistoryOutput `json:"-", xml:"-"`
}

type metadataDescribeAlarmHistoryOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeAlarmsForMetricInput struct {
	// The list of dimensions associated with the metric.
	Dimensions []*Dimension `type:"list"`

	// The name of the metric.
	MetricName *string `type:"string" required:"true"`

	// The namespace of the metric.
	Namespace *string `type:"string" required:"true"`

	// The period in seconds over which the statistic is applied.
	Period *int64 `type:"integer"`

	// The statistic for the metric.
	Statistic *string `type:"string"`

	// The unit for the metric.
	Unit *string `type:"string"`

	metadataDescribeAlarmsForMetricInput `json:"-", xml:"-"`
}

type metadataDescribeAlarmsForMetricInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the DescribeAlarmsForMetric action.
type DescribeAlarmsForMetricOutput struct {
	// A list of information for each alarm with the specified metric.
	MetricAlarms []*MetricAlarm `type:"list"`

	metadataDescribeAlarmsForMetricOutput `json:"-", xml:"-"`
}

type metadataDescribeAlarmsForMetricOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeAlarmsInput struct {
	// The action name prefix.
	ActionPrefix *string `type:"string"`

	// The alarm name prefix. AlarmNames cannot be specified if this parameter is
	// specified.
	AlarmNamePrefix *string `type:"string"`

	// A list of alarm names to retrieve information for.
	AlarmNames []*string `type:"list"`

	// The maximum number of alarm descriptions to retrieve.
	MaxRecords *int64 `type:"integer"`

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string `type:"string"`

	// The state value to be used in matching alarms.
	StateValue *string `type:"string"`

	metadataDescribeAlarmsInput `json:"-", xml:"-"`
}

type metadataDescribeAlarmsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the DescribeAlarms action.
type DescribeAlarmsOutput struct {
	// A list of information for the specified alarms.
	MetricAlarms []*MetricAlarm `type:"list"`

	// A string that marks the start of the next batch of returned results.
	NextToken *string `type:"string"`

	metadataDescribeAlarmsOutput `json:"-", xml:"-"`
}

type metadataDescribeAlarmsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The Dimension data type further expands on the identity of a metric using
// a Name, Value pair.
//
// For examples that use one or more dimensions, see PutMetricData.
type Dimension struct {
	// The name of the dimension.
	Name *string `type:"string" required:"true"`

	// The value representing the dimension measurement
	Value *string `type:"string" required:"true"`

	metadataDimension `json:"-", xml:"-"`
}

type metadataDimension struct {
	SDKShapeTraits bool `type:"structure"`
}

// The DimensionFilter data type is used to filter ListMetrics results.
type DimensionFilter struct {
	// The dimension name to be matched.
	Name *string `type:"string" required:"true"`

	// The value of the dimension to be matched.
	Value *string `type:"string"`

	metadataDimensionFilter `json:"-", xml:"-"`
}

type metadataDimensionFilter struct {
	SDKShapeTraits bool `type:"structure"`
}

type DisableAlarmActionsInput struct {
	// The names of the alarms to disable actions for.
	AlarmNames []*string `type:"list" required:"true"`

	metadataDisableAlarmActionsInput `json:"-", xml:"-"`
}

type metadataDisableAlarmActionsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DisableAlarmActionsOutput struct {
	metadataDisableAlarmActionsOutput `json:"-", xml:"-"`
}

type metadataDisableAlarmActionsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EnableAlarmActionsInput struct {
	// The names of the alarms to enable actions for.
	AlarmNames []*string `type:"list" required:"true"`

	metadataEnableAlarmActionsInput `json:"-", xml:"-"`
}

type metadataEnableAlarmActionsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EnableAlarmActionsOutput struct {
	metadataEnableAlarmActionsOutput `json:"-", xml:"-"`
}

type metadataEnableAlarmActionsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetMetricStatisticsInput struct {
	// A list of dimensions describing qualities of the metric.
	Dimensions []*Dimension `type:"list"`

	// The time stamp to use for determining the last datapoint to return. The value
	// specified is exclusive; results will include datapoints up to the time stamp
	// specified.
	EndTime *time.Time `type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The name of the metric, with or without spaces.
	MetricName *string `type:"string" required:"true"`

	// The namespace of the metric, with or without spaces.
	Namespace *string `type:"string" required:"true"`

	// The granularity, in seconds, of the returned datapoints. Period must be at
	// least 60 seconds and must be a multiple of 60. The default value is 60.
	Period *int64 `type:"integer" required:"true"`

	// The time stamp to use for determining the first datapoint to return. The
	// value specified is inclusive; results include datapoints with the time stamp
	// specified.
	StartTime *time.Time `type:"timestamp" timestampFormat:"iso8601" required:"true"`

	// The metric statistics to return. For information about specific statistics
	// returned by GetMetricStatistics, go to Statistics (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/index.html?CHAP_TerminologyandKeyConcepts.html#Statistic)
	// in the Amazon CloudWatch Developer Guide.
	//
	//  Valid Values: Average | Sum | SampleCount | Maximum | Minimum
	Statistics []*string `type:"list" required:"true"`

	// The unit for the metric.
	Unit *string `type:"string"`

	metadataGetMetricStatisticsInput `json:"-", xml:"-"`
}

type metadataGetMetricStatisticsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the GetMetricStatistics action.
type GetMetricStatisticsOutput struct {
	// The datapoints for the specified metric.
	Datapoints []*Datapoint `type:"list"`

	// A label describing the specified metric.
	Label *string `type:"string"`

	metadataGetMetricStatisticsOutput `json:"-", xml:"-"`
}

type metadataGetMetricStatisticsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

const (
	HistoryItemTypeAction              = "Action"
	HistoryItemTypeConfigurationUpdate = "ConfigurationUpdate"
	HistoryItemTypeStateUpdate         = "StateUpdate"
)

type ListMetricsInput struct {
	// A list of dimensions to filter against.
	Dimensions []*DimensionFilter `type:"list"`

	// The name of the metric to filter against.
	MetricName *string `type:"string"`

	// The namespace to filter against.
	Namespace *string `type:"string"`

	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string `type:"string"`

	metadataListMetricsInput `json:"-", xml:"-"`
}

type metadataListMetricsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The output for the ListMetrics action.
type ListMetricsOutput struct {
	// A list of metrics used to generate statistics for an AWS account.
	Metrics []*Metric `type:"list"`

	// A string that marks the start of the next batch of returned results.
	NextToken *string `type:"string"`

	metadataListMetricsOutput `json:"-", xml:"-"`
}

type metadataListMetricsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

// The Metric data type contains information about a specific metric. If you
// call ListMetrics, Amazon CloudWatch returns information contained by this
// data type.
//
//  The example in the Examples section publishes two metrics named buffers
// and latency. Both metrics are in the examples namespace. Both metrics have
// two dimensions, InstanceID and InstanceType.
type Metric struct {
	// A list of dimensions associated with the metric.
	Dimensions []*Dimension `type:"list"`

	// The name of the metric.
	MetricName *string `type:"string"`

	// The namespace of the metric.
	Namespace *string `type:"string"`

	metadataMetric `json:"-", xml:"-"`
}

type metadataMetric struct {
	SDKShapeTraits bool `type:"structure"`
}

// The MetricAlarm data type represents an alarm. You can use PutMetricAlarm
// to create or update an alarm.
type MetricAlarm struct {
	// Indicates whether actions should be executed during any changes to the alarm's
	// state.
	ActionsEnabled *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the alarm.
	AlarmARN *string `locationName:"AlarmArn" type:"string"`

	// The list of actions to execute when this alarm transitions into an ALARM
	// state from any other state. Each action is specified as an Amazon Resource
	// Number (ARN). Currently the only actions supported are publishing to an Amazon
	// SNS topic and triggering an Auto Scaling policy.
	AlarmActions []*string `type:"list"`

	// The time stamp of the last update to the alarm configuration. Amazon CloudWatch
	// uses Coordinated Universal Time (UTC) when returning time stamps, which do
	// not accommodate seasonal adjustments such as daylight savings time. For more
	// information, see Time stamps (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp)
	// in the Amazon CloudWatch Developer Guide.
	AlarmConfigurationUpdatedTimestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The description for the alarm.
	AlarmDescription *string `type:"string"`

	// The name of the alarm.
	AlarmName *string `type:"string"`

	// The arithmetic operation to use when comparing the specified Statistic and
	// Threshold. The specified Statistic value is used as the first operand.
	ComparisonOperator *string `type:"string"`

	// The list of dimensions associated with the alarm's associated metric.
	Dimensions []*Dimension `type:"list"`

	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *int64 `type:"integer"`

	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource
	// Number (ARN). Currently the only actions supported are publishing to an Amazon
	// SNS topic or triggering an Auto Scaling policy.
	//
	// The current WSDL lists this attribute as UnknownActions.
	InsufficientDataActions []*string `type:"list"`

	// The name of the alarm's metric.
	MetricName *string `type:"string"`

	// The namespace of alarm's associated metric.
	Namespace *string `type:"string"`

	// The list of actions to execute when this alarm transitions into an OK state
	// from any other state. Each action is specified as an Amazon Resource Number
	// (ARN). Currently the only actions supported are publishing to an Amazon SNS
	// topic and triggering an Auto Scaling policy.
	OKActions []*string `type:"list"`

	// The period in seconds over which the statistic is applied.
	Period *int64 `type:"integer"`

	// A human-readable explanation for the alarm's state.
	StateReason *string `type:"string"`

	// An explanation for the alarm's state in machine-readable JSON format
	StateReasonData *string `type:"string"`

	// The time stamp of the last update to the alarm's state. Amazon CloudWatch
	// uses Coordinated Universal Time (UTC) when returning time stamps, which do
	// not accommodate seasonal adjustments such as daylight savings time. For more
	// information, see Time stamps (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp)
	// in the Amazon CloudWatch Developer Guide.
	StateUpdatedTimestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The state value for the alarm.
	StateValue *string `type:"string"`

	// The statistic to apply to the alarm's associated metric.
	Statistic *string `type:"string"`

	// The value against which the specified statistic is compared.
	Threshold *float64 `type:"double"`

	// The unit of the alarm's associated metric.
	Unit *string `type:"string"`

	metadataMetricAlarm `json:"-", xml:"-"`
}

type metadataMetricAlarm struct {
	SDKShapeTraits bool `type:"structure"`
}

// The MetricDatum data type encapsulates the information sent with PutMetricData
// to either create a new metric or add new values to be aggregated into an
// existing metric.
type MetricDatum struct {
	// A list of dimensions associated with the metric. Note, when using the Dimensions
	// value in a query, you need to append .member.N to it (e.g., Dimensions.member.N).
	Dimensions []*Dimension `type:"list"`

	// The name of the metric.
	MetricName *string `type:"string" required:"true"`

	// A set of statistical values describing the metric.
	StatisticValues *StatisticSet `type:"structure"`

	// The time stamp used for the metric. If not specified, the default value is
	// set to the time the metric data was received. Amazon CloudWatch uses Coordinated
	// Universal Time (UTC) when returning time stamps, which do not accommodate
	// seasonal adjustments such as daylight savings time. For more information,
	// see Time stamps (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp)
	// in the Amazon CloudWatch Developer Guide.
	Timestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The unit of the metric.
	Unit *string `type:"string"`

	// The value for the metric.
	//
	// Although the Value parameter accepts numbers of type Double, Amazon CloudWatch
	// truncates values with very large exponents. Values with base-10 exponents
	// greater than 126 (1 x 10^126) are truncated. Likewise, values with base-10
	// exponents less than -130 (1 x 10^-130) are also truncated.
	Value *float64 `type:"double"`

	metadataMetricDatum `json:"-", xml:"-"`
}

type metadataMetricDatum struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutMetricAlarmInput struct {
	// Indicates whether or not actions should be executed during any changes to
	// the alarm's state.
	ActionsEnabled *bool `type:"boolean"`

	// The list of actions to execute when this alarm transitions into an ALARM
	// state from any other state. Each action is specified as an Amazon Resource
	// Number (ARN). Currently the only action supported is publishing to an Amazon
	// SNS topic or an Amazon Auto Scaling policy.
	AlarmActions []*string `type:"list"`

	// The description for the alarm.
	AlarmDescription *string `type:"string"`

	// The descriptive name for the alarm. This name must be unique within the user's
	// AWS account
	AlarmName *string `type:"string" required:"true"`

	// The arithmetic operation to use when comparing the specified Statistic and
	// Threshold. The specified Statistic value is used as the first operand.
	ComparisonOperator *string `type:"string" required:"true"`

	// The dimensions for the alarm's associated metric.
	Dimensions []*Dimension `type:"list"`

	// The number of periods over which data is compared to the specified threshold.
	EvaluationPeriods *int64 `type:"integer" required:"true"`

	// The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource
	// Number (ARN). Currently the only action supported is publishing to an Amazon
	// SNS topic or an Amazon Auto Scaling policy.
	InsufficientDataActions []*string `type:"list"`

	// The name for the alarm's associated metric.
	MetricName *string `type:"string" required:"true"`

	// The namespace for the alarm's associated metric.
	Namespace *string `type:"string" required:"true"`

	// The list of actions to execute when this alarm transitions into an OK state
	// from any other state. Each action is specified as an Amazon Resource Number
	// (ARN). Currently the only action supported is publishing to an Amazon SNS
	// topic or an Amazon Auto Scaling policy.
	OKActions []*string `type:"list"`

	// The period in seconds over which the specified statistic is applied.
	Period *int64 `type:"integer" required:"true"`

	// The statistic to apply to the alarm's associated metric.
	Statistic *string `type:"string" required:"true"`

	// The value against which the specified statistic is compared.
	Threshold *float64 `type:"double" required:"true"`

	// The unit for the alarm's associated metric.
	Unit *string `type:"string"`

	metadataPutMetricAlarmInput `json:"-", xml:"-"`
}

type metadataPutMetricAlarmInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutMetricAlarmOutput struct {
	metadataPutMetricAlarmOutput `json:"-", xml:"-"`
}

type metadataPutMetricAlarmOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutMetricDataInput struct {
	// A list of data describing the metric.
	MetricData []*MetricDatum `type:"list" required:"true"`

	// The namespace for the metric data.
	Namespace *string `type:"string" required:"true"`

	metadataPutMetricDataInput `json:"-", xml:"-"`
}

type metadataPutMetricDataInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PutMetricDataOutput struct {
	metadataPutMetricDataOutput `json:"-", xml:"-"`
}

type metadataPutMetricDataOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetAlarmStateInput struct {
	// The descriptive name for the alarm. This name must be unique within the user's
	// AWS account. The maximum length is 255 characters.
	AlarmName *string `type:"string" required:"true"`

	// The reason that this alarm is set to this specific state (in human-readable
	// text format)
	StateReason *string `type:"string" required:"true"`

	// The reason that this alarm is set to this specific state (in machine-readable
	// JSON format)
	StateReasonData *string `type:"string"`

	// The value of the state.
	StateValue *string `type:"string" required:"true"`

	metadataSetAlarmStateInput `json:"-", xml:"-"`
}

type metadataSetAlarmStateInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetAlarmStateOutput struct {
	metadataSetAlarmStateOutput `json:"-", xml:"-"`
}

type metadataSetAlarmStateOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

const (
	StandardUnitBits            = "Bits"
	StandardUnitBitsSecond      = "Bits/Second"
	StandardUnitBytes           = "Bytes"
	StandardUnitBytesSecond     = "Bytes/Second"
	StandardUnitCount           = "Count"
	StandardUnitCountSecond     = "Count/Second"
	StandardUnitGigabits        = "Gigabits"
	StandardUnitGigabitsSecond  = "Gigabits/Second"
	StandardUnitGigabytes       = "Gigabytes"
	StandardUnitGigabytesSecond = "Gigabytes/Second"
	StandardUnitKilobits        = "Kilobits"
	StandardUnitKilobitsSecond  = "Kilobits/Second"
	StandardUnitKilobytes       = "Kilobytes"
	StandardUnitKilobytesSecond = "Kilobytes/Second"
	StandardUnitMegabits        = "Megabits"
	StandardUnitMegabitsSecond  = "Megabits/Second"
	StandardUnitMegabytes       = "Megabytes"
	StandardUnitMegabytesSecond = "Megabytes/Second"
	StandardUnitMicroseconds    = "Microseconds"
	StandardUnitMilliseconds    = "Milliseconds"
	StandardUnitNone            = "None"
	StandardUnitPercent         = "Percent"
	StandardUnitSeconds         = "Seconds"
	StandardUnitTerabits        = "Terabits"
	StandardUnitTerabitsSecond  = "Terabits/Second"
	StandardUnitTerabytes       = "Terabytes"
	StandardUnitTerabytesSecond = "Terabytes/Second"
)

const (
	StateValueAlarm            = "ALARM"
	StateValueInsufficientData = "INSUFFICIENT_DATA"
	StateValueOk               = "OK"
)

const (
	StatisticAverage     = "Average"
	StatisticMaximum     = "Maximum"
	StatisticMinimum     = "Minimum"
	StatisticSampleCount = "SampleCount"
	StatisticSum         = "Sum"
)

// The StatisticSet data type describes the StatisticValues component of MetricDatum,
// and represents a set of statistics that describes a specific metric.
type StatisticSet struct {
	// The maximum value of the sample set.
	Maximum *float64 `type:"double" required:"true"`

	// The minimum value of the sample set.
	Minimum *float64 `type:"double" required:"true"`

	// The number of samples used for the statistic set.
	SampleCount *float64 `type:"double" required:"true"`

	// The sum of values for the sample set.
	Sum *float64 `type:"double" required:"true"`

	metadataStatisticSet `json:"-", xml:"-"`
}

type metadataStatisticSet struct {
	SDKShapeTraits bool `type:"structure"`
}