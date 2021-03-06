// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bjornm82/go-swagger-schema-registry/models"
)

// GetSchemaByVersionReader is a Reader for the GetSchemaByVersion structure.
type GetSchemaByVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemaByVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSchemaByVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSchemaByVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetSchemaByVersionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSchemaByVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSchemaByVersionOK creates a GetSchemaByVersionOK with default headers values
func NewGetSchemaByVersionOK() *GetSchemaByVersionOK {
	return &GetSchemaByVersionOK{}
}

/*GetSchemaByVersionOK handles this case with default header values.

successful operation
*/
type GetSchemaByVersionOK struct {
	Payload *models.Schema
}

func (o *GetSchemaByVersionOK) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions/{version}][%d] getSchemaByVersionOK  %+v", 200, o.Payload)
}

func (o *GetSchemaByVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Schema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemaByVersionNotFound creates a GetSchemaByVersionNotFound with default headers values
func NewGetSchemaByVersionNotFound() *GetSchemaByVersionNotFound {
	return &GetSchemaByVersionNotFound{}
}

/*GetSchemaByVersionNotFound handles this case with default header values.

Error code 40401 -- Subject not found
Error code 40402 -- Version not found
*/
type GetSchemaByVersionNotFound struct {
}

func (o *GetSchemaByVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions/{version}][%d] getSchemaByVersionNotFound ", 404)
}

func (o *GetSchemaByVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemaByVersionUnprocessableEntity creates a GetSchemaByVersionUnprocessableEntity with default headers values
func NewGetSchemaByVersionUnprocessableEntity() *GetSchemaByVersionUnprocessableEntity {
	return &GetSchemaByVersionUnprocessableEntity{}
}

/*GetSchemaByVersionUnprocessableEntity handles this case with default header values.

Error code 42202 -- Invalid version
*/
type GetSchemaByVersionUnprocessableEntity struct {
}

func (o *GetSchemaByVersionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions/{version}][%d] getSchemaByVersionUnprocessableEntity ", 422)
}

func (o *GetSchemaByVersionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSchemaByVersionInternalServerError creates a GetSchemaByVersionInternalServerError with default headers values
func NewGetSchemaByVersionInternalServerError() *GetSchemaByVersionInternalServerError {
	return &GetSchemaByVersionInternalServerError{}
}

/*GetSchemaByVersionInternalServerError handles this case with default header values.

Error code 50001 -- Error in the backend data store
*/
type GetSchemaByVersionInternalServerError struct {
}

func (o *GetSchemaByVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subjects/{subject}/versions/{version}][%d] getSchemaByVersionInternalServerError ", 500)
}

func (o *GetSchemaByVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
