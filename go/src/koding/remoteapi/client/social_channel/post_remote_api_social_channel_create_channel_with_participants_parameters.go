package social_channel

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

	"koding/remoteapi/models"
)

// NewPostRemoteAPISocialChannelCreateChannelWithParticipantsParams creates a new PostRemoteAPISocialChannelCreateChannelWithParticipantsParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelCreateChannelWithParticipantsParams() *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams {
	var ()
	return &PostRemoteAPISocialChannelCreateChannelWithParticipantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelCreateChannelWithParticipantsParamsWithTimeout creates a new PostRemoteAPISocialChannelCreateChannelWithParticipantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelCreateChannelWithParticipantsParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams {
	var ()
	return &PostRemoteAPISocialChannelCreateChannelWithParticipantsParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelCreateChannelWithParticipantsParamsWithContext creates a new PostRemoteAPISocialChannelCreateChannelWithParticipantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelCreateChannelWithParticipantsParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams {
	var ()
	return &PostRemoteAPISocialChannelCreateChannelWithParticipantsParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelCreateChannelWithParticipantsParams contains all the parameters to send to the API endpoint
for the post remote API social channel create channel with participants operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelCreateChannelWithParticipantsParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social channel create channel with participants params
func (o *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel create channel with participants params
func (o *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel create channel with participants params
func (o *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel create channel with participants params
func (o *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel create channel with participants params
func (o *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel create channel with participants params
func (o *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelCreateChannelWithParticipantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
