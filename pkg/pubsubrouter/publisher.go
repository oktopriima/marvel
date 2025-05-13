package pubsubrouter

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/oktopriima/marvel/pkg/config"
	"github.com/oktopriima/marvel/pkg/pubsubrouter/session"
	"time"
)

type publisher struct {
	cfg config.AppConfig
}

type Publisher interface {
	Publish(ctx context.Context, topic, path, message string) (string, error)
	PublishWithAttribute(ctx context.Context, topic, message string, attribute map[string]string) (string, error)
}

func NewPublisher(cfg config.AppConfig) Publisher {
	return &publisher{
		cfg: cfg,
	}
}

func (p *publisher) Publish(ctx context.Context, topic, path, message string) (string, error) {
	client := p.getClient(ctx)
	defer client.Close()

	t := client.Topic(topic)
	exists, err := t.Exists(ctx)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", fmt.Errorf("topic %s not found", topic)
	}

	result := client.Topic(topic).Publish(ctx, &pubsub.Message{
		Data:        []byte(message),
		PublishTime: time.Now(),
		Attributes: map[string]string{
			"path": path,
		},
	})
	id, err := result.Get(ctx)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (p *publisher) PublishWithAttribute(ctx context.Context, topic, message string, attribute map[string]string) (string, error) {
	client := p.getClient(ctx)
	defer client.Close()

	result := client.Topic(topic).Publish(ctx, &pubsub.Message{
		Data:        []byte(message),
		PublishTime: time.Now(),
		Attributes:  attribute,
	})

	id, err := result.Get(ctx)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (p *publisher) getClient(ctx context.Context) *pubsub.Client {
	cfg := p.cfg.PubSub
	sess := session.New(ctx, &session.Config{
		Type:                    cfg.Type,
		ProjectID:               cfg.ProjectId,
		PrivateKeyID:            cfg.PrivateKeyID,
		PrivateKey:              cfg.PrivateKey,
		ClientEmail:             cfg.ClientEmail,
		ClientID:                cfg.ClientID,
		AuthURI:                 cfg.AuthURI,
		TokenURI:                cfg.TokenURI,
		AuthProviderX509CertURL: cfg.AuthProviderX509CertURL,
		ClientX509CertURL:       cfg.ClientX509CertURL,
	})

	client, err := pubsub.NewClient(sess.Context(), sess.GetConfig().ProjectID, sess.Option()...)
	if err != nil {
		panic(err)
	}

	return client
}
