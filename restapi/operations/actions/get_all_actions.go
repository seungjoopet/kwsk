// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/projectodd/kwsk/models"
)

// GetAllActionsHandlerFunc turns a function with the right signature into a get all actions handler
type GetAllActionsHandlerFunc func(GetAllActionsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllActionsHandlerFunc) Handle(params GetAllActionsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetAllActionsHandler interface for that can handle valid get all actions params
type GetAllActionsHandler interface {
	Handle(GetAllActionsParams, *models.Principal) middleware.Responder
}

// NewGetAllActions creates a new http.Handler for the get all actions operation
func NewGetAllActions(ctx *middleware.Context, handler GetAllActionsHandler) *GetAllActions {
	return &GetAllActions{Context: ctx, Handler: handler}
}

/*GetAllActions swagger:route GET /namespaces/{namespace}/actions Actions getAllActions

Get all actions

Get all actions

*/
type GetAllActions struct {
	Context *middleware.Context
	Handler GetAllActionsHandler
}

func (o *GetAllActions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllActionsParams()

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
