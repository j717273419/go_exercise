package main

// 1.建立mysql连接
// 2.建立mysql数据库
// 3.建立mysql表
// 4.插入数据
// 5.查询数据,单条，多条数据查询
// 6.更新数据
// 7.删除数据
// 8.关闭mysql连接
// 9.事务
import (
	"database/sql"
	"log"

	// "time"
	"fmt"
)

func main() {
	fmt.Println("hi")
}

// 1.建立mysql连接
func Connet(name, password string) (db *sql.DB, err error) {
	// 1.建立mysql连接
	db, err = sql.Open(name, password)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

// 2.mysql 查询
func Query(db *sql.DB, sql string) (rows *sql.Rows, err error) {
	// mysql 查询
	rows, err = db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	return rows, err
}

// 3.mysql 插入
func Insert(db *sql.DB, username string) (result sql.Result, err error) {
	// mysql 插入
	result, err = db.Exec("insert into tb_user (username) values(?)", username)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

// 4.mysql 更新
func Update(db *sql.DB, id int, sql string) (result sql.Result, err error) {
	// mysql 更新
	result, err = db.Exec("update tb_user set username = ? where id = ?", sql, id)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

// 5.mysql 删除
func Delete(db *sql.DB, id int) (result sql.Result, err error) {
	// mysql 删除
	result, err = db.Exec("delete from tb_user where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
