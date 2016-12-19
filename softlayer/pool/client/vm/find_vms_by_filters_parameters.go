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

	"github.com/cloudfoundry/bosh-softlayer-cpi/softlayer/pool/models"
)

// NewFindVmsByFiltersParams creates a new FindVmsByFiltersParams object
// with the default values initialized.
func NewFindVmsByFiltersParams() *FindVmsByFiltersParams {
	var ()
	return &FindVmsByFiltersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindVmsByFiltersParamsWithTimeout creates a new FindVmsByFiltersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindVmsByFiltersParamsWithTimeout(timeout time.Duration) *FindVmsByFiltersParams {
	var ()
	return &FindVmsByFiltersParams{

		timeout: timeout,
	}
}

// NewFindVmsByFiltersParamsWithContext creates a new FindVmsByFiltersParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindVmsByFiltersParamsWithContext(ctx context.Context) *FindVmsByFiltersParams {
	var ()
	return &FindVmsByFiltersParams{

		Context: ctx,
	}
}

/*FindVmsByFiltersParams contains all the parameters to send to the API endpoint
for the find vms by filters operation typically these are written to a http.Request
*/
type FindVmsByFiltersParams struct {

	/*Body
	  Vm object that needs to be added to the pool

	*/
	Body *models.VMFilter

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the find vms by filters params
func (o *FindVmsByFiltersParams) WithTimeout(timeout time.Duration) *FindVmsByFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find vms by filters params
func (o *FindVmsByFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find vms by filters params
func (o *FindVmsByFiltersParams) WithContext(ctx context.Context) *FindVmsByFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find vms by filters params
func (o *FindVmsByFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the find vms by filters params
func (o *FindVmsByFiltersParams) WithBody(body *models.VMFilter) *FindVmsByFiltersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the find vms by filters params
func (o *FindVmsByFiltersParams) SetBody(body *models.VMFilter) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FindVmsByFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
