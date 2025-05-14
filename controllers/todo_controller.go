package controllers

import (
	"go_todo/config"
	"go_todo/models"

	"github.com/gofiber/fiber/v2"
)

// CreateTodo: สร้างรายการ Todo ใหม่
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	// อ่านข้อมูลจาก body ของ request
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// ตรวจสอบว่า User มีอยู่จริงหรือไม่ โดยเช็คจาก userID
	var user models.User
	if err := config.DB.First(&user, todo.UserID).Error; err != nil {
		// หากไม่พบผู้ใช้จะส่ง error กลับ
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// บันทึกรายการ Todo ลงในฐานข้อมูล
	config.DB.Create(&todo)
	return c.Status(201).JSON(todo)
}

// GetTodos: ดึงรายการ Todo ทั้งหมด
func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo

	// ดึงข้อมูลรายการ Todo ทั้งหมดจากฐานข้อมูล
	config.DB.Find(&todos)
	return c.JSON(todos)
}

// GetTodo: ดึงรายการ Todo ตาม ID
func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	// ค้นหารายการ Todo ตาม ID ที่ได้รับจาก URL params
	if err := config.DB.First(&todo, id).Error; err != nil {
		// หากไม่พบ Todo จะส่ง error กลับ
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}
	return c.JSON(todo)
}

// GetTodosByUser: ดึงรายการ Todo ตาม userID
func GetTodosByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	var todos []models.Todo

	// ค้นหารายการ Todo ที่เป็นของ user ตาม userID
	config.DB.Where("user_id = ?", userID).Find(&todos)
	return c.JSON(todos)
}

// UpdateTodo: อัปเดตรายการ Todo ตาม ID
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	// ค้นหารายการ Todo ตาม ID ที่ได้รับจาก URL params
	if err := config.DB.First(&todo, id).Error; err != nil {
		// หากไม่พบ Todo จะส่ง error กลับ
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	// อ่านข้อมูลใหม่จาก body ของ request เพื่ออัปเดต
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// บันทึกข้อมูลที่อัปเดตลงในฐานข้อมูล
	config.DB.Save(&todo)
	return c.JSON(todo)
}

// DeleteTodo: ลบรายการ Todo ตาม ID
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	// ค้นหารายการ Todo ตาม ID ที่ได้รับจาก URL params
	if err := config.DB.First(&todo, id).Error; err != nil {
		// หากไม่พบ Todo จะส่ง error กลับ
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	// ลบรายการ Todo ออกจากฐานข้อมูล
	config.DB.Delete(&todo)
	return c.SendStatus(204) // ส่ง status 204 No Content หลังจากลบสำเร็จ
}
