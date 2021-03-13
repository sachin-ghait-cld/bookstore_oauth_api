package rest

import (
	"encoding/json"
	"time"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/domain/users"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/utils/errors"
)

var (
	userRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8090",
		Timeout: 100 * time.Millisecond,
	}
)

// RestUsersRepository interface
type RestUsersRepository interface {
	Login(string, string) (*users.User, *errors.RestErr)
}

// NewRepository DbRepository
func NewRepository() RestUsersRepository {
	return &userRepository{}
}

type userRepository struct {
}

// Login Get token By ID
func (r *userRepository) Login(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	resp := userRestClient.Post("/login", request)
	if resp == nil || resp.Response == nil {
		return nil, errors.NewInternalServerError("Invalid response when trying to get user")
	}
	if resp.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(resp.Bytes(), &restErr)
		if err != nil {

			return nil, errors.NewInternalServerError("Invalid error interface")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(resp.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("invalid user interface")
	}
	return &user, nil
}
