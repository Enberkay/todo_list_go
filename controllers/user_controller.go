package controllers

import (
	"go_todo/config"
	"go_todo/models"

	"github.com/gofiber/fiber/v2"
)

// CreateUser: สร้างผู้ใช้งานใหม่
func CreateUser(c *fiber.Ctx) error {
	// สร้างตัวแปร user จาก model
	user := new(models.User)

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// ตรวจสอบว่า email ซ้ำไหม
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email already exists"})
	}

	// บันทึกข้อมูลผู้ใช้ใหม่ในฐานข้อมูล
	config.DB.Create(&user)
	return c.Status(201).JSON(user)
}

// GetUsers: ดึงข้อมูลผู้ใช้ทั้งหมด
func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	// ดึงข้อมูลผู้ใช้ทั้งหมดจากฐานข้อมูล
	if err := config.DB.Preload("Todos").Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.JSON(users)
}

// GetUser: ดึงข้อมูลผู้ใช้ตาม ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// หา user จากฐานข้อมูล
	if err := config.DB.Preload("Todos").First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

// UpdateUser: อัปเดตข้อมูลผู้ใช้
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// หา user ตาม ID
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// อ่านข้อมูลจาก request body เพื่ออัปเดต
	var updateData models.User
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	// อัปเดตข้อมูลของ user
	user.Name = updateData.Name
	user.Email = updateData.Email
	user.Password = updateData.Password

	// บันทึกข้อมูลที่อัปเดตในฐานข้อมูล
	config.DB.Save(&user)

	return c.JSON(user)
}

// DeleteUser: ลบผู้ใช้ตาม ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// หา user ที่ต้องการลบ
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// ลบ user จากฐานข้อมูล
	config.DB.Delete(&user)

	return c.SendStatus(204)
}
