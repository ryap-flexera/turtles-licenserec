package design

import (
	. "goa.design/goa/dsl"
	//cors "goa.design/plugins/cors/dsl"
)

var _ = API("License Reconcile API", func() {
	Title("License Reconcile API Service")
	Description("This Service handles License Reconcile related actions")

	// Server describes a single process listening for client requests. The DSL
	// defines the set of services that the server hosts as well as hosts details.
	Server("LicenseAPIServer", func() {
		Description("frontAPIServer hosts the services.")
		Services("licenseRec")

		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8080/")
		})
	})

})

var LicenseObject = Type("LicenseObject", func() {
	TypeName("LicenseObject")
	Attribute("LicenseID", Int)
	Required("LicenseID")
})

var LicenseConsumption = Type("LicenseConsumption", func() {
	TypeName("LicenseConsumption")
	Attribute("LicenseID", Int)
	Required("LicenseID")

	Attribute("ConsumptionValue", Int)
	Required("ConsumptionValue")
})
