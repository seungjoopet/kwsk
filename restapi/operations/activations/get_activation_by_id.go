// Code generated by go-swagger; DO NOT EDIT.

package activations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/projectodd/kwsk/models"
)

// GetActivationByIDHandlerFunc turns a function with the right signature into a get activation by Id handler
type GetActivationByIDHandlerFunc func(GetActivationByIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetActivationByIDHandlerFunc) Handle(params GetActivationByIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetActivationByIDHandler interface for that can handle valid get activation by Id params
type GetActivationByIDHandler interface {
	Handle(GetActivationByIDParams, *models.Principal) middleware.Responder
}

// NewGetActivationByID creates a new http.Handler for the get activation by Id operation
func NewGetActivationByID(ctx *middleware.Context, handler GetActivationByIDHandler) *GetActivationByID {
	return &GetActivationByID{Context: ctx, Handler: handler}
}

/*GetActivationByID swagger:route GET /namespaces/{namespace}/activations/{activationid} Activations getActivationById

Get activation information

Get activation information.

*/
type GetActivationByID struct {
	Context *middleware.Context
	Handler GetActivationByIDHandler
}

func (o *GetActivationByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetActivationByIDParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
