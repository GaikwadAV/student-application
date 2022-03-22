package main

import (
	"RESTAPI/controller"
)

//go:generate swagger generate spec -m -o ./swagger.json
func main() {

	controller.Handlerequest()

}
