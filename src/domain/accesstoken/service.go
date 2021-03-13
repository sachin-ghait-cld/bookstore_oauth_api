package accesstoken

import (
	"strings"

	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/repository/db"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/repository/rest"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/utils/rest_errors"
)

// Service to specify methods
type Service interface {
	GetByID(string) (*AccessToken, *rest_errors.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *rest_errors.RestErr)
	UpdateExpirationTime(AccessToken) *rest_errors.RestErr
}

// Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *rest_errors.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *rest_errors.RestErr)
	UpdateExpirationTime(AccessToken) *rest_errors.RestErr
}
type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

// NewService get instance of service
func NewService(restRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: restRepo,
		dbRepo:        dbRepo,
	}
}

// GetByID func
func (s *service) GetByID(accessTokenID string) (*AccessToken, *rest_errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, rest_errors.NewBadRequestError("accessTokenId not valid")
	}
	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request AccessTokenRequest) (*AccessToken, *rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := GetNewAccessToken(user.ID)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
