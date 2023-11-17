package main

import (
	"aqaryint/src/controllers"
	"aqaryint/src/database"
	"aqaryint/src/repositories"
	"aqaryint/src/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	//config.Appconfig = config.GetConfig()
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	studantRepository := repositories.NewStudantRepository(database.DB)
	studantService := services.NewStudentService(studantRepository)
	studantController := controllers.NewStudentController(studantService)

	r := gin.Default()

	r.POST("/users", studantController.CreateStudent)
	r.GET("/users/:id", studantController.GetStudent)
	r.PUT("/users/:id", studantController.UpdateStudent)
	r.DELETE("/users/:id", studantController.DeleteStudent)

	r.Run(":8080")

}
