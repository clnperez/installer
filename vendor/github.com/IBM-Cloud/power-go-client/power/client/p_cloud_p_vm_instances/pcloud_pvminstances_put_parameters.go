// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPvminstancesPutParams creates a new PcloudPvminstancesPutParams object
// with the default values initialized.
func NewPcloudPvminstancesPutParams() *PcloudPvminstancesPutParams {
	var ()
	return &PcloudPvminstancesPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesPutParamsWithTimeout creates a new PcloudPvminstancesPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPvminstancesPutParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesPutParams {
	var ()
	return &PcloudPvminstancesPutParams{

		timeout: timeout,
	}
}

// NewPcloudPvminstancesPutParamsWithContext creates a new PcloudPvminstancesPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPvminstancesPutParamsWithContext(ctx context.Context) *PcloudPvminstancesPutParams {
	var ()
	return &PcloudPvminstancesPutParams{

		Context: ctx,
	}
}

// NewPcloudPvminstancesPutParamsWithHTTPClient creates a new PcloudPvminstancesPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPvminstancesPutParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesPutParams {
	var ()
	return &PcloudPvminstancesPutParams{
		HTTPClient: client,
	}
}

/*PcloudPvminstancesPutParams contains all the parameters to send to the API endpoint
for the pcloud pvminstances put operation typically these are written to a http.Request
*/
type PcloudPvminstancesPutParams struct {

	/*Body
	  Parameters to update a PCloud PVM Instance

	*/
	Body *models.PVMInstanceUpdate
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PvmInstanceID
	  PCloud PVM Instance ID

	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) WithContext(ctx context.Context) *PcloudPvminstancesPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) WithBody(body *models.PVMInstanceUpdate) *PcloudPvminstancesPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) SetBody(body *models.PVMInstanceUpdate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesPutParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances put params
func (o *PcloudPvminstancesPutParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
