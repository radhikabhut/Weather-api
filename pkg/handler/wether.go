package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
	"wether-api/pkg/redis"
	"wether-api/pkg/resty"

	"github.com/gin-gonic/gin"
)

func GetData(ctx *gin.Context) {

	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "city name is required",
		})
		return
	}

	data, err := redis.Get(city)
	if err != nil {
		slog.Error("error while getting value from redis", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "failed",
		})
		return
	}
	if data != "" {
		ctx.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    data,
		})
		return
	}

	apiURL := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=us&key=3JG24BBFW5XUF8RFAYYHTWHEW&contentType=json", city)
	res, err := resty.Get(apiURL, map[string]string{}, map[string]string{})
	if err != nil {
		slog.Error("error while getting value from resty", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "failed to fetch data from API",
		})
		return
	}

	fmt.Println(res.String())

	err = redis.Set(city, res.String(), 1*time.Hour)

	if err != nil {
		slog.Error("error while setting value to redis", "error", err.Error())
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"success": false,
			"message": "internal server error",
		})
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"success": true,
		"data":    res.String(),
	})
}
