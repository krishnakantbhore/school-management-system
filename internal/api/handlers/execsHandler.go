package handlers

import "net/http"

func ExecsHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		res.Write([]byte("Hello from Execs inside GET Method"))
	case http.MethodPost:
		res.Write([]byte("Hello from Execs inside POST Method"))
	case http.MethodPut:
		res.Write([]byte("Hello from Execs inside PUT Method"))
	case http.MethodDelete:
		res.Write([]byte("Hello from Execs inside DELETE Method"))
	default:
		res.Write([]byte("you use " + req.Method + " Method"))

	}
}