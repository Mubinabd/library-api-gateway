package handler

import (
	"github.com/Mubinabd/library-api-gateway/clients"

)

type HandlerStruct struct {
	Clients  clients.Clients
}

func NewHandlerStruct() *HandlerStruct {
	return &HandlerStruct{
		Clients:*clients.NewClients(),   
	}
}