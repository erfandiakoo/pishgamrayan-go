# Pishgamrayan sms package for golang

[![Tests](https://github.com/erfandiakoo/pishgamrayan-go/actions/workflows/go.yml/badge.svg)](https://github.com/erfandiakoo/pishgamrayan-go/actions/workflows/go.yml)

## Installation
```sh
go get github.com/erfandiakoo/pishgamrayan-go
```

## Useage

```go
package main
import (
	"fmt"
	"net/url"
	"github.com/erfandiakoo/pishgamrayan-go"
)
func main() {
	api := pishgamrayan.New(" your apikey ")
	sender := ""                 
	receptor := []string{"", ""}
	message := "Hello Go!" 
	if res, err := api.Message.Send(sender, receptor, message, nil); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			fmt.Println(err.Error())
		case *kavenegar.HTTPError:
			fmt.Println(err.Error())
		default:
			fmt.Println(err.Error())
		}
	} else {
		for _, r := range res {
			fmt.Println("MessageID 	= ", r.MessageID)
			fmt.Println("Status    	= ", r.Status)
			//...
		}
	}
}
```