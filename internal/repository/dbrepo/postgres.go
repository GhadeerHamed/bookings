package dbrepo

import (
	"context"
	"github.com/ghadeerhamed/bookings/internal/models"
	"time"
)

func (m *postgresRepo) AllUsers() bool {
	return true
}

//InsertReservation inserts a reservation into a database
func (m *postgresRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	stmt := `INSERT INTO reservations
				(first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
				VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

//InsertRoomRestriction inserts a room restriction into the database
func (m *postgresRepo) InsertRoomRestriction(room_es models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO room_restrictions
				(start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
				VALUES ($1,$2,$3,$4,$5,$6,$7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		room_es.StartDate,
		room_es.EndDate,
		room_es.RoomID,
		room_es.ReservationID,
		room_es.RestrictionID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
