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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSchemaByVersionParams creates a new GetSchemaByVersionParams object
// with the default values initialized.
func NewGetSchemaByVersionParams() *GetSchemaByVersionParams {
	var ()
	return &GetSchemaByVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchemaByVersionParamsWithTimeout creates a new GetSchemaByVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSchemaByVersionParamsWithTimeout(timeout time.Duration) *GetSchemaByVersionParams {
	var ()
	return &GetSchemaByVersionParams{

		timeout: timeout,
	}
}

// NewGetSchemaByVersionParamsWithContext creates a new GetSchemaByVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSchemaByVersionParamsWithContext(ctx context.Context) *GetSchemaByVersionParams {
	var ()
	return &GetSchemaByVersionParams{

		Context: ctx,
	}
}

// NewGetSchemaByVersionParamsWithHTTPClient creates a new GetSchemaByVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSchemaByVersionParamsWithHTTPClient(client *http.Client) *GetSchemaByVersionParams {
	var ()
	return &GetSchemaByVersionParams{
		HTTPClient: client,
	}
}

/*GetSchemaByVersionParams contains all the parameters to send to the API endpoint
for the get schema by version operation typically these are written to a http.Request
*/
type GetSchemaByVersionParams struct {

	/*Deleted*/
	Deleted *bool
	/*Subject
	  Name of the Subject

	*/
	Subject string
	/*Version
	  Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string "latest". "latest" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get schema by version params
func (o *GetSchemaByVersionParams) WithTimeout(timeout time.Duration) *GetSchemaByVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schema by version params
func (o *GetSchemaByVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schema by version params
func (o *GetSchemaByVersionParams) WithContext(ctx context.Context) *GetSchemaByVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schema by version params
func (o *GetSchemaByVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schema by version params
func (o *GetSchemaByVersionParams) WithHTTPClient(client *http.Client) *GetSchemaByVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schema by version params
func (o *GetSchemaByVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleted adds the deleted to the get schema by version params
func (o *GetSchemaByVersionParams) WithDeleted(deleted *bool) *GetSchemaByVersionParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the get schema by version params
func (o *GetSchemaByVersionParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithSubject adds the subject to the get schema by version params
func (o *GetSchemaByVersionParams) WithSubject(subject string) *GetSchemaByVersionParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the get schema by version params
func (o *GetSchemaByVersionParams) SetSubject(subject string) {
	o.Subject = subject
}

// WithVersion adds the version to the get schema by version params
func (o *GetSchemaByVersionParams) WithVersion(version string) *GetSchemaByVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get schema by version params
func (o *GetSchemaByVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemaByVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool
		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {
			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}

	}

	// path param subject
	if err := r.SetPathParam("subject", o.Subject); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
