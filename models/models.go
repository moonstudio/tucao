package models

import (
	_ "fmt"
	"github.com/Unknwon/com"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	_ "strconv"
	"time"
)

const (
	_DB_NAME        = "data/tucao.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type TuCao struct {
	Id         int64
	Content    string
	CreateTime time.Time `orm:"index"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(TuCao))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddTuCao(content string) error {
	o := orm.NewOrm()
	t := &TuCao{Content: content, CreateTime: time.Now()}
	var err error
	_, err = o.Insert(t)
	if err != nil {
		return err
	}
	return nil
}
func GetAll() ([]*TuCao, error) {
	o := orm.NewOrm()
	cates := make([]*TuCao, 0)
	qs := o.QueryTable("tuCao")
	//倒序排序
	_, err := qs.OrderBy("-id").All(&cates)
	return cates, err
}
func Count() int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("tuCao").Count()
	if err != nil {
		return 0
	}
	return cnt
}
