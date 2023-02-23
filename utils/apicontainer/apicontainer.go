package apicontainer

var (
	data = make(map[string]APIHandler)
)

func Register(handlers map[string]APIHandler) {
	for key, handler := range handlers {
		data[key] = handler
	}
}

func Get(path string) (handler APIHandler, ok bool) {
	handler, ok = data[path]
	return
}
