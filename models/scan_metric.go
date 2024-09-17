package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScanMetric struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    string             `json:"userId,omitempty" bson:"userId,omitempty" validate:"required"`
	DeviceID  string             `json:"deviceId,omitempty" bson:"deviceId,omitempty" validate:"required"`
	ContentID string             `json:"contentId,omitempty" bson:"contentId,omitempty" validate:"required"`
}
