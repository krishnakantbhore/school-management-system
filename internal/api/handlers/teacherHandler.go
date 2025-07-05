package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school_management_system/internal/model"
	"school_management_system/internal/repo/sqlconnect"
)
func saveTeacher(res http.ResponseWriter,req *http.Request){
	teacherData:=[]model.Teacher{}
 

	err:=json.NewDecoder(req.Body).Decode(&teacherData)

	if err!=nil{
		// res.Write([]byte("invalid Data"))
		fmt.Println("invalid Data")
	}

	fmt.Println(teacherData)
	db:=sqlconnect.SqlConnector();
	st,err:=db.Prepare("INSERT INTO teacher(firstName,lastName,class, subject,email) VALUES(?,?,?,?,?)")
	defer db.Close()
	
	if err!=nil{
		// res.Write([]byte("Internal Server Error"))
		fmt.Println("Internal server error")
	}

	for _,teacher:=range teacherData{
		_,err=st.Exec(teacher.FirstName,teacher.LastName,teacher.Class,teacher.Subject,teacher.Email)
	}
	
	defer st.Close()
	if err!=nil{
		// res.Write([]byte("Internal Server Error : statement execute failed..."))
		fmt.Println("Internal Server Error : statement execute failed...")
	}
	res.Write([]byte("Data save successfully..."))



}
func TeacherHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		res.Write([]byte("Hello from Teacher inside GET Method"))
	case http.MethodPost:
		saveTeacher(res,req)
	case http.MethodPut:
		res.Write([]byte("Hello from Teacher inside PUT Method"))
	case http.MethodDelete:
		res.Write([]byte("Hello from Teacher inside DELETE Method"))
	default:
		res.Write([]byte("you use " + req.Method + " Method"))

	}
}