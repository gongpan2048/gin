//数据库配置文件
package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/pan?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	SqlDB.SetMaxIdleConns(20) //设置最大空闲连接数
	SqlDB.SetMaxOpenConns(20) //设置最大连接数

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
