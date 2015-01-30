// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package importexport provides a client for AWS Import/Export.
package importexport

import (
	"net/http"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/endpoints"
)

import (
	"encoding/xml"
	"io"
)

// ImportExport is a client for AWS Import/Export.
type ImportExport struct {
	client *aws.QueryClient
}

type ImportExportAPI interface {
	CancelJob(req *CancelJobInput) (resp *CancelJobOutput, err error)
	CreateJob(req *CreateJobInput) (resp *CreateJobOutput, err error)
	GetStatus(req *GetStatusInput) (resp *GetStatusOutput, err error)
	ListJobs(req *ListJobsInput) (resp *ListJobsOutput, err error)
	UpdateJob(req *UpdateJobInput) (resp *UpdateJobOutput, err error)
}

// New returns a new ImportExport client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *ImportExport {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("importexport", region)

	return &ImportExport{
		client: &aws.QueryClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2010-06-01",
		},
	}
}

// CancelJob this operation cancels a specified job. Only the job owner can
// cancel it. The operation fails if the job has already started or is
// complete.
func (c *ImportExport) CancelJob(req *CancelJobInput) (resp *CancelJobResult, err error) {
	resp = &CancelJobResult{}
	err = c.client.Do("CancelJob", "POST", "/?Operation=CancelJob", req, resp)
	return
}

// CreateJob this operation initiates the process of scheduling an upload
// or download of your data. You include in the request a manifest that
// describes the data transfer specifics. The response to the request
// includes a job ID, which you can use in other operations, a signature
// that you use to identify your storage device, and the address where you
// should ship your storage device.
func (c *ImportExport) CreateJob(req *CreateJobInput) (resp *CreateJobResult, err error) {
	resp = &CreateJobResult{}
	err = c.client.Do("CreateJob", "POST", "/?Operation=CreateJob", req, resp)
	return
}

// GetStatus this operation returns information about a job, including
// where the job is in the processing pipeline, the status of the results,
// and the signature value associated with the job. You can only return
// information about jobs you own.
func (c *ImportExport) GetStatus(req *GetStatusInput) (resp *GetStatusResult, err error) {
	resp = &GetStatusResult{}
	err = c.client.Do("GetStatus", "POST", "/?Operation=GetStatus", req, resp)
	return
}

// ListJobs this operation returns the jobs associated with the requester.
// AWS Import/Export lists the jobs in reverse chronological order based on
// the date of creation. For example if Job Test1 was created 2009Dec30 and
// Test2 was created 2010Feb05, the ListJobs operation would return Test2
// followed by Test1.
func (c *ImportExport) ListJobs(req *ListJobsInput) (resp *ListJobsResult, err error) {
	resp = &ListJobsResult{}
	err = c.client.Do("ListJobs", "POST", "/?Operation=ListJobs", req, resp)
	return
}

// UpdateJob you use this operation to change the parameters specified in
// the original manifest file by supplying a new manifest file. The
// manifest file attached to this request replaces the original manifest
// file. You can only use the operation after a CreateJob request but
// before the data transfer starts and you can only use it on jobs you own.
func (c *ImportExport) UpdateJob(req *UpdateJobInput) (resp *UpdateJobResult, err error) {
	resp = &UpdateJobResult{}
	err = c.client.Do("UpdateJob", "POST", "/?Operation=UpdateJob", req, resp)
	return
}

// CancelJobInput is undocumented.
type CancelJobInput struct {
	JobID aws.StringValue `query:"JobId" xml:"JobId"`
}

// CancelJobOutput is undocumented.
type CancelJobOutput struct {
	Success aws.BooleanValue `query:"Success" xml:"CancelJobResult>Success"`
}

// CreateJobInput is undocumented.
type CreateJobInput struct {
	JobType          aws.StringValue  `query:"JobType" xml:"JobType"`
	Manifest         aws.StringValue  `query:"Manifest" xml:"Manifest"`
	ManifestAddendum aws.StringValue  `query:"ManifestAddendum" xml:"ManifestAddendum"`
	ValidateOnly     aws.BooleanValue `query:"ValidateOnly" xml:"ValidateOnly"`
}

