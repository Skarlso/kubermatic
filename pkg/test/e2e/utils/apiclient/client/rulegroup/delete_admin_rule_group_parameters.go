// Code generated by go-swagger; DO NOT EDIT.

package rulegroup

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
)

// NewDeleteAdminRuleGroupParams creates a new DeleteAdminRuleGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAdminRuleGroupParams() *DeleteAdminRuleGroupParams {
	return &DeleteAdminRuleGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAdminRuleGroupParamsWithTimeout creates a new DeleteAdminRuleGroupParams object
// with the ability to set a timeout on a request.
func NewDeleteAdminRuleGroupParamsWithTimeout(timeout time.Duration) *DeleteAdminRuleGroupParams {
	return &DeleteAdminRuleGroupParams{
		timeout: timeout,
	}
}

// NewDeleteAdminRuleGroupParamsWithContext creates a new DeleteAdminRuleGroupParams object
// with the ability to set a context for a request.
func NewDeleteAdminRuleGroupParamsWithContext(ctx context.Context) *DeleteAdminRuleGroupParams {
	return &DeleteAdminRuleGroupParams{
		Context: ctx,
	}
}

// NewDeleteAdminRuleGroupParamsWithHTTPClient creates a new DeleteAdminRuleGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAdminRuleGroupParamsWithHTTPClient(client *http.Client) *DeleteAdminRuleGroupParams {
	return &DeleteAdminRuleGroupParams{
		HTTPClient: client,
	}
}

/* DeleteAdminRuleGroupParams contains all the parameters to send to the API endpoint
   for the delete admin rule group operation.

   Typically these are written to a http.Request.
*/
type DeleteAdminRuleGroupParams struct {

	// RulegroupID.
	RuleGroupID string

	// SeedName.
	SeedName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete admin rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdminRuleGroupParams) WithDefaults() *DeleteAdminRuleGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete admin rule group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdminRuleGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) WithTimeout(timeout time.Duration) *DeleteAdminRuleGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) WithContext(ctx context.Context) *DeleteAdminRuleGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) WithHTTPClient(client *http.Client) *DeleteAdminRuleGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRuleGroupID adds the rulegroupID to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) WithRuleGroupID(rulegroupID string) *DeleteAdminRuleGroupParams {
	o.SetRuleGroupID(rulegroupID)
	return o
}

// SetRuleGroupID adds the rulegroupId to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) SetRuleGroupID(rulegroupID string) {
	o.RuleGroupID = rulegroupID
}

// WithSeedName adds the seedName to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) WithSeedName(seedName string) *DeleteAdminRuleGroupParams {
	o.SetSeedName(seedName)
	return o
}

// SetSeedName adds the seedName to the delete admin rule group params
func (o *DeleteAdminRuleGroupParams) SetSeedName(seedName string) {
	o.SeedName = seedName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAdminRuleGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rulegroup_id
	if err := r.SetPathParam("rulegroup_id", o.RuleGroupID); err != nil {
		return err
	}

	// path param seed_name
	if err := r.SetPathParam("seed_name", o.SeedName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
