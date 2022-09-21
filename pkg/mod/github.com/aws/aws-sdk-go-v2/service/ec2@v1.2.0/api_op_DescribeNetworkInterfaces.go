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

// Describes one or more of your network interfaces.
func (c *Client) DescribeNetworkInterfaces(ctx context.Context, params *DescribeNetworkInterfacesInput, optFns ...func(*Options)) (*DescribeNetworkInterfacesOutput, error) {
	if params == nil {
		params = &DescribeNetworkInterfacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNetworkInterfaces", params, optFns, addOperationDescribeNetworkInterfacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNetworkInterfacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeNetworkInterfaces.
type DescribeNetworkInterfacesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// One or more filters.
	//
	// * addresses.private-ip-address - The private IPv4
	// addresses associated with the network interface.
	//
	// * addresses.primary - Whether
	// the private IPv4 address is the primary IP address associated with the network
	// interface.
	//
	// * addresses.association.public-ip - The association ID returned when
	// the network interface was associated with the Elastic IP address (IPv4).
	//
	// *
	// addresses.association.owner-id - The owner ID of the addresses associated with
	// the network interface.
	//
	// * association.association-id - The association ID
	// returned when the network interface was associated with an IPv4 address.
	//
	// *
	// association.allocation-id - The allocation ID returned when you allocated the
	// Elastic IP address (IPv4) for your network interface.
	//
	// * association.ip-owner-id
	// - The owner of the Elastic IP address (IPv4) associated with the network
	// interface.
	//
	// * association.public-ip - The address of the Elastic IP address
	// (IPv4) bound to the network interface.
	//
	// * association.public-dns-name - The
	// public DNS name for the network interface (IPv4).
	//
	// * attachment.attachment-id -
	// The ID of the interface attachment.
	//
	// * attachment.attach-time - The time that
	// the network interface was attached to an instance.
	//
	// *
	// attachment.delete-on-termination - Indicates whether the attachment is deleted
	// when an instance is terminated.
	//
	// * attachment.device-index - The device index to
	// which the network interface is attached.
	//
	// * attachment.instance-id - The ID of
	// the instance to which the network interface is attached.
	//
	// *
	// attachment.instance-owner-id - The owner ID of the instance to which the network
	// interface is attached.
	//
	// * attachment.status - The status of the attachment
	// (attaching | attached | detaching | detached).
	//
	// * availability-zone - The
	// Availability Zone of the network interface.
	//
	// * description - The description of
	// the network interface.
	//
	// * group-id - The ID of a security group associated with
	// the network interface.
	//
	// * group-name - The name of a security group associated
	// with the network interface.
	//
	// * ipv6-addresses.ipv6-address - An IPv6 address
	// associated with the network interface.
	//
	// * mac-address - The MAC address of the
	// network interface.
	//
	// * network-interface-id - The ID of the network interface.
	//
	// *
	// owner-id - The AWS account ID of the network interface owner.
	//
	// *
	// private-ip-address - The private IPv4 address or addresses of the network
	// interface.
	//
	// * private-dns-name - The private DNS name of the network interface
	// (IPv4).
	//
	// * requester-id - The alias or AWS account ID of the principal or
	// service that created the network interface.
	//
	// * requester-managed - Indicates
	// whether the network interface is being managed by an AWS service (for example,
	// AWS Management Console, Auto Scaling, and so on).
	//
	// * source-dest-check -
	// Indicates whether the network interface performs source/destination checking. A
	// value of true means checking is enabled, and false means checking is disabled.
	// The value must be false for the network interface to perform network address
	// translation (NAT) in your VPC.
	//
	// * status - The status of the network interface.
	// If the network interface is not attached to an instance, the status is
	// available; if a network interface is attached to an instance the status is
	// in-use.
	//
	// * subnet-id - The ID of the subnet for the network interface.
	//
	// * tag: -
	// The key/value combination of a tag assigned to the resource. Use the tag key in
	// the filter name and the tag value as the filter value. For example, to find all
	// resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	//
	// * tag-key - The
	// key of a tag assigned to the resource. Use this filter to find all resources
	// assigned a tag with a specific key, regardless of the tag value.
	//
	// * vpc-id - The
	// ID of the VPC for the network interface.
	Filters []types.Filter

	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// You cannot specify this parameter and the network interface IDs parameter in the
	// same request.
	MaxResults int32

	// One or more network interface IDs. Default: Describes all your network
	// interfaces.
	NetworkInterfaceIds []string

	// The token to retrieve the next page of results.
	NextToken *string
}

// Contains the output of DescribeNetworkInterfaces.
type DescribeNetworkInterfacesOutput struct {

	// Information about one or more network interfaces.
	NetworkInterfaces []types.NetworkInterface

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeNetworkInterfacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeNetworkInterfaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNetworkInterfaces{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNetworkInterfaces(options.Region), middleware.Before); err != nil {
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

// DescribeNetworkInterfacesAPIClient is a client that implements the
// DescribeNetworkInterfaces operation.
type DescribeNetworkInterfacesAPIClient interface {
	DescribeNetworkInterfaces(context.Context, *DescribeNetworkInterfacesInput, ...func(*Options)) (*DescribeNetworkInterfacesOutput, error)
}

var _ DescribeNetworkInterfacesAPIClient = (*Client)(nil)

// DescribeNetworkInterfacesPaginatorOptions is the paginator options for
// DescribeNetworkInterfaces
type DescribeNetworkInterfacesPaginatorOptions struct {
	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// You cannot specify this parameter and the network interface IDs parameter in the
	// same request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNetworkInterfacesPaginator is a paginator for DescribeNetworkInterfaces
type DescribeNetworkInterfacesPaginator struct {
	options   DescribeNetworkInterfacesPaginatorOptions
	client    DescribeNetworkInterfacesAPIClient
	params    *DescribeNetworkInterfacesInput
	nextToken *string
	firstPage bool
}

// NewDescribeNetworkInterfacesPaginator returns a new
// DescribeNetworkInterfacesPaginator
func NewDescribeNetworkInterfacesPaginator(client DescribeNetworkInterfacesAPIClient, params *DescribeNetworkInterfacesInput, optFns ...func(*DescribeNetworkInterfacesPaginatorOptions)) *DescribeNetworkInterfacesPaginator {
	if params == nil {
		params = &DescribeNetworkInterfacesInput{}
	}

	options := DescribeNetworkInterfacesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNetworkInterfacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNetworkInterfacesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeNetworkInterfaces page.
func (p *DescribeNetworkInterfacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNetworkInterfacesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeNetworkInterfaces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeNetworkInterfaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNetworkInterfaces",
	}
}