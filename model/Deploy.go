package model

import (
	"ginblog/utils/errmsg"
)

type Deploy struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Logo string `gorm:"type:varchar(20)" json:"logo"`
	Desc string `gorm:"type:varchar(200)" json:"desc"`
}

// GetDeploy 获取 博客配置
func GetDeploy() (Profile, int) {
	var profile Profile
	err = db.Where("ID = ?", "1").First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCSE
}

// UpdateProfile 更新博客配置
func UpdateDeploy(data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", "1").Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
