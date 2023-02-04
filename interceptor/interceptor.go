package interceptor

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/a-r-g-v/grpcerror"
)

type Translator interface {
	Translate(err error) error
}

func DefaultTranslator() *MapTranslator {
	return &MapTranslator{TranslateMap: DefaultTranslateMap()}
}

type MapTranslator struct {
	TranslateMap map[codes.Code]codes.Code
}

func DefaultTranslateMap() map[codes.Code]codes.Code {
	return map[codes.Code]codes.Code{
		codes.Canceled:           codes.Canceled,
		codes.DeadlineExceeded:   codes.DeadlineExceeded,
		codes.Unknown:            codes.Unknown,
		codes.InvalidArgument:    codes.Internal,
		codes.NotFound:           codes.Internal,
		codes.AlreadyExists:      codes.Internal,
		codes.PermissionDenied:   codes.Internal,
		codes.ResourceExhausted:  codes.Internal,
		codes.FailedPrecondition: codes.Internal,
		codes.Aborted:            codes.Internal,
		codes.OutOfRange:         codes.Internal,
		codes.Unimplemented:      codes.Internal,
		codes.Internal:           codes.Internal,
		codes.Unavailable:        codes.Internal,
		codes.DataLoss:           codes.Internal,
		codes.Unauthenticated:    codes.Internal,
	}
}

func (nt *MapTranslator) Translate(err error) error {
	var ts grpcerror.Translated
	if errors.As(err, &ts) {
		v := ts.TranslatedStatus()
		msg := fmt.Sprintf("code = %v. desc = %v. err = %v", v.Code(), v.Message(), err)
		return status.Error(v.Code(), msg)
	}

	var gs grpcerror.GRPCStatus
	if errors.As(err, &gs) {
		v := gs.GRPCStatus()

		translated, ok := nt.TranslateMap[v.Code()]
		if ok {
			msg := fmt.Sprintf("error auto translated. err: %v", err)
			return status.Error(translated, msg)
		}
	}

	return status.Error(codes.Unknown, "unknown")
}

func TranslateUnaryServerInterceptor(t Translator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		return resp, t.Translate(err)
	}
}
