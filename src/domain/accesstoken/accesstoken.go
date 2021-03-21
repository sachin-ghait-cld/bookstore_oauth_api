package accesstoken

import (
	"fmt"
	"strings"
	"time"

	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/utils/crypto_utils"
	"github.com/sachin-ghait-cld/bookstore_utils-go/rest_errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

// AccessToken struct
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// Validate AccessToken
func (at *AccessToken) Validate() rest_errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("accessTokenId not valid")
	}
	if at.UserID <= 0 {
		return rest_errors.NewBadRequestError("user_id not valid")
	}
	if at.ClientID <= 0 {
		return rest_errors.NewBadRequestError("client_id not valid")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

// AccessTokenRequest struct
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
	// Used for password grant type
	Username string `json:"username"`
	Password string `json:"password"`
	// Used for client_credentials grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate AccessTokenRequest
func (atr *AccessTokenRequest) Validate() rest_errors.RestErr {
	switch atr.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
		break
	default:
		return rest_errors.NewBadRequestError("invalid grant type")
	}
	//TODO: Validate parameters for each grant_type
	return nil
}

// GetNewAccessToken func
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired receiver
func (at AccessToken) IsExpired() bool {
	return time.Now().UTC().After(time.Unix(at.Expires, 0))
}

// Generate AccessToken
func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
