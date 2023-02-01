package model

import "time"

type Notification struct {
	ID           string     `json:"id,omitempty" bson:"id,omitempty"`
	User         string     `json:"user,omitempty" bson:"user,omitempty"`
	Action       string     `json:"action,omitempty" bson:"action,omitempty"`
	Status       string     `json:"status,omitempty" bson:"status,omitempty"`
	CreationDate *time.Time `json:"creationDate,omitempty" bson:"creationDate,omitempty"`
}
