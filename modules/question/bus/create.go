package questionbus

import (
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
)

type CreateQuestionStore interface {
	Create(ctx context.Context, data *questionmodel.Question) error
}

type createQuestionBus struct {
	store CreateQuestionStore
}

func NewCreateQuestionBus(store CreateQuestionStore) *createQuestionBus {
	return &createQuestionBus{store: store}
}

func (bus *createQuestionBus) CreateQuestion(ctx context.Context, data *questionmodel.Question) error {
	err := bus.store.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
