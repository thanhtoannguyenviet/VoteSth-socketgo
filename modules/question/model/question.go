package questionmodel

import "VoteSth-socketgo/common"

const EntityName = "question"

type Question struct {
	Id       int        `json:"-" gorm:"column:id"`
	FakeId   common.UID `json:"id" gorm:"-"`
	Question string     `json:"question" gorm:"column:question"`
}

func (Question) TableName() string {
	return "question"
}
func (data *Question) Mask() {
	data.FakeId = common.NewUID(uint32(data.Id), common.DbTypeQuestion, common.ShareID)
}
