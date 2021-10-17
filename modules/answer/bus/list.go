package answerbus

import (
	"VoteSth-socketgo/common"
	answermodel "VoteSth-socketgo/modules/answer/model"
	"context"
)

type ListAnswerStore interface {
	List(ctx context.Context, conditions map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]answermodel.Answer, error)
}

type listAnswerBus struct {
	store ListAnswerStore
}

func NewListAnswerBus(store ListAnswerStore) *listAnswerBus {
	return &listAnswerBus{store: store}
}

func (bus *listAnswerBus) ListAnswer(ctx context.Context, paging *common.Paging) ([]answermodel.Answer, error) {
	result, err := bus.store.List(ctx, nil, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(answermodel.EntityName, err)
	}
	return result, err
}
