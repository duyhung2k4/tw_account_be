package cron_handle_impl

import (
	"account-service/config"
	cron_handle "account-service/cronHandle"
	"account-service/model"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type cronHandle struct {
	db *gorm.DB
}

func (c *cronHandle) CheckTimeSaveInfoRegister() {
	var listSaveInfo []*model.SaveRegister
	var listSaveInfoRemove []*model.SaveRegister

	errFind := c.db.Model(&model.SaveRegister{}).Find(&listSaveInfo).Error
	if errFind != nil {
		log.Println(errFind)
		return
	}

	timeNow := time.Now()
	for _, info := range listSaveInfo {
		ok := timeNow.Before(info.FinishAt)
		if ok {
			listSaveInfoRemove = append(listSaveInfoRemove, info)
		}
	}

	errDelete := c.db.Model(&model.SaveRegister{}).Delete(&listSaveInfoRemove).Error
	if errDelete != nil {
		log.Println(errDelete)
	}
}

func cronInit() cron_handle.Cron {
	return &cronHandle{
		db: config.GetDB(),
	}
}

func CronRun() {
	handleCron := cronInit()

	c := cron.New()

	c.AddFunc("@every 1m", handleCron.CheckTimeSaveInfoRegister)

	c.Start()
	log.Println("Start Cron")
}
