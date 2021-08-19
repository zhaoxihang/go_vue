package model

type Friends struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Portrait string `gorm:"type:varchar(20)" json:"portrait"`
	Desc     string `gorm:"type:varchar(200)" json:"desc"`
}

func getFriendsLink() {

}
