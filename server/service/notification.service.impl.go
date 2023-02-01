package service

import (
	"context"

	"github.com/golang-angular-notification/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationServiceImpl struct {
	notificationCollection *mongo.Collection
	ctx                    context.Context
}

func NewNotificationService(ctx context.Context, collection *mongo.Collection) NotificationService {
	return &NotificationServiceImpl{
		notificationCollection: collection,
		ctx:                    ctx,
	}
}

func (noti *NotificationServiceImpl) GetNotification(id string) (*model.Notification, error) {
	var notification *model.Notification
	filter := bson.M{"id": id}
	err := noti.notificationCollection.FindOne(noti.ctx, filter).Decode(&notification)

	return notification, err
}

func (noti *NotificationServiceImpl) CreateNotification(notification *model.Notification) error {
	_, err := noti.notificationCollection.InsertOne(noti.ctx, notification)
	return err
}
func (noti *NotificationServiceImpl) GetNotifications() ([]*model.Notification, error) {
	var notifications = []*model.Notification{}
	cursor, err := noti.notificationCollection.Find(noti.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(noti.ctx) {
		var entity model.Notification
		err := cursor.Decode(&entity)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, &entity)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(noti.ctx)

	return notifications, nil
}

func (noti *NotificationServiceImpl) UpdateNotification(notification *model.Notification) error {
	filter := bson.M{"id": notification.ID}
	result := noti.notificationCollection.FindOneAndReplace(noti.ctx, filter, notification)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
