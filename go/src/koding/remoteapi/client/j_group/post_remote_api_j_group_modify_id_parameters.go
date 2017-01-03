package j_group

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

// NewPostRemoteAPIJGroupModifyIDParams creates a new PostRemoteAPIJGroupModifyIDParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupModifyIDParams() *PostRemoteAPIJGroupModifyIDParams {
	var ()
	return &PostRemoteAPIJGroupModifyIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupModifyIDParamsWithTimeout creates a new PostRemoteAPIJGroupModifyIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupModifyIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupModifyIDParams {
	var ()
	return &PostRemoteAPIJGroupModifyIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupModifyIDParamsWithContext creates a new PostRemoteAPIJGroupModifyIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupModifyIDParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupModifyIDParams {
	var ()
	return &PostRemoteAPIJGroupModifyIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupModifyIDParams contains all the parameters to send to the API endpoint
for the post remote API j group modify ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupModifyIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j group modify ID params
func (o *PostRemoteAPIJGroupModifyIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupModifyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group modify ID params
func (o *PostRemoteAPIJGroupModifyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group modify ID params
func (o *PostRemoteAPIJGroupModifyIDParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupModifyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group modify ID params
func (o *PostRemoteAPIJGroupModifyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j group modify ID params
func (o *PostRemoteAPIJGroupModifyIDParams) WithID(id string) *PostRemoteAPIJGroupModifyIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j group modify ID params
func (o *PostRemoteAPIJGroupModifyIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupModifyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
