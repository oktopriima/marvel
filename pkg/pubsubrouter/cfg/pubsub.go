package cfg

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/oktopriima/marvel/pkg/pubsubrouter/session"
)

type clientPubSub struct {
	sessPubSub *pubsub.Client
	ctx        context.Context
}

const (
	MessageAttributeNameRoute = `path`
)

func NewClient(credential session.Contract) (*clientPubSub, error) {
	sess, err := pubsub.NewClient(credential.Context(), credential.GetConfig().ProjectID, credential.Option()...)
	if err != nil {
		return nil, err
	}
	return &clientPubSub{
		ctx:        credential.Context(),
		sessPubSub: sess,
	}, nil
}

func (c *clientPubSub) Client() *pubsub.Client {
	return c.sessPubSub
}
