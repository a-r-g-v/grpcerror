package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

var allCodes = map[string]codes.Code{
	"Canceled":           codes.Canceled,
	"DeadlineExceeded":   codes.DeadlineExceeded,
	"Unknown":            codes.Unknown,
	"InvalidArgument":    codes.InvalidArgument,
	"NotFound":           codes.NotFound,
	"AlreadyExists":      codes.AlreadyExists,
	"PermissionDenied":   codes.PermissionDenied,
	"ResourceExhausted":  codes.ResourceExhausted,
	"FailedPrecondition": codes.FailedPrecondition,
	"Aborted":            codes.Aborted,
	"OutOfRange":         codes.OutOfRange,
	"Unimplemented":      codes.Unimplemented,
	"Internal":           codes.Internal,
	"Unavailable":        codes.Unavailable,
	"DataLoss":           codes.DataLoss,
	"Unauthenticated":    codes.Unauthenticated,
}

var allDetails = map[string]proto.Message{
	"RetryInfo":           &errdetails.RetryInfo{},
	"DebugInfo":           &errdetails.DebugInfo{},
	"QuotaFailure":        &errdetails.QuotaFailure{},
	"ErrorInfo":           &errdetails.ErrorInfo{},
	"PreconditionFailure": &errdetails.PreconditionFailure{},
	"BadRequest":          &errdetails.BadRequest{},
	"RequestInfo":         &errdetails.RequestInfo{},
	"ResourceInfo":        &errdetails.ResourceInfo{},
	"Help":                &errdetails.Help{},
	"LocalizedMessage":    &errdetails.LocalizedMessage{},
}

func main() {
	fmt.Printf(`
// Code generated by genhelper.go; DO NOT EDIT.
package grpcerror
import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
`)
	for name := range allCodes {
		fmt.Printf(`func %s(msg string, aps ...AppendDetail) *status.Status {
			return statusNew(codes.%s, msg, aps...)
		}
`, name, name)

		fmt.Printf(`func %sError(msg string, aps ...AppendDetail) error {
			return %s(msg, aps...).Err()
		}
`, name, name)
	}

	for name := range allDetails {
		fmt.Printf(`func %s(ri *errdetails.%s) AppendDetail {
	return appendDetail("%s", ri)
}
`, name, name, name)
	}
}