package services

import (
	"context"
	"fmt"
	"time"

	pb "github.com/dzoniops/common/pkg/notification"
	"github.com/dzoniops/notification-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedNotificationServiceServer
	DB mongo.Client
}

func (s *Server) CreateUserPreferences(c context.Context, req *pb.UserPreferences) (*pb.UserPreferences, error) {
	userPreferences := models.UserPreferences{
		UserId:               req.UserId,
		CreateNewReservation: true,
		CancelReservation:    true,
		RateHost:             true,
		RateAccommodation:    true,
		ReservationAnswer:    true,
	}
	_, err := s.DB.Database("notification_db").Collection("user_preferences").InsertOne(c, userPreferences)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to insert user prefences")
	}
	return req, nil
}

func (s *Server) UpdateUserPreferences(c context.Context, req *pb.UserPreferences) (*pb.UserPreferences, error) {
	coll := s.DB.Database("notification_db").Collection("user_preferences")
	userPreferences := models.UserPreferences{
		CreateNewReservation: req.CreateNewReservation,
		CancelReservation:    req.CancelReservation,
		RateHost:             req.RateHost,
		RateAccommodation:    req.RateAccommodation,
		ReservationAnswer:    req.ReservationAnswer,
	}
	filter := bson.D{{Key: "user_id", Value: req.UserId}}

	data, err := bson.Marshal(userPreferences)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to marshal request")
	}
	update := bson.D{{Key: "$set", Value: data}}
	_, err = coll.UpdateOne(c, filter, update)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to marshal request")
	}
	return req, nil
}

func (s *Server) RequestReservation(c context.Context, req *pb.RequestReservationNotification) (*pb.NotificationResponse, error) {
	coll := s.DB.Database("notification_db").Collection("request_reservation")
	message := fmt.Sprintf("You have new request for %s Accommodation", req.Accommodation)

	notification := models.Notification{
		UserId:    req.HostId,
		Message:   message,
		Status:    models.SENT,
		CreatedAt: time.Now(),
	}
	_, err := coll.InsertOne(c, notification)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to save notification")
	}
	return &pb.NotificationResponse{
		UserId:  notification.UserId,
		Message: message,
	}, nil
}

