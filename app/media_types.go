// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: Application Media Types
//
// Command:
// $ goagen
// --design=github.com/kovetskiy/goatest/design
// --out=$(GOPATH)/src/github.com/kovetskiy/goatest
// --version=v1.3.0

package app

// SomeData media type (default view)
//
// Identifier: application/vnd.some.data+json; view=default
type SomeData struct {
	Value int `form:"value" json:"value" xml:"value"`
}