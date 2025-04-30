package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quantinium3/lucy/cmd/database"
	"github.com/quantinium3/lucy/cmd/database/model"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db *gorm.DB
}

func UserService(db *database.Database) *Service {
	return &Service{db: db.DB}
}

func (s *Service) CreateUser(c echo.Context) error {
	user := new(model.User)

	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to hash the password",
		})
	}

	user.Password = string(hash)

	err = s.db.Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to create user",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "error",
		"message": "User Create Successfully",
		"data":    user,
	})
}
