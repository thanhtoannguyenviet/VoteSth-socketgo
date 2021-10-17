package answerstorage

import (
	"VoteSth-socketgo/common"
	answermodel "VoteSth-socketgo/modules/answer/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, answer *answermodel.Answer) error {
	db := s.db
	if err := db.Create(&answer).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
