// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen github.com/dragonator/goa-framework-service/design

package client

import (
	"encoding/json"
	"fmt"

	calc "github.com/dragonator/goa-framework-service/gen/calc"
	calcpb "github.com/dragonator/goa-framework-service/gen/grpc/calc/pb"
)

// BuildMultiplyPayload builds the payload for the calc multiply endpoint from
// CLI flags.
func BuildMultiplyPayload(calcMultiplyMessage string) (*calc.MultiplyPayload, error) {
	var err error
	var message calcpb.MultiplyRequest
	{
		if calcMultiplyMessage != "" {
			err = json.Unmarshal([]byte(calcMultiplyMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"a\": 8399553735696626949,\n      \"b\": 360622074634248926\n   }'")
			}
		}
	}
	v := &calc.MultiplyPayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}
