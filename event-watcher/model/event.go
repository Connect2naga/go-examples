// Package model contains ...
package model

/*
Author : Nagarjuna S
Date : 2/7/22 1:44 PM
Project : go-examples
File : event.go
*/

// Event represent an event got from k8s api server
// Events from different endpoints need to be casted to Event
// before being able to be handled by handler
type Event struct {
	Namespace string
	Kind      string
	Component string
	Host      string
	Reason    string
	Status    string
	Name      string
}
