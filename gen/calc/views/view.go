// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc views
//
// Command:
// $ goa gen github.com/dragonator/goa-framework-service/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Multiplyresponse is the viewed result type that is projected based on a view.
type Multiplyresponse struct {
	// Type to project
	Projected *MultiplyresponseView
	// View to render
	View string
}

// MultiplyresponseView is a type that runs validations on a projected type.
type MultiplyresponseView struct {
	// Result of multiplication
	Multiple *int
}

var (
	// MultiplyresponseMap is a map indexing the attribute names of
	// Multiplyresponse by view name.
	MultiplyresponseMap = map[string][]string{
		"default": {
			"multiple",
		},
	}
)

// ValidateMultiplyresponse runs the validations defined on the viewed result
// type Multiplyresponse.
func ValidateMultiplyresponse(result *Multiplyresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateMultiplyresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateMultiplyresponseView runs the validations defined on
// MultiplyresponseView using the "default" view.
func ValidateMultiplyresponseView(result *MultiplyresponseView) (err error) {
	if result.Multiple == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("multiple", "result"))
	}
	return
}
