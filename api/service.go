package api

import (
	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"

	"github.com/arrivets/nride-fcm/store"
)

type Service struct {
	config  Config
	handler Handler
}

func NewService(
	config Config,
	store store.IStore,
	firebaseApp *firebase.App,
) Service {
	return Service{
		config:  config,
		handler: NewAPIHandler(store, firebaseApp),
	}
}

func (s Service) Run() {
	router := gin.Default()

	router.POST("/users", s.handler.PostUser)
	router.DELETE("/users/:id", s.handler.DeleteUser)
	router.GET("/users/:id", s.handler.GetUser)

	router.POST("/notifications", s.handler.PostNotification)

	router.Run(s.config.BindAddress)
}
