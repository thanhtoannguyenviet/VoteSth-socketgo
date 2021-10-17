package questionstorage

import (
	"VoteSth-socketgo/common"
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) Get(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*questionmodel.Question, error) {
	var rs questionmodel.Question
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
