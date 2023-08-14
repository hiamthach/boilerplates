package authprovider

import (
	"encoding/json"
	"errors"

	"aidanwoods.dev/go-paseto"
)

type ClientTokenProvider struct {
	IClientAuthProvider
	config                TokenProviderOpts
	v4SymmetricKey        paseto.V4SymmetricKey
	v4AsymmetricSecretKey paseto.V4AsymmetricSecretKey
	v4AsymmetricPublicKey paseto.V4AsymmetricPublicKey
}

// New Token Maker
func NewClientTokenProvider(config TokenProviderOpts) (IClientAuthProvider, error) {
	publicAsymmetricHexKey := config.PublicKey

	if len(publicAsymmetricHexKey) != 64 {
		return nil, errors.New("invalid symmetricKey hex key size: must be exactly 64 characters")
	}

	asymmetricPublicKey, errNewPublicKey := paseto.NewV4AsymmetricPublicKeyFromHex(publicAsymmetricHexKey)
	if errNewPublicKey != nil {
		return nil, errors.New("new secret key from hex error, " + errNewPublicKey.Error())
	}

	maker := &ClientTokenProvider{
		config:                config,
		v4AsymmetricPublicKey: asymmetricPublicKey,
	}

	return maker, nil
}

// Verify Token and return Payload
func (authFeat *ClientTokenProvider) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	parser := paseto.NewParserForValidNow()
	tokenFromSign, err := parser.ParseV4Public(authFeat.v4AsymmetricPublicKey, token, nil)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(tokenFromSign.ClaimsJSON(), &payload); err != nil {
		return nil, err
	}

	return payload, nil
}
