package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//User 用户表
type User struct {
	Id       int
	Name     string
	PassWord string
	About    string `orm:"default(' ')"`
}

//Article 文章表
type Article struct {
	Id             int
	ArticleName    string
	ArticleContent *ArticleContent `orm:"rel(one)"`
	ArticleType    *ArticleType    `orm:"rel(fk)"`
	CreateTime     time.Time       `orm:"auto_now_add;type(datetime)"`
}

//ArticleContent 文章内容表
type ArticleContent struct {
	Id      int
	Article *Article `orm:"reverse(one);"`
	Content string   `orm:"type(text);default(' ')"`
}

//ArticleType 文章类型表
type ArticleType struct {
	Id       int
	TypeName string
	Article  []*Article `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(User), new(Article), new(ArticleContent), new(ArticleType))
	sqlname, _ := beego.GetConfig("String", "sqlname", "myblogbygo")
	dbname, _ := beego.GetConfig("String", "dbname", "mysql")
	sqluser, _ := beego.GetConfig("String", "sqluser", "root")
	sqlpass, _ := beego.GetConfig("String", "sqlpass", "123")
	sqlhost, _ := beego.GetConfig("String", "sqlhost", "127.0.0.1")
	sqlport, _ := beego.GetConfig("String", "sqlport", "3306")
	verification := "%s:%s@tcp(%s:%s)/%s?charset=utf8"
	verificationStr := fmt.Sprintf(verification, sqluser, sqlpass, sqlhost, sqlport, dbname)
	verificationStr += "&loc=Asia%2FShanghai"
	orm.RegisterDriver(sqlname.(string), orm.DRMySQL)
	orm.RegisterDataBase("default", sqlname.(string), verificationStr)
}