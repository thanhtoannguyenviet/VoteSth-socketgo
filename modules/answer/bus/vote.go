package answerbus

import "context"

type VoteAnswerStore interface {
	Vote(ctx context.Context, id int) error
}
type voteAnswerBus struct {
	store VoteAnswerStore
}

func NewVoteAnswerBus(store VoteAnswerStore) *voteAnswerBus {
	return &voteAnswerBus{store: store}
}

func (bus *voteAnswerBus) VoteAnswer(ctx context.Context, id int) error {
	err := bus.store.Vote(ctx, id)
	return err
}
