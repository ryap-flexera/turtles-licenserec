package design

import . "goa.design/goa/dsl"

var _ = Service("licenseRec", func() {
	Description("This service handles License Reconcile based actions.")

	HTTP(func() {
		Path("/api/licenses/")
	})

	Method("GetLicense", func() {
		Description("Returns license count")
		Result(Int)
		Payload(LicenseObject)
		HTTP(func() {
			GET("/get/{LicenseID}")
			Param("LicenseID", Int)
			Response(StatusOK)
		})
	})

	Method("UpdateDeviceLicense", func() {
		Description("Increments device license count given the LicenseID")
		Result(Int)
		Payload(LicenseObject)
		HTTP(func() {
			POST("/post/{LicenseID}")
			Param("LicenseID", Int)
			Response(StatusOK)
		})
	})

	Method("UpdateDeviceLicenseWithValue", func() {
		Description("Increments license count given value")
		Result(Int)
		Payload(LicenseConsumption)
		HTTP(func() {
			POST("/post/{LicenseID}/{ConsumptionValue}")
			Param("LicenseID", Int)
			Param("ConsumptionValue", Int)
			Response(StatusOK)
		})
	})
})
