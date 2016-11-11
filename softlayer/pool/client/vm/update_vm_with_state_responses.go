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

// UpdateVMWithStateReader is a Reader for the UpdateVMWithState structure.
type UpdateVMWithStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVMWithStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateVMWithStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateVMWithStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewUpdateVMWithStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewUpdateVMWithStateOK creates a UpdateVMWithStateOK with default headers values
func NewUpdateVMWithStateOK() *UpdateVMWithStateOK {
	return &UpdateVMWithStateOK{}
}

/*UpdateVMWithStateOK handles this case with default header values.

state updated successfully
*/
type UpdateVMWithStateOK struct {
	Payload string
}

func (o *UpdateVMWithStateOK) Error() string {
	return fmt.Sprintf("[PUT /vms/{cid}][%d] updateVmWithStateOK  %+v", 200, o.Payload)
}

func (o *UpdateVMWithStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVMWithStateNotFound creates a UpdateVMWithStateNotFound with default headers values
func NewUpdateVMWithStateNotFound() *UpdateVMWithStateNotFound {
	return &UpdateVMWithStateNotFound{}
}

/*UpdateVMWithStateNotFound handles this case with default header values.

vm not found
*/
type UpdateVMWithStateNotFound struct {
}

func (o *UpdateVMWithStateNotFound) Error() string {
	return fmt.Sprintf("[PUT /vms/{cid}][%d] updateVmWithStateNotFound ", 404)
}

func (o *UpdateVMWithStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateVMWithStateDefault creates a UpdateVMWithStateDefault with default headers values
func NewUpdateVMWithStateDefault(code int) *UpdateVMWithStateDefault {
	return &UpdateVMWithStateDefault{
		_statusCode: code,
	}
}

/*UpdateVMWithStateDefault handles this case with default header values.

unexpected error
*/
type UpdateVMWithStateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update Vm with state default response
func (o *UpdateVMWithStateDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVMWithStateDefault) Error() string {
	return fmt.Sprintf("[PUT /vms/{cid}][%d] updateVmWithState default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateVMWithStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
