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

func GetCurrentProject(c echo.Context) error {
	req, err := http.NewRequest("GET", utils.Config("WAKATIME_URI") + "/api/v1/users/current/stats/last_7_days", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to create request",
		})
	}

	req.Header.Set("Authorization", "Basic "+utils.Config("WAKATIME_APIKEY"))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to fetch current Project",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Error reading response",
		})
	}

	var result struct {
		Data struct {
			Projects []any `json:"projects"`
		} `json:"data"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to unmarshal data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "Successfully fetched recently played songs",
		"data":    result.Data.Projects[0],
	})
}

func GetOperatingSystems(c echo.Context) error {
	req, err := http.NewRequest("GET", utils.Config("WAKATIME_URI") + "/api/v1/users/current/stats/last_7_days", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to create request",
		})
	}

	req.Header.Set("Authorization", "Basic "+utils.Config("WAKATIME_APIKEY"))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to fetch current Project",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Error reading response",
		})
	}

	var result struct {
		Data struct {
			OperatingSystems []any `json:"operating_systems"`
		} `json:"data"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to unmarshal data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "Successfully fetched recently played songs",
		"data":    result.Data.OperatingSystems,
	})
}

func GetMachine(c echo.Context) error {
	req, err := http.NewRequest("GET", utils.Config("WAKATIME_URI") + "/api/v1/users/current/stats/last_7_days", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to create request",
		})
	}

	req.Header.Set("Authorization", "Basic "+utils.Config("WAKATIME_APIKEY"))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to fetch current Project",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Error reading response",
		})
	}

	var result struct {
		Data struct {
			Machines []any `json:"machines"`
		} `json:"data"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to unmarshal data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "Successfully fetched recently played songs",
		"data":    result.Data.Machines,
	})
}

func GetLanguages(c echo.Context) error {
	req, err := http.NewRequest("GET", utils.Config("WAKATIME_URI") + "/api/v1/users/current/stats/last_7_days", nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to create request",
		})
	}

	req.Header.Set("Authorization", "Basic "+utils.Config("WAKATIME_APIKEY"))
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to fetch current Project",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Error reading response",
		})
	}

	var result struct {
		Data struct {
			Languages []any `json:"languages"`
		} `json:"data"`
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "Failed to unmarshal data",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "Successfully fetched recently played songs",
		"data":    result.Data.Languages,
	})
}
