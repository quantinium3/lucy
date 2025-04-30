package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/quantinium3/lucy/cmd/database"
	"github.com/quantinium3/lucy/cmd/database/model"
	"gorm.io/gorm"
)

type Peripheral struct {
	db *gorm.DB
}

func PeripheralService(db *database.Database) *Peripheral {
	return &Peripheral{db: db.DB}
}

func (s *Peripheral) GetStats(c echo.Context) error {
	userId := c.Param("id")
	var stats model.Peripheral

	err := s.db.First(&stats, "userId = ?", userId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to query database",
		})
	}

	if stats.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "error",
			"message": "Stats not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "successfully fetched stats",
		"data": model.Peripheral{
			UserId:      stats.UserId,
			Keypress:    stats.Keypress,
			LeftClick:   stats.LeftClick,
			RightClick:  stats.RightClick,
			MouseTravel: stats.MouseTravel,
		},
	})
}

func (p *Peripheral) IncrementKeyStats(c echo.Context) error {
	userId := c.Param("id")
	statsReq := new(model.Peripheral)
	if err := c.Bind(statsReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "error",
			"message": "Invalid body",
		})
	}

	var stats model.Peripheral

	err := p.db.First(&stats, "userId = ?", userId).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to query database",
		})
	}

	if stats.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "error",
			"message": "Stats not found",
		})
	}

	stats.Keypress = statsReq.Keypress
	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "successfully updated keypress",
	})
}

