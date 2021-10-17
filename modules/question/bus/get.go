package questionbus

import (
	"VoteSth-socketgo/common"
	questionmodel "VoteSth-socketgo/modules/question/model"
	"context"
)

type GetQuestionStore interface {
	Get(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*questionmodel.Question, error)
}

type getQuestionBus struct {
	store GetQuestionStore
}

func NewGetQuestionBus(store GetQuestionStore) *getQuestionBus {
	return &getQuestionBus{store: store}
}

func (bus *getQuestionBus) GetQuestion(ctx context.Context, id int) (*questionmodel.Question, error) {
	data, err := bus.store.Get(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(questionmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(questionmodel.EntityName, err)
	}
	return data, err
}
