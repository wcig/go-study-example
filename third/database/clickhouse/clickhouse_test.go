package clickhouse

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/ClickHouse/clickhouse-go/v2"
)

func TestFirst(t *testing.T) {
	// 连接 ClickHouse 数据库
	// db, err := sql.Open("clickhouse", "tcp://localhost:9000/default?username=default&password=")
	db, err := sql.Open("clickhouse", "tcp://localhost:9000/default")
	if err != nil {
		fmt.Println("Failed to connect to ClickHouse:", err)
		return
	}
	defer func() {
		if dbCloseErr := db.Close(); dbCloseErr != nil {
			log.Fatal(err)
		}
	}()

	// 创建表结构
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS my_table2 (
			id UInt32,
			name String,
			age UInt8
		) ENGINE = MergeTree()
		ORDER BY id;
	`)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	// 插入数据
	_, err = db.Exec("INSERT INTO my_table2 (id, name, age) VALUES (?, ?, ?)", 1, "Alice", 30)
	if err != nil {
		fmt.Println("Failed to insert data:", err)
		return
	}

	// 查询数据
	rows, err := db.Query("SELECT * FROM my_table2")
	if err != nil {
		fmt.Println("Failed to query data:", err)
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	// 打印查询结果
	for rows.Next() {
		var id uint32
		var name string
		var age uint8
		if err = rows.Scan(&id, &name, &age); err != nil {
			fmt.Println("Failed to scan row:", err)
			return
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// 检查是否有错误发生
	if err = rows.Err(); err != nil {
		fmt.Println("Error occurred during rows iteration:", err)
		return
	}
}
