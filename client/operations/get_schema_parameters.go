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

// NewGetSchemaParams creates a new GetSchemaParams object
// with the default values initialized.
func NewGetSchemaParams() *GetSchemaParams {
	var (
		fetchMaxIDDefault = bool(false)
	)
	return &GetSchemaParams{
		FetchMaxID: &fetchMaxIDDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchemaParamsWithTimeout creates a new GetSchemaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSchemaParamsWithTimeout(timeout time.Duration) *GetSchemaParams {
	var (
		fetchMaxIDDefault = bool(false)
	)
	return &GetSchemaParams{
		FetchMaxID: &fetchMaxIDDefault,

		timeout: timeout,
	}
}

// NewGetSchemaParamsWithContext creates a new GetSchemaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSchemaParamsWithContext(ctx context.Context) *GetSchemaParams {
	var (
		fetchMaxIdDefault = bool(false)
	)
	return &GetSchemaParams{
		FetchMaxID: &fetchMaxIdDefault,

		Context: ctx,
	}
}

// NewGetSchemaParamsWithHTTPClient creates a new GetSchemaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSchemaParamsWithHTTPClient(client *http.Client) *GetSchemaParams {
	var (
		fetchMaxIdDefault = bool(false)
	)
	return &GetSchemaParams{
		FetchMaxID: &fetchMaxIdDefault,
		HTTPClient: client,
	}
}

/*GetSchemaParams contains all the parameters to send to the API endpoint
for the get schema operation typically these are written to a http.Request
*/
type GetSchemaParams struct {

	/*FetchMaxID*/
	FetchMaxID *bool
	/*Format*/
	Format *string
	/*ID
	  Globally unique identifier of the schema

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get schema params
func (o *GetSchemaParams) WithTimeout(timeout time.Duration) *GetSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schema params
func (o *GetSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schema params
func (o *GetSchemaParams) WithContext(ctx context.Context) *GetSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schema params
func (o *GetSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schema params
func (o *GetSchemaParams) WithHTTPClient(client *http.Client) *GetSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schema params
func (o *GetSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFetchMaxID adds the fetchMaxID to the get schema params
func (o *GetSchemaParams) WithFetchMaxID(fetchMaxID *bool) *GetSchemaParams {
	o.SetFetchMaxID(fetchMaxID)
	return o
}

// SetFetchMaxID adds the fetchMaxId to the get schema params
func (o *GetSchemaParams) SetFetchMaxID(fetchMaxID *bool) {
	o.FetchMaxID = fetchMaxID
}

// WithFormat adds the format to the get schema params
func (o *GetSchemaParams) WithFormat(format *string) *GetSchemaParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get schema params
func (o *GetSchemaParams) SetFormat(format *string) {
	o.Format = format
}

// WithID adds the id to the get schema params
func (o *GetSchemaParams) WithID(id int32) *GetSchemaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get schema params
func (o *GetSchemaParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FetchMaxID != nil {

		// query param fetchMaxId
		var qrFetchMaxID bool
		if o.FetchMaxID != nil {
			qrFetchMaxID = *o.FetchMaxID
		}
		qFetchMaxID := swag.FormatBool(qrFetchMaxID)
		if qFetchMaxID != "" {
			if err := r.SetQueryParam("fetchMaxId", qFetchMaxID); err != nil {
				return err
			}
		}

	}

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
