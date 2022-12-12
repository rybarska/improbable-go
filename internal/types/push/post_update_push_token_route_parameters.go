// Code generated by go-swagger; DO NOT EDIT.

package push

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"improbable-module/internal/types"
)

// NewPostUpdatePushTokenRouteParams creates a new PostUpdatePushTokenRouteParams object
// no default values defined in spec.
func NewPostUpdatePushTokenRouteParams() PostUpdatePushTokenRouteParams {

	return PostUpdatePushTokenRouteParams{}
}

// PostUpdatePushTokenRouteParams contains all the bound params for the post update push token route operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostUpdatePushTokenRoute
type PostUpdatePushTokenRouteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Payload *types.PostUpdatePushTokenPayload
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostUpdatePushTokenRouteParams() beforehand.
func (o *PostUpdatePushTokenRouteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body types.PostUpdatePushTokenPayload
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("payload", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Payload = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostUpdatePushTokenRouteParams) Validate(formats strfmt.Registry) error {
	var res []error

	// Payload
	// Required: false

	// body is validated in endpoint
	//if err := o.Payload.Validate(formats); err != nil {
	//  res = append(res, err)
	//}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
