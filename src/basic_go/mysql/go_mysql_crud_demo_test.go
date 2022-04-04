package main

import (
	"testing"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func TestConnect(t *testing.T) {
	// 1.建立mysql连接
	name := "mysql"
	password := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	_, err := Connet(name, password)
	t.Log(err, "连接信息")
}

func TestQuery(t *testing.T) {
	name := "mysql"
	password := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	db, err := Connet(name, password)
	if err != nil {
		t.Error(err)
	}
	rows, err := Query(db, "select * from tb_user")
	if err != nil {
		t.Error(err)
	}
	for rows.Next() {
		var id int
		var name string
		var createtime time.Time
		err = rows.Scan(&id, &name, &createtime)
		if err != nil {
			t.Error(err)
		}
		t.Log(id, name, createtime)
	}
}

func TestInsert(t *testing.T) {
	name := "mysql"
	password := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	db, err := Connet(name, password)
	if err != nil {
		t.Error(err)
	}
	result, err := Insert(db, "Trump")
	if err != nil {
		t.Error(err)
	}
	t.Log(result.LastInsertId())
}

func TestUpdate(t *testing.T) {
	name := "mysql"
	password := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	db, err := Connet(name, password)
	if err != nil {
		t.Error(err)
	}
	result, err := Update(db, 1, "Biden")
	if err != nil {
		t.Error(err)
	}
	t.Log(result.RowsAffected())
}

func TestDelete(t *testing.T) {
	name := "mysql"
	password := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	db, err := Connet(name, password)
	if err != nil {
		t.Error(err)
	}
	result, err := Delete(db, 1)
	if err != nil {
		t.Error(err)
	}	
	t.Log(result.RowsAffected())
}
