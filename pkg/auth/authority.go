package auth

import "time"

type Authority interface {
	NewToken(identity Identity, duration time.Duration) (string, int64, error)
	VerifyToken(token string) (identity *Identity, expireAt int64, err error)
}
