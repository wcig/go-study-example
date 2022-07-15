package sql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryRows(t *testing.T) {
	db, err := sql.Open("mysql", "bang:bang@tcp(127.0.0.1:3306)/test")
	assert.Nil(t, err)
	defer db.Close()

	rawSql := "select * from user"
	rows, err := db.Query(rawSql)
	assert.Nil(t, err)
	assert.NotNil(t, rows)
}

func TestQueryRowsColumn(t *testing.T) {
	db, err := sql.Open("mysql", "bang:bang@tcp(127.0.0.1:3306)/test")
	assert.Nil(t, err)
	defer db.Close()

	rawSql := "select * from user"
	rows, err := db.Query(rawSql)
	assert.Nil(t, err)
	assert.NotNil(t, rows)

	columns, err := rows.Columns()
	assert.Nil(t, err)
	for i, column := range columns {
		fmt.Printf("column index: %d, name: %s\n", i, column)
	}
	// Output:
	// column index: 0, name: id
	// column index: 1, name: name
	// column index: 2, name: phone
	// column index: 3, name: password
	// column index: 4, name: create_time
}

func TestQueryRowsColumnType(t *testing.T) {
	db, err := sql.Open("mysql", "bang:bang@tcp(127.0.0.1:3306)/test")
	assert.Nil(t, err)
	defer db.Close()

	rawSql := "select * from user"
	rows, err := db.Query(rawSql)
	assert.Nil(t, err)
	assert.NotNil(t, rows)

	columnTypes, err := rows.ColumnTypes()
	assert.Nil(t, err)
	for i, columnType := range columnTypes {
		columnName := columnType.Name()
		dataTypeName := columnType.DatabaseTypeName()
		length, lengthOk := columnType.Length()
		precision, scale, decimalSizeOk := columnType.DecimalSize()
		nullable, nullableOk := columnType.Nullable()
		fmt.Printf("column index: %d, "+
			"name: %s, dataTypeName: %s, "+
			"length: %d, lengthOk: %t, "+
			"precision: %d, scale: %d, decimalSizeOk: %t, "+
			"nullable: %t, nullableOk: %t, "+
			"scanType: %s\n",
			i,
			columnName, dataTypeName,
			length, lengthOk,
			precision, scale, decimalSizeOk,
			nullable, nullableOk,
			columnType.ScanType(),
		)
	}
	// Output:
	// column index: 0, name: id, dataTypeName: INT, length: 0, lengthOk: false, precision: 0, scale: 0, decimalSizeOk: false, nullable: false, nullableOk: true, scanType: uint32
	// column index: 1, name: name, dataTypeName: VARCHAR, length: 0, lengthOk: false, precision: 0, scale: 0, decimalSizeOk: false, nullable: false, nullableOk: true, scanType: sql.RawBytes
	// column index: 2, name: phone, dataTypeName: VARCHAR, length: 0, lengthOk: false, precision: 0, scale: 0, decimalSizeOk: false, nullable: false, nullableOk: true, scanType: sql.RawBytes
	// column index: 3, name: password, dataTypeName: VARCHAR, length: 0, lengthOk: false, precision: 0, scale: 0, decimalSizeOk: false, nullable: false, nullableOk: true, scanType: sql.RawBytes
	// column index: 4, name: create_time, dataTypeName: BIGINT, length: 0, lengthOk: false, precision: 0, scale: 0, decimalSizeOk: false, nullable: false, nullableOk: true, scanType: uint64
}
