package http

import (
	"net/http"

	"github.com/gsp-lang/gsp/core"
)

func HandleFunc(pattern core.Any, _handler func(core.Any, core.Any) core.Any) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_handler(w, r)
	}
	http.HandleFunc(pattern.(string), handler)
}

func FileServer(root core.Any) core.Any {
	return http.FileServer(http.Dir(root.(string)))
}

func ListenAndServe(port, _handler core.Any) {
	var handler http.Handler
	if _handler != nil {
		handler = _handler.(http.Handler)
	}
	http.ListenAndServe(port.(string), handler)
}
