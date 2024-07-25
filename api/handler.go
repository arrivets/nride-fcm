package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"

	"github.com/gin-gonic/gin"

	"github.com/arrivets/nride-fcm/store"
)

type Handler struct {
	store       store.IStore
	firebaseApp *firebase.App
}

func NewAPIHandler(store store.IStore, firebaseApp *firebase.App) Handler {
	return Handler{
		store:       store,
		firebaseApp: firebaseApp,
	}
}

func (h Handler) PostUser(c *gin.Context) {
	var request SubscribeUserRequest

	if err := c.BindJSON(&request); err != nil {
		return
	}

	user := h.store.AddUser(request.ID, request.Token)

	log.Printf("@@@ Added user %v\n", request.ID)

	c.IndentedJSON(http.StatusCreated, user)
}

func (h Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	h.store.DeleteUser(userID)

	c.IndentedJSON(http.StatusOK, userID)
}

func (h Handler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user := h.store.GetUser(userID)
	if user == nil {
		c.IndentedJSON(http.StatusNotFound, userID)
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func (h Handler) PostNotification(c *gin.Context) {
	var notification Notification

	if err := c.BindJSON(&notification); err != nil {
		log.Printf("@@@ ERROR parsing notification: %v\n", err)
		return
	}

	user := h.store.GetUser(notification.DestinationID)
	if user == nil {
		log.Printf("@@@ ERROR user not found. ID: %v\n", notification.DestinationID)
		c.IndentedJSON(http.StatusNotFound, notification.DestinationID)
		return
	}

	err := h.sendToToken(user.Token, notification.Title, notification.Body)
	if err != nil {
		log.Printf("@@@ ERROR sending fcm request: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, notification)
}

func (h *Handler) sendToToken(token string, title string, body string) error {
	ctx := context.Background()
	client, err := h.firebaseApp.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("error getting Messaging client: %v", err)
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	if _, err := client.Send(ctx, message); err != nil {
		return err
	}

	return nil
}
