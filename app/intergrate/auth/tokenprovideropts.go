package authprovider

type TokenProviderOpts struct {
	PublicKey                    string
	PrivateKey                   string
	AccessTokenDurationInMinute  int64
	RefreshTokenDurationInMinute int64
	Issuer                       string
}
