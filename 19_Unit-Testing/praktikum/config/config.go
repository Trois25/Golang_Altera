package config

import(
	"praktikum/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := models.Config{

		DB_Username: "root",

		DB_Password: "",

		DB_Port: "3306",

		DB_Host: "localhost",

		DB_Name: "crud_go",
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",

		config.DB_Username,

		config.DB_Password,

		config.DB_Host,

		config.DB_Port,

		config.DB_Name,
	)

	var e error
	DB, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	InitialMigration()

}

func InitialMigration() {

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.UserResponse{})

}

func InitDBTest() {

	config := models.ConfigTest{

		DB_USER_TEST : "root",
		DB_PASS_TEST : "",
		DB_HOST_TEST : "localhost",
		DB_PORT_TEST : "3306",
		DB_NAME_TEST : "crud_go_test",
	}
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER_TEST, config.DB_PASS_TEST, config.DB_HOST_TEST, config.DB_PORT_TEST, config.DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	//DB.Migrator().DropTable(&models.User{})
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
	DB.Migrator().DropTable(&models.Book{})
	DB.AutoMigrate(&models.Book{})
}