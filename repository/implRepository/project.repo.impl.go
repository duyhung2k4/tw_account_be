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

func (p *projectRepository) GetProjectCreaterById(id uint, credentialId uint) (project *model.Project, err error) {
	var simpleProject *model.Project

	errProject := p.db.
		Model(&model.Project{}).
		Where("id = ? AND creater_id = ?", id, credentialId).
		First(&simpleProject).Error

	if errProject != nil {
		return nil, errProject
	}

	return simpleProject, nil
}

func (p *projectRepository) GetProjectByCreaterId(id uint) (projects []model.Project, err error) {
	var listProject []model.Project

	errProjects := p.db.
		Model(&model.Project{}).
		Preload("Creater").
		Where("creater_id = ?", id).Find(&listProject).Error

	return listProject, errProjects
}

func (p *projectRepository) GetProjectJoined(id uint) (projects []model.Project, err error) {
	var listProject []model.Project
	var listProjectProfile []model.ProjectProfile

	errProjectProfile := p.db.
		Model(&model.ProjectProfile{}).
		Preload("Project").
		Preload("Project.Creater").
		Where("profile_id = ?", id).Find(&listProjectProfile).Error

	if errProjectProfile != nil && errProjectProfile.Error() != message_error.RECORD_NOT_FOUND {
		return listProject, errProjectProfile
	}

	for _, p := range listProjectProfile {
		listProject = append(listProject, *p.Project)
	}

	return listProject, nil
}

func (p *projectRepository) GetProjectJoinedById(id uint, credentialId uint) (project *model.Project, err error) {
	var simpleProjectProfile *model.ProjectProfile
	var simpleProject *model.Project

	errProject := p.db.
		Model(&model.ProjectProfile{}).
		Preload("Project").
		Where("project_id = ? AND profile_id = ?", id, credentialId).
		First(&simpleProjectProfile).Error
	if errProject != nil {
		return nil, errProject
	}

	simpleProject = simpleProjectProfile.Project
	return simpleProject, nil
}

func (p *projectRepository) CreateProject(req request.NewProjectRequest) (project *model.Project, err error) {
	var newProject = &model.Project{
		CreaterId: req.CreaterId,
		Name:      req.Name,
		Code:      "",
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

func (p *projectRepository) DeleteProject(req request.DeleteProjectRequest) (err error) {
	var projectDelete = &model.Project{
		Id:        req.Id,
		CreaterId: req.CreaterId,
	}

	errDelete := p.db.Model(&model.Project{}).Delete(&projectDelete).Error

	return errDelete
}

func (p *projectRepository) CheckProjectOfCredential(credentialId uint, projectId uint) (ok bool, err error) {
	var project *model.Project

	errProject := p.db.Model(&model.Project{}).Where("id = ? AND creater_id = ?", projectId, credentialId).First(&project).Error
	if errProject != nil && errProject.Error() != message_error.RECORD_NOT_FOUND {
		return false, errProject
	}

	if project.Id == 0 || project == nil {
		return false, nil
	}

	return true, nil
}

func ProjectRepositoryInit() repository.ProjectRepository {
	return &projectRepository{
		db: config.GetDB(),
	}
}
