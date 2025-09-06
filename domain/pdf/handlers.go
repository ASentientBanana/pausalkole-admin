package pdf

import (
	"github.com/asentientbanana/pausalkole-admin/domain/invoice"
	"github.com/asentientbanana/pausalkole-admin/domain/pdf/templates"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetInvoicePdf(context *gin.Context, db *gorm.DB, id string) {
	in, err := invoice.GetCompleteInvoiceByID(db, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	templates.GenerateDefaultInvoicePdf(in)
	context.JSON(http.StatusOK, gin.H{"invoice": in})
}
