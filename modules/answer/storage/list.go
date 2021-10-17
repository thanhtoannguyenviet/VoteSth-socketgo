package answerstorage

import (
	"VoteSth-socketgo/common"
	answermodel "VoteSth-socketgo/modules/answer/model"
	"context"
)

func (s *sqlStore) List(ctx context.Context, conditions map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]answermodel.Answer, error) {
	var rs []answermodel.Answer
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	db = db.Table(answermodel.Answer{}.TableName()).Where(conditions)
	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Order("id desc").Find(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return rs, nil
}
