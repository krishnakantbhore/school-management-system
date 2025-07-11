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

func SaveStudent(res http.ResponseWriter, req *http.Request){
	var studentData []model.Student

	err:=json.NewDecoder(req.Body).Decode(&studentData)

	if(err != nil){
		fmt.Println("error occured during decode")
		return
	}

	db:=sqlconnect.SqlConnector();
	defer db.Close()

	for _,student:=range studentData{
	_,err =	db.Exec("INSERT INTO student (firstName,lastName,class,email) VALUES (?,?,?,?)",student.FirstName,student.LastName,student.Class,student.Email)

	if(err!=nil){
		fmt.Print("error occured during save",err) 
		cError:=strings.Split(err.Error(), ":")
		http.Error(res,cError[1],http.StatusInternalServerError)
		return
	}
	}
	 
	responce:=struct {
      Status string `json:"staus"`
			Count  int     `json:"count"`
			Data [] model.Student `json:"data"`
	}{
		Status: "ok",
		Count: len(studentData),
		Data:studentData,
	}

	res.Header().Set("Content/Type","application/json")
	json.NewEncoder(res).Encode(responce)


}

func GetAllStudentData(res http.ResponseWriter, req *http.Request) {
	 db:=sqlconnect.SqlConnector();
	 defer db.Close()
   
	 var result []model.Student

	 stmt,err:=db.Query("SELECT id,firstName,lastName,class,email FROM STUDENT")
	 if err!=nil{
		http.Error(res,"NO DATA FOUND",http.StatusNotFound)
		return
	 }

	 for stmt.Next(){
		var student model.Student
		err=stmt.Scan(&student.Id,&student.FirstName,&student.LastName,&student.Class,&student.Email);
		if err!=nil{
			http.Error(res,"NO DATA FOUND",http.StatusNotFound)
			return
		 }
		result=append(result, student)
	 }

	 responce:=struct {
		Status string `json:"staus"`
		Count  int     `json:"count"`
		Data [] model.Student `json:"data"`
}{
	Status: "ok",
	Count: len(result),
	Data:result,
}

res.Header().Set("Content/Type","application/json")
json.NewEncoder(res).Encode(responce)

	 
}



func GetStudent(res http.ResponseWriter,req *http.Request){

	value:=req.PathValue("id")
	id,err:=strconv.Atoi(value)
	if err!=nil{
		http.Error(res,"",http.StatusInternalServerError)
		return
	}

	var studentdData model.Student
 
	db:=sqlconnect.SqlConnector();

	defer db.Close()

	err=db.QueryRow("SELECT id,firstName,lastName,class,email FROM STUDENT WHERE id=?",id).Scan(&studentdData.Id,&studentdData.FirstName,&studentdData.LastName,&studentdData.Class,&studentdData.Email)
	if err!=nil{
		if strings.Contains(err.Error(),"no rows in result set")	{
		http.Error(res,"INVALID ID",http.StatusNotFound)
		return
	}
	return
}

	responce:=struct {
		Status string `json:"status"`
		Data model.Student `json:"data"`
	 
}{
	Status: "Ok", 
	Data: studentdData,
 
}
res.Header().Set("Content/Type","application/json")
json.NewEncoder(res).Encode(responce)



}


func UpdateStudent(res http.ResponseWriter,req *http.Request){

	value:=req.PathValue("id")
	id,err:=strconv.Atoi(value)
	if err!=nil{
		http.Error(res,"",http.StatusInternalServerError)
		return
	}

	var updatedData model.Student

	json.NewDecoder(req.Body).Decode(&updatedData)

	db:=sqlconnect.SqlConnector();

	defer db.Close()

	_,err=db.Exec("UPDATE student set firstName=?,lastName=?,class=?,email=? where id=?",updatedData.FirstName,updatedData.LastName,updatedData.Class,updatedData.Email,id)

	if err!=nil{
		http.Error(res,"UPDATE FAILED",http.StatusInternalServerError)
		return
	}

	responce:=struct {
		Status string `json:"staus"`
	 
}{
	Status: "update Successfully", 
 
}
res.Header().Set("Content/Type","application/json")
json.NewEncoder(res).Encode(responce)



}


func DeleteStudent(res http.ResponseWriter,req *http.Request){

	value:=req.PathValue("id")
	id,err:=strconv.Atoi(value)
	if err!=nil{
		http.Error(res,"",http.StatusInternalServerError)
		return
	}

	
	db:=sqlconnect.SqlConnector();

	defer db.Close()

	_,err=db.Exec(" DELETE FROM student where id=?" ,id)

	if err!=nil{
		http.Error(res,"UPDATE FAILED",http.StatusInternalServerError)
		return
	}

	responce:=struct {
		Status string `json:"staus"`
	 
}{
	Status: "DELETE Successfully", 
 
}
res.Header().Set("Content/Type","application/json")
json.NewEncoder(res).Encode(responce)



}