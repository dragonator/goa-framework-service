// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc gRPC server encoders and decoders
//
// Command:
// $ goa gen github.com/dragonator/goa-framework-service/design

package server

import (
	"context"

	calc "github.com/dragonator/goa-framework-service/gen/calc"
	calcviews "github.com/dragonator/goa-framework-service/gen/calc/views"
	calcpb "github.com/dragonator/goa-framework-service/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeMultiplyResponse encodes responses from the "calc" service "multiply"
// endpoint.
func EncodeMultiplyResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	vres, ok := v.(*calcviews.Multiplyresponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calcviews.Multiplyresponse", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoMultiplyResponse(result)
	return resp, nil
}

// DecodeMultiplyRequest decodes requests sent to "calc" service "multiply"
// endpoint.
func DecodeMultiplyRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *calcpb.MultiplyRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.MultiplyRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "multiply", "*calcpb.MultiplyRequest", v)
		}
	}
	var payload *calc.MultiplyPayload
	{
		payload = NewMultiplyPayload(message)
	}
	return payload, nil
}
