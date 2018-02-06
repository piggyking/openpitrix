// Code generated by go-swagger; DO NOT EDIT.

package app_runtime_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DeleteAppRuntimeReader is a Reader for the DeleteAppRuntime structure.
type DeleteAppRuntimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAppRuntimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteAppRuntimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAppRuntimeOK creates a DeleteAppRuntimeOK with default headers values
func NewDeleteAppRuntimeOK() *DeleteAppRuntimeOK {
	return &DeleteAppRuntimeOK{}
}

/*DeleteAppRuntimeOK handles this case with default header values.

DeleteAppRuntimeOK delete app runtime o k
*/
type DeleteAppRuntimeOK struct {
	Payload models.ProtobufEmpty
}

func (o *DeleteAppRuntimeOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/appruntimes/{id}][%d] deleteAppRuntimeOK  %+v", 200, o.Payload)
}

func (o *DeleteAppRuntimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}