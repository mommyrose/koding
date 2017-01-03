package j_stack_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJStackTemplateGenerateStackIDParams creates a new PostRemoteAPIJStackTemplateGenerateStackIDParams object
// with the default values initialized.
func NewPostRemoteAPIJStackTemplateGenerateStackIDParams() *PostRemoteAPIJStackTemplateGenerateStackIDParams {
	var ()
	return &PostRemoteAPIJStackTemplateGenerateStackIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJStackTemplateGenerateStackIDParamsWithTimeout creates a new PostRemoteAPIJStackTemplateGenerateStackIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJStackTemplateGenerateStackIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJStackTemplateGenerateStackIDParams {
	var ()
	return &PostRemoteAPIJStackTemplateGenerateStackIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJStackTemplateGenerateStackIDParamsWithContext creates a new PostRemoteAPIJStackTemplateGenerateStackIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJStackTemplateGenerateStackIDParamsWithContext(ctx context.Context) *PostRemoteAPIJStackTemplateGenerateStackIDParams {
	var ()
	return &PostRemoteAPIJStackTemplateGenerateStackIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJStackTemplateGenerateStackIDParams contains all the parameters to send to the API endpoint
for the post remote API j stack template generate stack ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJStackTemplateGenerateStackIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j stack template generate stack ID params
func (o *PostRemoteAPIJStackTemplateGenerateStackIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJStackTemplateGenerateStackIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j stack template generate stack ID params
func (o *PostRemoteAPIJStackTemplateGenerateStackIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j stack template generate stack ID params
func (o *PostRemoteAPIJStackTemplateGenerateStackIDParams) WithContext(ctx context.Context) *PostRemoteAPIJStackTemplateGenerateStackIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j stack template generate stack ID params
func (o *PostRemoteAPIJStackTemplateGenerateStackIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j stack template generate stack ID params
func (o *PostRemoteAPIJStackTemplateGenerateStackIDParams) WithID(id string) *PostRemoteAPIJStackTemplateGenerateStackIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j stack template generate stack ID params
func (o *PostRemoteAPIJStackTemplateGenerateStackIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJStackTemplateGenerateStackIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
