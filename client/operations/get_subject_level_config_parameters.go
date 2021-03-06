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

// NewGetSubjectLevelConfigParams creates a new GetSubjectLevelConfigParams object
// with the default values initialized.
func NewGetSubjectLevelConfigParams() *GetSubjectLevelConfigParams {
	var ()
	return &GetSubjectLevelConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubjectLevelConfigParamsWithTimeout creates a new GetSubjectLevelConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubjectLevelConfigParamsWithTimeout(timeout time.Duration) *GetSubjectLevelConfigParams {
	var ()
	return &GetSubjectLevelConfigParams{

		timeout: timeout,
	}
}

// NewGetSubjectLevelConfigParamsWithContext creates a new GetSubjectLevelConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubjectLevelConfigParamsWithContext(ctx context.Context) *GetSubjectLevelConfigParams {
	var ()
	return &GetSubjectLevelConfigParams{

		Context: ctx,
	}
}

// NewGetSubjectLevelConfigParamsWithHTTPClient creates a new GetSubjectLevelConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubjectLevelConfigParamsWithHTTPClient(client *http.Client) *GetSubjectLevelConfigParams {
	var ()
	return &GetSubjectLevelConfigParams{
		HTTPClient: client,
	}
}

/*GetSubjectLevelConfigParams contains all the parameters to send to the API endpoint
for the get subject level config operation typically these are written to a http.Request
*/
type GetSubjectLevelConfigParams struct {

	/*DefaultToGlobal*/
	DefaultToGlobal *bool
	/*Subject*/
	Subject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subject level config params
func (o *GetSubjectLevelConfigParams) WithTimeout(timeout time.Duration) *GetSubjectLevelConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subject level config params
func (o *GetSubjectLevelConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subject level config params
func (o *GetSubjectLevelConfigParams) WithContext(ctx context.Context) *GetSubjectLevelConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subject level config params
func (o *GetSubjectLevelConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subject level config params
func (o *GetSubjectLevelConfigParams) WithHTTPClient(client *http.Client) *GetSubjectLevelConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subject level config params
func (o *GetSubjectLevelConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDefaultToGlobal adds the defaultToGlobal to the get subject level config params
func (o *GetSubjectLevelConfigParams) WithDefaultToGlobal(defaultToGlobal *bool) *GetSubjectLevelConfigParams {
	o.SetDefaultToGlobal(defaultToGlobal)
	return o
}

// SetDefaultToGlobal adds the defaultToGlobal to the get subject level config params
func (o *GetSubjectLevelConfigParams) SetDefaultToGlobal(defaultToGlobal *bool) {
	o.DefaultToGlobal = defaultToGlobal
}

// WithSubject adds the subject to the get subject level config params
func (o *GetSubjectLevelConfigParams) WithSubject(subject string) *GetSubjectLevelConfigParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the get subject level config params
func (o *GetSubjectLevelConfigParams) SetSubject(subject string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubjectLevelConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DefaultToGlobal != nil {

		// query param defaultToGlobal
		var qrDefaultToGlobal bool
		if o.DefaultToGlobal != nil {
			qrDefaultToGlobal = *o.DefaultToGlobal
		}
		qDefaultToGlobal := swag.FormatBool(qrDefaultToGlobal)
		if qDefaultToGlobal != "" {
			if err := r.SetQueryParam("defaultToGlobal", qDefaultToGlobal); err != nil {
				return err
			}
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
