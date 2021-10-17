package answerstorage

import (
	"VoteSth-socketgo/common"
	answermodel "VoteSth-socketgo/modules/answer/model"
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) Get(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*answermodel.Answer, error) {
	var rs answermodel.Answer
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Where(conditions).First(&rs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntity(questionmodel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &rs, nil
}
