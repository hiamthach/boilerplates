package authprovider

// Maker is an interface for managing tokens
type IClientAuthProvider interface {
	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
