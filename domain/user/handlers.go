package user

import (
	"github.com/asentientbanana/pausalkole-admin/errors"
	"net/http"

	security2 "github.com/asentientbanana/pausalkole-admin/domain/security"
	dto2 "github.com/asentientbanana/pausalkole-admin/domain/user/dto"
	"github.com/asentientbanana/pausalkole-admin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// HandleRegister godoc
// @Summary      Register
// @Description  Register user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        credentials  body  dto.RegisterDTO  false  "User credentials"
// @Success      200  {object}  dto.UserResponseDto
// @Failure      400
// @Router       /user/register [post]
func HandleRegister(c *gin.Context, db *gorm.DB) {
	var json dto2.RegisterDTO

	if c.ShouldBind(&json) != nil {
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
		c.JSON(http.StatusInternalServerError, errors.CreateInternalServerError())
		return
	}

	encrypted, err := security2.HashPassword(json.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.CreateInternalServerError())
	}

	tx := db.Create(&models.Users{
		ID:       id,
		Email:    json.Email,
		Password: encrypted,
	})

	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, errors.CreateInternalServerError())
		return
	}

	token, err := security2.GenerateJwtToken(id.String())

	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.CreateInternalServerError())
	}

	var userModel models.Users

	if tx.First(&userModel, "id = ?", id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusAccepted,
		dto2.UserResponseDto{
			User:  dto2.UserResponse{ID: userModel.ID.String(), Email: userModel.Email},
			Token: token,
		})
}
