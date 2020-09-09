package auth

import "time"

// Authority is the interface that issue and verify token.
type Authority interface {
	Issuer
	Verifier
	Revoker
}

type Issuer interface {
	Issue(identity *Identity, duration time.Duration) (string, int64, error)
}

type Verifier interface {
	Verify(token string) (*Identity, int64, error)
}

type Revoker interface {
	Revoke(token string) error
}
