// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCdPolicyPolicyIDReader is a Reader for the DeleteCdPolicyPolicyID structure.
type DeleteCdPolicyPolicyIDReader struct {
	Formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCdPolicyPolicyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCdPolicyPolicyIDNoContent()
		if err := result.readResponse(response, consumer, o.Formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCdPolicyPolicyIDNoContent creates a DeleteCdPolicyPolicyIDNoContent with default headers values
func NewDeleteCdPolicyPolicyIDNoContent() *DeleteCdPolicyPolicyIDNoContent {
	return &DeleteCdPolicyPolicyIDNoContent{}
}

/*DeleteCdPolicyPolicyIDNoContent handles this case with default header values.

Success
*/
type DeleteCdPolicyPolicyIDNoContent struct {
}

func (o *DeleteCdPolicyPolicyIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /cdPolicy/{policyId}][%d] deleteCdPolicyPolicyIdNoContent ", 204)
}

func (o *DeleteCdPolicyPolicyIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
