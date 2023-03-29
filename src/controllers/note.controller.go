package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateNoteHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": "note"}})
}
