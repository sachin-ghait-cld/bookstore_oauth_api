package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/services/access_token"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/utils/rest_errors"
)

type AccessTokenHandler interface {
	GetByID(*gin.Context)
	CreateToken(*gin.Context)
	UpdateToken(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

// NewHandler contains service
func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

// GetByID Get token By ID
func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accesstoken, err := handler.service.GetByID(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accesstoken)
}

// GetByID Get token By ID
func (handler *accessTokenHandler) CreateToken(c *gin.Context) {
	var tokenreq accesstoken.AccessTokenRequest
	if err := c.ShouldBindJSON(&tokenreq); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	accessToken, createErr := handler.service.Create(tokenreq)
	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

// GetByID Get token By ID
func (handler *accessTokenHandler) UpdateToken(c *gin.Context) {
	var token accesstoken.AccessToken
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	err := handler.service.UpdateExpirationTime(token)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, token)
}
