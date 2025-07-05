package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school_management_system/internal/model"
	"school_management_system/internal/repo/sqlconnect"
)

// this is only for tasting purpose
func saveEntry(res http.ResponseWriter,req *http.Request){
	person:=model.Person{}

	err:=json.NewDecoder(req.Body).Decode(&person)
	if err!=nil{
		fmt.Println("error occured during decode")
	}
	db:=sqlconnect.SqlConnector();
	st,err:=db.Prepare("INSERT INTO testgoapi (name,lastName,age) values(?,?,?)")
	if err!=nil{
		fmt.Println("error occured during prepare a value")
	}

	_,err=st.Exec(person.Name,person.LastName,person.Age)

	if err!=nil{
		fmt.Println("error occured during execute a statement")
	}
	res.Write([]byte("data save successfully"))

}

func RootHandler(res http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		res.Write([]byte("Hello from root inside GET Method"))
	case http.MethodPost:
		saveEntry(res,req)
	case http.MethodPut:
		res.Write([]byte("Hello from root inside PUT Method"))
	case http.MethodDelete:
		res.Write([]byte("Hello from root inside DELETE Method"))
	default:
		res.Write([]byte("you use " + req.Method + " Method"))

	}

}