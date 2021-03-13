package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expirationTime should be 24 hours")

}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "Brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, 0 == at.UserID, "New token should not have user id associated")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring in 3 hours, should not be expired")
}
