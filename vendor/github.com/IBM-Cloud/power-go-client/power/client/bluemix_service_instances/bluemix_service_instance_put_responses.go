// Code generated by go-swagger; DO NOT EDIT.

package bluemix_service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// BluemixServiceInstancePutReader is a Reader for the BluemixServiceInstancePut structure.
type BluemixServiceInstancePutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BluemixServiceInstancePutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBluemixServiceInstancePutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBluemixServiceInstancePutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBluemixServiceInstancePutOK creates a BluemixServiceInstancePutOK with default headers values
func NewBluemixServiceInstancePutOK() *BluemixServiceInstancePutOK {
	return &BluemixServiceInstancePutOK{}
}

/*BluemixServiceInstancePutOK handles this case with default header values.

OK
*/
type BluemixServiceInstancePutOK struct {
	Payload *models.ServiceInstance
}

func (o *BluemixServiceInstancePutOK) Error() string {
	return fmt.Sprintf("[PUT /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstancePutOK  %+v", 200, o.Payload)
}

func (o *BluemixServiceInstancePutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBluemixServiceInstancePutBadRequest creates a BluemixServiceInstancePutBadRequest with default headers values
func NewBluemixServiceInstancePutBadRequest() *BluemixServiceInstancePutBadRequest {
	return &BluemixServiceInstancePutBadRequest{}
}

/*BluemixServiceInstancePutBadRequest handles this case with default header values.

Bad Request
*/
type BluemixServiceInstancePutBadRequest struct {
	Payload *models.Error
}

func (o *BluemixServiceInstancePutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /bluemix_v1/service_instances/{instance_id}][%d] bluemixServiceInstancePutBadRequest  %+v", 400, o.Payload)
}

func (o *BluemixServiceInstancePutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
