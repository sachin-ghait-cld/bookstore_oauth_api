package rest

import (
	"net/http"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8090/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"the@email.com","password":"password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := userRepository{}
	user, err := repository.Login("the@email.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid response when trying to get user", err.Message)
}

func TestLoginUserInvalidUserInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8090/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"the@email.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login creds","status":"404","error":"not_found"}`,
	})

	repository := userRepository{}
	user, err := repository.Login("the@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface", err.Message)
}
func TestLoginUserInvalidCreds(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8090/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"the@email.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login creds","status":404,"error":"not_found"}`,
	})

	repository := userRepository{}
	user, err := repository.Login("the@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login creds", err.Message)
}

func TestLoginUserInvalidJsonResp(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8090/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"the@email.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": "7",
		"firstName": "user",
		"lastName": "name",
		"email": "user@name.com"}`,
	})

	repository := userRepository{}
	user, err := repository.Login("the@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid user interface", err.Message)
}
func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8090/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"the@email.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id": 7,
		"firstName": "user",
		"lastName": "name",
		"email": "user@name.com"}`,
	})

	repository := userRepository{}
	user, err := repository.Login("the@email.com", "password")

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, "user@name.com", user.Email)
	assert.EqualValues(t, "user", user.FirstName)
	assert.EqualValues(t, "name", user.LastName)
	assert.EqualValues(t, 7, user.ID)
}
