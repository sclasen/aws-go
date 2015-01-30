// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package lambda provides a client for Amazon Lambda.
package lambda

import (
	"net/http"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/endpoints"
)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// Lambda is a client for Amazon Lambda.
type Lambda struct {
	client *aws.RestClient
}

type LambdaAPI interface {
	AddEventSource(req *AddEventSourceRequest) (resp *EventSourceConfiguration, err error)
	DeleteFunction(req *DeleteFunctionRequest) (err error)
	GetEventSource(req *GetEventSourceRequest) (resp *EventSourceConfiguration, err error)
	GetFunction(req *GetFunctionRequest) (resp *GetFunctionResponse, err error)
	GetFunctionConfiguration(req *GetFunctionConfigurationRequest) (resp *FunctionConfiguration, err error)
	InvokeAsync(req *InvokeAsyncRequest) (resp *InvokeAsyncResponse, err error)
	ListEventSources(req *ListEventSourcesRequest) (resp *ListEventSourcesResponse, err error)
	ListFunctions(req *ListFunctionsRequest) (resp *ListFunctionsResponse, err error)
	RemoveEventSource(req *RemoveEventSourceRequest) (err error)
	UpdateFunctionConfiguration(req *UpdateFunctionConfigurationRequest) (resp *FunctionConfiguration, err error)
	UploadFunction(req *UploadFunctionRequest) (resp *FunctionConfiguration, err error)
}

// New returns a new Lambda client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *Lambda {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("lambda", region)

	return &Lambda{
		client: &aws.RestClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2014-11-11",
		},
	}
}

