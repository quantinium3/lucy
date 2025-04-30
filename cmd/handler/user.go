package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/quantinium3/lucy/cmd/database"
	"github.com/quantinium3/lucy/cmd/database/model"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	db *gorm.DB
}

func UserService(db *database.Database) *User {
	return &User{db: db.DB}
}

func (s *User) CreateUser(c echo.Context) error {
	user := new(model.User)
	peripheral := new(model.Peripheral)

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

	peripheral.Keypress = 0
	peripheral.LeftClick = 0
	peripheral.RightClick = 0
	peripheral.MouseTravel = 0.0
	if err = s.db.Create(&peripheral).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to create peripheral stats",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "error",
		"message": "User Create Successfully",
		"data": model.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func (s *User) GetUser(c echo.Context) error {
	userId := c.Param("id")

	var user model.User

	err := s.db.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to query database",
		})
	}

	if user.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "error",
			"message": "No user found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "user fetched successfully",
		"data": model.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func (s *User) DeleteUser(c echo.Context) error {
	userId := c.Param("id")
	var user model.User

	err := s.db.First(&user, "id = ?", userId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to query database",
		})
	}

	if user.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "error",
			"message": "User not found",
		})
	}

	err = s.db.Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "User not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "User successfully deleted",
		"data": model.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}
