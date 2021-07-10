package repository

import "github.com/ghadeerhamed/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)

	InsertRoomRestriction(res models.RoomRestriction) error
}
