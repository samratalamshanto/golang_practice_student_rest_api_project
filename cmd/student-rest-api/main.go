package main

import (
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/database"
	"github.com/samratalamshanto/student_rest_api_project/cmd/student-rest-api/routers"
	"log"
)

func main() {

	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	r := routers.SetupRoutes()

	err := r.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}
}
