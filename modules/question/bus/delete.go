package questionbus

import (
	"VoteSth-socketgo/common"
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
	"errors"
)

type DeleteQuestionStore interface {
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*questionmodel.Question, error)
}

type deleteQuestionBus struct {
	store DeleteQuestionStore
}

func NewDeleteQuestionBus(store DeleteQuestionStore) *deleteQuestionBus {
	return &deleteQuestionBus{store: store}
}

func (bus *deleteQuestionBus) DeleteQuestion(ctx context.Context, id int) error {
	data, err := bus.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if data != nil {
		bus.store.Delete(ctx, id)
	} else {
		return common.ErrEntityDeleted(questionmodel.EntityName, errors.New("Cannot find question for delete"))
	}
	return nil
}
