package questionbus

import (
	"VoteSth-socketgo/common"
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
)

type ListQuestionStore interface {
	List(ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]questionmodel.Question, error)
}

type listQuestionBus struct {
	store ListQuestionStore
}

func NewListQuestionBus(store ListQuestionStore) *listQuestionBus {
	return &listQuestionBus{store: store}
}

func (bus *listQuestionBus) ListQuestion(ctx context.Context, paging *common.Paging) ([]questionmodel.Question, error) {
	result, err := bus.store.List(ctx, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(questionmodel.EntityName, err)
	}
	return result, nil
}
