// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/MoveAddressToVpcRequest
type MoveAddressToVpcInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The Elastic IP address.
	//
	// PublicIp is a required field
	PublicIp *string `locationName:"publicIp" type:"string" required:"true"`
}

// String returns the string representation
func (s MoveAddressToVpcInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MoveAddressToVpcInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MoveAddressToVpcInput"}

	if s.PublicIp == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicIp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/MoveAddressToVpcResult
type MoveAddressToVpcOutput struct {
	_ struct{} `type:"structure"`

	// The allocation ID for the Elastic IP address.
	AllocationId *string `locationName:"allocationId" type:"string"`

	// The status of the move of the IP address.
	Status Status `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s MoveAddressToVpcOutput) String() string {
	return awsutil.Prettify(s)
}

const opMoveAddressToVpc = "MoveAddressToVpc"

// MoveAddressToVpcRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Moves an Elastic IP address from the EC2-Classic platform to the EC2-VPC
// platform. The Elastic IP address must be allocated to your account for more
// than 24 hours, and it must not be associated with an instance. After the
// Elastic IP address is moved, it is no longer available for use in the EC2-Classic
// platform, unless you move it back using the RestoreAddressToClassic request.
// You cannot move an Elastic IP address that was originally allocated for use
// in the EC2-VPC platform to the EC2-Classic platform.
//
//    // Example sending a request using MoveAddressToVpcRequest.
//    req := client.MoveAddressToVpcRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/MoveAddressToVpc
func (c *Client) MoveAddressToVpcRequest(input *MoveAddressToVpcInput) MoveAddressToVpcRequest {
	op := &aws.Operation{
		Name:       opMoveAddressToVpc,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MoveAddressToVpcInput{}
	}

	req := c.newRequest(op, input, &MoveAddressToVpcOutput{})
	return MoveAddressToVpcRequest{Request: req, Input: input, Copy: c.MoveAddressToVpcRequest}
}

// MoveAddressToVpcRequest is the request type for the
// MoveAddressToVpc API operation.
type MoveAddressToVpcRequest struct {
	*aws.Request
	Input *MoveAddressToVpcInput
	Copy  func(*MoveAddressToVpcInput) MoveAddressToVpcRequest
}

// Send marshals and sends the MoveAddressToVpc API request.
func (r MoveAddressToVpcRequest) Send(ctx context.Context) (*MoveAddressToVpcResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &MoveAddressToVpcResponse{
		MoveAddressToVpcOutput: r.Request.Data.(*MoveAddressToVpcOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// MoveAddressToVpcResponse is the response type for the
// MoveAddressToVpc API operation.
type MoveAddressToVpcResponse struct {
	*MoveAddressToVpcOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// MoveAddressToVpc request.
func (r *MoveAddressToVpcResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
