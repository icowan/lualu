package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"log"
)

var o orm.Ormer
var err error
func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	var mode string
	mode = beego.AppConfig.String("runmode")

	host := beego.AppConfig.String(mode + "::mysqlurls")
	pass := beego.AppConfig.String(mode + "::mysqlpass")
	user := beego.AppConfig.String(mode + "::mysqluser")
	db_name := beego.AppConfig.String(mode + "::mysqldb")
	port := beego.AppConfig.String(mode + "::mysqlport")

	log.Println(host, user, pass, db_name, port)
	//注册默认数据库
	err := orm.RegisterDataBase("default", "mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+db_name+"?charset=utf8&timeout=10s&collation=utf8mb4_general_ci", 30, 30) //密码为空格式

	if err != nil {
		log.Fatal(err)
	}

	orm.RegisterModel(new(User), new(Post), new(Image), new(Recommend))

	orm.RunSyncdb("default", false, true)

	o = orm.NewOrm()
}
