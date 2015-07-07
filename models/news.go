package models

import (
	"time"
)

type News struct {
	Id         int64
	Title      string
	Content    string
	PictureUrl string
	HrefUrl    string
	SubDate    time.Time
	Status     int
	OrderBy    int
}

func GetNewsCount() (int64, error) {
	news := new(News)
	count, err := Xorm.Count(news)
	return count, err
}

func GetNewsLimit(s int, p int) ([]News, error) {
	news := make([]News, 0)
	err := Xorm.Limit(s, s*p).Find(&news)
	return news, err
}

func InsertNews(news *News) error {
	_, err := Xorm.Insert(news)
	return err
}

func UpdateNews(id int64, news *News) error {
	_, err := Xorm.Id(id).Cols("status", "order_by").Update(news)
	return err
}

func DeleteNews(id int64) error {
	_, err := Xorm.Id(id).Delete(new(News))
	return err
}

func SelectNews(id int64) (*News, error) {
	news := new(News)
	_, err := Xorm.Id(id).Get(news)
	return news, err
}

func GetNewsHomeList() ([]News, error) {
	news := make([]News, 0)
	err := Xorm.Where("status=1 and order_by=1").Find(&news)
	return news, err
}
