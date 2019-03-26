package licensereconcileapi

import (
	"context"
	"log"
	"strconv"

	licenserec "github.com/flexera/turtles-licenserec/license_rec_api/gen/license_rec"
)

var licenses = make(map[int]int)

// licenseRec service example implementation.
// The example methods log the requests and return zero values.
type licenseRecsrvc struct {
	logger *log.Logger
}

// NewLicenseRec returns the licenseRec service implementation.
func NewLicenseRec(logger *log.Logger) licenserec.Service {
	return &licenseRecsrvc{logger}
}

// Returns license count
func (s *licenseRecsrvc) GetLicense(ctx context.Context, p *licenserec.LicenseObject) (res int, err error) {
	s.logger.Print("Getting license with ID: " + strconv.Itoa(p.LicenseID))

	var temp = licenses[p.LicenseID]

	if temp == 0 {
		s.logger.Print("License doesn't exist. Adding to storage.")
		licenses[p.LicenseID] = 0
		s.logger.Print("Incrementing license.")
		licenses[p.LicenseID]++
	}

	return licenses[p.LicenseID], nil
}

// Increments license count given the LicenseID
func (s *licenseRecsrvc) UpdateLicense(ctx context.Context, p *licenserec.LicenseObject) (res int, err error) {
	s.logger.Print("Updating license")

	var temp = licenses[p.LicenseID]

	if temp == 0 {
		s.logger.Print("License doesn't exist. Adding to storage.")
		licenses[p.LicenseID] = 0
		s.logger.Print("Incrementing license.")
		licenses[p.LicenseID]++

		return
	}

	licenses[p.LicenseID]++

	return
}
