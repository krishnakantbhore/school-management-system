package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"school_management_system/internal/api/routers"
	"school_management_system/internal/repo/sqlconnect"

	"github.com/joho/godotenv"
)
 
func main() {
	router:=routers.Router()

	sqlconnect.SqlConnector()
	err:= godotenv.Load();
	if err!= nil{
		log.Println(err)
	}
	 
	fmt.Println("server run on port ",os.Getenv("PORT"))
	
	err=http.ListenAndServe(os.Getenv("PORT"),router)
	if err!= nil{
		log.Println(err)
	}
}