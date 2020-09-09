package jwtauthority

import (
	"github.com/miraclew/tao/pkg/auth"
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// ErrTokenExpired indicate expired token
	ErrTokenExpired = errors.New("token is either expired or not active yet")
)

// authority encode/decode token
type authority struct {
	algorithm string // signature and hash algorithm
	secret    string // secret for signature signing and verification. can be replaced with certificate.
	expireIn  int64
}

func (a *authority) Revoke(token string) error {
	return nil
}

func (a *authority) Issue(identity *auth.Identity, duration time.Duration) (string, int64, error) {
	if duration == 0 {
		duration = time.Duration(a.expireIn) * time.Second
	}
	expiresAt := time.Now().Add(duration).Unix()
	claims := myClaims{
		DeviceID: identity.DeviceID,
		UserID:   identity.UserID,
		Roles:    identity.Roles,
		Internal: identity.Internal,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(a.algorithm), claims)
	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresAt, nil
}

func (a *authority) Verify(tokenStr string) (identity *auth.Identity, expireAt int64, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(a.secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				err = errors.New("that's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				err = ErrTokenExpired
			} else {
				err = fmt.Errorf("couldn't handle this token:%s", err.Error())
			}
		} else {
			err = fmt.Errorf("couldn't handle this token:%s", err.Error())
		}
		return
	}

	if !token.Valid {
		err = fmt.Errorf("token invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		identity = &auth.Identity{}
		if claims["device_id"] != nil {
			identity.DeviceID = claims["device_id"].(string)
		}

		var uid float64
		uid, ok = claims["user_id"].(float64)

		if claims["internal"] != nil {
			identity.Internal = claims["internal"].(string)
		} else {
			if !ok {
				err = errors.New("user_id not valid")
				return
			}
		}

		identity.UserID = int64(uid)
		var _roles []interface{}
		_roles, ok = claims["roles"].([]interface{})
		if ok {
			for _, role := range _roles {
				identity.Roles = append(identity.Roles, role.(string))
			}
		}
		expireAt = int64(claims["exp"].(float64))
	} else {
		err = errors.New("get claims failed or token invalid")
		return
	}
	return
}

// New create new codec
func New(algorithm string, secret string, expireIn int64) auth.Authority {
	return &authority{
		algorithm: algorithm,
		secret:    secret,
		expireIn:  expireIn,
	}
}

type myClaims struct {
	UserID   int64    `json:"user_id,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	DeviceID string   `json:"device_id,omitempty"`
	Internal string   `json:"internal,omitempty"`
	jwt.StandardClaims
}
