// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	models "github.com/bjornm82/go-swagger-schema-registry/models"
)

// NewUpdateSubjectLevelConfigParams creates a new UpdateSubjectLevelConfigParams object
// with the default values initialized.
func NewUpdateSubjectLevelConfigParams() *UpdateSubjectLevelConfigParams {
	var ()
	return &UpdateSubjectLevelConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSubjectLevelConfigParamsWithTimeout creates a new UpdateSubjectLevelConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSubjectLevelConfigParamsWithTimeout(timeout time.Duration) *UpdateSubjectLevelConfigParams {
	var ()
	return &UpdateSubjectLevelConfigParams{

		timeout: timeout,
	}
}

// NewUpdateSubjectLevelConfigParamsWithContext creates a new UpdateSubjectLevelConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSubjectLevelConfigParamsWithContext(ctx context.Context) *UpdateSubjectLevelConfigParams {
	var ()
	return &UpdateSubjectLevelConfigParams{

		Context: ctx,
	}
}

// NewUpdateSubjectLevelConfigParamsWithHTTPClient creates a new UpdateSubjectLevelConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSubjectLevelConfigParamsWithHTTPClient(client *http.Client) *UpdateSubjectLevelConfigParams {
	var ()
	return &UpdateSubjectLevelConfigParams{
		HTTPClient: client,
	}
}

/*UpdateSubjectLevelConfigParams contains all the parameters to send to the API endpoint
for the update subject level config operation typically these are written to a http.Request
*/
type UpdateSubjectLevelConfigParams struct {

	/*Body
	  Config Update Request

	*/
	Body *models.ConfigUpdateRequest
	/*Subject
	  Name of the Subject

	*/
	Subject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) WithTimeout(timeout time.Duration) *UpdateSubjectLevelConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) WithContext(ctx context.Context) *UpdateSubjectLevelConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) WithHTTPClient(client *http.Client) *UpdateSubjectLevelConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) WithBody(body *models.ConfigUpdateRequest) *UpdateSubjectLevelConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) SetBody(body *models.ConfigUpdateRequest) {
	o.Body = body
}

// WithSubject adds the subject to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) WithSubject(subject string) *UpdateSubjectLevelConfigParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the update subject level config params
func (o *UpdateSubjectLevelConfigParams) SetSubject(subject string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSubjectLevelConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param subject
	if err := r.SetPathParam("subject", o.Subject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
