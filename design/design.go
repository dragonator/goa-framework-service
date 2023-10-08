package design

import (
	dsl "goa.design/goa/v3/dsl"
)

var _ = dsl.API("calc", func() {
	dsl.Title("Calculator Service")
	dsl.Description("Service for multiplying numbers, a Goa teaser")
	dsl.Server("calc", func() {
		dsl.Host("localhost", func() {
			dsl.URI("http://localhost:8000")
			dsl.URI("grpc://localhost:8080")
		})
	})
})

var MultiplyPayload = dsl.Type("MultiplyPayload", func() {
	dsl.Field(1, "a", dsl.Int, "Left operand")
	dsl.Field(2, "b", dsl.Int, "Right operand")

	dsl.Required("a")
	dsl.Required("b")
})

var MultiplyResponse = dsl.Type("MultiplyResponse", func() {
	dsl.Field(1, "multiple", dsl.Int, "Result of multiplication")

	dsl.Required("multiple")
})

var _ = dsl.Service("calc", func() {
	dsl.Description("The calc service performs operations on numbers.")

	dsl.Method("multiply", func() {
		dsl.Payload(MultiplyPayload)

		dsl.Result(MultiplyResponse)

		dsl.HTTP(func() {
			dsl.POST("/multiply")
		})

		dsl.GRPC(func() {
			// suppress sonarlint warning
		})
	})

	dsl.Files("/openapi.json", "./gen/http/openapi.json")
})
