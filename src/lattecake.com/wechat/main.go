package main

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"net/http"
	"io/ioutil"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beego/logs"
	"fmt"
	"reflect"
	"strconv"
	"time"
	"lattecake.com/wechat/core"
)

var config *core.Config

func init() {
	config = core.Instance()

	switch config.Env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	}
}

func main() {

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	logs.SetLogger(logs.AdapterFile, `{"filename":"`+config.LogFileName+`","daily":true}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()

	router.POST("/server/", WechatHandler)
	router.Any("/", WechatHandler)

	router.Run(":8088")
}

type TuringMessage struct {
	Code json.Number `json:"code"`
	Text string `json:"text"`
	Url  string `json:"url"`
	List []interface{} `json:"list"`
}

type WechatUser struct {
	Id        int `json:"id"`
	OpenId    string `json:"open_id"`
	NickName  string `json:"nickname"`
	Sex       string `json:"sex"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Avatar    string `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func WechatHandler(c *gin.Context) {
	config := &wechat.Config{
		AppID:          config.Wechat.AppId,
		AppSecret:      config.Wechat.AppSecret,
		Token:          config.Wechat.Token,
		EncodingAESKey: config.Wechat.EncodeAESKey,
	}
	wx := wechat.NewWechat(config)

	server := wx.GetServer(c.Request, c.Writer)

	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		text := message.NewText(msg.Content)
		var news *message.News
		msgType := message.MsgTypeText
		switch msg.MsgType {
		case message.MsgTypeEvent:
			text = message.NewText("欢迎关注 嘟嘟噜! 接下来我们一起努力吧...(*^_^*)。(偷偷告诉你，主人不在的时候可以跟 \"嘟嘟噜\" 聊天哦~_~;)")
		case message.MsgTypeText:

			msgCh := make(chan TuringMessage)

			// 没啥效果
			go func() {
				resp, err := http.Get("http://www.tuling123.com/openapi/api?key=01166bad56349c2a938687dbe7cce848&info=" + msg.Content + "&userid=" + msg.FromUserName)
				if err != nil {
					text = message.NewText(err.Error())
					return
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					text = message.NewText(err.Error())
					return
				}
				var turingMsg TuringMessage
				err = json.Unmarshal(body, &turingMsg)
				if err != nil {
					text = message.NewText(err.Error())
					return
				}
				msgCh <- turingMsg
			}()

			resp := <-msgCh
			close(msgCh)

			number, _ := resp.Code.Int64()

			switch number {
			case int64(302000):
				msgType = message.MsgTypeNews
				articles := NewArticles(resp.List, "article")
				news = message.NewNews(articles)
			case int64(308000):
				msgType = message.MsgTypeNews
				articles := NewArticles(resp.List, "news")
				news = message.NewNews(articles)
			case int64(200000):
				text = message.NewText(resp.Url)
			default:
				text = message.NewText(resp.Text)
			}
		case message.MsgTypeVoice:
			text = message.NewText("好的,我听到你的声音了!")
		case message.MsgTypeLocation:
			text = message.NewText("好的,知道你的位置啦!")
		default:
			text = message.NewText("好的,我知道了!")
		}

		var msgData interface{}

		if msgType == message.MsgTypeNews {
			msgData = news
		} else {
			msgData = text
		}

		// 没啥用

		err := SaveUser(msg, text.Content, c.Request.URL.Query().Get("signature"), c.Request.URL.Query().Get("nonce"))
		if err != nil {
			logs.Error(err)
		}

		return &message.Reply{msgType, msgData}
	})

	err := server.Serve()

	if err != nil {
		logs.Error(err)
		return
	}
	server.Send()
}

func NewArticles(list []interface{}, newsType string) []*message.Article {

	response := make([]*message.Article, len(list))

	for k, item := range list {
		itemMap := item.(map[string]interface{})
		if newsType == "article" {
			response = append(response, message.NewArticle(itemMap["article"].(string), itemMap["article"].(string), "", itemMap["detailurl"].(string)))
		} else if newsType == "news" {
			response = append(response, message.NewArticle(itemMap["name"].(string), itemMap["info"].(string), itemMap["icon"].(string), itemMap["url"].(string)))
		}
		if k == 4 {
			break
		}
	}

	return response
}

func SaveUser(msg message.MixMessage, reply string, sign string, nonce string) error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.Mysql.User, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.Database, config.Mysql.Charset))
	if err != nil {
		logs.Critical(err)
	}
	//defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(0)
	db.Ping()

	dateTime := time.Now().Format("2006-01-02 15:04:05")

	userCh := make(chan bool, 1)
	msgCh := make(chan bool, 1)

	go func() {
		columns := []string{
			"id",
			"open_id",
			"nickname",
			"sex",
			"province",
			"city",
			"country",
			"avatar",
			"created_at",
			"updated_at",
		}

		var wechatUser WechatUser

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for j := range values {
			scanArgs[j] = &values[j]
		}

		err = db.QueryRow("SELECT id FROM wechat_users WHERE open_id = ?", msg.ToUserName).Scan(scanArgs...)

		if err != nil {
			logs.Error(err, "select err")
			userCh <- false
			return
		}

		record := make(map[string]interface{})

		for i, col := range values {
			if col != nil {
				if reflect.TypeOf(col).String() != "[]uint8" {
					record[columns[i]] = col
				} else {
					if columns[i] == "id" {
						idVal, _ := strconv.Atoi(string(col.([]byte)))
						record[columns[i]] = idVal
					} else {
						record[columns[i]] = string(col.([]byte))
					}

				}
			}
		}

		if record == nil {
			userCh <- false
			return
		}

		jsonData, _ := json.Marshal(record)

		json.Unmarshal(jsonData, &wechatUser)

		if wechatUser.Id < 1 {
			stmt, err := db.Prepare("INSERT INTO wechat_users(`open_id`, `created_at`, `updated_at`) VALUES (?, ?, ?)")
			if err != nil {
				logs.Critical(err)
				userCh <- false
				return
			}
			defer stmt.Close()
			_, err = stmt.Exec(msg.ToUserName, dateTime, dateTime)
			if err != nil {
				logs.Error(err, "wechat_users")
				userCh <- false
				return
			}
		}
		userCh <- true
	}()

	go func() {
		stmt, err := db.Prepare("INSERT INTO messages(`nonce`, `to_username`, `from_username`, `msg_type`, `signature`, `msg_id`, `reply`, `content`, `pic_url`, `media_id`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			logs.Critical(err)
			msgCh <- false
			return
		}
		defer stmt.Close()

		res, err := stmt.Exec(nonce, msg.ToUserName, msg.FromUserName, string(msg.MsgType), sign, msg.MsgID, reply, msg.Content, msg.PicURL, msg.MediaID, dateTime, dateTime)
		if err != nil {
			logs.Error(err, "messages")
			msgCh <- false
			return
		}

		logs.Info(res)
		msgCh <- true
	}()

	<-userCh
	<-msgCh
	close(userCh)
	close(msgCh)
	return err
}
