package models

//config "aqaryint/aqaryint/src/configs"

//gorm "github.com/jinzhu/gorm"
//postgres _ "github.com/jinzhu/gorm/dialects/postgres"

//"fmt"

//var db *gorm.DB

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city"`
	Class rune   `json:"class"`
	Fit   bool   `json:fit`
}

// func initDataBase() {

// 	username := config.Appconfig.GetString("admin")
// 	password := config.Appconfig.GetString("123")
// 	dbName := config.Appconfig.GetString("demo_db")
// 	dbHost := config.Appconfig.GetString("localhost")

// 	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
// 	fmt.Println(dbUri)

// 	conn, err := gorm.Open("postgres", dbUri)
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	db = conn
// 	db.Debug().AutoMigrate(&User{}, &Task{})
// }

// func GetDB() *gorm.DB {
// 	return db
// }
