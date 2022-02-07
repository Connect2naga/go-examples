// Package event_watcher contains ...
package main

import (
	"github.com/connect2naga/go-examples/event-watcher/controller"
	"github.com/connect2naga/go-examples/event-watcher/handler"
	"log"
)

/*
Author : Nagarjuna S
Date : 2/7/22 1:13 PM
Project : go-examples
File : event_watcher.go
*/


type EventWatcher struct {

}

func (ew *EventWatcher)Start(){
	eventHandler := new(handler.PrintHandler)
	if err := eventHandler.Init(); err != nil {
		log.Fatal(err)
	}
	controller.Start(eventHandler)
}
func main() {
	new(EventWatcher).Start()
}