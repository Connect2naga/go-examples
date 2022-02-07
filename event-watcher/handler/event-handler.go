// Package handler contains ...
package handler

import (
	"context"

	"github.com/connect2naga/go-examples/event-watcher/model"
	"github.com/connect2naga/logger/logging"
)

/*
Author : Nagarjuna S
Date : 2/7/22 1:46 PM
Project : go-examples
File : event-handler.go
*/

// Handler is implemented by any handler.
// The Handle method is used to process event
type Handler interface {
	Init() error
	Handle(e model.Event)
}

type PrintHandler struct {
	logger logging.Logger
}

func NewPrintHandler() Handler {
	return &PrintHandler{}
}
func (p *PrintHandler) Init() error {
	p.logger = logging.NewLogger()
	return nil
}

func (p *PrintHandler) Handle(e model.Event) {
	p.logger.Infof(context.TODO(), "Event received : %+v", e)
}
