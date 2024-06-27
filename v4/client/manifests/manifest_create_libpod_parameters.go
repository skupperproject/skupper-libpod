// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/skupperproject/skupper-libpod/v4/models"
)

// NewManifestCreateLibpodParams creates a new ManifestCreateLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManifestCreateLibpodParams() *ManifestCreateLibpodParams {
	return &ManifestCreateLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManifestCreateLibpodParamsWithTimeout creates a new ManifestCreateLibpodParams object
// with the ability to set a timeout on a request.
func NewManifestCreateLibpodParamsWithTimeout(timeout time.Duration) *ManifestCreateLibpodParams {
	return &ManifestCreateLibpodParams{
		timeout: timeout,
	}
}

// NewManifestCreateLibpodParamsWithContext creates a new ManifestCreateLibpodParams object
// with the ability to set a context for a request.
func NewManifestCreateLibpodParamsWithContext(ctx context.Context) *ManifestCreateLibpodParams {
	return &ManifestCreateLibpodParams{
		Context: ctx,
	}
}

// NewManifestCreateLibpodParamsWithHTTPClient creates a new ManifestCreateLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewManifestCreateLibpodParamsWithHTTPClient(client *http.Client) *ManifestCreateLibpodParams {
	return &ManifestCreateLibpodParams{
		HTTPClient: client,
	}
}

/*
ManifestCreateLibpodParams contains all the parameters to send to the API endpoint

	for the manifest create libpod operation.

	Typically these are written to a http.Request.
*/
type ManifestCreateLibpodParams struct {

	/* All.

	   add all contents if given list
	*/
	All *bool

	/* Images.

	     One or more names of an image or a manifest list. Repeat parameter as needed.

	Support for multiple images, as of version 4.0.0
	Alias of `image` is support for compatibility with < 4.0.0
	Response status code is 200 with < 4.0.0 for compatibility

	*/
	Images string

	/* Name.

	   manifest list or index name to create
	*/
	Name string

	/* Options.

	   options for new manifest
	*/
	Options *models.ManifestModifyOptions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the manifest create libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManifestCreateLibpodParams) WithDefaults() *ManifestCreateLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the manifest create libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManifestCreateLibpodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the manifest create libpod params
func (o *ManifestCreateLibpodParams) WithTimeout(timeout time.Duration) *ManifestCreateLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the manifest create libpod params
func (o *ManifestCreateLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the manifest create libpod params
func (o *ManifestCreateLibpodParams) WithContext(ctx context.Context) *ManifestCreateLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the manifest create libpod params
func (o *ManifestCreateLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the manifest create libpod params
func (o *ManifestCreateLibpodParams) WithHTTPClient(client *http.Client) *ManifestCreateLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the manifest create libpod params
func (o *ManifestCreateLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the manifest create libpod params
func (o *ManifestCreateLibpodParams) WithAll(all *bool) *ManifestCreateLibpodParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the manifest create libpod params
func (o *ManifestCreateLibpodParams) SetAll(all *bool) {
	o.All = all
}

// WithImages adds the images to the manifest create libpod params
func (o *ManifestCreateLibpodParams) WithImages(images string) *ManifestCreateLibpodParams {
	o.SetImages(images)
	return o
}

// SetImages adds the images to the manifest create libpod params
func (o *ManifestCreateLibpodParams) SetImages(images string) {
	o.Images = images
}

// WithName adds the name to the manifest create libpod params
func (o *ManifestCreateLibpodParams) WithName(name string) *ManifestCreateLibpodParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the manifest create libpod params
func (o *ManifestCreateLibpodParams) SetName(name string) {
	o.Name = name
}

// WithOptions adds the options to the manifest create libpod params
func (o *ManifestCreateLibpodParams) WithOptions(options *models.ManifestModifyOptions) *ManifestCreateLibpodParams {
	o.SetOptions(options)
	return o
}

// SetOptions adds the options to the manifest create libpod params
func (o *ManifestCreateLibpodParams) SetOptions(options *models.ManifestModifyOptions) {
	o.Options = options
}

// WriteToRequest writes these params to a swagger request
func (o *ManifestCreateLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool

		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {

			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}
	}

	// query param images
	qrImages := o.Images
	qImages := qrImages
	if qImages != "" {

		if err := r.SetQueryParam("images", qImages); err != nil {
			return err
		}
	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {

		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}
	if o.Options != nil {
		if err := r.SetBodyParam(o.Options); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
