package model

import "time"

type Actors struct {
	Id         int       `db:"id"`
	Name       string    `db:"name"`
	Created_At time.Time `db:"created_at"`
	Updated_At time.Time `db:"updated_at"`
}
