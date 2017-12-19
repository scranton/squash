// Code generated by go-swagger; DO NOT EDIT.

package debugattachment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// NewPatchDebugAttachmentParams creates a new PatchDebugAttachmentParams object
// with the default values initialized.
func NewPatchDebugAttachmentParams() PatchDebugAttachmentParams {
	var ()
	return PatchDebugAttachmentParams{}
}

// PatchDebugAttachmentParams contains all the bound params for the patch debug attachment operation
// typically these are obtained from a http.Request
//
// swagger:parameters patchDebugAttachment
type PatchDebugAttachmentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*DebugAttachment object
	  Required: true
	  In: body
	*/
	Body *models.DebugAttachment
	/*ID of config to return
	  Required: true
	  In: path
	*/
	DebugAttachmentID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PatchDebugAttachmentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.DebugAttachment
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}

	} else {
		res = append(res, errors.Required("body", "body"))
	}

	rDebugAttachmentID, rhkDebugAttachmentID, _ := route.Params.GetOK("debugAttachmentId")
	if err := o.bindDebugAttachmentID(rDebugAttachmentID, rhkDebugAttachmentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchDebugAttachmentParams) bindDebugAttachmentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.DebugAttachmentID = raw

	return nil
}