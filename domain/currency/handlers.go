package currency

import (
	"github.com/asentientbanana/pausalkole-admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetCurrencies(c *gin.Context, db *gorm.DB) {
	var currencies []models.InvoiceCurrencies

	if db.Find(&currencies).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Problem getting currencies",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"currencies": currencies,
	})
}
