// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// LicenseAPIServer HTTP client CLI support package
//
// Command:
// $ goa gen github.com/ryap-flexera/turtles-licenserec/license_rec_api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	licenserecc "github.com/ryap-flexera/turtles-licenserec/license_rec_api/gen/http/license_rec/client"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `license-rec (get-license|update-device-license|update-device-license-with-value)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` license-rec get-license --licenseid 4847839135972539348` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		licenseRecFlags = flag.NewFlagSet("license-rec", flag.ContinueOnError)

		licenseRecGetLicenseFlags         = flag.NewFlagSet("get-license", flag.ExitOnError)
		licenseRecGetLicenseLicenseIDFlag = licenseRecGetLicenseFlags.String("licenseid", "REQUIRED", "")

		licenseRecUpdateDeviceLicenseFlags         = flag.NewFlagSet("update-device-license", flag.ExitOnError)
		licenseRecUpdateDeviceLicenseLicenseIDFlag = licenseRecUpdateDeviceLicenseFlags.String("licenseid", "REQUIRED", "")

		licenseRecUpdateDeviceLicenseWithValueFlags                = flag.NewFlagSet("update-device-license-with-value", flag.ExitOnError)
		licenseRecUpdateDeviceLicenseWithValueLicenseIDFlag        = licenseRecUpdateDeviceLicenseWithValueFlags.String("licenseid", "REQUIRED", "")
		licenseRecUpdateDeviceLicenseWithValueConsumptionValueFlag = licenseRecUpdateDeviceLicenseWithValueFlags.String("consumption-value", "REQUIRED", "")
	)
	licenseRecFlags.Usage = licenseRecUsage
	licenseRecGetLicenseFlags.Usage = licenseRecGetLicenseUsage
	licenseRecUpdateDeviceLicenseFlags.Usage = licenseRecUpdateDeviceLicenseUsage
	licenseRecUpdateDeviceLicenseWithValueFlags.Usage = licenseRecUpdateDeviceLicenseWithValueUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "license-rec":
			svcf = licenseRecFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "license-rec":
			switch epn {
			case "get-license":
				epf = licenseRecGetLicenseFlags

			case "update-device-license":
				epf = licenseRecUpdateDeviceLicenseFlags

			case "update-device-license-with-value":
				epf = licenseRecUpdateDeviceLicenseWithValueFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "license-rec":
			c := licenserecc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-license":
				endpoint = c.GetLicense()
				data, err = licenserecc.BuildGetLicensePayload(*licenseRecGetLicenseLicenseIDFlag)
			case "update-device-license":
				endpoint = c.UpdateDeviceLicense()
				data, err = licenserecc.BuildUpdateDeviceLicensePayload(*licenseRecUpdateDeviceLicenseLicenseIDFlag)
			case "update-device-license-with-value":
				endpoint = c.UpdateDeviceLicenseWithValue()
				data, err = licenserecc.BuildUpdateDeviceLicenseWithValuePayload(*licenseRecUpdateDeviceLicenseWithValueLicenseIDFlag, *licenseRecUpdateDeviceLicenseWithValueConsumptionValueFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// license-recUsage displays the usage of the license-rec command and its
// subcommands.
func licenseRecUsage() {
	fmt.Fprintf(os.Stderr, `This service handles License Reconcile based actions.
Usage:
    %s [globalflags] license-rec COMMAND [flags]

COMMAND:
    get-license: Returns license count
    update-device-license: Increments device license count given the LicenseID
    update-device-license-with-value: Increments license count given value

Additional help:
    %s license-rec COMMAND --help
`, os.Args[0], os.Args[0])
}
func licenseRecGetLicenseUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] license-rec get-license -licenseid INT

Returns license count
    -licenseid INT: 

Example:
    `+os.Args[0]+` license-rec get-license --licenseid 4847839135972539348
`, os.Args[0])
}

func licenseRecUpdateDeviceLicenseUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] license-rec update-device-license -licenseid INT

Increments device license count given the LicenseID
    -licenseid INT: 

Example:
    `+os.Args[0]+` license-rec update-device-license --licenseid 4326563986109237020
`, os.Args[0])
}

func licenseRecUpdateDeviceLicenseWithValueUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] license-rec update-device-license-with-value -licenseid INT -consumption-value INT

Increments license count given value
    -licenseid INT: 
    -consumption-value INT: 

Example:
    `+os.Args[0]+` license-rec update-device-license-with-value --licenseid 2716133480549925254 --consumption-value 3495748079444956790
`, os.Args[0])
}
