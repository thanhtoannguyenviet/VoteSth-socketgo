package questionstorage

import (
	"VoteSth-socketgo/common"
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *questionmodel.Question) error {
	db := s.db
	if err := db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
