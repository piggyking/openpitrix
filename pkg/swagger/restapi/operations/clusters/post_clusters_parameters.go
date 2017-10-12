// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"openpitrix.io/openpitrix/pkg/swagger/models"
)

// NewPostClustersParams creates a new PostClustersParams object
// with the default values initialized.
func NewPostClustersParams() PostClustersParams {
	var ()
	return PostClustersParams{}
}

// PostClustersParams contains all the bound params for the post clusters operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostClusters
type PostClustersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The cluster to create.
	  In: body
	*/
	Cluster *models.Cluster
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostClustersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Cluster
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("cluster", "body", "", err))
		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Cluster = &body
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}