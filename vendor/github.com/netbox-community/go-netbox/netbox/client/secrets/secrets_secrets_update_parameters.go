// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netbox-community/go-netbox/netbox/models"
)

// NewSecretsSecretsUpdateParams creates a new SecretsSecretsUpdateParams object
// with the default values initialized.
func NewSecretsSecretsUpdateParams() *SecretsSecretsUpdateParams {
	var ()
	return &SecretsSecretsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretsUpdateParamsWithTimeout creates a new SecretsSecretsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsSecretsUpdateParamsWithTimeout(timeout time.Duration) *SecretsSecretsUpdateParams {
	var ()
	return &SecretsSecretsUpdateParams{

		timeout: timeout,
	}
}

// NewSecretsSecretsUpdateParamsWithContext creates a new SecretsSecretsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsSecretsUpdateParamsWithContext(ctx context.Context) *SecretsSecretsUpdateParams {
	var ()
	return &SecretsSecretsUpdateParams{

		Context: ctx,
	}
}

// NewSecretsSecretsUpdateParamsWithHTTPClient creates a new SecretsSecretsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsSecretsUpdateParamsWithHTTPClient(client *http.Client) *SecretsSecretsUpdateParams {
	var ()
	return &SecretsSecretsUpdateParams{
		HTTPClient: client,
	}
}

/*SecretsSecretsUpdateParams contains all the parameters to send to the API endpoint
for the secrets secrets update operation typically these are written to a http.Request
*/
type SecretsSecretsUpdateParams struct {

	/*Data*/
	Data *models.WritableSecret
	/*ID
	  A unique integer value identifying this secret.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithTimeout(timeout time.Duration) *SecretsSecretsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithContext(ctx context.Context) *SecretsSecretsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithHTTPClient(client *http.Client) *SecretsSecretsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithData(data *models.WritableSecret) *SecretsSecretsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetData(data *models.WritableSecret) {
	o.Data = data
}

// WithID adds the id to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) WithID(id int64) *SecretsSecretsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secrets update params
func (o *SecretsSecretsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}