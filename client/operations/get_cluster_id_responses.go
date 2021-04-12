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

// GetClusterIDReader is a Reader for the GetClusterID structure.
type GetClusterIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetClusterIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClusterIDOK creates a GetClusterIDOK with default headers values
func NewGetClusterIDOK() *GetClusterIDOK {
	return &GetClusterIDOK{}
}

/*GetClusterIDOK handles this case with default header values.

successful operation
*/
type GetClusterIDOK struct {
	Payload *models.ServerClusterID
}

func (o *GetClusterIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/metadata/id][%d] getClusterIdOK  %+v", 200, o.Payload)
}

func (o *GetClusterIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServerClusterID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterIDInternalServerError creates a GetClusterIDInternalServerError with default headers values
func NewGetClusterIDInternalServerError() *GetClusterIDInternalServerError {
	return &GetClusterIDInternalServerError{}
}

/*GetClusterIDInternalServerError handles this case with default header values.

Error code 50001 -- Error in the backend data store

*/
type GetClusterIDInternalServerError struct {
}

func (o *GetClusterIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/metadata/id][%d] getClusterIdInternalServerError ", 500)
}

func (o *GetClusterIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}