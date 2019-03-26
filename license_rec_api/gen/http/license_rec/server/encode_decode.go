// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// licenseRec HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/flexera/turtles-licenserec/license_rec_api/design

package server

import (
	"context"
	"net/http"
	"strconv"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// EncodeGetLicenseResponse returns an encoder for responses returned by the
// licenseRec GetLicense endpoint.
func EncodeGetLicenseResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetLicenseRequest returns a decoder for requests sent to the
// licenseRec GetLicense endpoint.
func DecodeGetLicenseRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			licenseID int
			err       error

			params = mux.Vars(r)
		)
		{
			licenseIDRaw := params["LicenseID"]
			v, err2 := strconv.ParseInt(licenseIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("licenseID", licenseIDRaw, "integer"))
			}
			licenseID = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetLicenseLicenseObject(licenseID)

		return payload, nil
	}
}

// EncodeUpdateLicenseResponse returns an encoder for responses returned by the
// licenseRec UpdateLicense endpoint.
func EncodeUpdateLicenseResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateLicenseRequest returns a decoder for requests sent to the
// licenseRec UpdateLicense endpoint.
func DecodeUpdateLicenseRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			licenseID int
			err       error

			params = mux.Vars(r)
		)
		{
			licenseIDRaw := params["LicenseID"]
			v, err2 := strconv.ParseInt(licenseIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("licenseID", licenseIDRaw, "integer"))
			}
			licenseID = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateLicenseLicenseObject(licenseID)

		return payload, nil
	}
}