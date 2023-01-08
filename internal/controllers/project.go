package controllers

import (
	"errors"

	"github.com/slashbaseide/slashbase/internal/dao"
	"github.com/slashbaseide/slashbase/internal/models"
)

type ProjectController struct{}

func (ProjectController) CreateProject(projectName string) (*models.Project, error) {
	var project *models.Project

	if len(strings.TrimSpace(projectName)) > 0 {
		fmt.Println("testing")
		project = models.NewProject(projectName)
		err := dao.Project.CreateProject(project)
		if err != nil {
			return nil, errors.New("there was some problem")
		}
	} else {
		return nil, errors.New("Project name cannot be empty.")
	}

	return project, nil
}

func (ProjectController) GetProjects() (*[]models.Project, error) {

	projects, err := dao.Project.GetAllProjects()
	if err != nil {
		return nil, errors.New("there was some problem")
	}
	return projects, nil
}

func (ProjectController) DeleteProject(id string) error {

	project, err := dao.Project.GetProject(id)
	if err != nil {
		return errors.New("could not find project")
	}

	allDBsInProject, err := dao.DBConnection.GetDBConnectionsByProject(project.ID)
	if err != nil {
		return errors.New("there was some problem")
	}

	for _, dbConn := range allDBsInProject {
		err := dao.DBConnection.DeleteDBConnectionById(dbConn.ID)
		if err != nil {
			return errors.New("there was some problem deleting db `" + dbConn.Name + "` in the project")
		}
	}

	err = dao.Project.DeleteProject(project.ID)
	if err != nil {
		return errors.New("there was some problem deleting the project")
	}

	return nil
}
