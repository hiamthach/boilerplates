package middleware

import (
	"context"
	"errors"
	"go-microservices/app/config"
	authprovider "go-microservices/app/intergrate/go/utils/auth"
	cacheUtils "go-microservices/app/intergrate/go/utils/cache"
	"log"
	"net/http"
	"strings"

	"github.com/redis/go-redis/v9"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

type auth struct{}

var Auth auth

var redisDbFss *redis.Client

func init() {
	redisDbFss = redis.NewClient(&redis.Options{
		Addr:     config.Get().Redis.ServerForceSS,
		Password: config.Get().Redis.PasswordForceSS,
		DB:       config.Get().Redis.DbNumberForceSS,
	})
	if cmd := redisDbFss.Ping(context.Background()); cmd == nil && cmd.Err() != nil {
		log.Println("failed redis fss connection: ", config.Get().Redis.ServerForceSS)
		return
	}
	log.Println("succeed redis fss connection: ", config.Get().Redis.ServerForceSS)
}

// Auth Token middleware
func (a *auth) AuthToken(w http.ResponseWriter, r *http.Request) (err error) {
	if isPublicPath(r.URL.Path) {
		return nil
	}

	if isDevPath(r.URL.Path) {
		return nil
	}

	authHeader := r.Header.Get(authorizationHeaderKey)
	if len(authHeader) == 0 {
		err = errors.New("authorization header is not provided")
		return err
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		err = errors.New("authorization header is not provided")
		return err
	}

	token := fields[1]
	_, errTkn := a.ValidateToken(token)
	if errTkn != nil {
		err = errors.New("unauthenticated")
		return err
	}

	return nil
}

// ValidateToken function
func (a *auth) ValidateToken(token string) (payload *authprovider.Payload, err error) {
	// Create a token provider to verify the token with public key
	tokenProvider, err := authprovider.NewClientTokenProvider(authprovider.TokenProviderOpts{
		PublicKey: config.Get().Auth.PublicKeyHex,
	})
	if err != nil {
		return nil, err
	}

	payload, err = tokenProvider.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	if payload != nil && len(payload.SessionId) > 0 {
		var redisHelper = cacheUtils.RedisHelper{}
		redisHelper.Init(nil, redisDbFss)
		isLogout := redisHelper.IsLogoutSession(payload.SessionId)
		if isLogout {
			err = errors.New("unauthenticated")
			return nil, err
		}
	}

	return payload, nil
}

func (a *auth) GetPayload(token string) (payload *authprovider.Payload, err error) {
	tokenProvider, err := authprovider.NewClientTokenProvider(authprovider.TokenProviderOpts{
		PublicKey: config.Get().Auth.PublicKeyHex,
	})
	if err != nil {
		return nil, err
	}

	payload, err = tokenProvider.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
