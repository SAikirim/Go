// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the specified inline policy document that is embedded with the
// specified IAM role. Policies returned by this operation are URL-encoded
// compliant with RFC 3986 (https://tools.ietf.org/html/rfc3986). You can use a URL
// decoding method to convert the policy back to plain JSON text. For example, if
// you use Java, you can use the decode method of the java.net.URLDecoder utility
// class in the Java SDK. Other languages and SDKs provide similar functionality.
// An IAM role can also have managed policies attached to it. To retrieve a managed
// policy document that is attached to a role, use GetPolicy to determine the
// policy's default version, then use GetPolicyVersion to retrieve the policy
// document. For more information about policies, see Managed policies and inline
// policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide. For more information about roles, see Using roles to
// delegate permissions and federate identities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html).
func (c *Client) GetRolePolicy(ctx context.Context, params *GetRolePolicyInput, optFns ...func(*Options)) (*GetRolePolicyOutput, error) {
	if params == nil {
		params = &GetRolePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRolePolicy", params, optFns, addOperationGetRolePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRolePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRolePolicyInput struct {

	// The name of the policy document to get. This parameter allows (through its regex
	// pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of
	// upper and lowercase alphanumeric characters with no spaces. You can also include
	// any of the following characters: _+=,.@-
	//
	// This member is required.
	PolicyName *string

	// The name of the role associated with the policy. This parameter allows (through
	// its regex pattern (http://wikipedia.org/wiki/regex)) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	RoleName *string
}

// Contains the response to a successful GetRolePolicy request.
type GetRolePolicyOutput struct {

	// The policy document. IAM stores policies in JSON format. However, resources that
	// were created using AWS CloudFormation templates can be formatted in YAML. AWS
	// CloudFormation always converts a YAML policy to JSON format before submitting it
	// to IAM.
	//
	// This member is required.
	PolicyDocument *string

	// The name of the policy.
	//
	// This member is required.
	PolicyName *string

	// The role the policy is associated with.
	//
	// This member is required.
	RoleName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRolePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetRolePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetRolePolicy{}, middleware.After)
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
	if err = addOpGetRolePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRolePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRolePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetRolePolicy",
	}
}
