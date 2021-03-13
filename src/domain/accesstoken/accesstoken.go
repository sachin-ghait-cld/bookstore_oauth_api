package accesstoken

import (
	"strings"
	"time"

	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// GetNewAccessToken func
func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired receiver
func (at AccessToken) IsExpired() bool {
	return time.Now().UTC().After(time.Unix(at.Expires, 0))
}

// Validate AccessToken
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("accessTokenId not valid")
	}
	if at.UserID <= 0 {
		return errors.NewBadRequestError("user_id not valid")
	}
	if at.ClientID <= 0 {
		return errors.NewBadRequestError("client_id not valid")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}
