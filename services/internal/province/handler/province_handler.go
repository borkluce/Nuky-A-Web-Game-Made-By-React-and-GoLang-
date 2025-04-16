package handler

import (
	"net/http"
	"strconv"

	"your_project/internal/province/model"

	"github.com/gin-gonic/gin"
)

var provinces = []model.Province{
	{ID: 1, ProvinceName: "Berlin", ProvinceColorHex: "#FF0000"},
	{ID: 2, ProvinceName: "Paris", ProvinceColorHex: "#00FF00"},
	{ID: 3, ProvinceName: "Rome", ProvinceColorHex: "#0000FF"},
}

// Dummy reCAPTCHA and permission checker
func canAttack(c *gin.Context) bool {
	// TODO: Replace with actual recaptcha verification and user permission check
	// Example: Check session or token-based context
	return true
}

// GET /provinces
func GetAllProvinces(c *gin.Context) {
	if !canAttack(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can't perform this action"})
		return
	}
	c.JSON(http.StatusOK, provinces)
}

// POST /attack/:id
func Attack(c *gin.Context) {
	if !canAttack(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not allowed"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid province ID"})
		return
	}

	for i := range provinces {
		if provinces[i].ID == id {
			provinces[i].AttackCount++
			c.JSON(http.StatusOK, gin.H{"message": "Attack registered"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Province not found"})
}

// POST /support/:id
func Support(c *gin.Context) {
	if !canAttack(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not allowed"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid province ID"})
		return
	}

	for i := range provinces {
		if provinces[i].ID == id {
			provinces[i].SupportCount++
			c.JSON(http.StatusOK, gin.H{"message": "Support registered"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Province not found"})
}
