package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct{
	id int  `json:"id"`
	Body string  `json:"Body`
	completed bool `json:"Completed"`
}
const server=(":4000")
func main(){
	todos := [] Todo{}
	fmt.Println("Starting server at",server)
	app:=fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"Messsage":"This is TESTING"})
	})

	app.Post("api/detail", func(c *fiber.Ctx) error {
		todo := &Todo{}
		c.BodyParser(todo)

		if err := c.BodyParser(todo); err != nil{
			return err
		}

		if todo.Body==""{
			return c.Status(400).JSON(fiber.Map{"Error" :"Body is empty"})
		}
		todo.id =len(todos)+1
		todos=append(todos,*todo)
		
		return c.Status(201).JSON(todo)
	})

	log.Fatal(app.Listen(server))
}


