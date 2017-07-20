package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"strings"
)

func init() {

}

type Post struct {
	Id          int64 `orm:"auto;unique;column(id)"`
	Title       string `orm:"index;size(255);column(title)"`
	Description string `orm:"null;size(255);column(description)"`
	Content     string `orm:"size(65535);column(content);type(text)"`
	IsMarkdown  int `orm:"default(0);column(is_rarkdown)"`
	Status      int `orm:"default(0);column(status)"`
	ReadNum     int64 `orm:"default(0);column(read_num)"`
	Reviews     int64 `orm:"default(0);column(reviews)"`
	PushTime    time.Time `orm:"null;column(push_time);type(datetime)"`
	CreatedAt   time.Time `orm:"null;column(created_at);type(datetime)"`
	Action      int `orm:"columnh(action)"`
}

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("mysql", "root:admin@tcp(192.168.0.103:3306)/superCong?charset=utf8&timeout=10s&collation=utf8mb4_general_ci")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//ExecPost()

	//ExecImages()

	ExecOldImages()
}

func ExecOldImages()  {
	rows, err := db.Query("SELECT id,image,createdAt FROM post WHERE image IS NOT NULL AND image != \"\"")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var image string
		var id int
		var createdAt string
		err := rows.Scan(&id, &image, &createdAt)
		if err != nil {
			panic(err.Error())
		}
		var post_id int
		err = db.QueryRow("SELECT post_id,image_name FROM images WHERE image_name = ?", image).Scan(&post_id)
		if err != nil {
			log.Println(err.Error())
			db.QueryRow("DELETE FROM images WHERE image_name = ?", image)
		}

		stmt, err := db.Prepare("INSERT INTO images(image_name, extension, image_path, real_path, image_time, image_status, image_size, md5, client_original_mame, post_id) VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}

		createdTime := strings.Replace(createdAt[0:7], "-", "/", 2)

		log.Println(createdTime)

		res, err := stmt.Exec(image, "jpeg", "uploads/images/" + createdTime, "/mnt/storage/uploads/images/" + createdTime + "/" + image, createdAt, 0, "", strings.Split(image, ".")[0], strings.Split(image, ".")[0],  id)

		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	}
}

func ExecImages() {
	rows, err := db.Query("SELECT * FROM image")
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			if string(columns[i]) == "id" {
				var id int
				db.QueryRow("SELECT id FROM images WHERE id = ?", value).Scan(&id)

				if id == 0 {
					stmt, err := db.Prepare("INSERT INTO images(id, image_name, extension, image_path, real_path, image_time, image_status, image_size, md5, client_original_mame, post_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
					if err != nil {
						log.Fatal(err)
					}
					if err != nil {
						log.Fatal(err)
					}
					postId := "0"
					if string(values[3]) != "" {
						postId = string(values[3])
					}

					res, err := stmt.Exec(string(values[0]), string(values[1]), string(values[2]), string(values[4]), string(values[5]), string(values[6]), string(values[7]), string(values[8]), string(values[9]), string(values[10]),  postId)

					if err != nil {
						log.Fatal(err)
					}
					lastId, err := res.LastInsertId()
					if err != nil {
						log.Fatal(err)
					}
					rowCnt, err := res.RowsAffected()
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
				}
				break
			}
		}
	}
}

func ExecPost() {

	var (
		id          int
		title       string
		description string
		content     string
		createdAt   string
		image       interface{}
		authorId    int
		categoryId  interface{}
		isMarkdown  interface{}
		readNum     int
		action      int
		oldId       interface{}
		source      interface{}
		status      int
		reviews     interface{}
		pushTime    interface{}
		modified    interface{}
	)

	rows, err := db.Query("SELECT id, title, description, content, createdAt, image, authorId, categoryId, isMarkdown, readNum, action, oldId, source, status, reviews, pushTime, modified FROM post")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&id,
			&title,
			&description,
			&content,
			&createdAt,
			&image,
			&authorId,
			&categoryId,
			&isMarkdown,
			&readNum,
			&action,
			&oldId,
			&source,
			&status,
			&reviews,
			&pushTime,
			&modified)
		if err != nil {
			log.Fatal(err)
		}

		var created_at string
		//post := Post{}
		db.QueryRow("SELECT created_at FROM posts WHERE id = ?", id).Scan(&created_at)
		log.Println(created_at)
		if len(created_at) == 0 {
			stmt, err := db.Prepare("INSERT INTO posts(id, title, description, content, is_rarkdown, status, read_num, reviews, push_time, created_at, action) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}

			res, err := stmt.Exec(id, title, description, content, isMarkdown, status, readNum, reviews, pushTime, createdAt, action)

			if err != nil {
				log.Fatal(err)
			}
			lastId, err := res.LastInsertId()
			if err != nil {
				log.Fatal(err)
			}
			rowCnt, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
		}
	}
}
