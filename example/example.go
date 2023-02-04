package example

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"github.com/a-r-g-v/grpcerror"
)

func mayOccurGRPCError() (interface{}, error) {
	return nil, nil
}

func Test() (interface{}, error) {
	resp, err := mayOccurGRPCError()
	if err != nil {
		return nil, grpcerror.InvalidArgumentError("invalid argument", grpcerror.BadRequest(
			&errdetails.BadRequest{
				FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{
						Field:       "name",
						Description: "len(name) must be [1, 32]",
					},
				},
			}))
	}
	_ = resp
}
