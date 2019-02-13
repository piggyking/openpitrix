// Code generated by go-swagger; DO NOT EDIT.

package access_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// NewCanDoParams creates a new CanDoParams object
// with the default values initialized.
func NewCanDoParams() *CanDoParams {
	var ()
	return &CanDoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCanDoParamsWithTimeout creates a new CanDoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCanDoParamsWithTimeout(timeout time.Duration) *CanDoParams {
	var ()
	return &CanDoParams{

		timeout: timeout,
	}
}

// NewCanDoParamsWithContext creates a new CanDoParams object
// with the default values initialized, and the ability to set a context for a request
func NewCanDoParamsWithContext(ctx context.Context) *CanDoParams {
	var ()
	return &CanDoParams{

		Context: ctx,
	}
}

// NewCanDoParamsWithHTTPClient creates a new CanDoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCanDoParamsWithHTTPClient(client *http.Client) *CanDoParams {
	var ()
	return &CanDoParams{
		HTTPClient: client,
	}
}

/*CanDoParams contains all the parameters to send to the API endpoint
for the can do operation typically these are written to a http.Request
*/
type CanDoParams struct {

	/*Body*/
	Body *models.OpenpitrixCanDoRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the can do params
func (o *CanDoParams) WithTimeout(timeout time.Duration) *CanDoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the can do params
func (o *CanDoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the can do params
func (o *CanDoParams) WithContext(ctx context.Context) *CanDoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the can do params
func (o *CanDoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the can do params
func (o *CanDoParams) WithHTTPClient(client *http.Client) *CanDoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the can do params
func (o *CanDoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the can do params
func (o *CanDoParams) WithBody(body *models.OpenpitrixCanDoRequest) *CanDoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the can do params
func (o *CanDoParams) SetBody(body *models.OpenpitrixCanDoRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CanDoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}