package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"

)
type apiConfig struct {
	DB *database.queries
}
func main() {
	fmt.Println("Hello, World! ")
	godotenv.Load(".env")
	
	portString:= os.Getenv("PORT") 
	if portString ==""{
		log.Fatal("PORT is not found in the environment")
	}


	dbURL:= os.Getenv("DB_URL") 
	if dbURL ==""{
		log.Fatal("DB_URL is not found in the environment")
	}
conn, err := sql.Open("postgres", dbURL)
if err!=nil{
	log.Fatal("Cannot connect to database:", err)
}

apicfg:= apiConfig{
DB: database.New(conn),
}
	fmt.Println("Port is set to:", portString)
}


	