package questionstorage

import (
	"VoteSth-socketgo/common"
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db
	if err := db.Table(questionmodel.Question{}.TableName()).
		Delete(questionmodel.Question{}, id).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
