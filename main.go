//go:generate goagen bootstrap -d github.com/kovetskiy/goatest/design
package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/kovetskiy/goatest/app"
)

func main() {
	service := goa.New("some")

	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	app.MountSomeController(service, NewSomeController(service))

	err := service.ListenAndServe(":8080")
	if err != nil {
		panic(err)
	}
}
