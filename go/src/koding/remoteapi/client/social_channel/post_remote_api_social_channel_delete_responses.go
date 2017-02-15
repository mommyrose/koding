package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPISocialChannelDeleteReader is a Reader for the PostRemoteAPISocialChannelDelete structure.
type PostRemoteAPISocialChannelDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPISocialChannelDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPISocialChannelDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPISocialChannelDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPISocialChannelDeleteOK creates a PostRemoteAPISocialChannelDeleteOK with default headers values
func NewPostRemoteAPISocialChannelDeleteOK() *PostRemoteAPISocialChannelDeleteOK {
	return &PostRemoteAPISocialChannelDeleteOK{}
}

/*PostRemoteAPISocialChannelDeleteOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPISocialChannelDeleteOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPISocialChannelDeleteOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.delete][%d] postRemoteApiSocialChannelDeleteOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPISocialChannelDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPISocialChannelDeleteUnauthorized creates a PostRemoteAPISocialChannelDeleteUnauthorized with default headers values
func NewPostRemoteAPISocialChannelDeleteUnauthorized() *PostRemoteAPISocialChannelDeleteUnauthorized {
	return &PostRemoteAPISocialChannelDeleteUnauthorized{}
}

/*PostRemoteAPISocialChannelDeleteUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPISocialChannelDeleteUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPISocialChannelDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/SocialChannel.delete][%d] postRemoteApiSocialChannelDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPISocialChannelDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
