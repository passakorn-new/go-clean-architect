package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/passakorn-new/resume/firebase"
	_informationHandler "github.com/passakorn-new/resume/information/delivery/http"
	_informationRepo "github.com/passakorn-new/resume/information/repository"
	_informationUcase "github.com/passakorn-new/resume/information/usecase"
	"github.com/passakorn-new/resume/middleware"
)

func main() {
	ctx := context.Background()
	firebase_app, _ := firebase.ConnectFirebase()
	client, err := firebase_app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	g := gin.Default()
	middL := middleware.InitMiddleware()
	g.Use(middL.CORS())
	ar := _informationRepo.NewFirebaseInformation(client)
	au := _informationUcase.NewInformationUsecase(ar)
	_informationHandler.NewInformationHandlerHandler(g, au)
	g.Run()
}
