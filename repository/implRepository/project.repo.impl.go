package impl_repository

import (
	"account-service/config"
	"account-service/dto/request"
	message_error "account-service/messageError"
	"account-service/model"
	"account-service/repository"
	"strconv"

	"gorm.io/gorm"
)

type projectRepository struct {
	db *gorm.DB
}

func (p *projectRepository) GetProjectById(id uint) (project *model.Project, err error) {
	var simpleProject *model.Project

	errProject := p.db.Model(&model.Project{}).Where("id = ?", id).First(&simpleProject).Error

	if errProject != nil && errProject.Error() != message_error.RECORD_NOT_FOUND {
		return nil, errProject
	}

	return simpleProject, nil
}

func (p *projectRepository) GetProjectByCreaterId(id uint) (projects []model.Project, err error) {
	var listProject []model.Project

	errProjects := p.db.Model(&model.Project{}).Where("creater_id = ?", id).Find(&listProject).Error

	return listProject, errProjects
}

func (p *projectRepository) GetProjectJoined(id uint) (projects []model.Project, err error) {
	var listProject []model.Project
	var listProjectProfile []model.ProjectProfile

	errProjectProfile := p.db.Model(&model.ProjectProfile{}).Where("profile_id = ?", id).Find(&listProjectProfile).Error

	if errProjectProfile != nil && errProjectProfile.Error() != message_error.RECORD_NOT_FOUND {
		return listProject, errProjectProfile
	}

	for _, p := range listProjectProfile {
		listProject = append(listProject, *p.Project)
	}

	return listProject, nil
}

func (p *projectRepository) CreateProject(req request.NewProjectRequest) (project *model.Project, err error) {
	var newProject = &model.Project{
		CreaterId: req.CreaterId,
		Name:      req.Name,
	}

	errCreate := p.db.Model(&model.Project{}).Create(&newProject).Error
	if errCreate != nil {
		return nil, errCreate
	}

	newProject.Code =
		strconv.Itoa(int(newProject.Id)) + "_" +
			strconv.Itoa(int(newProject.CreaterId)) + "_" +
			newProject.Name

	errSetCodeProject := p.db.Model(&model.Project{}).Where("id = ?", newProject.Id).Updates(&newProject).Error
	if errSetCodeProject != nil {
		return nil, errSetCodeProject
	}

	return newProject, nil
}

func (p *projectRepository) DeleteProject(id uint) (err error) {
	var projectDelete = &model.Project{
		Id: id,
	}

	errDelete := p.db.Model(&model.Project{}).Delete(&projectDelete).Error

	return errDelete
}

func ProjectRepositoryInit() repository.ProjectRepository {
	return &projectRepository{
		db: config.GetDB(),
	}
}