// CreateJobOutput is undocumented.
type CreateJobOutput struct {
	AWSShippingAddress    aws.StringValue `query:"AwsShippingAddress" xml:"CreateJobResult>AwsShippingAddress"`
	JobID                 aws.StringValue `query:"JobId" xml:"CreateJobResult>JobId"`
	JobType               aws.StringValue `query:"JobType" xml:"CreateJobResult>JobType"`
	Signature             aws.StringValue `query:"Signature" xml:"CreateJobResult>Signature"`
	SignatureFileContents aws.StringValue `query:"SignatureFileContents" xml:"CreateJobResult>SignatureFileContents"`
	WarningMessage        aws.StringValue `query:"WarningMessage" xml:"CreateJobResult>WarningMessage"`
}

// GetStatusInput is undocumented.
type GetStatusInput struct {
	JobID aws.StringValue `query:"JobId" xml:"JobId"`
}

// GetStatusOutput is undocumented.
type GetStatusOutput struct {
	AWSShippingAddress    aws.StringValue  `query:"AwsShippingAddress" xml:"GetStatusResult>AwsShippingAddress"`
	Carrier               aws.StringValue  `query:"Carrier" xml:"GetStatusResult>Carrier"`
	CreationDate          time.Time        `query:"CreationDate" xml:"GetStatusResult>CreationDate"`
	CurrentManifest       aws.StringValue  `query:"CurrentManifest" xml:"GetStatusResult>CurrentManifest"`
	ErrorCount            aws.IntegerValue `query:"ErrorCount" xml:"GetStatusResult>ErrorCount"`
	JobID                 aws.StringValue  `query:"JobId" xml:"GetStatusResult>JobId"`
	JobType               aws.StringValue  `query:"JobType" xml:"GetStatusResult>JobType"`
	LocationCode          aws.StringValue  `query:"LocationCode" xml:"GetStatusResult>LocationCode"`
	LocationMessage       aws.StringValue  `query:"LocationMessage" xml:"GetStatusResult>LocationMessage"`
	LogBucket             aws.StringValue  `query:"LogBucket" xml:"GetStatusResult>LogBucket"`
	LogKey                aws.StringValue  `query:"LogKey" xml:"GetStatusResult>LogKey"`
	ProgressCode          aws.StringValue  `query:"ProgressCode" xml:"GetStatusResult>ProgressCode"`
	ProgressMessage       aws.StringValue  `query:"ProgressMessage" xml:"GetStatusResult>ProgressMessage"`
	Signature             aws.StringValue  `query:"Signature" xml:"GetStatusResult>Signature"`
	SignatureFileContents aws.StringValue  `query:"SignatureFileContents" xml:"GetStatusResult>SignatureFileContents"`
	TrackingNumber        aws.StringValue  `query:"TrackingNumber" xml:"GetStatusResult>TrackingNumber"`
}

// Job is undocumented.
type Job struct {
	CreationDate time.Time        `query:"CreationDate" xml:"CreationDate"`
	IsCanceled   aws.BooleanValue `query:"IsCanceled" xml:"IsCanceled"`
	JobID        aws.StringValue  `query:"JobId" xml:"JobId"`
	JobType      aws.StringValue  `query:"JobType" xml:"JobType"`
}

// Possible values for ImportExport.
const (
	JobTypeExport = "Export"
	JobTypeImport = "Import"
)

// ListJobsInput is undocumented.
type ListJobsInput struct {
	Marker  aws.StringValue  `query:"Marker" xml:"Marker"`
	MaxJobs aws.IntegerValue `query:"MaxJobs" xml:"MaxJobs"`
}

// ListJobsOutput is undocumented.
type ListJobsOutput struct {
	IsTruncated aws.BooleanValue `query:"IsTruncated" xml:"ListJobsResult>IsTruncated"`
	Jobs        []Job            `query:"Jobs.member" xml:"ListJobsResult>Jobs>member"`
}

// UpdateJobInput is undocumented.
type UpdateJobInput struct {
	JobID        aws.StringValue  `query:"JobId" xml:"JobId"`
	JobType      aws.StringValue  `query:"JobType" xml:"JobType"`
	Manifest     aws.StringValue  `query:"Manifest" xml:"Manifest"`
	ValidateOnly aws.BooleanValue `query:"ValidateOnly" xml:"ValidateOnly"`
}

