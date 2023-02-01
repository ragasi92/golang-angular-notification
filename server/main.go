package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang-angular-notification/server/controller"
	"github.com/golang-angular-notification/server/model"
	"github.com/golang-angular-notification/server/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var notificationController *controller.NotificationController
var users map[int64]string
var actions map[int64]string

func init() {
	fmt.Println("Starting the application...")
	ctx := context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:3306")
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("Error connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	notificationDB := mongoclient.Database("notificationdb").Collection("notifications")
	notificationService := service.NewNotificationService(ctx, notificationDB)
	notificationController = controller.New(notificationService)
	initializeDummyData()
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Add("Access-Control-Allow-Headers:", "content-type")
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200/")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}


func main() {
	go jobCreateNewNotification()
	router := mux.NewRouter()
	router.Use(CORS)
	router.HandleFunc("/notifications", notificationController.GetNotifications).Methods("GET")
	router.HandleFunc("/notifications/{id}", notificationController.UpdateNotification).Methods("PUT")
	router.HandleFunc("/notifications", notificationController.OptionHandler).Methods("OPTIONS")
	router.HandleFunc("/notifications/{id}", notificationController.UpdateNotification).Methods("OPTIONS")
	fmt.Println("----------Application ready------------")
	http.ListenAndServe(":12345", router)
}

func jobCreateNewNotification() {
	for {
		time.Sleep(1 * time.Minute)
		randUser := rand.Int63n(5)
		randAction := rand.Int63n(3)
		fmt.Println("------Creating notification----------")

		user := users[randUser]
		action := actions[randAction]
		now := time.Now()
		tweet := model.Notification{
			ID:           uuid.New().String(),
			User:         user,
			Action:       action,
			Status:       "UNREAD",
			CreationDate: &now,
		}

		notificationController.CreateNotification(&tweet)
	}
}

func initializeDummyData() {
	users = map[int64]string{
		0: "Renee P",
		1: "Fernando M",
		2: "Sofia P",
		3: "Gaby P",
		4: "Diego L",
		5: "Bill W",
	}

	actions = map[int64]string{
		0: "start following you",
		1: "liked your video",
		2: "liked your story",
		3: "liked your picture",
	}
}
