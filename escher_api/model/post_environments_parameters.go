// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostEnvironmentsParams creates a new PostEnvironmentsParams object
// with the default values initialized.
func NewPostEnvironmentsParams() *PostEnvironmentsParams {
	var ()
	return &PostEnvironmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEnvironmentsParamsWithTimeout creates a new PostEnvironmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEnvironmentsParamsWithTimeout(timeout time.Duration) *PostEnvironmentsParams {
	var ()
	return &PostEnvironmentsParams{

		timeout: timeout,
	}
}

// NewPostEnvironmentsParamsWithContext creates a new PostEnvironmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostEnvironmentsParamsWithContext(ctx context.Context) *PostEnvironmentsParams {
	var ()
	return &PostEnvironmentsParams{

		Context: ctx,
	}
}

// NewPostEnvironmentsParamsWithHTTPClient creates a new PostEnvironmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostEnvironmentsParamsWithHTTPClient(client *http.Client) *PostEnvironmentsParams {
	var ()
	return &PostEnvironmentsParams{
		HTTPClient: client,
	}
}

/*PostEnvironmentsParams contains all the parameters to send to the API endpoint
for the post environments operation typically these are written to a http.Request
*/
type PostEnvironmentsParams struct {

	/*Body
	  Environment definition

	*/
	Body *Environment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post environments params
func (o *PostEnvironmentsParams) WithTimeout(timeout time.Duration) *PostEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post environments params
func (o *PostEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post environments params
func (o *PostEnvironmentsParams) WithContext(ctx context.Context) *PostEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post environments params
func (o *PostEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post environments params
func (o *PostEnvironmentsParams) WithHTTPClient(client *http.Client) *PostEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post environments params
func (o *PostEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post environments params
func (o *PostEnvironmentsParams) WithBody(body *Environment) *PostEnvironmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post environments params
func (o *PostEnvironmentsParams) SetBody(body *Environment) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
