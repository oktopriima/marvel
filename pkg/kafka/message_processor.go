package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/oktopriima/marvel/pkg/validates"
	"reflect"
	"time"
)

type MessageProcessorFunc func(*MessageDecoder)

// MessageProcessor contract message consumer processor
type MessageProcessor interface {
	Processor(decoder *MessageDecoder) error
}

// MessageDecoder decoder message data  on topic
type MessageDecoder struct {
	Body      []byte
	Key       []byte
	Message   string
	Error     string
	Source    *SourceData
	Topic     string
	Partition int32
	TimeStamp time.Time
	Offset    int64
	Commit    func(*MessageDecoder)
}

// Cast decode kafka message byte to struct
func (decoder *MessageDecoder) Cast(out interface{}) error {
	if reflect.TypeOf(out).Kind() != reflect.Ptr {
		return fmt.Errorf("%s", "output destination cannot addressable")
	}
	err := json.Unmarshal(decoder.Body, out)
	if err != nil {
		return err
	}
	validator := validates.New()
	return validator.Request(out)
}

// MessageEncoder message encoder  publish message to kafka
type MessageEncoder interface {
	Encode() ([]byte, error)
	Key() string
	Length() int
}
