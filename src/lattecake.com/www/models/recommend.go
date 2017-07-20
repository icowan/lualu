package models

import (
	"time"
	"log"
)

type Recommend struct {
	Id          int64 `orm:"auto;unique;column(id)"`
	Description string `orm:"null;size(1024);column(description)"`
	Status      int `orm:"default(0);column(status)"`
	CreatedAt   time.Time `orm:"null;column(created_at);type(datetime)"`
	Action      int `orm:"column(action)"`
	Score       int `orm:"column(score);null;default(0)"`

	User *User `orm:"index;rel(fk);null;on_delete(set_null)"` // 推荐人
	Post *Post `orm:"index;rel(fk);null;on_delete(set_null)"`
}

func (c *Recommend) TableName() string {
	return "recommend"
}

func (c *Recommend) TableEngine() string {
	return "INNODB"
}

func (c *Recommend) StarPost() (Recommend, error) {
	var recommend Recommend
	err := o.Raw("SELECT * FROM recommend WHERE status = 0 AND action = 0 ORDER BY id DESC LIMIT 1").QueryRow(&recommend)
	if err != nil {
		return recommend, err
	}

	if recommend.User != nil {
		err = o.Read(recommend.User)
		if err != nil {
			log.Println(err)
		}
	}
	if recommend.Post != nil {
		err = o.Read(recommend.Post)
		if err != nil {
			log.Println(err)
		}
		_, err = o.LoadRelated(recommend.Post, "images")
		if err != nil {
			log.Println(err)
		}
	}

	return recommend, nil
}

func (c *Recommend) HomePost(action int, limit ...int) ([]Recommend, error) {
	if limit == nil {
		limit = []int{3}
	}

	var recommends []Recommend
	//var posts Posts
	_, err := o.Raw("SELECT * FROM recommend WHERE status = 0 AND action = ? ORDER BY id DESC LIMIT ?", action, limit[0]).QueryRows(&recommends)
	if err != nil {
		return recommends, err
	}

	if recommends != nil {
		for k, recommend := range recommends {
			o.Read(recommends[k].Post)
			if recommend.User != nil {
				err = o.Read(recommends[k].User)
				if err != nil {
					log.Println(err)
				}
			}

			_, err = o.LoadRelated(recommends[k].Post, "images")
			if err != nil {
				log.Println(err)
			}
		}
	}

	return recommends, nil
}
