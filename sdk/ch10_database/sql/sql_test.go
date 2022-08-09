package sql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库准备
// CREATE SCHEMA `test` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ;
// CREATE TABLE `test`.`user` (
//  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
//  `name` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
//  `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号',
//  `password` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
//  `create_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
//  PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';
// INSERT INTO `test`.`user` (`id`, `name`, `phone`, `password`, `create_time`) VALUES ('1', 'tom', '111', '111', '1623580604');
// INSERT INTO `test`.`user` (`id`, `name`, `phone`, `password`, `create_time`) VALUES ('2', 'jerry', '222', '222', '1623580627');

type User struct {
	Id         int
	Name       string
	Phone      string
	Password   string
	CreateTime int64
}

// 连接数据库，创建DB
func TestCreateDB(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

// DB选项
func TestDBOptions(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(2)                  // 最大连接数
	db.SetMaxIdleConns(10)                 // 最大空闲连接数
	db.SetConnMaxLifetime(time.Minute)     // 连接可以复用的最大时间
	db.SetConnMaxIdleTime(time.Minute * 5) // 连接处于空闲的最大时间
}

// 创建Conn连接
func TestCreateConn(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
}

// 创建Stmt
func TestCreateStmt(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stmt, err := conn.PrepareContext(ctx, "select name from user where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
}

// stmt执行sql
func TestStmtExecSql(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stmt, err := conn.PrepareContext(ctx, "INSERT INTO `user` (`name`, `phone`, `password`, `create_time`) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	u := User{
		Name:       "han",
		Phone:      "333",
		Password:   "333",
		CreateTime: time.Now().Unix(),
	}
	result, err := stmt.Exec(u.Name, u.Phone, u.Password, u.CreateTime)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
	fmt.Println(result.LastInsertId())
	// output:
	// 1 <nil>
	// 3 <nil>
}

// 获取结果
func TestQueryAndParseRows(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stmt, err := conn.PrepareContext(ctx, "select * from user")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var (
		id         int
		name       string
		phone      string
		password   string
		createTime int64
	)
	for rows.Next() {
		if err := rows.Scan(&id, &name, &phone, &password, &createTime); err != nil {
			panic(err)
		}
		fmt.Printf("id:%d, name:%s, phone:%s, password:%s, createTime:%d\n",
			id, name, phone, password, createTime)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	// output:
	// id:1, name:tom, phone:111, password:111, createTime:1623580604
	// id:2, name:jerry, phone:222, password:222, createTime:1623580627
	// id:3, name:han, phone:333, password:333, createTime:1623581256
}

// 创建Tx有两种方式：DB创建和Conn创建（这里采取Conn方式）
func TestTx(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	tx, err := conn.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	tx.Exec("select * from user where id = ?", 1)
	tx.Exec("update user set phone = ? where id = ?", "111111", 1)
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}
