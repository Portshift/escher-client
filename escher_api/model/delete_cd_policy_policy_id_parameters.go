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

// NewDeleteCdPolicyPolicyIDParams creates a new DeleteCdPolicyPolicyIDParams object
// with the default values initialized.
func NewDeleteCdPolicyPolicyIDParams() *DeleteCdPolicyPolicyIDParams {
	var ()
	return &DeleteCdPolicyPolicyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCdPolicyPolicyIDParamsWithTimeout creates a new DeleteCdPolicyPolicyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCdPolicyPolicyIDParamsWithTimeout(timeout time.Duration) *DeleteCdPolicyPolicyIDParams {
	var ()
	return &DeleteCdPolicyPolicyIDParams{

		timeout: timeout,
	}
}

// NewDeleteCdPolicyPolicyIDParamsWithContext creates a new DeleteCdPolicyPolicyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCdPolicyPolicyIDParamsWithContext(ctx context.Context) *DeleteCdPolicyPolicyIDParams {
	var ()
	return &DeleteCdPolicyPolicyIDParams{

		Context: ctx,
	}
}

// NewDeleteCdPolicyPolicyIDParamsWithHTTPClient creates a new DeleteCdPolicyPolicyIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCdPolicyPolicyIDParamsWithHTTPClient(client *http.Client) *DeleteCdPolicyPolicyIDParams {
	var ()
	return &DeleteCdPolicyPolicyIDParams{
		HTTPClient: client,
	}
}

/*DeleteCdPolicyPolicyIDParams contains all the parameters to send to the API endpoint
for the delete cd policy policy ID operation typically these are written to a http.Request
*/
type DeleteCdPolicyPolicyIDParams struct {

	/*PolicyID*/
	PolicyID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cd policy policy ID params
func (o *DeleteCdPolicyPolicyIDParams) WithTimeout(timeout time.Duration) *DeleteCdPolicyPolicyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cd policy policy ID params
func (o *DeleteCdPolicyPolicyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cd policy policy ID params
func (o *DeleteCdPolicyPolicyIDParams) WithContext(ctx context.Context) *DeleteCdPolicyPolicyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cd policy policy ID params
func (o *DeleteCdPolicyPolicyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cd policy policy ID params
func (o *DeleteCdPolicyPolicyIDParams) WithHTTPClient(client *http.Client) *DeleteCdPolicyPolicyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cd policy policy ID params
func (o *DeleteCdPolicyPolicyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyID adds the policyID to the delete cd policy policy ID params
func (o *DeleteCdPolicyPolicyIDParams) WithPolicyID(policyID strfmt.UUID) *DeleteCdPolicyPolicyIDParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the delete cd policy policy ID params
func (o *DeleteCdPolicyPolicyIDParams) SetPolicyID(policyID strfmt.UUID) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCdPolicyPolicyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param policyId
	if err := r.SetPathParam("policyId", o.PolicyID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
