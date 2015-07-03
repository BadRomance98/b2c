package models

import (
	"time"
)

type Media struct {
	Id         int64
	PictureUrl string
	HrefUrl    string
	SubDate    time.Time
	Status     int
	OrderBy    int
}

func GetMediaCount() (int64, error) {
	media := new(Media)
	count, err := Xorm.Count(media)
	return count, err
}

func GetMediaLimit(p int) ([]Media, error) {
	media := make([]Media, 0)
	err := Xorm.Limit(10, 10*p).Find(&media)
	return media, err
}

func InsertMedia(media *Media) error {
	_, err := Xorm.Insert(media)
	return err
}

func UpdateMedia(id int64, media *Media) error {
	_, err := Xorm.Id(id).Cols("status", "order_by").Update(media)
	return err
}

func DeleteMedia(id int64) error {
	_, err := Xorm.Id(id).Delete(new(Media))
	return err
}

func SelectMedia(id int64) (*Media, error) {
	media := new(Media)
	_, err := Xorm.Id(id).Get(media)
	return media, err
}
