// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetStationsID(params *GetStationsIDParams) (*GetStationsIDOK, error)

	GetStopsIDArrivals(params *GetStopsIDArrivalsParams) (*GetStopsIDArrivalsOK, error)

	GetStopsIDDepartures(params *GetStopsIDDeparturesParams) (*GetStopsIDDeparturesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetStationsID returns a stop station from db stations

  Returns a stop/station from [db-stations](https://npmjs.com/package/db-stations).
*/
func (a *Client) GetStationsID(params *GetStationsIDParams) (*GetStationsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStationsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStationsID",
		Method:             "GET",
		PathPattern:        "/stations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStationsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStationsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStationsID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStopsIDArrivals fetches arrivals at a stop station

  Works like `/stops/{id}/departures`, except that it uses [`hafasClient.arrivals()`](https://github.com/public-transport/hafas-client/blob/5/docs/arrivals.md) to **query arrivals at a stop/station**.
*/
func (a *Client) GetStopsIDArrivals(params *GetStopsIDArrivalsParams) (*GetStopsIDArrivalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStopsIDArrivalsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStopsIDArrivals",
		Method:             "GET",
		PathPattern:        "/stops/{id}/arrivals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStopsIDArrivalsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStopsIDArrivalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStopsIDArrivals: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetStopsIDDepartures fetches departures at a stop station

  Uses [`hafasClient.departures()`](https://github.com/public-transport/hafas-client/blob/5/docs/departures.md) to **query departures at a stop/station**.
*/
func (a *Client) GetStopsIDDepartures(params *GetStopsIDDeparturesParams) (*GetStopsIDDeparturesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStopsIDDeparturesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStopsIDDepartures",
		Method:             "GET",
		PathPattern:        "/stops/{id}/departures",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStopsIDDeparturesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetStopsIDDeparturesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetStopsIDDepartures: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
