// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc gRPC server types
//
// Command:
// $ goa gen github.com/dragonator/goa-framework-service/design

package server

import (
	calc "github.com/dragonator/goa-framework-service/gen/calc"
	calcviews "github.com/dragonator/goa-framework-service/gen/calc/views"
	calcpb "github.com/dragonator/goa-framework-service/gen/grpc/calc/pb"
)

// NewMultiplyPayload builds the payload of the "multiply" endpoint of the
// "calc" service from the gRPC request type.
func NewMultiplyPayload(message *calcpb.MultiplyRequest) *calc.MultiplyPayload {
	v := &calc.MultiplyPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewProtoMultiplyResponse builds the gRPC response type from the result of
// the "multiply" endpoint of the "calc" service.
func NewProtoMultiplyResponse(result *calcviews.MultiplyresponseView) *calcpb.MultiplyResponse {
	message := &calcpb.MultiplyResponse{
		Multiple: int32(*result.Multiple),
	}
	return message
}
