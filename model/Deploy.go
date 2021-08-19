package model

import (
	"ginblog/utils/errmsg"
)

type Deploy struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Logo     string `gorm:"type:varchar(200)" json:"logo"`
	Abstract string `gorm:"type:varchar(200)" json:"abstract"`
	Details  string `gorm:"type:varchar(500)" json:"details"`
}

// GetDeploy 获取 博客配置
func GetDeploy() (Deploy, int) {
	var deploy Deploy
	err = db.Where("ID = ?", "1").First(&deploy).Error
	if err != nil {
		return deploy, errmsg.ERROR
	}
	return deploy, errmsg.SUCCSE
}

// UpdateProfile 更新博客配置
func UpdateDeploy(data *Deploy) int {
	var deploy Deploy
	err = db.Model(&deploy).Where("ID = ?", "1").Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
