package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"slashbase.com/backend/daos"
	"slashbase.com/backend/middlewares"
	"slashbase.com/backend/models"
	"slashbase.com/backend/utils"
	"slashbase.com/backend/views"
)

type TeamController struct{}

var teamDao daos.TeamDao

func (tc TeamController) CreateTeam(c *gin.Context) {
	var createBody struct {
		Name string `json:"name"`
	}
	c.BindJSON(&createBody)
	authUser := middlewares.GetAuthUser(c)
	team := models.NewTeam(authUser, createBody.Name)
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
}

func (tc TeamController) GetTeams(c *gin.Context) {
	authUser := middlewares.GetAuthUser(c)
	teamMembers, err := teamDao.GetTeamMembersForUser(authUser.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	teamViews := []views.TeamView{}
	for _, t := range *teamMembers {
		teamViews = append(teamViews, views.BuildTeam(&t.Team))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teamViews,
	})
}

func (tc TeamController) GetTeamMembers(c *gin.Context) {
	teamID := c.Param("teamId")
	authUserTeamIds := middlewares.GetAuthUserTeamIds(c)
	if !utils.ContainsString(*authUserTeamIds, teamID) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   errors.New("not allowed"),
		})
		return
	}
	teamMembers, err := teamDao.GetTeamMembers(teamID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	teamMemberViews := []views.TeamMemberView{}
	for _, t := range *teamMembers {
		teamMemberViews = append(teamMemberViews, views.BuildTeamMember(&t))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    teamMemberViews,
	})
}
