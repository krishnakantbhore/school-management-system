package routers

import (
	"net/http"
	"school_management_system/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	// root handler
	mux.HandleFunc("/", handlers.RootHandler)

	// teacher handler
	mux.HandleFunc("/teachers/", handlers.TeacherHandler)


	// student handler
	mux.HandleFunc("GET /students", handlers.GetAllStudentData)
	mux.HandleFunc("GET /student/{id}", handlers.GetStudent)
	mux.HandleFunc("POST /saveStudent", handlers.SaveStudent)
	mux.HandleFunc("PUT /updateStudent/{id}", handlers.UpdateStudent)
	mux.HandleFunc("DELETE /deleteStudent/{id}", handlers.DeleteStudent)



	mux.HandleFunc("/execs", handlers.ExecsHandler)
	return mux
}