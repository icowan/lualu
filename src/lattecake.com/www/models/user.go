package models

import (
	"time"
	"log"
)

type User struct {
	Id        int `orm:"unique;column(id)"`
	Email     string `orm:"index;size(255);column(email)"`
	Username  string `orm:"index;size(255);column(username)"`
	Password  string `orm:"size(255);column(password)"`
	CreatedAt time.Time `orm:"type(datetime)"`

	Posts []*Post `orm:"reverse(many)"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) TableEngine() string {
	return "INNODB"
}

func (u *User) TableUnique() [][]string {
	return [][]string{
		[]string{"Email", "Username"},
	}
}

func (u *User) PostsAll()  {
	user := User{}

	var err error

	err = o.Read(&user)

	num, _ := o.LoadRelated(&user, "Posts")

	log.Println(num, err)

	for _, post := range user.Posts {
		log.Println(post)
	}
}