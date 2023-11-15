package xorm

import (
	"fmt"
	"log"
	"testing"
	"time"

	"xorm.io/xorm"
)

var xe *xorm.Engine

func initXorm() {
	const sourceFormat = "%s:%s@tcp(%s:%s)/%s?interpolateParams=False&charset=utf8mb4&parseTime=True&loc=Local"
	dataSourceName := fmt.Sprintf(sourceFormat, username, password, host, port, dbname)
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Fatalf(">> init xorm engine err: %v", err)
	}
	engine.SetMaxOpenConns(10)
	engine.SetMaxIdleConns(3)
	engine.ShowSQL(true)
	if err = engine.Ping(); err != nil {
		log.Fatalf(">> xorm ping err: %v", err)
	}
	xe = engine
}

func closeXorm() {
	if xe != nil {
		if err := xe.Close(); err != nil {
			log.Fatalf(">> xorm close err: %v", err)
		}
	}
}

func TestXorm(t *testing.T) {
	initXorm()
	defer closeXorm()
}

func TestCURD(t *testing.T) {
	initXorm()
	defer closeXorm()

	// insert
	iu := &XormUser{
		Name:       "tom",
		Age:        12,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	in, ie := xe.Insert(iu)
	if ie != nil {
		log.Fatalf(">> xorm insert err: %v", ie)
	}
	log.Printf(">> xorm insert success: %d, %v", in, iu)

	// update
	time.Sleep(time.Second)
	up := map[string]interface{}{
		"name":        "jerry",
		"update_time": time.Now().Unix(),
	}
	un, ue := xe.Table((&XormUser{}).TableName()).Where("id = ?", iu.Id).Update(up)
	if ue != nil {
		log.Fatalf(">> xorm update err: %v", ue)
	}
	log.Printf(">> xorm update success: %d", un)

	// select
	var su XormUser
	sn, se := xe.Where("id = ?", iu.Id).Get(&su)
	if se != nil {
		log.Fatalf(">> xorm select err: %v", se)
	}
	log.Printf(">> xorm select success: %t, %v", sn, su)

	// delete
	dn, de := xe.Where("id = ?", iu.Id).Delete(&XormUser{})
	if de != nil {
		log.Fatalf(">> xorm delete err: %v", de)
	}
	log.Printf(">> xorm delete success: %d", dn)

	// Output:
	// [xorm] [info]  2023/11/15 21:09:11.977635 PING DATABASE mysql
	// [xorm] [info]  2023/11/15 21:09:11.985024 [SQL] INSERT INTO `xorm_user` (`name`,`age`,`create_time`,`update_time`) VALUES (?,?,?,?) [tom 12 1700053751 1700053751] - 1.615041ms
	// 2023/11/15 21:09:11 >> xorm insert success: 1, &{1 tom 12 1700053751 1700053751}
	// [xorm] [info]  2023/11/15 21:09:12.987703 [SQL] UPDATE `xorm_user` SET `name` = ?, `update_time` = ? WHERE (id = ?) [jerry 1700053752 1] - 1.41775ms
	// 2023/11/15 21:09:12 >> xorm update success: 1
	// [xorm] [info]  2023/11/15 21:09:12.988275 [SQL] SELECT `id`, `name`, `age`, `create_time`, `update_time` FROM `xorm_user` WHERE (id = ?) LIMIT 1 [1] - 469.5µs
	// 2023/11/15 21:09:12 >> xorm select success: true, {1 jerry 12 1700053751 1700053752}
	// [xorm] [info]  2023/11/15 21:09:12.989649 [SQL] DELETE FROM `xorm_user` WHERE (id = ?) [1] - 1.184292ms
	// 2023/11/15 21:09:12 >> xorm delete success: 1
}

type XormUser struct {
	Id         int64  `json:"id,omitempty" xorm:"'id' not null pk autoincr comment('主键ID') BIGINT(20)"`
	Name       string `json:"name,omitempty" xorm:"'name' comment('名称') VARCHAR(60)"`
	Age        int    `json:"age,omitempty" xorm:"'age' comment('年龄') TINYINT(4)"`
	CreateTime int64  `json:"create_time,omitempty" xorm:"'create_time' comment('创建时间') BIGINT(20)"`
	UpdateTime int64  `json:"update_time,omitempty" xorm:"'update_time' comment('更新时间') BIGINT(20)"`
}

func (XormUser) TableName() string {
	return "xorm_user"
}
