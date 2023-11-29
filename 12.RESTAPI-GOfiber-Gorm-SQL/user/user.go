package user

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/fiberdb?charset=utf8mb4&parseTime=True&loc=Local"
	//user : root , pw : root , dbname : fiberdb

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected")
	db.AutoMigrate(&User{})
	DB = db

}

// Get all users
func GetUsers(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)
	return c.JSON(&users)
}

// Get all users
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.Find(&user, id)
	return c.JSON(&user)
}

// Get all users
func SaveUser(c *fiber.Ctx) error {
	// user := new(User)
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&user)
	return c.JSON(&user)
}

// Get all users
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	DB.Delete(&user)
	return c.SendString("User is deleted")
}

// Get all users
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&user)
	return c.JSON(&user)
}
