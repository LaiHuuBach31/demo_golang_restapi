package Controllers

import (
	db "demo_api/Config"
	models "demo_api/Models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetCashierList(c *fiber.Ctx) error {
	var cashiers []models.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashiers).Count(&count)

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Query cashier list successfully",
			"data":    cashiers,
		},
	)
}

func GetCashierDetail(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Select("id, name", "create_at", "update_at").Where("id = ?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createAt"] = cashier.CreateAt
	cashierData["updateAt"] = cashier.UpdateAt

	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
				"error":   map[string]interface{}{},
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Query cashier by id successfully",
			"data":    cashier,
		},
	)
}

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Invalid data",
			},
		)
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier Name is required",
			},
		)
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier Passcode is required",
			},
		)
	}

	cashier := models.Cashier{
		Name:     data["name"],
		Passcode: data["passcode"],
		CreateAt: time.Time{},
		UpdateAt: time.Time{},
	}

	db.DB.Create(&cashier)
	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier add successfully",
			"data":    cashier,
		},
	)
}

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Find(&cashier, "id=?", cashierId)

	if cashier.Name == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not foun",
			},
		)
	}

	var updateCashier models.Cashier
	err := c.BodyParser(&updateCashier)
	if err != nil {
		return err
	}

	if updateCashier.Name == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier name not required",
			},
		)
	}

	cashier.Name = updateCashier.Name

	db.DB.Save(&cashier)

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier update successfully",
			"data":    cashier,
		},
	)
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"success": false,
				"message": "Cashier not found",
			},
		)
	}

	db.DB.Where("id=?", cashierId).Delete(&cashier)

	return c.Status(200).JSON(
		fiber.Map{
			"success": true,
			"message": "Cashier deleted successfully",
		},
	)
}
