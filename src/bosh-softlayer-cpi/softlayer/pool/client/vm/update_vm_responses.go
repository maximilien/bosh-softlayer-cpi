package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"bosh-softlayer-cpi/softlayer/pool/models"
)

// UpdateVMReader is a Reader for the UpdateVM structure.
type UpdateVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateVMNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateVMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewUpdateVMOK creates a UpdateVMOK with default headers values
func NewUpdateVMOK() *UpdateVMOK {
	return &UpdateVMOK{}
}

/*UpdateVMOK handles this case with default header values.

update vm in the pool
*/
type UpdateVMOK struct {
	Payload string
}

func (o *UpdateVMOK) Error() string {
	return fmt.Sprintf("[PUT /vms][%d] updateVmOK  %+v", 200, o.Payload)
}

func (o *UpdateVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMNotFound creates a UpdateVMNotFound with default headers values
func NewUpdateVMNotFound() *UpdateVMNotFound {
	return &UpdateVMNotFound{}
}

/*UpdateVMNotFound handles this case with default header values.

vm not found
*/
type UpdateVMNotFound struct {
}

func (o *UpdateVMNotFound) Error() string {
	return fmt.Sprintf("[PUT /vms][%d] updateVmNotFound ", 404)
}

func (o *UpdateVMNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVMDefault creates a UpdateVMDefault with default headers values
func NewUpdateVMDefault(code int) *UpdateVMDefault {
	return &UpdateVMDefault{
		_statusCode: code,
	}
}

/*UpdateVMDefault handles this case with default header values.

unexpected error
*/
type UpdateVMDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update Vm default response
func (o *UpdateVMDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVMDefault) Error() string {
	return fmt.Sprintf("[PUT /vms][%d] updateVm default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateVMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
