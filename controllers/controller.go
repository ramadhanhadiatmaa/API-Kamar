package controllers

import (
	"apikamar/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var kamar []models.Kamar

	models.DB.Find(&kamar)

	return c.Status(fiber.StatusOK).JSON(kamar)

}

func Show (c *fiber.Ctx) error {

	id := c.Params("id")

	var kamar models.Kamar

	if err := models.DB.Model(&kamar).Where("id = ?", id).First(&kamar).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Kamar dengan id " + id + " tidak ditemukan",
			})
		}

			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Server sedang mengalami gangguan",
			})
	}

	return c.JSON(kamar)
}

func Create (c *fiber.Ctx) error {

	var kamar models.Kamar

	if err := c.BodyParser(&kamar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	kamar.Updated = time.Now()

	kamar.Total = 47
	kamar.Tersedia = kamar.Markusnicu + kamar.Markusvvip + kamar.Markusvip + kamar.Lukas + kamar.Maria + kamar.Fransiskus + kamar.Matius + kamar.Teresia + kamar.Teresiatiga + kamar.Yosef + kamar.Klara + kamar.Egidio + kamar.Yohanes
	kamar.Terisi = kamar.Total - kamar.Tersedia

	if err := models.DB.Create(&kamar).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data kamar berhasil dibuat",
	})
}

func Update (c *fiber.Ctx) error {
	id := c.Params("id")

	var kamar models.Kamar

	if err := c.BodyParser(&kamar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	kamar.Updated = time.Now()

	kamar.Total = 47
	kamar.Tersedia = kamar.Markusnicu + kamar.Markusvvip + kamar.Markusvip + kamar.Lukas + kamar.Maria + kamar.Fransiskus + kamar.Matius + kamar.Teresia + kamar.Teresiatiga + kamar.Yosef + kamar.Klara + kamar.Egidio + kamar.Yohanes
	kamar.Terisi = kamar.Total - kamar.Tersedia

	if models.DB.Where("id = ?", id).Updates(&kamar).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Kamar dengan id " + id + " tidak ditemukan",
		})
	} 

	return c.JSON(fiber.Map{
		"message": "Data berhasil diperbaharui",
	})
}

func Delete (c *fiber.Ctx) error {

	id := c.Params("id")

	var kamar models.Kamar

	if models.DB.Where("id = ?", id).Delete(&kamar).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Kamar dengan id " + id + " tidak dapat dihapus",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus data kamar",
	})
}

