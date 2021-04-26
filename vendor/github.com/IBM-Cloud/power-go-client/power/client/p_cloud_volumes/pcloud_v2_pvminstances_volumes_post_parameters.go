// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

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

// NewPcloudV2PvminstancesVolumesPostParams creates a new PcloudV2PvminstancesVolumesPostParams object
// with the default values initialized.
func NewPcloudV2PvminstancesVolumesPostParams() *PcloudV2PvminstancesVolumesPostParams {
	var ()
	return &PcloudV2PvminstancesVolumesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV2PvminstancesVolumesPostParamsWithTimeout creates a new PcloudV2PvminstancesVolumesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudV2PvminstancesVolumesPostParamsWithTimeout(timeout time.Duration) *PcloudV2PvminstancesVolumesPostParams {
	var ()
	return &PcloudV2PvminstancesVolumesPostParams{

		timeout: timeout,
	}
}

// NewPcloudV2PvminstancesVolumesPostParamsWithContext creates a new PcloudV2PvminstancesVolumesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudV2PvminstancesVolumesPostParamsWithContext(ctx context.Context) *PcloudV2PvminstancesVolumesPostParams {
	var ()
	return &PcloudV2PvminstancesVolumesPostParams{

		Context: ctx,
	}
}

// NewPcloudV2PvminstancesVolumesPostParamsWithHTTPClient creates a new PcloudV2PvminstancesVolumesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudV2PvminstancesVolumesPostParamsWithHTTPClient(client *http.Client) *PcloudV2PvminstancesVolumesPostParams {
	var ()
	return &PcloudV2PvminstancesVolumesPostParams{
		HTTPClient: client,
	}
}

/*PcloudV2PvminstancesVolumesPostParams contains all the parameters to send to the API endpoint
for the pcloud v2 pvminstances volumes post operation typically these are written to a http.Request
*/
type PcloudV2PvminstancesVolumesPostParams struct {

	/*Body
	  Parameter to attach volumes to a PVMInstance

	*/
	Body *models.VolumesAttach
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

// WithTimeout adds the timeout to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) WithTimeout(timeout time.Duration) *PcloudV2PvminstancesVolumesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) WithContext(ctx context.Context) *PcloudV2PvminstancesVolumesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) WithHTTPClient(client *http.Client) *PcloudV2PvminstancesVolumesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) WithBody(body *models.VolumesAttach) *PcloudV2PvminstancesVolumesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) SetBody(body *models.VolumesAttach) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV2PvminstancesVolumesPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudV2PvminstancesVolumesPostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud v2 pvminstances volumes post params
func (o *PcloudV2PvminstancesVolumesPostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV2PvminstancesVolumesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
