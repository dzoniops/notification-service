package models



type UserPreferences struct {
	UserId               int64 `bson:"user_id"`
	CreateNewReservation bool  `bson:"create_new_reservation"`
	CancelReservation    bool  `bson:"cancel_reservation"`
	RateHost             bool  `bson:"rate_host"`
	RateAccommodation    bool  `bson:"rate_accommodation"`
	ReservationAnswer    bool  `bson:"reservation_answer"`
}
