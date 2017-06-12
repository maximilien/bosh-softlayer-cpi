package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"bosh-softlayer-cpi/softlayer/pool/models"
)

// NewOrderVMByFilterParams creates a new OrderVMByFilterParams object
// with the default values initialized.
func NewOrderVMByFilterParams() *OrderVMByFilterParams {
	var ()
	return &OrderVMByFilterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOrderVMByFilterParamsWithTimeout creates a new OrderVMByFilterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOrderVMByFilterParamsWithTimeout(timeout time.Duration) *OrderVMByFilterParams {
	var ()
	return &OrderVMByFilterParams{

		timeout: timeout,
	}
}

// NewOrderVMByFilterParamsWithContext creates a new OrderVMByFilterParams object
// with the default values initialized, and the ability to set a context for a request
func NewOrderVMByFilterParamsWithContext(ctx context.Context) *OrderVMByFilterParams {
	var ()
	return &OrderVMByFilterParams{

		Context: ctx,
	}
}

/*OrderVMByFilterParams contains all the parameters to send to the API endpoint
for the order Vm by filter operation typically these are written to a http.Request
*/
type OrderVMByFilterParams struct {

	/*Body
	  vm filter to order

	*/
	Body *models.VMFilter

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the order Vm by filter params
func (o *OrderVMByFilterParams) WithTimeout(timeout time.Duration) *OrderVMByFilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the order Vm by filter params
func (o *OrderVMByFilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the order Vm by filter params
func (o *OrderVMByFilterParams) WithContext(ctx context.Context) *OrderVMByFilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the order Vm by filter params
func (o *OrderVMByFilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the order Vm by filter params
func (o *OrderVMByFilterParams) WithBody(body *models.VMFilter) *OrderVMByFilterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the order Vm by filter params
func (o *OrderVMByFilterParams) SetBody(body *models.VMFilter) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *OrderVMByFilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.VMFilter)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
