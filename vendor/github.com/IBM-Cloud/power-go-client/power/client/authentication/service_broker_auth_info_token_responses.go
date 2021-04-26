// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerAuthInfoTokenReader is a Reader for the ServiceBrokerAuthInfoToken structure.
type ServiceBrokerAuthInfoTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerAuthInfoTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceBrokerAuthInfoTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewServiceBrokerAuthInfoTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceBrokerAuthInfoTokenOK creates a ServiceBrokerAuthInfoTokenOK with default headers values
func NewServiceBrokerAuthInfoTokenOK() *ServiceBrokerAuthInfoTokenOK {
	return &ServiceBrokerAuthInfoTokenOK{}
}

/*ServiceBrokerAuthInfoTokenOK handles this case with default header values.

OK
*/
type ServiceBrokerAuthInfoTokenOK struct {
	Payload *models.TokenExtra
}

func (o *ServiceBrokerAuthInfoTokenOK) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenExtra)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerAuthInfoTokenInternalServerError creates a ServiceBrokerAuthInfoTokenInternalServerError with default headers values
func NewServiceBrokerAuthInfoTokenInternalServerError() *ServiceBrokerAuthInfoTokenInternalServerError {
	return &ServiceBrokerAuthInfoTokenInternalServerError{}
}

/*ServiceBrokerAuthInfoTokenInternalServerError handles this case with default header values.

Internal Server Error
*/
type ServiceBrokerAuthInfoTokenInternalServerError struct {
	Payload *models.Error
}

func (o *ServiceBrokerAuthInfoTokenInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/v1/info/token][%d] serviceBrokerAuthInfoTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceBrokerAuthInfoTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
