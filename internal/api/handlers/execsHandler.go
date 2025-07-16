package handlers

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"school_management_system/internal/model"
	"school_management_system/internal/repo/sqlconnect"
	"school_management_system/pkg/utils"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
)

func SaveExecs(res http.ResponseWriter, req *http.Request){
	var execsData []model.Execs

	err:=json.NewDecoder(req.Body).Decode(&execsData)

	// data,err:=io.ReadAll(req.Body);
	// json.Unmarshal(data,&execsData)

	if(err != nil){
		fmt.Println("error occured during decode")
		return
	}

	db:=sqlconnect.SqlConnector();
	defer db.Close()

 

 

	for _,execs:=range execsData{
		// incode password
			salt:=make([]byte,16)
	_,err=rand.Read(salt)

	
	if err!=nil{
		fmt.Println("err occured during creating salt",err);
		return
	}

	passKey:=argon2.IDKey([]byte(execs.Password),salt,1,64*1024,4,32);
	fmt.Println(salt)
	fmt.Println(passKey)

	incodedSalt:=base64.StdEncoding.EncodeToString(salt);
	incodedPassword:=base64.StdEncoding.EncodeToString(passKey);

	fmt.Println(incodedSalt)
	fmt.Println(incodedPassword)

	execs.Password=fmt.Sprintf("%v.%v",incodedSalt,incodedPassword)

	_,err =	db.Exec("INSERT INTO execs (firstName,lastName,email,role,username,password) VALUES (?,?,?,?,?,?)",execs.FirstName,execs.LastName,execs.Email,execs.Role,execs.UserName,execs.Password)

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
			Data [] model.Execs `json:"data"`
	}{
		Status: "ok",
		Count: len(execsData),
		Data:execsData,
	}

	res.Header().Set("Content/Type","application/json")
	json.NewEncoder(res).Encode(responce)


}

func GetAllExecsData(res http.ResponseWriter, req *http.Request) {
	 db:=sqlconnect.SqlConnector();
	 defer db.Close()
   
	 var result []model.Execs

	 stmt,err:=db.Query("SELECT id,firstName,lastName,email,role,userName,password  FROM execs")
	 if err!=nil{
		fmt.Println(err)
		http.Error(res,"NO DATA FOUND",http.StatusNotFound)
		return
	 }

	 for stmt.Next(){
		var execs model.Execs
		err=stmt.Scan(&execs.Id,&execs.FirstName,&execs.LastName,&execs.Email,&execs.Role,&execs.UserName,&execs.Password);
		if err!=nil{
			fmt.Println(err)
			http.Error(res,"NO DATA FOUND",http.StatusNotFound)
			return
		 }
		result=append(result, execs)
	 }

	 responce:=struct {
		Status string `json:"staus"`
		Count  int     `json:"count"`
		Data [] model.Execs `json:"data"`
}{
	Status: "ok",
	Count: len(result),
	Data:result,
}

res.Header().Set("Content/Type","application/json")
json.NewEncoder(res).Encode(responce)

	 
}



func GetExec(res http.ResponseWriter,req *http.Request){

	value:=req.PathValue("id")
	id,err:=strconv.Atoi(value)
	if err!=nil{
		http.Error(res,"",http.StatusInternalServerError)
		return
	}

	var execs model.Execs
 
	db:=sqlconnect.SqlConnector();

	defer db.Close()

	err=db.QueryRow("SELECT id,firstName,lastName,email,role,username,password FROM execs WHERE id=?",id).Scan(&execs.Id,&execs.FirstName,&execs.LastName,&execs.Email,&execs.Role,&execs.UserName,&execs.Password);
	if err!=nil{
		if strings.Contains(err.Error(),"no rows in result set")	{
		http.Error(res,"INVALID ID",http.StatusNotFound)
		return
	}
	return
}

	responce:=struct {
		Status string `json:"status"`
		Data model.Execs `json:"data"`
	 
}{
	Status: "Ok", 
	Data: execs,
 
}
res.Header().Set("Content/Type","application/json")
json.NewEncoder(res).Encode(responce)



}


