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

// NewPostCiPolicyParams creates a new PostCiPolicyParams object
// with the default values initialized.
func NewPostCiPolicyParams() *PostCiPolicyParams {
	var ()
	return &PostCiPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCiPolicyParamsWithTimeout creates a new PostCiPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCiPolicyParamsWithTimeout(timeout time.Duration) *PostCiPolicyParams {
	var ()
	return &PostCiPolicyParams{

		timeout: timeout,
	}
}

// NewPostCiPolicyParamsWithContext creates a new PostCiPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCiPolicyParamsWithContext(ctx context.Context) *PostCiPolicyParams {
	var ()
	return &PostCiPolicyParams{

		Context: ctx,
	}
}

// NewPostCiPolicyParamsWithHTTPClient creates a new PostCiPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCiPolicyParamsWithHTTPClient(client *http.Client) *PostCiPolicyParams {
	var ()
	return &PostCiPolicyParams{
		HTTPClient: client,
	}
}

/*PostCiPolicyParams contains all the parameters to send to the API endpoint
for the post ci policy operation typically these are written to a http.Request
*/
type PostCiPolicyParams struct {

	/*Body*/
	Body *CiPolicy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ci policy params
func (o *PostCiPolicyParams) WithTimeout(timeout time.Duration) *PostCiPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ci policy params
func (o *PostCiPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ci policy params
func (o *PostCiPolicyParams) WithContext(ctx context.Context) *PostCiPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ci policy params
func (o *PostCiPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ci policy params
func (o *PostCiPolicyParams) WithHTTPClient(client *http.Client) *PostCiPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ci policy params
func (o *PostCiPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post ci policy params
func (o *PostCiPolicyParams) WithBody(body *CiPolicy) *PostCiPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post ci policy params
func (o *PostCiPolicyParams) SetBody(body *CiPolicy) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostCiPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
