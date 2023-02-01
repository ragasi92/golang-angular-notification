package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-angular-notification/server/model"
	"github.com/golang-angular-notification/server/service"
	"github.com/gorilla/mux"
)

type NotificationController struct {
	notificationService service.NotificationService
}

func New(service service.NotificationService) *NotificationController {
	return &NotificationController{
		notificationService: service,
	}
}

func (nc *NotificationController) CreateNotification(newNotification *model.Notification) error {

	err := nc.notificationService.CreateNotification(newNotification)
	if err != nil {
		return err
	}

	return nil
}

func (nc *NotificationController) GetNotifications(response http.ResponseWriter, request *http.Request) {
	fmt.Println("----------Get notifications------------")
	response.Header().Set("content-type", "application/json")
	var notifications []*model.Notification

	notifications, err := nc.notificationService.GetNotifications()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(response).Encode(notifications)
}

func (nc *NotificationController) UpdateNotification(response http.ResponseWriter, request *http.Request) {
	fmt.Println("----------Update notification------------")
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	notificationID := params["id"]

	notification, err := nc.notificationService.GetNotification(notificationID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	notification.Status = "READ"
	if err := nc.notificationService.UpdateNotification(notification); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(notification)
}

func (nc *NotificationController) OptionHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("----------OPTIONS handler------------")
	response.Header().Set("content-type", "application/json")
	json.NewEncoder(response)
}
