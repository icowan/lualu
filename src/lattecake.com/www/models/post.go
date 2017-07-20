package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"time"
	"log"
)

type Posts []Post

type Post struct {
	Id          int64 `orm:"auto;unique;column(id)"`
	Title       string `orm:"index;size(255);column(title)"`
	Description string `orm:"null;size(1024);column(description)"`
	Content     string `orm:"size(65535);column(content);type(text)"`
	IsMarkdown  int `orm:"null;default(0);column(is_rarkdown)"`
	Status      int `orm:"default(0);column(status)"`
	ReadNum     int64 `orm:"default(0);column(read_num)"`
	Reviews     int64 `orm:"default(0);column(reviews)"`
	PushTime    time.Time `orm:"null;column(push_time);type(datetime)"`
	CreatedAt   time.Time `orm:"null;column(created_at);type(datetime)"`
	Action      int `orm:"column(action)"`

	User   *User `orm:"index;rel(fk);null;on_delete(set_null)"`
	Images []*Image `orm:"reverse(many)"`
}

func (p *Post) TableName() string {
	return "posts"
}

func (p *Post) TableEngine() string {
	return "INNODB"
}

func (p *Post) Find(id int64) (Post, error) {

	var post Post
	var err error
	post = Post{Id: id, Status: 1}
	err = o.Read(&post)
	if err != nil {
		return post, errors.New("post is not found")
	}

	if post.User != nil {
		err = o.Read(post.User)
		if err != nil {
			log.Println(err)
		}
	}

	_, err = o.LoadRelated(&post, "images")
	if err != nil {
		log.Println(err)
	}

	//post.ReadNum = post.ReadNum + 1
	//o.Update(post)

	return post, nil
}

func (p *Post) Popular(limit int, id ...int) (Posts, error) {
	posts := Posts{}

	var err error
	_, err = o.Raw("SELECT id, title, description, image FROM post WHERE status = 1 ORDER BY readNum DESC limit ?", limit).QueryRows(&posts)

	if err != nil {
		return posts, err
	}

	for k, _ := range posts {
		_, err = o.LoadRelated(&posts[k], "images")
		if err != nil {
			log.Println(err)
		}
	}

	return posts, nil
}

func (p *Post) LastPosts(action int, limit ...int) (Posts, error) {
	posts := Posts{}

	if limit == nil {
		limit = []int{1}
	}

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id, title, description").From("posts").Where("action = ?").OrderBy("id").Desc().Limit(limit[0])

	sql := qb.String()

	o.Raw(sql, action).QueryRows(&posts)

	for k, _ := range posts {
		_, err = o.LoadRelated(&posts[k], "images")
		if err != nil {
			log.Println(err)
		}
	}

	return posts, nil
}

func (p *Post) Update(post Post) int64 {
	if num, err := o.Update(&post); err == nil {
		return num
	}
	log.Println(err)
	return int64(0)
}