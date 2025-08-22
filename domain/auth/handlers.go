package auth

import (
	"github.com/asentientbanana/pausalkole-admin/domain/auth/dto"
	security2 "github.com/asentientbanana/pausalkole-admin/domain/security"
	"github.com/asentientbanana/pausalkole-admin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func HandleLogin(c *gin.Context, db *gorm.DB) {
	var json dto.LoginDto
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Users

	if db.First(&user, "email = ?", json.Email).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	token, err := security2.GenerateJwtToken(user.ID.String())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dto.UserResponseDto{
		User:  dto.UserResponse{ID: user.ID.String(), Email: json.Email},
		Token: token,
	})
}

func HandleRegister(c *gin.Context, db *gorm.DB) {
	var json dto.RegisterDTO

	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Problem binding",
			"req":   json,
		})
		return
	}

	if json.Password != json.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not match"})
		return
	}

	id, err := uuid.NewV7()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	encrypted, err := security2.HashPassword(json.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}

	tx := db.Create(&models.Users{
		ID:       id,
		Email:    json.Email,
		Password: encrypted,
	})

	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	token, err := security2.GenerateJwtToken(id.String())

	var userModel models.Users

	tx.First(&userModel, "id = ?", id)

	c.JSON(200,
		dto.UserResponseDto{
			User:  dto.UserResponse{ID: userModel.ID.String(), Email: userModel.Email},
			Token: token,
		})
}
