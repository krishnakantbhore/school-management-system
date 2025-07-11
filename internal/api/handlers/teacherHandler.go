package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school_management_system/internal/model"
	"school_management_system/internal/repo/sqlconnect"
	"strconv"
	"strings"
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
		if err!=nil{
			// res.Write([]byte("Internal Server Error : statement execute failed..."))
			fmt.Println("Internal Server Error : statement execute failed...")
		}
	}
	
	defer st.Close()
	
	res.Write([]byte("Data save successfully..."))



}

func getTeacher(res http.ResponseWriter,req *http.Request){
 

	path:=strings.TrimPrefix(req.URL.Path,"/teachers/")
	id:=strings.TrimSuffix(path,"/")
	firstName:=""
	lastName:=""

	db:=sqlconnect.SqlConnector();
	teacher:=model.Teacher{}
	
	fmt.Println("path: ",path)
	fmt.Println("id ",id)

	if id==""{
		firstName=req.URL.Query().Get("firstName")
	 
		lastName=req.URL.Query().Get("lastName")
		 
		query:="SELECT firstName,lastName,class,subject,email FROM teacher WHERE 1=1"
		var args []interface{}
		if firstName!=""{
			query+=" AND firstName = ?"
			args=append(args,firstName)
		}

		if lastName!=""{
			query+=" AND lastName = ?"
			args=append(args,lastName)
		}
		row,err:=db.Query(query,args...)
		if err!=nil{
			fmt.Println("error occured during Query")
			return
		}
		defer row.Close()

		var teacherList [] model.Teacher
         for row.Next(){
					 var t model.Teacher
					 row.Scan(&t.FirstName,&t.LastName,&t.Class,&t.Subject,&t.Email)
					 teacherList=append(teacherList, t)
				 }

     res.Header().Set("Content-type","application/json")
		 json.NewEncoder(res).Encode(teacherList)
return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
			http.Error(res, "Invalid ID", http.StatusBadRequest)
			return
	}


	err=db.QueryRow("SELECT firstname, lastname, class, subject, email FROM teacher WHERE id=?",idInt).Scan(&teacher.FirstName,&teacher.LastName,&teacher.Class,&teacher.Subject,&teacher.Email)

	if err!=nil{
		res.Write([]byte("no data found"))
		return
	}


	res.Header().Set("Content-Type","application/json")
 err=json.NewEncoder(res).Encode(teacher)
 if err!=nil{
	fmt.Print("error occureed during Encode")
	return
 }
	 
}

func updateTeacher(res http.ResponseWriter,req *http.Request){
	teacher:= model.Teacher{}
   
	 path:=strings.TrimPrefix(req.URL.Path,"/teachers/")
	 id,err:=strconv.Atoi(path)

	 if err!=nil{
		http.Error(res,"invalid id",http.StatusBadRequest)
	 }

	 db:=sqlconnect.SqlConnector()

	 err=db.QueryRow("SELECT firstName,lastName,class,subject,email from teacher where id=?",id).Scan(&teacher.FirstName,&teacher.LastName,&teacher.Class,&teacher.Subject,&teacher.Email)

	 if err!=nil{
		http.Error(res,"Not Found",http.StatusNotFound)
		return
	 }

	 newTeacher:=model.Teacher{}

	 json.NewDecoder(req.Body).Decode(&newTeacher)

	 db.Exec("update teacher set firstName=?,lastName=?,class=?,subject=?,email=? where id=?",newTeacher.FirstName,newTeacher.LastName,newTeacher.Class,newTeacher.Subject,newTeacher.Email,id);

	 res.Header().Set("COntent-Type","application/json")
	 json.NewEncoder(res).Encode(newTeacher)


}

func deleteTeacher(res http.ResponseWriter,req *http.Request){
	path:=strings.TrimPrefix(req.URL.Path,"/teachers/")
	id,err:=strconv.Atoi(path)

	if err!=nil{
	 http.Error(res,"invalid id",http.StatusBadRequest)
	}

	db:=sqlconnect.SqlConnector()

	result,err := db.Exec("DELETE FROM teacher WHERE id=?",id)
	if err!=nil{
		http.Error(res,"DATA NOT FOUND",http.StatusNotFound)
		return
	}
	 rowEffected,err:=result.RowsAffected()
	 if(err!=nil){
		http.Error(res,"Internal Server Error",http.StatusInternalServerError)
		return
	 }

	 if rowEffected==0{
		http.Error(res,"DATA NOT FOUND",http.StatusNotFound)
		return
	 }

	res.Write([]byte("Delete Succesfully"))
}
func TeacherHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		 getTeacher(res,req)
	case http.MethodPost:
		saveTeacher(res,req)
	case http.MethodPut:
		//  PUT handler
		updateTeacher(res,req)
	case http.MethodDelete:
		 deleteTeacher(res,req)
	default:
		res.Write([]byte("you use " + req.Method + " Method"))

	}
}