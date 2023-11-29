// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package bcmdataexports

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An error on the server occurred during the processing of your request. Try
	// again later.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified Amazon Resource Name (ARN) in the request doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// You've reached the limit on the number of resources you can create, or exceeded
	// the size of an individual resource.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints specified by an Amazon Web Services
	// service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalServerException":       newErrorInternalServerException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"ValidationException":           newErrorValidationException,
}
