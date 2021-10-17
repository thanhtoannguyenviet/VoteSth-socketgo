package questionstorage

import (
	"VoteSth-socketgo/common"
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
)

func (s *sqlStore) List(ctx context.Context, conditions map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]questionmodel.Question, error) {
	var rs []questionmodel.Question

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(questionmodel.Question{}.TableName()).Where(conditions)

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase64(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := db.Order("id desc").Find(&rs).Error; err != nil {
		return nil, err
	}

	return rs, nil
}
