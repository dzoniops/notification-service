package models

import "time"

type Notification struct {
	UserId    int64              `bson:"user_id"`
	Message   string             `bson:"message"`
	Status    NotificationStatus `bson:"status"`
	CreatedAt time.Time          `bson:"created_at"`
}

type NotificationStatus int32

const (
	UNSPECIFIED NotificationStatus = 0
	SENT        NotificationStatus = 1
	READ        NotificationStatus = 2
)
