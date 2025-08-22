package entity

import (
	"github.com/asentientbanana/pausalkole-admin/domain/entity/dto"
	"github.com/asentientbanana/pausalkole-admin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func AddEntity(context *gin.Context, db *gorm.DB) {
	var json dto.AddEntityDto

	if err := context.ShouldBind(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := models.Entity{}
	entity.ID = uuid.New()
	entity.Name = json.Name

	for _, field := range json.Fields {
		entity.Fields = append(entity.Fields, models.EntityField{
			ID:        uuid.New(),
			Field:     field.Field,
			Value:     field.Value,
			IsVisible: field.IsVisible,
		})
	}

	tx := db.Create(&entity)
	if tx.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"entity": json})
}

func UpdateEntity(context *gin.Context, db *gorm.DB) {

	var json dto.UpdateEntityDto
	if err := context.ShouldBind(&json); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if db.Model(&models.Entity{}).Where("id = ?", json.ID).Updates(json).Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"entity": models.Entity{}})
}

func GetEntities(context *gin.Context, db *gorm.DB) {
	var entities []models.Entity
	tx := db.Preload("Fields").Find(&entities)
	if tx.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"entities": entities})
}

func DeleteEntity(context *gin.Context, db *gorm.DB, id string) {
	err := db.Where("id = ?", id).Unscoped().Delete(&models.Entity{}).Error
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"entity": id})
}
