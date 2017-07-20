package models

import (
	"time"
)

type Image struct {
	Id          int64 `orm:"unique;column(id)"`
	ImageName       string `orm:"size(255);column(image_name)"`
	Extension string `orm:"null;size(32);column(extension)"`
	//PostId     int64 `orm:"column(post_id);default(0)"`
	ImagePath       string `orm:"null;size(255);column(image_path)"`
	RealPath    string `orm:"null;size(255);column(real_path)"`
	ImageTime  time.Time `orm:"null;column(image_time);type(datetime)"`
	ImageStatus  int `orm:"null;default(0);column(image_status)"`
	ImageSize  string `orm:"null;default(0);column(image_size)"`
	Md5  string `orm:"null;size(255);default(0);column(md5)"`
	ClientOriginalName  string `orm:"null;size(255);default(0);column(client_original_mame)"`

	Post *Post `orm:"index;rel(fk)"`
}

func (c *Image) TableName() string {
	return "images"
}

func (c *Image) TableEngine() string {
	return "INNODB"
}