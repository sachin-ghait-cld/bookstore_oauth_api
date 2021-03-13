package db

import (
	"github.com/gocql/gocql"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/clients/cassandra"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/utils/errors"
)

const (
	queryGetAccessToken    = "Select access_token, user_id, client_id, expires from access_tokens where access_token =?;"
	queryCreateAccessToken = "Insert into access_tokens(access_token, user_id, client_id, expires) values( ?,?,?,?);"
	queryUpdateAccessToken = "Update access_tokens set expires=? where access_token=?;"
)

// NewRepository DbRepository
func NewRepository() DbRepository {
	return &dbRepository{}
}

// DbRepository interface
type DbRepository interface {
	GetByID(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

// GetByID Get token By ID
func (r *dbRepository) GetByID(id string) (*accesstoken.AccessToken, *errors.RestErr) {

	var result accesstoken.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewInternalServerError("no access token found for given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(
		queryCreateAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(
		queryUpdateAccessToken,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
