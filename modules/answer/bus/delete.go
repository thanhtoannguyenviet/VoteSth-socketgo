package answerbus

import (
	"VoteSth-socketgo/common"
	answermodel "VoteSth-socketgo/modules/answer/model"
	"context"
	"errors"
)

type DeleteAnswerStore interface {
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*answermodel.Answer, error)
}

type deleteAnswerBus struct {
	store DeleteAnswerStore
}

func NewDeleteAnswerBus(store DeleteAnswerStore) *deleteAnswerBus {
	return &deleteAnswerBus{store: store}
}

func (bus *deleteAnswerBus) DeleteAnswer(ctx context.Context, id int) error {
	data, err := bus.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data != nil {
		bus.store.Delete(ctx, id)
	} else {
		return common.ErrEntityDeleted(answermodel.EntityName, errors.New("Cannot find question for delete"))
	}
	return nil
}
