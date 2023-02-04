package grpcerror

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCStatus interface {
	GRPCStatus() *status.Status
}

type Translated interface {
	TranslatedStatus() *status.Status
}

type AppendDetail func(*status.Status) *status.Status

func appendDetail(detailName string, ri proto.Message) AppendDetail {
	return func(s *status.Status) *status.Status {
		details, err := s.WithDetails(ri)
		if err != nil {
			return status.New(s.Code(), fmt.Sprintf("convertion(%s) failed(%v): %v", detailName, err, s.Err()))
		}
		return details
	}
}

func statusNew(code codes.Code, msg string, aps ...AppendDetail) *status.Status {
	s := status.New(code, msg)
	for _, ap := range aps {
		s = ap(s)
	}
	return s
}

type TranslatedError struct {
	Translated *status.Status
	Original   error
}

func (te *TranslatedError) TranslatedStatus() *status.Status {
	return te.Translated
}

func (te *TranslatedError) GRPCStatus() *status.Status {
	return te.Translated
}

func (te *TranslatedError) Error() string {
	return fmt.Sprintf("error translated as %v. orignal = %v", te.Translated.Code(), te.Original)
}

func Translate(err error, st *status.Status) *TranslatedError {
	return &TranslatedError{
		Translated: st,
		Original:   err,
	}
}

func AsIs(st *status.Status) *TranslatedError {
	return &TranslatedError{
		Translated: st,
		Original:   st.Err(),
	}
}
