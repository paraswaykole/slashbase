package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/daos"
	"slashbase.com/backend/middlewares"
	"slashbase.com/backend/models"
	"slashbase.com/backend/views"
)

type TeamController struct{}

var teamDao daos.TeamDao

func (tc TeamController) CreateTeam(c *gin.Context) {
	var createCmd struct {
		Name string `json:"name"`
	}
	c.BindJSON(&createCmd)
	authUser := middlewares.GetAuthUser(c)
	team := models.NewTeam(authUser, createCmd.Name)
	err := teamDao.CreateTeam(team)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    views.BuildTeam(team),
	})
	return
}

func (tc TeamController) GetTeams(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	teamViews := []views.TeamView{}
	for _, t := range authUser.Teams {
		teamViews = append(teamViews, views.BuildTeam(&t))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teamViews,
	})
	return
}
