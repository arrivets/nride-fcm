package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"

	"github.com/arrivets/nride-fcm/api"
	"github.com/arrivets/nride-fcm/store"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	gac := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if gac == "" {
		panic("Environment variable GOOGLE_APPLICATION_CREDENTIALS not set")
	}
}

func main() {
	store := store.NewInmemStore()

	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	apiService := api.NewService(
		api.Config{
			BindAddress: "0.0.0.0:8081",
		},
		store,
		firebaseApp,
	)

	apiService.Run()
}
