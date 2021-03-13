package accesstoken

import (
	"strings"

	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/utils/errors"
)

// Service to specify methods
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

// Repository interface
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
type service struct {
	repository Repository
}

// NewService get instance of service
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

// GetByID func
func (s *service) GetByID(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("accessTokenId not valid")
	}
	return s.repository.GetByID(accessTokenId)
}

func (s *service) Create(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
