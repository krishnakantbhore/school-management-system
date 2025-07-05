package routers

import (
	"net/http"
	"school_management_system/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)

	mux.HandleFunc("/teachers", handlers.TeacherHandler)

	mux.HandleFunc("/students", handlers.StudentHandler)

	mux.HandleFunc("/execs", handlers.ExecsHandler)
	return mux
}