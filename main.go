package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type StoryJira struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Printf("Hello Thuy Nhwe")

	app := fiber.New()

	stories := []StoryJira{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Thuy oi"})
	})

	//create a story
	app.Post("/api/stories", func(c *fiber.Ctx) error {
		story := &StoryJira{}

		if err := c.BodyParser(story); err != nil {
			return err
		}

		if story.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Story Jira body is required"})
		}

		story.ID = len(stories) + 1
		stories = append(stories, *story)

		return c.Status(201).JSON(story)
	})

	//update a story
	app.Patch("/api/stories/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, story := range stories {
			if fmt.Sprint(story.ID) == id {
				stories[i].Completed = true
				return c.Status(201).JSON(stories[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Story is not found"})
	})

	log.Fatal(app.Listen(":4000"))
}
