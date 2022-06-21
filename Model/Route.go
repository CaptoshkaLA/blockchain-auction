package Model

import (
	"net/http"
	"time"
)

/*
	Route is a structure for routing
*/
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

/*
	ApiResponse is a structure for http response
*/
type ApiResponse struct {
	Name   string
	Status string
	Time   time.Time
}
