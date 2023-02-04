package interceptor_test

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/a-r-g-v/grpcerror"
	"github.com/a-r-g-v/grpcerror/interceptor"
)

func Test_MapTranslator(t *testing.T) {
	type testCase struct {
		in           error
		translateMap map[codes.Code]codes.Code

		want codes.Code
	}
	for name, tt := range map[string]testCase{
		"non-grpc error must be Unknown": {
			in: fmt.Errorf("aaa"),

			want: codes.Unknown,
		},
		"grpc errors that exists in translateMap must be auto translated": {
			in: grpcerror.InvalidArgumentError("invalid argument"),
			translateMap: map[codes.Code]codes.Code{
				codes.InvalidArgument: codes.Internal,
			},

			want: codes.Internal,
		},
		"wrapped grpc errors that exists in translateMap must be auto translated": {
			in: fmt.Errorf("xxx.yyy failed: %w", grpcerror.InvalidArgumentError("invalid argument")),
			translateMap: map[codes.Code]codes.Code{
				codes.InvalidArgument: codes.Internal,
			},

			want: codes.Internal,
		},
		"grpc errors that doesn't exist in translateMap must be Unknown": {
			in: grpcerror.InvalidArgumentError("invalid argument"),

			want: codes.Unknown,
		},
		"translated error must be keeps translated error": {
			in:   grpcerror.Translate(grpcerror.InvalidArgumentError("invalid_argument"), grpcerror.NotFound("not found")),
			want: codes.NotFound,
		},
		"wrapped translated error must be keeps translated error": {
			in:   fmt.Errorf("%w", grpcerror.Translate(grpcerror.InvalidArgumentError("invalid_argument"), grpcerror.NotFound("not found"))),
			want: codes.NotFound,
		},
	} {
		t.Run(name, func(t *testing.T) {
			mt := &interceptor.MapTranslator{TranslateMap: tt.translateMap}
			if err := mt.Translate(tt.in); status.Code(err) != tt.want {
				t.Errorf("mismatch want %v but got %v", tt.want, status.Code(err))
			}
		})
	}
}

func Test_DefaultTranslator(t *testing.T) {
	type testCase struct {
		in error

		want codes.Code
	}
	for name, tt := range map[string]testCase{
		"Canceled": {
			in:   grpcerror.CanceledError("Canceled"),
			want: codes.Canceled,
		},
		"DeadlineExceeded": {
			in:   grpcerror.DeadlineExceededError("DeadlineExceeded"),
			want: codes.DeadlineExceeded,
		},
		"Unknown": {
			in:   grpcerror.UnknownError("Unknown"),
			want: codes.Unknown,
		},
		"InvalidArgument": {
			in:   grpcerror.InvalidArgumentError("InvalidArgument"),
			want: codes.Internal,
		},
		"NotFound": {
			in:   grpcerror.NotFoundError("NotFound"),
			want: codes.Internal,
		},
		"AlreadyExists": {
			in:   grpcerror.AlreadyExistsError("AlreadyExists"),
			want: codes.Internal,
		},
		"PermissionDenied": {
			in:   grpcerror.PermissionDeniedError("PermissionDenied"),
			want: codes.Internal,
		},
		"ResourceExhausted": {
			in:   grpcerror.ResourceExhaustedError("ResourceExhausted"),
			want: codes.Internal,
		},
		"FailedPrecondition": {
			in:   grpcerror.FailedPreconditionError("FailedPrecondition"),
			want: codes.Internal,
		},
		"Aborted": {
			in:   grpcerror.AbortedError("Aborted"),
			want: codes.Internal,
		},
		"OutOfRange": {
			in:   grpcerror.OutOfRangeError("OutOfRange"),
			want: codes.Internal,
		},
		"Unimplemented": {
			in:   grpcerror.UnimplementedError("Unimplemented"),
			want: codes.Internal,
		},
		"Internal": {
			in:   grpcerror.InternalError("Internal"),
			want: codes.Internal,
		},
		"Unavailable": {
			in:   grpcerror.UnavailableError("Unavailable"),
			want: codes.Internal,
		},
		"DataLoss": {
			in:   grpcerror.DataLossError("DataLoss"),
			want: codes.Internal,
		},
		"Unauthenticated": {
			in:   grpcerror.UnauthenticatedError("Unauthenticated"),
			want: codes.Internal,
		},
	} {
		t.Run(name, func(t *testing.T) {
			if err := interceptor.DefaultTranslator().Translate(tt.in); status.Code(err) != tt.want {
				t.Errorf("mismatch want %v but got %v", tt.want, status.Code(err))
			}
		})
	}
}

func Test_TranslateUnaryServerInterceptor(t *testing.T) {
	type testCase struct {
		handlerFunc func(ctx context.Context, req interface{}) (interface{}, error)

		wantErrorCode codes.Code
	}
	for name, tt := range map[string]testCase{
		"ok": {
			handlerFunc: func(ctx context.Context, req interface{}) (interface{}, error) {
				return "", nil
			},
			wantErrorCode: codes.OK,
		},
		"not translated error": {
			handlerFunc: func(ctx context.Context, req interface{}) (interface{}, error) {
				return "", grpcerror.InvalidArgumentError("invalid argument")
			},
			wantErrorCode: codes.Internal,
		},
		"translated error": {
			handlerFunc: func(ctx context.Context, req interface{}) (interface{}, error) {
				return "", grpcerror.Translate(grpcerror.InvalidArgumentError("invalid argument"), grpcerror.FailedPrecondition("failed"))
			},
			wantErrorCode: codes.FailedPrecondition,
		},
	} {
		t.Run(name, func(t *testing.T) {
			ti := interceptor.TranslateUnaryServerInterceptor(interceptor.DefaultTranslator())

			unaryInfo := &grpc.UnaryServerInfo{
				FullMethod: "Test.UnaryMethod",
			}

			_, err := ti(context.Background(), "test", unaryInfo, tt.handlerFunc)
			if status.Code(err) != tt.wantErrorCode {
				t.Errorf("mismatch code. want %v but got %v", tt.wantErrorCode, status.Code(err))
			}
		})
	}

}
