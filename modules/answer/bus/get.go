package answerbus

import (
	"VoteSth-socketgo/common"
	answermodel "VoteSth-socketgo/modules/answer/model"
	"context"
)

type GetAnswerStore interface {
	Get(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*answermodel.Answer, error)
}

type getAnswerBus struct {
	store GetAnswerStore
}

func NewGetAnswerBus(store GetAnswerStore) *getAnswerBus {
	return &getAnswerBus{store: store}
}

func (bus *getAnswerBus) GetAnswer(ctx context.Context, id int) (*answermodel.Answer, error) {
	data, err := bus.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(answermodel.EntityName, err)
	}
	return data, err
}
