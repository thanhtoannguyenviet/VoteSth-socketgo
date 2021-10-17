package answerstorage

import (
	"VoteSth-socketgo/common"
	answermodel "VoteSth-socketgo/modules/answer/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) Vote(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(answermodel.Answer{}.TableName()).Where("id = ?", id).
		Update("vote", gorm.Expr("vote + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
