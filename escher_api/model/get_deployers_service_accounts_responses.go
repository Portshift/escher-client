// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetDeployersServiceAccountsReader is a Reader for the GetDeployersServiceAccounts structure.
type GetDeployersServiceAccountsReader struct {
	Formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeployersServiceAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeployersServiceAccountsOK()
		if err := result.readResponse(response, consumer, o.Formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeployersServiceAccountsOK creates a GetDeployersServiceAccountsOK with default headers values
func NewGetDeployersServiceAccountsOK() *GetDeployersServiceAccountsOK {
	return &GetDeployersServiceAccountsOK{}
}

/*GetDeployersServiceAccountsOK handles this case with default header values.

Success
*/
type GetDeployersServiceAccountsOK struct {
	Payload []*ServiceAccountInfo
}

func (o *GetDeployersServiceAccountsOK) Error() string {
	return fmt.Sprintf("[GET /deployers/serviceAccounts][%d] getDeployersServiceAccountsOK  %+v", 200, o.Payload)
}

func (o *GetDeployersServiceAccountsOK) GetPayload() []*ServiceAccountInfo {
	return o.Payload
}

func (o *GetDeployersServiceAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
