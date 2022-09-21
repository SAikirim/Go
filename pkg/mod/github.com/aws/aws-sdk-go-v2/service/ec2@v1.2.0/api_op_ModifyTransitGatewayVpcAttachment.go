// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified VPC attachment.
func (c *Client) ModifyTransitGatewayVpcAttachment(ctx context.Context, params *ModifyTransitGatewayVpcAttachmentInput, optFns ...func(*Options)) (*ModifyTransitGatewayVpcAttachmentOutput, error) {
	if params == nil {
		params = &ModifyTransitGatewayVpcAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyTransitGatewayVpcAttachment", params, optFns, addOperationModifyTransitGatewayVpcAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyTransitGatewayVpcAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTransitGatewayVpcAttachmentInput struct {

	// The ID of the attachment.
	//
	// This member is required.
	TransitGatewayAttachmentId *string

	// The IDs of one or more subnets to add. You can specify at most one subnet per
	// Availability Zone.
	AddSubnetIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The new VPC attachment options. You cannot modify the IPv6 options.
	Options *types.ModifyTransitGatewayVpcAttachmentRequestOptions

	// The IDs of one or more subnets to remove.
	RemoveSubnetIds []string
}

type ModifyTransitGatewayVpcAttachmentOutput struct {

	// Information about the modified attachment.
	TransitGatewayVpcAttachment *types.TransitGatewayVpcAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyTransitGatewayVpcAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyTransitGatewayVpcAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyTransitGatewayVpcAttachment{}, middleware.After)
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
	if err = addOpModifyTransitGatewayVpcAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTransitGatewayVpcAttachment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyTransitGatewayVpcAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyTransitGatewayVpcAttachment",
	}
}
