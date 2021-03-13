package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/clients/cassandra"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/domain/accesstoken"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/http"
	"github.com/sachin-ghait-cld/bookstore_oauth_api/src/repository/db"
)

var (
	router = gin.Default()
)

// StartApp Starts the App
func StartApp() {
	session := cassandra.GetSession()
	session.Close()
	atHandler := http.NewHandler(accesstoken.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.CreateToken)
	router.Run(":8091")
}
