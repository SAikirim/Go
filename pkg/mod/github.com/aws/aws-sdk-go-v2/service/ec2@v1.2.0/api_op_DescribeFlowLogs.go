// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more flow logs. To view the information in your flow logs (the
// log streams for the network interfaces), you must use the CloudWatch Logs
// console or the CloudWatch Logs API.
func (c *Client) DescribeFlowLogs(ctx context.Context, params *DescribeFlowLogsInput, optFns ...func(*Options)) (*DescribeFlowLogsOutput, error) {
	if params == nil {
		params = &DescribeFlowLogsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFlowLogs", params, optFns, addOperationDescribeFlowLogsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFlowLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFlowLogsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * deliver-log-status - The status of the logs delivery
	// (SUCCESS | FAILED).
	//
	// * log-destination-type - The type of destination to which
	// the flow log publishes data. Possible destination types include cloud-watch-logs
	// and s3.
	//
	// * flow-log-id - The ID of the flow log.
	//
	// * log-group-name - The name of
	// the log group.
	//
	// * resource-id - The ID of the VPC, subnet, or network
	// interface.
	//
	// * traffic-type - The type of traffic (ACCEPT | REJECT | ALL).
	//
	// *
	// tag: - The key/value combination of a tag assigned to the resource. Use the tag
	// key in the filter name and the tag value as the filter value. For example, to
	// find all resources that have a tag with the key Owner and the value TeamA,
	// specify tag:Owner for the filter name and TeamA for the filter value.
	//
	// * tag-key
	// - The key of a tag assigned to the resource. Use this filter to find all
	// resources assigned a tag with a specific key, regardless of the tag value.
	Filter []types.Filter

	// One or more flow log IDs. Constraint: Maximum of 1000 flow log IDs.
	FlowLogIds []string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeFlowLogsOutput struct {

	// Information about the flow logs.
	FlowLogs []types.FlowLog

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFlowLogsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeFlowLogs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFlowLogs{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFlowLogs(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeFlowLogsAPIClient is a client that implements the DescribeFlowLogs
// operation.
type DescribeFlowLogsAPIClient interface {
	DescribeFlowLogs(context.Context, *DescribeFlowLogsInput, ...func(*Options)) (*DescribeFlowLogsOutput, error)
}

var _ DescribeFlowLogsAPIClient = (*Client)(nil)

// DescribeFlowLogsPaginatorOptions is the paginator options for DescribeFlowLogs
type DescribeFlowLogsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFlowLogsPaginator is a paginator for DescribeFlowLogs
type DescribeFlowLogsPaginator struct {
	options   DescribeFlowLogsPaginatorOptions
	client    DescribeFlowLogsAPIClient
	params    *DescribeFlowLogsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFlowLogsPaginator returns a new DescribeFlowLogsPaginator
func NewDescribeFlowLogsPaginator(client DescribeFlowLogsAPIClient, params *DescribeFlowLogsInput, optFns ...func(*DescribeFlowLogsPaginatorOptions)) *DescribeFlowLogsPaginator {
	if params == nil {
		params = &DescribeFlowLogsInput{}
	}

	options := DescribeFlowLogsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFlowLogsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFlowLogsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeFlowLogs page.
func (p *DescribeFlowLogsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFlowLogsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeFlowLogs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeFlowLogs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeFlowLogs",
	}
}
