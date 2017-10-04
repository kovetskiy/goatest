package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var (
	SomeData = MediaType(
		"application/vnd.some.data+json",
		func() {
			Attributes(func() {
				Attribute("value", Integer)
				Required("value")
			})

			View("default", func() {
				Attribute("value")
			})
		},
	)
)

func init() {
	Resource("some", func() {
		BasePath("/")

		Action("get", func() {
			Description("Get some value")
			Routing(GET("/some"))

			Response(OK, func() {
				Media(SomeData, "default")
			})
		})
	})
}
