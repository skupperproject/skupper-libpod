// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/skupperproject/skupper-libpod/v4/models"
)

// PodStopLibpodReader is a Reader for the PodStopLibpod structure.
type PodStopLibpodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodStopLibpodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodStopLibpodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewPodStopLibpodNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewPodStopLibpodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPodStopLibpodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPodStopLibpodConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPodStopLibpodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPodStopLibpodOK creates a PodStopLibpodOK with default headers values
func NewPodStopLibpodOK() *PodStopLibpodOK {
	return &PodStopLibpodOK{}
}

/*
PodStopLibpodOK describes a response with status code 200, with default header values.

Stop pod
*/
type PodStopLibpodOK struct {
	Payload *models.PodStopReport
}

// IsSuccess returns true when this pod stop libpod o k response has a 2xx status code
func (o *PodStopLibpodOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pod stop libpod o k response has a 3xx status code
func (o *PodStopLibpodOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod stop libpod o k response has a 4xx status code
func (o *PodStopLibpodOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pod stop libpod o k response has a 5xx status code
func (o *PodStopLibpodOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pod stop libpod o k response a status code equal to that given
func (o *PodStopLibpodOK) IsCode(code int) bool {
	return code == 200
}

func (o *PodStopLibpodOK) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodOK  %+v", 200, o.Payload)
}

func (o *PodStopLibpodOK) String() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodOK  %+v", 200, o.Payload)
}

func (o *PodStopLibpodOK) GetPayload() *models.PodStopReport {
	return o.Payload
}

func (o *PodStopLibpodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodStopReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodStopLibpodNotModified creates a PodStopLibpodNotModified with default headers values
func NewPodStopLibpodNotModified() *PodStopLibpodNotModified {
	return &PodStopLibpodNotModified{}
}

/*
PodStopLibpodNotModified describes a response with status code 304, with default header values.

Pod already stopped
*/
type PodStopLibpodNotModified struct {
	Payload *PodStopLibpodNotModifiedBody
}

// IsSuccess returns true when this pod stop libpod not modified response has a 2xx status code
func (o *PodStopLibpodNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod stop libpod not modified response has a 3xx status code
func (o *PodStopLibpodNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this pod stop libpod not modified response has a 4xx status code
func (o *PodStopLibpodNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this pod stop libpod not modified response has a 5xx status code
func (o *PodStopLibpodNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this pod stop libpod not modified response a status code equal to that given
func (o *PodStopLibpodNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *PodStopLibpodNotModified) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodNotModified  %+v", 304, o.Payload)
}

func (o *PodStopLibpodNotModified) String() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodNotModified  %+v", 304, o.Payload)
}

func (o *PodStopLibpodNotModified) GetPayload() *PodStopLibpodNotModifiedBody {
	return o.Payload
}

func (o *PodStopLibpodNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodStopLibpodNotModifiedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodStopLibpodBadRequest creates a PodStopLibpodBadRequest with default headers values
func NewPodStopLibpodBadRequest() *PodStopLibpodBadRequest {
	return &PodStopLibpodBadRequest{}
}

/*
PodStopLibpodBadRequest describes a response with status code 400, with default header values.

Bad parameter in request
*/
type PodStopLibpodBadRequest struct {
	Payload *PodStopLibpodBadRequestBody
}

// IsSuccess returns true when this pod stop libpod bad request response has a 2xx status code
func (o *PodStopLibpodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod stop libpod bad request response has a 3xx status code
func (o *PodStopLibpodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod stop libpod bad request response has a 4xx status code
func (o *PodStopLibpodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pod stop libpod bad request response has a 5xx status code
func (o *PodStopLibpodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pod stop libpod bad request response a status code equal to that given
func (o *PodStopLibpodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PodStopLibpodBadRequest) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *PodStopLibpodBadRequest) String() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodBadRequest  %+v", 400, o.Payload)
}

func (o *PodStopLibpodBadRequest) GetPayload() *PodStopLibpodBadRequestBody {
	return o.Payload
}

func (o *PodStopLibpodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodStopLibpodBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodStopLibpodNotFound creates a PodStopLibpodNotFound with default headers values
func NewPodStopLibpodNotFound() *PodStopLibpodNotFound {
	return &PodStopLibpodNotFound{}
}

/*
PodStopLibpodNotFound describes a response with status code 404, with default header values.

No such pod
*/
type PodStopLibpodNotFound struct {
	Payload *PodStopLibpodNotFoundBody
}

// IsSuccess returns true when this pod stop libpod not found response has a 2xx status code
func (o *PodStopLibpodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod stop libpod not found response has a 3xx status code
func (o *PodStopLibpodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod stop libpod not found response has a 4xx status code
func (o *PodStopLibpodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pod stop libpod not found response has a 5xx status code
func (o *PodStopLibpodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pod stop libpod not found response a status code equal to that given
func (o *PodStopLibpodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PodStopLibpodNotFound) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodNotFound  %+v", 404, o.Payload)
}

func (o *PodStopLibpodNotFound) String() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodNotFound  %+v", 404, o.Payload)
}

func (o *PodStopLibpodNotFound) GetPayload() *PodStopLibpodNotFoundBody {
	return o.Payload
}

func (o *PodStopLibpodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodStopLibpodNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodStopLibpodConflict creates a PodStopLibpodConflict with default headers values
func NewPodStopLibpodConflict() *PodStopLibpodConflict {
	return &PodStopLibpodConflict{}
}

/*
PodStopLibpodConflict describes a response with status code 409, with default header values.

Stop pod
*/
type PodStopLibpodConflict struct {
	Payload *models.PodStopReport
}

// IsSuccess returns true when this pod stop libpod conflict response has a 2xx status code
func (o *PodStopLibpodConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod stop libpod conflict response has a 3xx status code
func (o *PodStopLibpodConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod stop libpod conflict response has a 4xx status code
func (o *PodStopLibpodConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pod stop libpod conflict response has a 5xx status code
func (o *PodStopLibpodConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pod stop libpod conflict response a status code equal to that given
func (o *PodStopLibpodConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PodStopLibpodConflict) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodConflict  %+v", 409, o.Payload)
}

func (o *PodStopLibpodConflict) String() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodConflict  %+v", 409, o.Payload)
}

func (o *PodStopLibpodConflict) GetPayload() *models.PodStopReport {
	return o.Payload
}

func (o *PodStopLibpodConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodStopReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodStopLibpodInternalServerError creates a PodStopLibpodInternalServerError with default headers values
func NewPodStopLibpodInternalServerError() *PodStopLibpodInternalServerError {
	return &PodStopLibpodInternalServerError{}
}

/*
PodStopLibpodInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PodStopLibpodInternalServerError struct {
	Payload *PodStopLibpodInternalServerErrorBody
}

// IsSuccess returns true when this pod stop libpod internal server error response has a 2xx status code
func (o *PodStopLibpodInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pod stop libpod internal server error response has a 3xx status code
func (o *PodStopLibpodInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pod stop libpod internal server error response has a 4xx status code
func (o *PodStopLibpodInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pod stop libpod internal server error response has a 5xx status code
func (o *PodStopLibpodInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pod stop libpod internal server error response a status code equal to that given
func (o *PodStopLibpodInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PodStopLibpodInternalServerError) Error() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *PodStopLibpodInternalServerError) String() string {
	return fmt.Sprintf("[POST /libpod/pods/{name}/stop][%d] podStopLibpodInternalServerError  %+v", 500, o.Payload)
}

func (o *PodStopLibpodInternalServerError) GetPayload() *PodStopLibpodInternalServerErrorBody {
	return o.Payload
}

func (o *PodStopLibpodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PodStopLibpodInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PodStopLibpodBadRequestBody pod stop libpod bad request body
swagger:model PodStopLibpodBadRequestBody
*/
type PodStopLibpodBadRequestBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod stop libpod bad request body
func (o *PodStopLibpodBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod stop libpod bad request body based on context it is used
func (o *PodStopLibpodBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodStopLibpodBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodStopLibpodBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PodStopLibpodBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PodStopLibpodInternalServerErrorBody pod stop libpod internal server error body
swagger:model PodStopLibpodInternalServerErrorBody
*/
type PodStopLibpodInternalServerErrorBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod stop libpod internal server error body
func (o *PodStopLibpodInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod stop libpod internal server error body based on context it is used
func (o *PodStopLibpodInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodStopLibpodInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodStopLibpodInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PodStopLibpodInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PodStopLibpodNotFoundBody pod stop libpod not found body
swagger:model PodStopLibpodNotFoundBody
*/
type PodStopLibpodNotFoundBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod stop libpod not found body
func (o *PodStopLibpodNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod stop libpod not found body based on context it is used
func (o *PodStopLibpodNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodStopLibpodNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodStopLibpodNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PodStopLibpodNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PodStopLibpodNotModifiedBody pod stop libpod not modified body
swagger:model PodStopLibpodNotModifiedBody
*/
type PodStopLibpodNotModifiedBody struct {

	// API root cause formatted for automated parsing
	// Example: API root cause
	Because string `json:"cause,omitempty"`

	// human error message, formatted for a human to read
	// Example: human error message
	Message string `json:"message,omitempty"`

	// http response code
	ResponseCode int64 `json:"response,omitempty"`
}

// Validate validates this pod stop libpod not modified body
func (o *PodStopLibpodNotModifiedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod stop libpod not modified body based on context it is used
func (o *PodStopLibpodNotModifiedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PodStopLibpodNotModifiedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PodStopLibpodNotModifiedBody) UnmarshalBinary(b []byte) error {
	var res PodStopLibpodNotModifiedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
