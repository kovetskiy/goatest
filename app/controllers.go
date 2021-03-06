// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: Application Controllers
//
// Command:
// $ goagen
// --design=github.com/kovetskiy/goatest/design
// --out=$(GOPATH)/src/github.com/kovetskiy/goatest
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// SomeController is the controller interface for the Some actions.
type SomeController interface {
	goa.Muxer
	Get(*GetSomeContext) error
}

// MountSomeController "mounts" a Some resource controller on the given service.
func MountSomeController(service *goa.Service, ctrl SomeController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetSomeContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/some", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Some", "action", "Get", "route", "GET /some")
}
