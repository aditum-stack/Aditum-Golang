package model

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

// 数据库配置
const (
	userName = "root"
	password = "wangshihao"
	ip       = "47.100.236.64"
	port     = "3306"
	dbName   = "golang-email"
)

// 数据表接口
type DataTable struct {
}

// 数据库连接池
var DB *sql.DB

// 初始化连接
func init() {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	// 打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	// 设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	// 设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	// 验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

// 使用事务执行数据库操作
func Transaction(invoke func(tx *sql.Tx, res sql.Result) (sql.Result, error)) (sql.Result, error) {
	// 开启事务
	tx, err := DB.Begin()
	if err != nil {
		msg := "open tx fail"
		fmt.Println(msg)
		return nil, errors.New(msg)
	}

	// 预设结果集
	var res sql.Result

	// 通过事务执行订制操作
	res, err = invoke(tx, res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 将事务提交
	tx.Commit()

	// 返回正确结果集
	return res, nil
}

// 数据表订制操作内容
func invoke(tx *sql.Tx, res sql.Result) (sql.Result, error) {
	// 准备sql语句
	stmt, err := tx.Prepare("sql")
	if err != nil {
		msg := "Prepare fail"
		fmt.Println(msg)
		return nil, errors.New(msg)
	}
	// 将参数传递到sql语句中并且执行
	res, err = stmt.Exec()
	if err != nil {
		msg := "Exec fail"
		fmt.Println(msg)
		return nil, errors.New(msg)
	}
	return res, nil
}
