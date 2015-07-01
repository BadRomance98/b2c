package models

import (
	"time"
)

type Media struct {
	Id         int64
	PictureUrl int
	HrefUrl    string
	SubDate    time.Time
	Status     int
	OrderBy    int64
}