func UpdateExecs(res http.ResponseWriter,req *http.Request){

	value:=req.PathValue("id")
	id,err:=strconv.Atoi(value)
	if err!=nil{
		http.Error(res,"",http.StatusInternalServerError)
		return
	}

	var updatedData model.Execs

	json.NewDecoder(req.Body).Decode(&updatedData)

	db:=sqlconnect.SqlConnector();

	defer db.Close()

	_,err=db.Exec("UPDATE execs set firstName=?,lastName=?,email=?,rol? where id=?",updatedData.FirstName,updatedData.LastName,updatedData.Email,updatedData.Role,id)

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


func DeleteExecs(res http.ResponseWriter,req *http.Request){

	value:=req.PathValue("id")
	id,err:=strconv.Atoi(value)
	if err!=nil{
		http.Error(res,"",http.StatusInternalServerError)
		return
	}

	
	db:=sqlconnect.SqlConnector();

	defer db.Close()

	_,err=db.Exec(" DELETE FROM execs where id=?" ,id)

	if err!=nil{
		http.Error(res,"Delete FAILED",http.StatusInternalServerError)
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

func Login(res http.ResponseWriter,req *http.Request){


	fmt.Println("LOGIN ROUTE ===================================")
	data :=model.Execs{}
 
 
	// data validation

	err:=json.NewDecoder(req.Body).Decode(&data);

	if err != nil {
		fmt.Println("err",err)
		return 
	}

	if data.UserName=="" || data.Password==""{
		http.Error(res,"invalid Data",http.StatusNotAcceptable)
		return
	}

	// Search user is exists

	db:=sqlconnect.SqlConnector()
	resp:=model.Execs{}

	err=db.QueryRow("select firstName,lastName,email,role,username,password from execs where username=?",data.UserName).Scan(&resp.FirstName,&resp.LastName,&resp.Email,&resp.Role,&resp.UserName,&resp.Password);
if err != nil {
	
	http.Error(res,"NO DATA FOUND",http.StatusNotFound)
	
	return 
}
  // isUser active
    

	// verify password

	encodedPassword:=strings.Split(resp.Password,".")
   encodedSalt:=encodedPassword[0]
   encodedPasskey:=encodedPassword[1]

	salt,err := base64.StdEncoding.DecodeString(encodedSalt)

	if err != nil {
		http.Error(res,"INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return 
	}
	passKey,err:= base64.StdEncoding.DecodeString(encodedPasskey)

	if err != nil {
		http.Error(res,"INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return 
	}
	hash := argon2.IDKey([]byte(data.Password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(passKey){
		fmt.Println(err,"error occured dirung count len")
			http.Error(res,"INVALID PASSWORD",http.StatusInternalServerError)
		return 
	}

	result:=subtle.ConstantTimeCompare(hash,passKey)

	if result!=1{
		fmt.Println(result,"error occured during comparison")
			http.Error(res,"INVALID PASSWORD",http.StatusInternalServerError)
		return 
	}


	// generate token
	stringToken,err:= utils.GenerateToken(resp.Id,resp.UserName)

	if err != nil {
		http.Error(res,"Failed to generate token",http.StatusInternalServerError)
		return 
	}

	http.SetCookie(res,&http.Cookie{
		Name: "bariar",
		Value: stringToken,
		Path: "/",
		Expires: time.Now().Add(10*time.Minute),
	})
	// send responce as a token
		res.Write([]byte(stringToken))
}

func Logout(res http.ResponseWriter,req *http.Request){

	http.SetCookie(res,&http.Cookie{
		Name: "bariar",
		Value: "",
		Expires: time.Now(),
	})


	res.Write([]byte("Logout Successfully"))
}