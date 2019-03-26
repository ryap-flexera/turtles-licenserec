// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// licenseRec client HTTP transport
//
// Command:
// $ goa gen github.com/flexera/turtles-licenserec/license_rec_api/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the licenseRec service endpoint HTTP clients.
type Client struct {
	// GetLicense Doer is the HTTP client used to make requests to the GetLicense
	// endpoint.
	GetLicenseDoer goahttp.Doer

	// UpdateLicense Doer is the HTTP client used to make requests to the
	// UpdateLicense endpoint.
	UpdateLicenseDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the licenseRec service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetLicenseDoer:      doer,
		UpdateLicenseDoer:   doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GetLicense returns an endpoint that makes HTTP requests to the licenseRec
// service GetLicense server.
func (c *Client) GetLicense() goa.Endpoint {
	var (
		decodeResponse = DecodeGetLicenseResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetLicenseRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetLicenseDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("licenseRec", "GetLicense", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateLicense returns an endpoint that makes HTTP requests to the licenseRec
// service UpdateLicense server.
func (c *Client) UpdateLicense() goa.Endpoint {
	var (
		decodeResponse = DecodeUpdateLicenseResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateLicenseRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateLicenseDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("licenseRec", "UpdateLicense", err)
		}
		return decodeResponse(resp)
	}
}
