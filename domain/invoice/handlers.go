package invoice

import (
	"github.com/asentientbanana/pausalkole-admin/domain/invoice/dto"
	"github.com/asentientbanana/pausalkole-admin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func AddInvoice(c *gin.Context, db *gorm.DB) {
	var json dto.AddInvoiceDto
	if err := c.ShouldBind(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice := models.Invoice{
		ID:            uuid.New(),
		DateCompleted: 0,
		Recipient:     json.Recipient,
		Agency:        json.Agency,
		Total:         json.Total,
		DateDue:       json.DateDue,
		Description:   json.Description,
		Currency:      json.Currency,
		Status:        json.Status,
		Items:         nil,
		UserID:        "0198b33f-b562-7ecd-8390-1313263aca53",
	}

	for _, item := range json.Items {
		invoice.Items = append(invoice.Items, models.InvoiceItem{
			//InvoiceID:   invoice.ID.String(),
			Quantity:    item.Quantity,
			Metric:      item.Metric,
			Description: item.Description,
			Amount:      item.Amount,
			ID:          uuid.New(),
		})
	}

	tx := db.Create(&invoice)
	if tx.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func DeleteInvoice(c *gin.Context, db *gorm.DB, id string) {
	tx := db.Delete(&models.Invoice{}, "id = ?", id)
	if tx.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted"})
}

func UpdateInvoice(c *gin.Context, db *gorm.DB) {
	var json dto.UpdateInvoiceDto
	if err := c.ShouldBind(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if db.Model(&models.Invoice{}).Where("id = ?", json.ID).Updates(json) != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invoice update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated Invoice"})
}
