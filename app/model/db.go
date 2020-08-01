package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GormWorker struct {
	dsn string
	Db  *gorm.DB
}

func NewGorm() GormWorker{
	dbw := GormWorker{dsn: "root:woshiytc@tcp(localhost:3306)/dingwei?charset=utf8mb4"}
	//支持下面几种DSN写法，具体看mysql服务端配置，常见为第2种
	//user@unix(/path/to/socket)/dbname?charset=utf8
	//user:password@tcp(localhost:5555)/dbname?charset=utf8
	//user:password@/dbname
	//user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

	dbtemp, _ := gorm.Open("mysql", dbw.dsn)

	dbw.Db = dbtemp
	return dbw
}