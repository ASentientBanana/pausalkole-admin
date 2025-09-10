package auth

import (
	"fmt"
	"github.com/asentientbanana/pausalkole-admin/domain/auth/dto"
	security2 "github.com/asentientbanana/pausalkole-admin/domain/security"
	dto2 "github.com/asentientbanana/pausalkole-admin/domain/user/dto"
	"github.com/asentientbanana/pausalkole-admin/errors"
	"github.com/asentientbanana/pausalkole-admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// HandleLogin godoc
// @Summary      Login
// @Description  Login existing user by email and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  dto.LoginDto  false  "User credentials"
// @Success      200  {object}  dto.UserResponseDto
// @Failure      400
// @Router       /auth/login [post]
func HandleLogin(c *gin.Context, db *gorm.DB) {
	var json dto.LoginDto
	if err := c.ShouldBind(&json); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var user models.Users

	if db.First(&user, "email = ?", json.Email).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	token, err := security2.GenerateJwtToken(user.ID.String())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.CreateInternalServerError()})
		return
	}

	c.JSON(200, dto2.UserResponseDto{
		User:  dto2.UserResponse{ID: user.ID.String(), Email: json.Email},
		Token: token,
	})
}
