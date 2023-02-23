package apicontainer

import "context"

type APIHandler interface {
	Call(context.Context) (interface{}, error)
}
