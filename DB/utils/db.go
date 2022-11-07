// 使用到database/sql包
package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"//只想导入不想用，所以在包前加个下划线
)

var(
	Db *sql.DB
	err error
)

func init() {
	// 有两个参数 第一个是driver驱动 第二个是 mysql的连接参数
	// 用户名:密码@tcp(ip地址:端口号)/数据库名称 密码为空也可以不填
//open是操作数据库的一个句柄
	Db, err = sql.Open("mysql", "root:y760069562@tcp(localhost:3306)/test")
	//如果连接本机数据库，则 tcp(localhost:3306) 可省略
	//Db, err = sql.Open("mysql", "root:y760069562@/test")
	//如果有错误则打印异常信息
	//if err != nil {
	//	panic(err.Error())
	//}
}
