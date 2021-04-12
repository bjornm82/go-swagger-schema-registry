// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetSchemaTypesReader is a Reader for the GetSchemaTypes structure.
type GetSchemaTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemaTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSchemaTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetSchemaTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSchemaTypesOK creates a GetSchemaTypesOK with default headers values
func NewGetSchemaTypesOK() *GetSchemaTypesOK {
	return &GetSchemaTypesOK{}
}

/*GetSchemaTypesOK handles this case with default header values.

successful operation
*/
type GetSchemaTypesOK struct {
	Payload []string
}

func (o *GetSchemaTypesOK) Error() string {
	return fmt.Sprintf("[GET /schemas/types][%d] getSchemaTypesOK  %+v", 200, o.Payload)
}

func (o *GetSchemaTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemaTypesInternalServerError creates a GetSchemaTypesInternalServerError with default headers values
func NewGetSchemaTypesInternalServerError() *GetSchemaTypesInternalServerError {
	return &GetSchemaTypesInternalServerError{}
}

/*GetSchemaTypesInternalServerError handles this case with default header values.

Error code 50001 -- Error in the backend data store

*/
type GetSchemaTypesInternalServerError struct {
}

func (o *GetSchemaTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schemas/types][%d] getSchemaTypesInternalServerError ", 500)
}

func (o *GetSchemaTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
