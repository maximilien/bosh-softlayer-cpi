package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/pool/models"
)

// FindVmsByDeploymentReader is a Reader for the FindVmsByDeployment structure.
type FindVmsByDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindVmsByDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindVmsByDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewFindVmsByDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewFindVmsByDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewFindVmsByDeploymentOK creates a FindVmsByDeploymentOK with default headers values
func NewFindVmsByDeploymentOK() *FindVmsByDeploymentOK {
	return &FindVmsByDeploymentOK{}
}

/*FindVmsByDeploymentOK handles this case with default header values.

successful operation
*/
type FindVmsByDeploymentOK struct {
	Payload *models.VmsResponse
}

func (o *FindVmsByDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /vms/findByDeployment][%d] findVmsByDeploymentOK  %+v", 200, o.Payload)
}

func (o *FindVmsByDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VmsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindVmsByDeploymentNotFound creates a FindVmsByDeploymentNotFound with default headers values
func NewFindVmsByDeploymentNotFound() *FindVmsByDeploymentNotFound {
	return &FindVmsByDeploymentNotFound{}
}

/*FindVmsByDeploymentNotFound handles this case with default header values.

vm not found
*/
type FindVmsByDeploymentNotFound struct {
}

func (o *FindVmsByDeploymentNotFound) Error() string {
	return fmt.Sprintf("[GET /vms/findByDeployment][%d] findVmsByDeploymentNotFound ", 404)
}

func (o *FindVmsByDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindVmsByDeploymentDefault creates a FindVmsByDeploymentDefault with default headers values
func NewFindVmsByDeploymentDefault(code int) *FindVmsByDeploymentDefault {
	return &FindVmsByDeploymentDefault{
		_statusCode: code,
	}
}

/*FindVmsByDeploymentDefault handles this case with default header values.

unexpected error
*/
type FindVmsByDeploymentDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find vms by deployment default response
func (o *FindVmsByDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *FindVmsByDeploymentDefault) Error() string {
	return fmt.Sprintf("[GET /vms/findByDeployment][%d] findVmsByDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *FindVmsByDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