// AddEventSource identifies an Amazon Kinesis stream as the event source
// for an AWS Lambda function. AWS Lambda invokes the specified function
// when records are posted to the stream. This is the pull model, where AWS
// Lambda invokes the function. For more information, go to AWS LambdaL How
// it Works in the AWS Lambda Developer Guide. This association between an
// Amazon Kinesis stream and an AWS Lambda function is called the event
// source mapping. You provide the configuration information (for example,
// which stream to read from and which AWS Lambda function to invoke) for
// the event source mapping in the request body. This operation requires
// permission for the iam:PassRole action for the IAM role. It also
// requires permission for the lambda:AddEventSource action.
func (c *Lambda) AddEventSource(req *AddEventSourceRequest) (resp *EventSourceConfiguration, err error) {
	resp = &EventSourceConfiguration{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/event-source-mappings/"

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// DeleteFunction deletes the specified Lambda function code and
// configuration. This operation requires permission for the
// lambda:DeleteFunction action.
func (c *Lambda) DeleteFunction(req *DeleteFunctionRequest) (err error) {
	// NRE

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}"

	if req.FunctionName != nil {
		uri = strings.Replace(uri, "{"+"FunctionName"+"}", aws.EscapePath(*req.FunctionName), -1)
		uri = strings.Replace(uri, "{"+"FunctionName+"+"}", aws.EscapePath(*req.FunctionName), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	return
}

// GetEventSource returns configuration information for the specified event
// source mapping (see AddEventSource This operation requires permission
// for the lambda:GetEventSource action.
func (c *Lambda) GetEventSource(req *GetEventSourceRequest) (resp *EventSourceConfiguration, err error) {
	resp = &EventSourceConfiguration{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/event-source-mappings/{UUID}"

	if req.UUID != nil {
		uri = strings.Replace(uri, "{"+"UUID"+"}", aws.EscapePath(*req.UUID), -1)
		uri = strings.Replace(uri, "{"+"UUID+"+"}", aws.EscapePath(*req.UUID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// GetFunction returns the configuration information of the Lambda function
// and a presigned URL link to the .zip file you uploaded with
// UploadFunction so you can download the .zip file. Note that the URL is
// valid for up to 10 minutes. The configuration information is the same
// information you provided as parameters when uploading the function. This
// operation requires permission for the lambda:GetFunction action.
func (c *Lambda) GetFunction(req *GetFunctionRequest) (resp *GetFunctionResponse, err error) {
	resp = &GetFunctionResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}"

	if req.FunctionName != nil {
		uri = strings.Replace(uri, "{"+"FunctionName"+"}", aws.EscapePath(*req.FunctionName), -1)
		uri = strings.Replace(uri, "{"+"FunctionName+"+"}", aws.EscapePath(*req.FunctionName), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// GetFunctionConfiguration returns the configuration information of the
// Lambda function. This the same information you provided as parameters
// when uploading the function by using UploadFunction This operation
// requires permission for the lambda:GetFunctionConfiguration operation.
func (c *Lambda) GetFunctionConfiguration(req *GetFunctionConfigurationRequest) (resp *FunctionConfiguration, err error) {
	resp = &FunctionConfiguration{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}/configuration"

	if req.FunctionName != nil {
		uri = strings.Replace(uri, "{"+"FunctionName"+"}", aws.EscapePath(*req.FunctionName), -1)
		uri = strings.Replace(uri, "{"+"FunctionName+"+"}", aws.EscapePath(*req.FunctionName), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// InvokeAsync submits an invocation request to AWS Lambda. Upon receiving
// the request, Lambda executes the specified function asynchronously. To
// see the logs generated by the Lambda function execution, see the
// CloudWatch logs console. This operation requires permission for the
// lambda:InvokeAsync action.
func (c *Lambda) InvokeAsync(req *InvokeAsyncRequest) (resp *InvokeAsyncResponse, err error) {
	resp = &InvokeAsyncResponse{}

	var body io.Reader
	var contentType string

	contentType = "application/json"
	b, err := json.Marshal(req.InvokeArgs)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}/invoke-async/"

	if req.FunctionName != nil {
		uri = strings.Replace(uri, "{"+"FunctionName"+"}", aws.EscapePath(*req.FunctionName), -1)
		uri = strings.Replace(uri, "{"+"FunctionName+"+"}", aws.EscapePath(*req.FunctionName), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	resp.Status = aws.Integer(httpResp.StatusCode)

	return
}

// ListEventSources returns a list of event source mappings. For each
// mapping, the API returns configuration information (see AddEventSource
// ). You can optionally specify filters to retrieve specific event source
// mappings. This operation requires permission for the
// lambda:ListEventSources action.
func (c *Lambda) ListEventSources(req *ListEventSourcesRequest) (resp *ListEventSourcesResponse, err error) {
	resp = &ListEventSourcesResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/event-source-mappings/"

	q := url.Values{}

	if req.EventSourceARN != nil {
		q.Set("EventSource", *req.EventSourceARN)
	}

	if req.FunctionName != nil {
		q.Set("FunctionName", *req.FunctionName)
	}

	if req.Marker != nil {
		q.Set("Marker", *req.Marker)
	}

	if req.MaxItems != nil {
		q.Set("MaxItems", strconv.Itoa(*req.MaxItems))
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// ListFunctions returns a list of your Lambda functions. For each
// function, the response includes the function configuration information.
// You must use GetFunction to retrieve the code for your function. This
// operation requires permission for the lambda:ListFunctions action.
func (c *Lambda) ListFunctions(req *ListFunctionsRequest) (resp *ListFunctionsResponse, err error) {
	resp = &ListFunctionsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/"

	q := url.Values{}

	if req.Marker != nil {
		q.Set("Marker", *req.Marker)
	}

	if req.MaxItems != nil {
		q.Set("MaxItems", strconv.Itoa(*req.MaxItems))
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// RemoveEventSource removes an event source mapping. This means AWS Lambda
// will no longer invoke the function for events in the associated source.
// This operation requires permission for the lambda:RemoveEventSource
// action.
func (c *Lambda) RemoveEventSource(req *RemoveEventSourceRequest) (err error) {
	// NRE

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/event-source-mappings/{UUID}"

	if req.UUID != nil {
		uri = strings.Replace(uri, "{"+"UUID"+"}", aws.EscapePath(*req.UUID), -1)
		uri = strings.Replace(uri, "{"+"UUID+"+"}", aws.EscapePath(*req.UUID), -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	return
}

// UpdateFunctionConfiguration updates the configuration parameters for the
// specified Lambda function by using the values provided in the request.
// You provide only the parameters you want to change. This operation must
// only be used on an existing Lambda function and cannot be used to update
// the function's code. This operation requires permission for the
// lambda:UpdateFunctionConfiguration action.
func (c *Lambda) UpdateFunctionConfiguration(req *UpdateFunctionConfigurationRequest) (resp *FunctionConfiguration, err error) {
	resp = &FunctionConfiguration{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}/configuration"

	if req.FunctionName != nil {
		uri = strings.Replace(uri, "{"+"FunctionName"+"}", aws.EscapePath(*req.FunctionName), -1)
		uri = strings.Replace(uri, "{"+"FunctionName+"+"}", aws.EscapePath(*req.FunctionName), -1)
	}

	q := url.Values{}

	if req.Description != nil {
		q.Set("Description", *req.Description)
	}

	if req.Handler != nil {
		q.Set("Handler", *req.Handler)
	}

	if req.MemorySize != nil {
		q.Set("MemorySize", strconv.Itoa(*req.MemorySize))
	}

	if req.Role != nil {
		q.Set("Role", *req.Role)
	}

	if req.Timeout != nil {
		q.Set("Timeout", strconv.Itoa(*req.Timeout))
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// UploadFunction creates a new Lambda function or updates an existing
// function. The function metadata is created from the request parameters,
// and the code for the function is provided by a .zip file in the request
// body. If the function name already exists, the existing Lambda function
// is updated with the new code and metadata. This operation requires
// permission for the lambda:UploadFunction action.
func (c *Lambda) UploadFunction(req *UploadFunctionRequest) (resp *FunctionConfiguration, err error) {
	resp = &FunctionConfiguration{}

	var body io.Reader
	var contentType string

	contentType = "application/json"
	b, err := json.Marshal(req.FunctionZip)
	if err != nil {
		return
	}
	body = bytes.NewReader(b)

	uri := c.client.Endpoint + "/2014-11-13/functions/{FunctionName}"

	if req.FunctionName != nil {
		uri = strings.Replace(uri, "{"+"FunctionName"+"}", aws.EscapePath(*req.FunctionName), -1)
		uri = strings.Replace(uri, "{"+"FunctionName+"+"}", aws.EscapePath(*req.FunctionName), -1)
	}

	q := url.Values{}

	if req.Description != nil {
		q.Set("Description", *req.Description)
	}

	if req.Handler != nil {
		q.Set("Handler", *req.Handler)
	}

	if req.MemorySize != nil {
		q.Set("MemorySize", strconv.Itoa(*req.MemorySize))
	}

	if req.Mode != nil {
		q.Set("Mode", *req.Mode)
	}

	if req.Role != nil {
		q.Set("Role", *req.Role)
	}

	if req.Runtime != nil {
		q.Set("Runtime", *req.Runtime)
	}

	if req.Timeout != nil {
		q.Set("Timeout", strconv.Itoa(*req.Timeout))
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("PUT", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	if e := json.NewDecoder(httpResp.Body).Decode(resp); e != nil && e != io.EOF {
		err = e
		return
	}

	return
}

// AddEventSourceRequest is undocumented.
type AddEventSourceRequest struct {
	BatchSize    aws.IntegerValue  `json:"BatchSize,omitempty"`
	EventSource  aws.StringValue   `json:"EventSource"`
	FunctionName aws.StringValue   `json:"FunctionName"`
	Parameters   map[string]string `json:"Parameters,omitempty"`
	Role         aws.StringValue   `json:"Role"`
}

// DeleteFunctionRequest is undocumented.
type DeleteFunctionRequest struct {
	FunctionName aws.StringValue `json:"-"`
}

// EventSourceConfiguration is undocumented.
type EventSourceConfiguration struct {
	BatchSize    aws.IntegerValue  `json:"BatchSize,omitempty"`
	EventSource  aws.StringValue   `json:"EventSource,omitempty"`
	FunctionName aws.StringValue   `json:"FunctionName,omitempty"`
	IsActive     aws.BooleanValue  `json:"IsActive,omitempty"`
	LastModified time.Time         `json:"LastModified,omitempty"`
	Parameters   map[string]string `json:"Parameters,omitempty"`
	Role         aws.StringValue   `json:"Role,omitempty"`
	Status       aws.StringValue   `json:"Status,omitempty"`
	UUID         aws.StringValue   `json:"UUID,omitempty"`
}

// FunctionCodeLocation is undocumented.
type FunctionCodeLocation struct {
	Location       aws.StringValue `json:"Location,omitempty"`
	RepositoryType aws.StringValue `json:"RepositoryType,omitempty"`
}

// FunctionConfiguration is undocumented.
type FunctionConfiguration struct {
	CodeSize        aws.LongValue    `json:"CodeSize,omitempty"`
	ConfigurationID aws.StringValue  `json:"ConfigurationId,omitempty"`
	Description     aws.StringValue  `json:"Description,omitempty"`
	FunctionARN     aws.StringValue  `json:"FunctionARN,omitempty"`
	FunctionName    aws.StringValue  `json:"FunctionName,omitempty"`
	Handler         aws.StringValue  `json:"Handler,omitempty"`
	LastModified    time.Time        `json:"LastModified,omitempty"`
	MemorySize      aws.IntegerValue `json:"MemorySize,omitempty"`
	Mode            aws.StringValue  `json:"Mode,omitempty"`
	Role            aws.StringValue  `json:"Role,omitempty"`
	Runtime         aws.StringValue  `json:"Runtime,omitempty"`
	Timeout         aws.IntegerValue `json:"Timeout,omitempty"`
}

// GetEventSourceRequest is undocumented.
type GetEventSourceRequest struct {
	UUID aws.StringValue `json:"-"`
}

// GetFunctionConfigurationRequest is undocumented.
type GetFunctionConfigurationRequest struct {
	FunctionName aws.StringValue `json:"-"`
}

// GetFunctionRequest is undocumented.
type GetFunctionRequest struct {
	FunctionName aws.StringValue `json:"-"`
}

// GetFunctionResponse is undocumented.
type GetFunctionResponse struct {
	Code          *FunctionCodeLocation  `json:"Code,omitempty"`
	Configuration *FunctionConfiguration `json:"Configuration,omitempty"`
}

// InvokeAsyncRequest is undocumented.
type InvokeAsyncRequest struct {
	FunctionName aws.StringValue `json:"-"`
	InvokeArgs   []byte          `json:"InvokeArgs"`
}

// InvokeAsyncResponse is undocumented.
type InvokeAsyncResponse struct {
	Status aws.IntegerValue `json:"-"`
}

// ListEventSourcesRequest is undocumented.
type ListEventSourcesRequest struct {
	EventSourceARN aws.StringValue  `json:"-"`
	FunctionName   aws.StringValue  `json:"-"`
	Marker         aws.StringValue  `json:"-"`
	MaxItems       aws.IntegerValue `json:"-"`
}

// ListEventSourcesResponse is undocumented.
type ListEventSourcesResponse struct {
	EventSources []EventSourceConfiguration `json:"EventSources,omitempty"`
	NextMarker   aws.StringValue            `json:"NextMarker,omitempty"`
}

// ListFunctionsRequest is undocumented.
type ListFunctionsRequest struct {
	Marker   aws.StringValue  `json:"-"`
	MaxItems aws.IntegerValue `json:"-"`
}

// ListFunctionsResponse is undocumented.
type ListFunctionsResponse struct {
	Functions  []FunctionConfiguration `json:"Functions,omitempty"`
	NextMarker aws.StringValue         `json:"NextMarker,omitempty"`
}

// Possible values for Lambda.
const (
	ModeEvent = "event"
)

// RemoveEventSourceRequest is undocumented.
type RemoveEventSourceRequest struct {
	UUID aws.StringValue `json:"-"`
}

// Possible values for Lambda.
const (
	RuntimeNodejs = "nodejs"
)

// UpdateFunctionConfigurationRequest is undocumented.
type UpdateFunctionConfigurationRequest struct {
	Description  aws.StringValue  `json:"-"`
	FunctionName aws.StringValue  `json:"-"`
	Handler      aws.StringValue  `json:"-"`
	MemorySize   aws.IntegerValue `json:"-"`
	Role         aws.StringValue  `json:"-"`
	Timeout      aws.IntegerValue `json:"-"`
}

// UploadFunctionRequest is undocumented.
type UploadFunctionRequest struct {
	Description  aws.StringValue  `json:"-"`
	FunctionName aws.StringValue  `json:"-"`
	FunctionZip  []byte           `json:"FunctionZip"`
	Handler      aws.StringValue  `json:"-"`
	MemorySize   aws.IntegerValue `json:"-"`
	Mode         aws.StringValue  `json:"-"`
	Role         aws.StringValue  `json:"-"`
	Runtime      aws.StringValue  `json:"-"`
	Timeout      aws.IntegerValue `json:"-"`
}

// avoid errors if the packages aren't referenced
var _ time.Time

var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
var _ json.RawMessage