// UpdateJobOutput is undocumented.
type UpdateJobOutput struct {
	Success        aws.BooleanValue `query:"Success" xml:"UpdateJobResult>Success"`
	WarningMessage aws.StringValue  `query:"WarningMessage" xml:"UpdateJobResult>WarningMessage"`
}

// CancelJobResult is a wrapper for CancelJobOutput.
type CancelJobResult struct {
	Success aws.BooleanValue `query:"Success" xml:"CancelJobResult>Success"`
}

// CreateJobResult is a wrapper for CreateJobOutput.
type CreateJobResult struct {
	AWSShippingAddress    aws.StringValue `query:"AwsShippingAddress" xml:"CreateJobResult>AwsShippingAddress"`
	JobID                 aws.StringValue `query:"JobId" xml:"CreateJobResult>JobId"`
	JobType               aws.StringValue `query:"JobType" xml:"CreateJobResult>JobType"`
	Signature             aws.StringValue `query:"Signature" xml:"CreateJobResult>Signature"`
	SignatureFileContents aws.StringValue `query:"SignatureFileContents" xml:"CreateJobResult>SignatureFileContents"`
	WarningMessage        aws.StringValue `query:"WarningMessage" xml:"CreateJobResult>WarningMessage"`
}

// GetStatusResult is a wrapper for GetStatusOutput.
type GetStatusResult struct {
	AWSShippingAddress    aws.StringValue  `query:"AwsShippingAddress" xml:"GetStatusResult>AwsShippingAddress"`
	Carrier               aws.StringValue  `query:"Carrier" xml:"GetStatusResult>Carrier"`
	CreationDate          time.Time        `query:"CreationDate" xml:"GetStatusResult>CreationDate"`
	CurrentManifest       aws.StringValue  `query:"CurrentManifest" xml:"GetStatusResult>CurrentManifest"`
	ErrorCount            aws.IntegerValue `query:"ErrorCount" xml:"GetStatusResult>ErrorCount"`
	JobID                 aws.StringValue  `query:"JobId" xml:"GetStatusResult>JobId"`
	JobType               aws.StringValue  `query:"JobType" xml:"GetStatusResult>JobType"`
	LocationCode          aws.StringValue  `query:"LocationCode" xml:"GetStatusResult>LocationCode"`
	LocationMessage       aws.StringValue  `query:"LocationMessage" xml:"GetStatusResult>LocationMessage"`
	LogBucket             aws.StringValue  `query:"LogBucket" xml:"GetStatusResult>LogBucket"`
	LogKey                aws.StringValue  `query:"LogKey" xml:"GetStatusResult>LogKey"`
	ProgressCode          aws.StringValue  `query:"ProgressCode" xml:"GetStatusResult>ProgressCode"`
	ProgressMessage       aws.StringValue  `query:"ProgressMessage" xml:"GetStatusResult>ProgressMessage"`
	Signature             aws.StringValue  `query:"Signature" xml:"GetStatusResult>Signature"`
	SignatureFileContents aws.StringValue  `query:"SignatureFileContents" xml:"GetStatusResult>SignatureFileContents"`
	TrackingNumber        aws.StringValue  `query:"TrackingNumber" xml:"GetStatusResult>TrackingNumber"`
}

// ListJobsResult is a wrapper for ListJobsOutput.
type ListJobsResult struct {
	IsTruncated aws.BooleanValue `query:"IsTruncated" xml:"ListJobsResult>IsTruncated"`
	Jobs        []Job            `query:"Jobs.member" xml:"ListJobsResult>Jobs>member"`
}

// UpdateJobResult is a wrapper for UpdateJobOutput.
type UpdateJobResult struct {
	Success        aws.BooleanValue `query:"Success" xml:"UpdateJobResult>Success"`
	WarningMessage aws.StringValue  `query:"WarningMessage" xml:"UpdateJobResult>WarningMessage"`
}

// avoid errors if the packages aren't referenced
var _ time.Time

var _ xml.Decoder
var _ = io.EOF
