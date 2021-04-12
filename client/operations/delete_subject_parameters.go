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

// NewDeleteSubjectParams creates a new DeleteSubjectParams object
// with the default values initialized.
func NewDeleteSubjectParams() *DeleteSubjectParams {
	var ()
	return &DeleteSubjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubjectParamsWithTimeout creates a new DeleteSubjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSubjectParamsWithTimeout(timeout time.Duration) *DeleteSubjectParams {
	var ()
	return &DeleteSubjectParams{

		timeout: timeout,
	}
}

// NewDeleteSubjectParamsWithContext creates a new DeleteSubjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSubjectParamsWithContext(ctx context.Context) *DeleteSubjectParams {
	var ()
	return &DeleteSubjectParams{

		Context: ctx,
	}
}

// NewDeleteSubjectParamsWithHTTPClient creates a new DeleteSubjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSubjectParamsWithHTTPClient(client *http.Client) *DeleteSubjectParams {
	var ()
	return &DeleteSubjectParams{
		HTTPClient: client,
	}
}

/*DeleteSubjectParams contains all the parameters to send to the API endpoint
for the delete subject operation typically these are written to a http.Request
*/
type DeleteSubjectParams struct {

	/*Permanent*/
	Permanent *bool
	/*Subject
	  the name of the subject

	*/
	Subject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete subject params
func (o *DeleteSubjectParams) WithTimeout(timeout time.Duration) *DeleteSubjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subject params
func (o *DeleteSubjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subject params
func (o *DeleteSubjectParams) WithContext(ctx context.Context) *DeleteSubjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subject params
func (o *DeleteSubjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subject params
func (o *DeleteSubjectParams) WithHTTPClient(client *http.Client) *DeleteSubjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subject params
func (o *DeleteSubjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPermanent adds the permanent to the delete subject params
func (o *DeleteSubjectParams) WithPermanent(permanent *bool) *DeleteSubjectParams {
	o.SetPermanent(permanent)
	return o
}

// SetPermanent adds the permanent to the delete subject params
func (o *DeleteSubjectParams) SetPermanent(permanent *bool) {
	o.Permanent = permanent
}

// WithSubject adds the subject to the delete subject params
func (o *DeleteSubjectParams) WithSubject(subject string) *DeleteSubjectParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the delete subject params
func (o *DeleteSubjectParams) SetSubject(subject string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Permanent != nil {

		// query param permanent
		var qrPermanent bool
		if o.Permanent != nil {
			qrPermanent = *o.Permanent
		}
		qPermanent := swag.FormatBool(qrPermanent)
		if qPermanent != "" {
			if err := r.SetQueryParam("permanent", qPermanent); err != nil {
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
