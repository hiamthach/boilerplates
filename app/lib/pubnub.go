package lib

import (
	"go-microservices/app/config"

	pubnub "github.com/pubnub/go"
)

var (
	pubnubClient *pubnub.PubNub
)

func GetPubNubClient() *pubnub.PubNub {
	if pubnubClient == nil {
		cfg := pubnub.NewConfig()
		cfg.SubscribeKey = config.Get().PubNub.SubscribeKey
		cfg.PublishKey = config.Get().PubNub.PublishKey
		cfg.SecretKey = config.Get().PubNub.SecretKey
		pubnubClient = pubnub.NewPubNub(cfg)
	}

	return pubnubClient
}
