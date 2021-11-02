package answermodel

import "VoteSth-socketgo/common"

const EntityName = "answer"

type Answer struct {
	Id         int        `json:"-" gorm:"column:id"`
	FakeId     common.UID `json:"id" gorm:"-"`
	Answer     string     `json:"answer" gorm:"column:answer"`
	Vote       int32      `json:"vote" gorm:"column:vote"`
	QuestionId int32      `json:"question_id" gorm:"column:question_id"`
}

func (Answer) TableName() string {
	return "answer"
}
func (data *Answer) Mask() {
	data.FakeId = common.NewUID(uint32(data.Id), common.DbTypeAnswer, common.ShareID)
}
