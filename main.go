package main

import (
	config "aqaryint/src/configs"
	"aqaryint/src/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
	postgres "gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	config.Appconfig = config.GetConfig()
	NewRouter()

}

func Init() {
	router := NewRouter()
	router.Run(config.Appconfig.GetString("server.port"))
}
func NewRouter() *gin.Engine {
	router := gin.New()
	//router.Handler()

	{
		router.POST("/limitedPostMethod", controllers.PostMethod)
		router.POST("/add-user", controllers.AddUser)//controller.CreateAccount
		router.POST("/createTables", controllers.CreateTable)
	}
	return router
}

func initDataBase() {

	username := config.Appconfig.GetString("database.username")
	password := config.Appconfig.GetString("database.password")
	dbName := config.Appconfig.GetString("database.name")
	dbHost := config.Appconfig.GetString("database.host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	db = conn

}

func GetDB() *gorm.DB {
	return db
}
