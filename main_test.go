package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/kovetskiy/goatest/client"
)

func TestErrorMessage(t *testing.T) {
	cli := client.New(nil)
	cli.Scheme = "http"
	cli.Host = "localhost:8080"

	response, err := cli.GetSome(context.TODO(), client.GetSomePath())
	fmt.Printf("XXXXXX err: %s\n", err)
	fmt.Printf("XXXXXX response: %#v\n", response)
	fmt.Printf("XXXXXX response.Status: %#v\n", response.Status)

	some, decodeErr := cli.DecodeSomeData(response)
	fmt.Printf("XXXXXX decodeErr: %#v\n", decodeErr)
	fmt.Printf("XXXXXX decodeErr: %#v\n", decodeErr.Error())
	fmt.Printf("XXXXXX some: %#v\n", some)
}
