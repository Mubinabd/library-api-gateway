package main

import (
	// "log"
	r "github.com/Mubinabd/library-api-gateway/api-gateway"
	"github.com/Mubinabd/library-api-gateway/api-gateway/handler"

)

func main() {

	engine := r.NewGin(handler.NewHandlerStruct())
	engine.Run(":8090")
}