package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/payment-backend/config"
	"github.com/payment-backend/models"
	"github.com/payment-backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var input LoginRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	var user models.User

	result := config.DB.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "email atau password salah",
		})
	}

	checkPass := utils.CheckPassword(user.Password, input.Password)

	if !checkPass {
		return c.Status(401).JSON(fiber.Map{
			"message": "email atau password salah",
		})
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal membuat token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "login berhasil",
		"token":   token,
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

func Register(c *fiber.Ctx) error {

	var input models.User

	if err := c.BodyParser(&input); err != nil {

		return c.Status(400).JSON(
			fiber.Map{
				"message": "invalid request",
			},
		)
	}

	hash, _ := utils.HashPassword(
		input.Password,
	)

	user := models.User{

		Name: input.Name,

		Email: input.Email,

		Password: hash,

		Role: "customer",
	}

	config.DB.Create(&user)

	return c.JSON(
		fiber.Map{
			"message": "register success",
		},
	)
}
