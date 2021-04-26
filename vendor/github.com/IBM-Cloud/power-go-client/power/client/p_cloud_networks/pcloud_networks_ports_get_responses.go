// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudNetworksPortsGetReader is a Reader for the PcloudNetworksPortsGet structure.
type PcloudNetworksPortsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudNetworksPortsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudNetworksPortsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPcloudNetworksPortsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudNetworksPortsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudNetworksPortsGetOK creates a PcloudNetworksPortsGetOK with default headers values
func NewPcloudNetworksPortsGetOK() *PcloudNetworksPortsGetOK {
	return &PcloudNetworksPortsGetOK{}
}

/*PcloudNetworksPortsGetOK handles this case with default header values.

OK
*/
type PcloudNetworksPortsGetOK struct {
	Payload *models.NetworkPort
}

func (o *PcloudNetworksPortsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudNetworksPortsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetNotFound creates a PcloudNetworksPortsGetNotFound with default headers values
func NewPcloudNetworksPortsGetNotFound() *PcloudNetworksPortsGetNotFound {
	return &PcloudNetworksPortsGetNotFound{}
}

/*PcloudNetworksPortsGetNotFound handles this case with default header values.

Not Found
*/
type PcloudNetworksPortsGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudNetworksPortsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudNetworksPortsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudNetworksPortsGetInternalServerError creates a PcloudNetworksPortsGetInternalServerError with default headers values
func NewPcloudNetworksPortsGetInternalServerError() *PcloudNetworksPortsGetInternalServerError {
	return &PcloudNetworksPortsGetInternalServerError{}
}

/*PcloudNetworksPortsGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudNetworksPortsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudNetworksPortsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/networks/{network_id}/ports/{port_id}][%d] pcloudNetworksPortsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudNetworksPortsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
