package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/internal/controllers"
)

type SettingRoutes struct{}

var settingController controllers.SettingController

func (sr SettingRoutes) GetSingleSetting(c *gin.Context) {

	name := c.Query("name")

	value, err := settingController.GetSingleSetting(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    value,
	})
}

func (sr SettingRoutes) UpdateSingleSetting(c *gin.Context) {

	var reqBody struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	c.BindJSON(&reqBody)

	err := settingController.UpdateSingleSetting(reqBody.Name, reqBody.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
