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

// NewGetKubernetesClustersKubernetesClusterNameParams creates a new GetKubernetesClustersKubernetesClusterNameParams object
// with the default values initialized.
func NewGetKubernetesClustersKubernetesClusterNameParams() *GetKubernetesClustersKubernetesClusterNameParams {
	var ()
	return &GetKubernetesClustersKubernetesClusterNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubernetesClustersKubernetesClusterNameParamsWithTimeout creates a new GetKubernetesClustersKubernetesClusterNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKubernetesClustersKubernetesClusterNameParamsWithTimeout(timeout time.Duration) *GetKubernetesClustersKubernetesClusterNameParams {
	var ()
	return &GetKubernetesClustersKubernetesClusterNameParams{

		timeout: timeout,
	}
}

// NewGetKubernetesClustersKubernetesClusterNameParamsWithContext creates a new GetKubernetesClustersKubernetesClusterNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKubernetesClustersKubernetesClusterNameParamsWithContext(ctx context.Context) *GetKubernetesClustersKubernetesClusterNameParams {
	var ()
	return &GetKubernetesClustersKubernetesClusterNameParams{

		Context: ctx,
	}
}

// NewGetKubernetesClustersKubernetesClusterNameParamsWithHTTPClient creates a new GetKubernetesClustersKubernetesClusterNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKubernetesClustersKubernetesClusterNameParamsWithHTTPClient(client *http.Client) *GetKubernetesClustersKubernetesClusterNameParams {
	var ()
	return &GetKubernetesClustersKubernetesClusterNameParams{
		HTTPClient: client,
	}
}

/*GetKubernetesClustersKubernetesClusterNameParams contains all the parameters to send to the API endpoint
for the get kubernetes clusters kubernetes cluster name operation typically these are written to a http.Request
*/
type GetKubernetesClustersKubernetesClusterNameParams struct {

	/*KubernetesClusterName
	  kubernetesClusterName

	*/
	KubernetesClusterName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kubernetes clusters kubernetes cluster name params
func (o *GetKubernetesClustersKubernetesClusterNameParams) WithTimeout(timeout time.Duration) *GetKubernetesClustersKubernetesClusterNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubernetes clusters kubernetes cluster name params
func (o *GetKubernetesClustersKubernetesClusterNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubernetes clusters kubernetes cluster name params
func (o *GetKubernetesClustersKubernetesClusterNameParams) WithContext(ctx context.Context) *GetKubernetesClustersKubernetesClusterNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubernetes clusters kubernetes cluster name params
func (o *GetKubernetesClustersKubernetesClusterNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubernetes clusters kubernetes cluster name params
func (o *GetKubernetesClustersKubernetesClusterNameParams) WithHTTPClient(client *http.Client) *GetKubernetesClustersKubernetesClusterNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubernetes clusters kubernetes cluster name params
func (o *GetKubernetesClustersKubernetesClusterNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKubernetesClusterName adds the kubernetesClusterName to the get kubernetes clusters kubernetes cluster name params
func (o *GetKubernetesClustersKubernetesClusterNameParams) WithKubernetesClusterName(kubernetesClusterName string) *GetKubernetesClustersKubernetesClusterNameParams {
	o.SetKubernetesClusterName(kubernetesClusterName)
	return o
}

// SetKubernetesClusterName adds the kubernetesClusterName to the get kubernetes clusters kubernetes cluster name params
func (o *GetKubernetesClustersKubernetesClusterNameParams) SetKubernetesClusterName(kubernetesClusterName string) {
	o.KubernetesClusterName = kubernetesClusterName
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubernetesClustersKubernetesClusterNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kubernetesClusterName
	if err := r.SetPathParam("kubernetesClusterName", o.KubernetesClusterName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
