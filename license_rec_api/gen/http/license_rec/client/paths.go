// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the licenseRec service.
//
// Command:
// $ goa gen github.com/ryap-flexera/turtles-licenserec/license_rec_api/design

package client

import (
	"fmt"
)

// GetLicenseLicenseRecPath returns the URL path to the licenseRec service GetLicense HTTP endpoint.
func GetLicenseLicenseRecPath(licenseID int) string {
	return fmt.Sprintf("/api/licenses/get/%v", licenseID)
}

// UpdateDeviceLicenseLicenseRecPath returns the URL path to the licenseRec service UpdateDeviceLicense HTTP endpoint.
func UpdateDeviceLicenseLicenseRecPath(licenseID int) string {
	return fmt.Sprintf("/api/licenses/post/%v", licenseID)
}

// UpdateDeviceLicenseWithValueLicenseRecPath returns the URL path to the licenseRec service UpdateDeviceLicenseWithValue HTTP endpoint.
func UpdateDeviceLicenseWithValueLicenseRecPath(licenseID int, consumptionValue int) string {
	return fmt.Sprintf("/api/licenses/post/%v/%v", licenseID, consumptionValue)
}
