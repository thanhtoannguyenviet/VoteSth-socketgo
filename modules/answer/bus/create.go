package answerbus

import (
	answermodel "VoteSth-socketgo/modules/answer/model"
	"context"
)

type CreateAnswerStore interface {
	Create(ctx context.Context, data *answermodel.Answer) error
}

type createAnswerBus struct {
	store CreateAnswerStore
}

func NewCreateAnswerBus(store CreateAnswerStore) *createAnswerBus {
	return &createAnswerBus{store: store}
}

func (bus *createAnswerBus) CreateAnswer(ctx context.Context, data *answermodel.Answer) error {
	err := bus.store.Create(ctx, data)
	return err
}
