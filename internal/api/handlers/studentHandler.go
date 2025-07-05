package handlers

import "net/http"

func StudentHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		res.Write([]byte("Hello from Student inside GET Method"))
	case http.MethodPost:
		res.Write([]byte("Hello from Student inside POST Method"))
	case http.MethodPut:
		res.Write([]byte("Hello from Student inside PUT Method"))
	case http.MethodDelete:
		res.Write([]byte("Hello from Student inside DELETE Method"))
	default:
		res.Write([]byte("you use " + req.Method + " Method"))

	}
}