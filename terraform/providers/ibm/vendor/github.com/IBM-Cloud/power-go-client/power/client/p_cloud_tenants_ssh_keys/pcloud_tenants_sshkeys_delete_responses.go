// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tenants_ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudTenantsSshkeysDeleteReader is a Reader for the PcloudTenantsSshkeysDelete structure.
type PcloudTenantsSshkeysDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTenantsSshkeysDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTenantsSshkeysDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTenantsSshkeysDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTenantsSshkeysDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudTenantsSshkeysDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTenantsSshkeysDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudTenantsSshkeysDeleteOK creates a PcloudTenantsSshkeysDeleteOK with default headers values
func NewPcloudTenantsSshkeysDeleteOK() *PcloudTenantsSshkeysDeleteOK {
	return &PcloudTenantsSshkeysDeleteOK{}
}

/* PcloudTenantsSshkeysDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTenantsSshkeysDeleteOK struct {
	Payload models.Object
}

func (o *PcloudTenantsSshkeysDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteOK  %+v", 200, o.Payload)
}
func (o *PcloudTenantsSshkeysDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysDeleteBadRequest creates a PcloudTenantsSshkeysDeleteBadRequest with default headers values
func NewPcloudTenantsSshkeysDeleteBadRequest() *PcloudTenantsSshkeysDeleteBadRequest {
	return &PcloudTenantsSshkeysDeleteBadRequest{}
}

/* PcloudTenantsSshkeysDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTenantsSshkeysDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudTenantsSshkeysDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysDeleteUnauthorized creates a PcloudTenantsSshkeysDeleteUnauthorized with default headers values
func NewPcloudTenantsSshkeysDeleteUnauthorized() *PcloudTenantsSshkeysDeleteUnauthorized {
	return &PcloudTenantsSshkeysDeleteUnauthorized{}
}

/* PcloudTenantsSshkeysDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTenantsSshkeysDeleteUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudTenantsSshkeysDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysDeleteGone creates a PcloudTenantsSshkeysDeleteGone with default headers values
func NewPcloudTenantsSshkeysDeleteGone() *PcloudTenantsSshkeysDeleteGone {
	return &PcloudTenantsSshkeysDeleteGone{}
}

/* PcloudTenantsSshkeysDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudTenantsSshkeysDeleteGone struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteGone  %+v", 410, o.Payload)
}
func (o *PcloudTenantsSshkeysDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTenantsSshkeysDeleteInternalServerError creates a PcloudTenantsSshkeysDeleteInternalServerError with default headers values
func NewPcloudTenantsSshkeysDeleteInternalServerError() *PcloudTenantsSshkeysDeleteInternalServerError {
	return &PcloudTenantsSshkeysDeleteInternalServerError{}
}

/* PcloudTenantsSshkeysDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTenantsSshkeysDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudTenantsSshkeysDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/tenants/{tenant_id}/sshkeys/{sshkey_name}][%d] pcloudTenantsSshkeysDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudTenantsSshkeysDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTenantsSshkeysDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
