package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/go-swagger/go-swagger/examples/generated/models"
)

// NewCreateUsersWithListInputParams creates a new CreateUsersWithListInputParams object
// with the default values initialized.
func NewCreateUsersWithListInputParams() CreateUsersWithListInputParams {
	var ()
	return CreateUsersWithListInputParams{}
}

// CreateUsersWithListInputParams contains all the bound params for the create users with list input operation
// typically these are obtained from a http.Request
//
// swagger:parameters createUsersWithListInput
type CreateUsersWithListInputParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*List of user object
	  In: body
	*/
	Body []*models.User
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateUsersWithListInputParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []*models.User
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			for _, io := range o.Body {
				if err := io.Validate(route.Formats); err != nil {
					res = append(res, err)
					break
				}
			}

			if len(res) == 0 {
				o.Body = body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
