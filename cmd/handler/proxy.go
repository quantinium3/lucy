package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quantinium3/lucy/cmd/utils"
)

func GetLastFMTracks(c echo.Context) error {
	resp, err := http.Get(fmt.Sprintf("%s?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=20", utils.Config("LASTFM_URI"), utils.Config("LASTFM_USERNAME"), utils.Config("LASTFM_APIKEY")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to fetch recently played songs",
		})
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to fetch recently played songs",
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return c.JSON(resp.StatusCode, map[string]any{
			"status":  "error",
			"message": "Failed to fetch recently played songs",
		})
	}

	var result struct {
		RecentTracks struct {
			Track []any `json:"track"`
		} `json:"recenttracks"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to parse response",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "Successfully fetched recently played songs",
		"data":    result.RecentTracks.Track,
	})
}
