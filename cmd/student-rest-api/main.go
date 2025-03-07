package main

import (
	"fmt"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/database"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/routers"
	"log"
)

func main() {

	database.ConnectDB()
	r := routers.SetupRoutes()

	fmt.Println("Hello World.")

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
