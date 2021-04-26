// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesVolumesSetbootPutReader is a Reader for the PcloudPvminstancesVolumesSetbootPut structure.
type PcloudPvminstancesVolumesSetbootPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVolumesSetbootPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudPvminstancesVolumesSetbootPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudPvminstancesVolumesSetbootPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPcloudPvminstancesVolumesSetbootPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesVolumesSetbootPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesVolumesSetbootPutOK creates a PcloudPvminstancesVolumesSetbootPutOK with default headers values
func NewPcloudPvminstancesVolumesSetbootPutOK() *PcloudPvminstancesVolumesSetbootPutOK {
	return &PcloudPvminstancesVolumesSetbootPutOK{}
}

/*PcloudPvminstancesVolumesSetbootPutOK handles this case with default header values.

OK
*/
type PcloudPvminstancesVolumesSetbootPutOK struct {
	Payload models.Object
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutOK  %+v", 200, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutBadRequest creates a PcloudPvminstancesVolumesSetbootPutBadRequest with default headers values
func NewPcloudPvminstancesVolumesSetbootPutBadRequest() *PcloudPvminstancesVolumesSetbootPutBadRequest {
	return &PcloudPvminstancesVolumesSetbootPutBadRequest{}
}

/*PcloudPvminstancesVolumesSetbootPutBadRequest handles this case with default header values.

Bad Request
*/
type PcloudPvminstancesVolumesSetbootPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutNotFound creates a PcloudPvminstancesVolumesSetbootPutNotFound with default headers values
func NewPcloudPvminstancesVolumesSetbootPutNotFound() *PcloudPvminstancesVolumesSetbootPutNotFound {
	return &PcloudPvminstancesVolumesSetbootPutNotFound{}
}

/*PcloudPvminstancesVolumesSetbootPutNotFound handles this case with default header values.

Not Found
*/
type PcloudPvminstancesVolumesSetbootPutNotFound struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutNotFound  %+v", 404, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesSetbootPutInternalServerError creates a PcloudPvminstancesVolumesSetbootPutInternalServerError with default headers values
func NewPcloudPvminstancesVolumesSetbootPutInternalServerError() *PcloudPvminstancesVolumesSetbootPutInternalServerError {
	return &PcloudPvminstancesVolumesSetbootPutInternalServerError{}
}

/*PcloudPvminstancesVolumesSetbootPutInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesVolumesSetbootPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}/setboot][%d] pcloudPvminstancesVolumesSetbootPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesVolumesSetbootPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
