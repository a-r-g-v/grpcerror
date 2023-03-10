// Code generated by genhelper.go; DO NOT EDIT.
package grpcerror

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AlreadyExists(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.AlreadyExists, msg, aps...)
}

func AlreadyExistsError(msg string, aps ...AppendDetail) error {
	return AlreadyExists(msg, aps...).Err()
}

func ResourceExhausted(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.ResourceExhausted, msg, aps...)
}

func ResourceExhaustedError(msg string, aps ...AppendDetail) error {
	return ResourceExhausted(msg, aps...).Err()
}

func FailedPrecondition(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.FailedPrecondition, msg, aps...)
}

func FailedPreconditionError(msg string, aps ...AppendDetail) error {
	return FailedPrecondition(msg, aps...).Err()
}

func Internal(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.Internal, msg, aps...)
}

func InternalError(msg string, aps ...AppendDetail) error {
	return Internal(msg, aps...).Err()
}

func Unavailable(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.Unavailable, msg, aps...)
}

func UnavailableError(msg string, aps ...AppendDetail) error {
	return Unavailable(msg, aps...).Err()
}

func DeadlineExceeded(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.DeadlineExceeded, msg, aps...)
}

func DeadlineExceededError(msg string, aps ...AppendDetail) error {
	return DeadlineExceeded(msg, aps...).Err()
}

func Unknown(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.Unknown, msg, aps...)
}

func UnknownError(msg string, aps ...AppendDetail) error {
	return Unknown(msg, aps...).Err()
}

func NotFound(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.NotFound, msg, aps...)
}

func NotFoundError(msg string, aps ...AppendDetail) error {
	return NotFound(msg, aps...).Err()
}

func PermissionDenied(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.PermissionDenied, msg, aps...)
}

func PermissionDeniedError(msg string, aps ...AppendDetail) error {
	return PermissionDenied(msg, aps...).Err()
}

func OutOfRange(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.OutOfRange, msg, aps...)
}

func OutOfRangeError(msg string, aps ...AppendDetail) error {
	return OutOfRange(msg, aps...).Err()
}

func Canceled(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.Canceled, msg, aps...)
}

func CanceledError(msg string, aps ...AppendDetail) error {
	return Canceled(msg, aps...).Err()
}

func InvalidArgument(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.InvalidArgument, msg, aps...)
}

func InvalidArgumentError(msg string, aps ...AppendDetail) error {
	return InvalidArgument(msg, aps...).Err()
}

func Aborted(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.Aborted, msg, aps...)
}

func AbortedError(msg string, aps ...AppendDetail) error {
	return Aborted(msg, aps...).Err()
}

func Unimplemented(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.Unimplemented, msg, aps...)
}

func UnimplementedError(msg string, aps ...AppendDetail) error {
	return Unimplemented(msg, aps...).Err()
}

func DataLoss(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.DataLoss, msg, aps...)
}

func DataLossError(msg string, aps ...AppendDetail) error {
	return DataLoss(msg, aps...).Err()
}

func Unauthenticated(msg string, aps ...AppendDetail) *status.Status {
	return statusNew(codes.Unauthenticated, msg, aps...)
}

func UnauthenticatedError(msg string, aps ...AppendDetail) error {
	return Unauthenticated(msg, aps...).Err()
}

func LocalizedMessage(ri *errdetails.LocalizedMessage) AppendDetail {
	return appendDetail("LocalizedMessage", ri)
}

func RetryInfo(ri *errdetails.RetryInfo) AppendDetail {
	return appendDetail("RetryInfo", ri)
}

func QuotaFailure(ri *errdetails.QuotaFailure) AppendDetail {
	return appendDetail("QuotaFailure", ri)
}

func PreconditionFailure(ri *errdetails.PreconditionFailure) AppendDetail {
	return appendDetail("PreconditionFailure", ri)
}

func BadRequest(ri *errdetails.BadRequest) AppendDetail {
	return appendDetail("BadRequest", ri)
}

func Help(ri *errdetails.Help) AppendDetail {
	return appendDetail("Help", ri)
}

func DebugInfo(ri *errdetails.DebugInfo) AppendDetail {
	return appendDetail("DebugInfo", ri)
}

func ErrorInfo(ri *errdetails.ErrorInfo) AppendDetail {
	return appendDetail("ErrorInfo", ri)
}

func RequestInfo(ri *errdetails.RequestInfo) AppendDetail {
	return appendDetail("RequestInfo", ri)
}

func ResourceInfo(ri *errdetails.ResourceInfo) AppendDetail {
	return appendDetail("ResourceInfo", ri)
}
