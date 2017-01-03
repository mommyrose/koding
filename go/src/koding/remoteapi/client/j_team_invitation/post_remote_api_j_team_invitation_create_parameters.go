package j_team_invitation

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

// NewPostRemoteAPIJTeamInvitationCreateParams creates a new PostRemoteAPIJTeamInvitationCreateParams object
// with the default values initialized.
func NewPostRemoteAPIJTeamInvitationCreateParams() *PostRemoteAPIJTeamInvitationCreateParams {
	var ()
	return &PostRemoteAPIJTeamInvitationCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJTeamInvitationCreateParamsWithTimeout creates a new PostRemoteAPIJTeamInvitationCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJTeamInvitationCreateParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJTeamInvitationCreateParams {
	var ()
	return &PostRemoteAPIJTeamInvitationCreateParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJTeamInvitationCreateParamsWithContext creates a new PostRemoteAPIJTeamInvitationCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJTeamInvitationCreateParamsWithContext(ctx context.Context) *PostRemoteAPIJTeamInvitationCreateParams {
	var ()
	return &PostRemoteAPIJTeamInvitationCreateParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJTeamInvitationCreateParams contains all the parameters to send to the API endpoint
for the post remote API j team invitation create operation typically these are written to a http.Request
*/
type PostRemoteAPIJTeamInvitationCreateParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j team invitation create params
func (o *PostRemoteAPIJTeamInvitationCreateParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJTeamInvitationCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j team invitation create params
func (o *PostRemoteAPIJTeamInvitationCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j team invitation create params
func (o *PostRemoteAPIJTeamInvitationCreateParams) WithContext(ctx context.Context) *PostRemoteAPIJTeamInvitationCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j team invitation create params
func (o *PostRemoteAPIJTeamInvitationCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j team invitation create params
func (o *PostRemoteAPIJTeamInvitationCreateParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJTeamInvitationCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j team invitation create params
func (o *PostRemoteAPIJTeamInvitationCreateParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJTeamInvitationCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
