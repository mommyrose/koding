package j_api_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJAPITokenCreateReader is a Reader for the PostRemoteAPIJAPITokenCreate structure.
type PostRemoteAPIJAPITokenCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJAPITokenCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJAPITokenCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJAPITokenCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJAPITokenCreateOK creates a PostRemoteAPIJAPITokenCreateOK with default headers values
func NewPostRemoteAPIJAPITokenCreateOK() *PostRemoteAPIJAPITokenCreateOK {
	return &PostRemoteAPIJAPITokenCreateOK{}
}

/*PostRemoteAPIJAPITokenCreateOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJAPITokenCreateOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJAPITokenCreateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JApiToken.create][%d] postRemoteApiJApiTokenCreateOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJAPITokenCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJAPITokenCreateUnauthorized creates a PostRemoteAPIJAPITokenCreateUnauthorized with default headers values
func NewPostRemoteAPIJAPITokenCreateUnauthorized() *PostRemoteAPIJAPITokenCreateUnauthorized {
	return &PostRemoteAPIJAPITokenCreateUnauthorized{}
}

/*PostRemoteAPIJAPITokenCreateUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJAPITokenCreateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJAPITokenCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JApiToken.create][%d] postRemoteApiJApiTokenCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJAPITokenCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
