package impl_repository

import (
	"account-service/config"
	"account-service/dto/request"
	"account-service/model"
	"account-service/repository"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func (t *taskRepository) GetTaskOfProject(projectId uint) (task []model.Task, err error) {
	var listTask []model.Task

	errListTask := t.db.
		Model(&model.Task{}).
		Preload("Creater").
		Preload("Project").
		Where("project_id = ?", projectId).
		Find(&listTask).Error

	return listTask, errListTask
}

func (t *taskRepository) CreateTask(task request.CreateTaskRequest) (newTask *model.Task, err error) {
	var createTask = &model.Task{
		CreaterId: task.CreaterId,
		ProjectId: task.ProjectId,
		Name:      task.Name,
		Level:     task.Level,
		Detail:    task.Detail,
		StartAt:   task.StartAt,
		FinishAt:  task.FinishAt,
		Status:    model.OPEN,
	}

	errCreate := t.db.Model(&model.Task{}).Create(&createTask).Error
	return createTask, errCreate
}

func (t *taskRepository) DeleteTask(taskId uint) (err error) {
	taskDelete := &model.Task{
		Id: taskId,
	}

	errDelete := t.db.Model(&model.Task{}).Delete(&taskDelete).Error
	return errDelete
}

func (t *taskRepository) UpdateStatusTask(req model.Task) (newTask *model.Task, err error) {
	taskUpdate := &req

	errUpdate := t.db.Model(&model.Task{}).Where("id = ?", req.Id).Updates(&taskUpdate).Error
	return taskUpdate, errUpdate
}

func (t *taskRepository) AddUserToTask(credentialId uint, taskId uint) (taskProfile *model.TaskProfile, err error) {
	newTaskProfile := &model.TaskProfile{
		TaskId:        taskId,
		ImplementerId: credentialId,
	}

	errCreate := t.db.Model(&model.TaskProfile{}).Create(&newTaskProfile).Error
	return newTaskProfile, errCreate
}

func (t *taskRepository) RemoveUserToTask(credentialId uint, taskId uint) (err error) {
	errDelete := t.db.
		Model(&model.TaskProfile{}).
		Where("task_id = ? AND implementer_id = ?", taskId, credentialId).
		Delete(&model.TaskProfile{}).Error
	return errDelete
}

func (t *taskRepository) GetAllUserOfTask() (taskProfiles []model.TaskProfile, err error) {
	var data []model.TaskProfile

	errData := t.db.Model(&model.TaskProfile{}).Find(&data).Error
	return data, errData
}

func TaskRepositoryInit() repository.TaskRepository {
	return &taskRepository{
		db: config.GetDB(),
	}
}
