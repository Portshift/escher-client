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

// NewPutDeployersDeployerIDParams creates a new PutDeployersDeployerIDParams object
// with the default values initialized.
func NewPutDeployersDeployerIDParams() *PutDeployersDeployerIDParams {
	var ()
	return &PutDeployersDeployerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDeployersDeployerIDParamsWithTimeout creates a new PutDeployersDeployerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDeployersDeployerIDParamsWithTimeout(timeout time.Duration) *PutDeployersDeployerIDParams {
	var ()
	return &PutDeployersDeployerIDParams{

		timeout: timeout,
	}
}

// NewPutDeployersDeployerIDParamsWithContext creates a new PutDeployersDeployerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDeployersDeployerIDParamsWithContext(ctx context.Context) *PutDeployersDeployerIDParams {
	var ()
	return &PutDeployersDeployerIDParams{

		Context: ctx,
	}
}

// NewPutDeployersDeployerIDParamsWithHTTPClient creates a new PutDeployersDeployerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDeployersDeployerIDParamsWithHTTPClient(client *http.Client) *PutDeployersDeployerIDParams {
	var ()
	return &PutDeployersDeployerIDParams{
		HTTPClient: client,
	}
}

/*PutDeployersDeployerIDParams contains all the parameters to send to the API endpoint
for the put deployers deployer ID operation typically these are written to a http.Request
*/
type PutDeployersDeployerIDParams struct {

	/*Body*/
	Body Deployer
	/*DeployerID*/
	DeployerID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) WithTimeout(timeout time.Duration) *PutDeployersDeployerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) WithContext(ctx context.Context) *PutDeployersDeployerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) WithHTTPClient(client *http.Client) *PutDeployersDeployerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) WithBody(body Deployer) *PutDeployersDeployerIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) SetBody(body Deployer) {
	o.Body = body
}

// WithDeployerID adds the deployerID to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) WithDeployerID(deployerID strfmt.UUID) *PutDeployersDeployerIDParams {
	o.SetDeployerID(deployerID)
	return o
}

// SetDeployerID adds the deployerId to the put deployers deployer ID params
func (o *PutDeployersDeployerIDParams) SetDeployerID(deployerID strfmt.UUID) {
	o.DeployerID = deployerID
}

// WriteToRequest writes these params to a swagger request
func (o *PutDeployersDeployerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param deployerId
	if err := r.SetPathParam("deployerId", o.DeployerID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
