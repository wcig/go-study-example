package xorm

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"xorm.io/xorm/names"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

const (
	driverName  = "mysql"
	saveDir     = "out"
	packageName = "entity"
)

// 根据数据库表生成Go结构体文件
func AutoGen(userName, password, host, port, dbName, dir string) {
	db := initDB(driverName, userName, password, host, port, dbName)
	tables := listTables(db)
	genFiles(tables)
}

// 初始化数据库
func initDB(driverName, userName, password, host, port, dbName string) *xorm.Engine {
	const sourceFormat = "%s:%s@tcp(%s:%s)/%s?interpolateParams=False&charset=utf8mb4"
	dataSourceName := fmt.Sprintf(sourceFormat, userName, password, host, port, dbName)
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = engine.Ping(); err != nil {
		panic(err)
	}
	return engine
}

// 列出指定数据库所有表信息
func listTables(engine *xorm.Engine) []*schemas.Table {
	tables, err := engine.DBMetas()
	if err != nil {
		panic(err)
	}
	return tables
}

// 生成go结构体文件+表名列名常量文件
func genFiles(tables []*schemas.Table) {
	dir := filepath.Join(saveDir, packageName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	genConstFile(tables, saveDir, packageName)
	genStructFiles(tables, saveDir, packageName)
}

// 生成go结构体文件
func genStructFiles(tables []*schemas.Table, saveDir, packageName string) {
	for _, table := range tables {
		mapper := &names.SnakeMapper{}
		tableName := table.Name
		structName := mapper.Table2Obj(tableName)

		fileName := filepath.Join(saveDir, packageName, fmt.Sprintf("%s.go", structName))
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		var filedLines []string
		cols := table.Columns()
		for _, col := range cols {
			filedName := mapper.Table2Obj(col.Name)
			filedTypeName := typestring(col)
			filedTagStr := genTag(table, col)
			filedLine := fmt.Sprintf("%s %s %s", filedName, filedTypeName, filedTagStr)
			filedLines = append(filedLines, filedLine)
		}
		packageLine := fmt.Sprintf("package %s", packageName)
		typePrefixLine := fmt.Sprintf("type %s struct{", structName)
		typeSuffixLine := "}"
		tableNameFuncLine := fmt.Sprintf("func (%s) TableName() string {\n return \"%s\" \n}", structName, tableName)

		var lines []string
		lines = append(lines, packageLine, typePrefixLine)
		lines = append(lines, filedLines...)
		lines = append(lines, typeSuffixLine, tableNameFuncLine)
		fileStr := strings.Join(lines, "\n")
		_, err = file.WriteString(formatGo(fileStr))
		if err != nil {
			panic(err)
		}
	}
}

// 表名列名常量文件
func genConstFile(tables []*schemas.Table, saveDir, packageName string) {
	var (
		tableNameList []string
		colNameList   []string
	)

	for _, table := range tables {
		tableNameList = append(tableNameList, table.Name)
		for _, col := range table.Columns() {
			colNameList = append(colNameList, col.Name)
		}
	}

	tableNameList = removeDuplicateSlice(tableNameList)
	colNameList = removeDuplicateSlice(colNameList)

	packageLine := fmt.Sprintf("package %s", packageName)
	constPrefixLine := "const ("
	constSuffixLine := ")"
	mapper := &names.SnakeMapper{}
	var tableLines string
	for _, tableName := range tableNameList {
		structName := mapper.Table2Obj(tableName)
		tableLines += fmt.Sprintf("Tbl%s = \"%s\"\n", structName, tableName)
	}
	var colLines string
	for _, colName := range colNameList {
		structName := mapper.Table2Obj(colName)
		colLines += fmt.Sprintf("Col%s = \"%s\"\n", structName, colName)
	}
	lines := []string{packageLine, constPrefixLine, tableLines, constSuffixLine, constPrefixLine, colLines, constSuffixLine}
	fileStr := strings.Join(lines, "\n")

	fileName := filepath.Join(saveDir, packageName, "ConstTableInfo.go")
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(formatGo(fileStr))
	if err != nil {
		panic(err)
	}
}

// 生成field tag
func genTag(table *schemas.Table, col *schemas.Column) string {
	isNameId := col.FieldName == "Id"
	isIdPk := isNameId && typestring(col) == "int64"

	var res []string
	if !col.Nullable {
		if !isIdPk {
			res = append(res, "not null")
		}
	}
	if col.IsPrimaryKey {
		res = append(res, "pk")
	}
	if col.Default != "" {
		res = append(res, "default "+col.Default)
	}
	if col.IsAutoIncrement {
		res = append(res, "autoincr")
	}

	/*if col.SQLType.IsTime() && include(created, col.Name) {
		res = append(res, "created")
	}

	if col.SQLType.IsTime() && include(updated, col.Name) {
		res = append(res, "updated")
	}

	if col.SQLType.IsTime() && include(deleted, col.Name) {
		res = append(res, "deleted")
	}*/

	if /*supportComment &&*/ col.Comment != "" {
		res = append(res, fmt.Sprintf("comment('%s')", col.Comment))
	}

	names := make([]string, 0, len(col.Indexes))
	for name := range col.Indexes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		index := table.Indexes[name]
		var uistr string
		if index.Type == schemas.UniqueType {
			uistr = "unique"
		} else if index.Type == schemas.IndexType {
			uistr = "index"
		}
		if len(index.Cols) > 1 {
			uistr += "(" + index.Name + ")"
		}
		res = append(res, uistr)
	}

	nstr := col.SQLType.Name
	if col.Length != 0 {
		if col.Length2 != 0 {
			nstr += fmt.Sprintf("(%v,%v)", col.Length, col.Length2)
		} else {
			nstr += fmt.Sprintf("(%v)", col.Length)
		}
	} else if len(col.EnumOptions) > 0 { //enum
		nstr += "("
		opts := ""

		enumOptions := make([]string, 0, len(col.EnumOptions))
		for enumOption := range col.EnumOptions {
			enumOptions = append(enumOptions, enumOption)
		}
		sort.Strings(enumOptions)

		for _, v := range enumOptions {
			opts += fmt.Sprintf(",'%v'", v)
		}
		nstr += strings.TrimLeft(opts, ",")
		nstr += ")"
	} else if len(col.SetOptions) > 0 { //enum
		nstr += "("
		opts := ""

		setOptions := make([]string, 0, len(col.SetOptions))
		for setOption := range col.SetOptions {
			setOptions = append(setOptions, setOption)
		}
		sort.Strings(setOptions)

		for _, v := range setOptions {
			opts += fmt.Sprintf(",'%v'", v)
		}
		nstr += strings.TrimLeft(opts, ",")
		nstr += ")"
	}
	res = append(res, nstr)

	var tags []string
	jsonTag := fmt.Sprintf(`json:"%s,omitempty"`, col.Name)
	tags = append(tags, jsonTag)
	if len(res) > 0 {
		xormTag := fmt.Sprintf(`xorm:"%s"`, strings.Join(res, " "))
		tags = append(tags, xormTag)
	}
	return fmt.Sprintf("`%s`", strings.Join(tags, " "))
}

func typestring(col *schemas.Column) string {
	st := col.SQLType
	t := schemas.SQLType2Type(st)
	s := t.String()
	if s == "[]uint8" {
		return "[]byte"
	}
	return s
}

// go文件格式化
func formatGo(src string) string {
	b, err := format.Source([]byte(src))
	if err != nil {
		panic(err)
	}
	return string(b)
}

// 切片去重
func removeDuplicateSlice(list []string) []string {
	var result []string
	m := make(map[string]struct{})
	for _, item := range list {
		if _, ok := m[item]; !ok {
			m[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
