package lib

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type IGPubSubHelper interface {
	NewClient() (*pubsub.Client, error)
	PubProto(topicID string, proto protoreflect.ProtoMessage) (string, error)
	PullProto(subID string, beforeAck func(room []byte) bool, afterHandle func(room []byte) bool) error
}

type GPubSubOpts struct {
	CertPath  string
	ProjectId string
}

type GPubSubHelper struct {
	IGPubSubHelper
	opts *GPubSubOpts
}

func (gps *GPubSubHelper) Init(opts *GPubSubOpts) {
	gps.opts = opts
}

func GetGPubSub(opts *GPubSubOpts) IGPubSubHelper {
	gps := &GPubSubHelper{}
	gps.Init(opts)
	return gps
}

func (gps *GPubSubHelper) PubProto(topicID string, state protoreflect.ProtoMessage) (string, error) {

	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%+v", errors.New(fmt.Sprintf("%v", r))))
		}
	}()

	ctx := context.Background()
	client, err := gps.NewClient()
	if err != nil {
		panic(err)
	}

	defer client.Close()
	t := client.Topic(topicID)
	msg, err := proto.Marshal(state)
	if err != nil {
		panic(err)
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: msg,
	})
	//retrun id publish
	return result.Get(ctx)
}

func (gps *GPubSubHelper) PullProto(subID string, beforeAck func(room []byte) bool, afterHandle func(room []byte) bool) error {

	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%+v", errors.New(fmt.Sprintf("Pull Proto has error :%v", r))))
		}
	}()

	client, err := gps.NewClient()
	if err != nil {
		panic(err)
	}

	defer client.Close()
	sub := client.Subscription(subID)
	sub.ReceiveSettings.Synchronous = false
	sub.ReceiveSettings.NumGoroutines = 16

	cctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	var mu sync.Mutex
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {

		mu.Lock()
		defer mu.Unlock()
		defer func() {
			if r := recover(); r != nil {
				fmt.Errorf("%+v", errors.New(fmt.Sprintf("Receive proto has error :%v", r)))
			}
		}()

		if beforeAck != nil {
			result := beforeAck(msg.Data)
			if result == true {
				msg.Ack()
			} else {
				msg.Nack()
			}
		}
		if afterHandle != nil {
			afterHandle(msg.Data)
		}
	})
	if err != nil {
		panic(err)
	}
	return nil
}
