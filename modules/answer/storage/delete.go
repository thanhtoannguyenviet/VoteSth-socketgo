package answerstorage

import (
	"VoteSth-socketgo/common"
	answermodel "VoteSth-socketgo/modules/answer/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	db := s.db
	if err := db.Delete(&answermodel.Answer{}, id).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
