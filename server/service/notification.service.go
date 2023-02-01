package service

import "github.com/golang-angular-notification/server/model"

type NotificationService interface {
	GetNotification(id string) (*model.Notification, error)
	CreateNotification(*model.Notification) error
	GetNotifications() ([]*model.Notification, error)
	UpdateNotification(*model.Notification) error
}
