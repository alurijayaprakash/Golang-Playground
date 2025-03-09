package main

import (
	"Restgofiber/user"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World")
}

func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	app.Get("/user/:id", user.GetUser)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)

}

func main() {
	fmt.Println("Welsome to REST API GO FIBER SIMPLE APP")
	user.InitialMigration() //DB connection & migration

	app := fiber.New()

	Routers(app)

	app.Get("/", HelloWorld)

	log.Fatal(app.Listen(":3000"))
}
