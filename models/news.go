package models

import (
	"strconv"
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
	OrderBy    int64
}

func GetNewsCount() (int64, error) {
	news := new(News)
	count, err := Xorm.Count(news)
	return count, err
}

func GetNewsLimit(p int) ([]News, error) {
	news := make([]News, 0)
	err := Xorm.Limit(10, 10*p).Find(&news)
	return news, err
}

func SaveOrUpdateNews(id string, title string, content string, hrefurl string, fileName string) error {
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		news := new(News)
		news.Title = title
		news.Content = content
		news.HrefUrl = hrefurl
		news.PictureUrl = fileName
		news.SubDate = time.Now()
		news.Status = 1

		err = InsertNews(news)
		return err
	}

	news := new(News)
	has, err := Xorm.Id(intId).Get(news)
	if err != nil {
		return err
	}
	if has {
		news.Title = title
		news.Content = content
		news.HrefUrl = hrefurl
		news.PictureUrl = fileName
		err = UpdateNews(intId, news)
		return err
	} else {
		return nil
	}

}

func InsertNews(news *News) error {
	_, err := Xorm.Insert(news)
	return err
}

func UpdateNews(id int64, news *News) error {
	_, err := Xorm.Id(id).Update(news)
	return err
}

func DeleteNews(id int64) error {
	_, err := Xorm.Id(id).Delete(new(User))
	return err
}
