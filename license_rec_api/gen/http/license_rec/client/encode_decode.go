// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// licenseRec HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/flexera/turtles-licenserec/license_rec_api/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	licenserec "github.com/flexera/turtles-licenserec/license_rec_api/gen/license_rec"
	goahttp "goa.design/goa/http"
)

// BuildGetLicenseRequest instantiates a HTTP request object with method and
// path set to call the "licenseRec" service "GetLicense" endpoint
func (c *Client) BuildGetLicenseRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		licenseID int
	)
	{
		p, ok := v.(*licenserec.LicenseObject)
		if !ok {
			return nil, goahttp.ErrInvalidType("licenseRec", "GetLicense", "*licenserec.LicenseObject", v)
		}
		licenseID = p.LicenseID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetLicenseLicenseRecPath(licenseID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("licenseRec", "GetLicense", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetLicenseResponse returns a decoder for responses returned by the
// licenseRec GetLicense endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeGetLicenseResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("licenseRec", "GetLicense", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("licenseRec", "GetLicense", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateLicenseRequest instantiates a HTTP request object with method and
// path set to call the "licenseRec" service "UpdateLicense" endpoint
func (c *Client) BuildUpdateLicenseRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		licenseID int
	)
	{
		p, ok := v.(*licenserec.LicenseObject)
		if !ok {
			return nil, goahttp.ErrInvalidType("licenseRec", "UpdateLicense", "*licenserec.LicenseObject", v)
		}
		licenseID = p.LicenseID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateLicenseLicenseRecPath(licenseID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("licenseRec", "UpdateLicense", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeUpdateLicenseResponse returns a decoder for responses returned by the
// licenseRec UpdateLicense endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeUpdateLicenseResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("licenseRec", "UpdateLicense", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("licenseRec", "UpdateLicense", resp.StatusCode, string(body))
		}
	}
}
