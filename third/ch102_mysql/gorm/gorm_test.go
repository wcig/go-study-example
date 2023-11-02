package gorm

import (
	"fmt"
	"log"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	host     = "localhost"
	port     = "3306"
	username = "root"
	password = "123456"
	dbname   = "test"

	dsn           = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	defaultLogger = logger.Default.LogMode(logger.Info)

	DB *gorm.DB
)

func initGorm() {
	dia := mysql.Dialector{Config: &mysql.Config{DSN: dsn}}
	cfg := &gorm.Config{Logger: defaultLogger}
	db, err := gorm.Open(dia, cfg)
	if err != nil {
		log.Fatalf(">> init gorm err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf(">> init gorm sql db err: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(1)
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf(">> init gorm sql db ping err: %v", err)
	}

	DB = db
	log.Println(">> init gorm success")
}

func closeGorm() {
	if DB != nil {
		sqlDB, _ := DB.DB()
		if err := sqlDB.Close(); err != nil {
			log.Fatalf(">> close gorm err: %v", err)
		}
	}
	log.Println(">> close gorm success")
}

func TestInitGorm(t *testing.T) {
	initGorm()
	defer closeGorm()
}

type GormUser struct {
	ID         int64  `json:"id" gorm:"id"`
	Name       string `json:"name" gorm:"age"`
	Age        int    `json:"age" gorm:"age"`
	CreateTime int64  `json:"create_time" gorm:"create_time"`
	UpdateTime int64  `json:"updateTime" gorm:"updateTime"`
}

func (gu *GormUser) TableName() string {
	return "gorm_user"
}

func TestCRUD(t *testing.T) {
	initGorm()
	defer closeGorm()

	// insert
	iu := &GormUser{
		Name:       "tom",
		Age:        12,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	ie := DB.Create(iu).Error
	if ie != nil {
		log.Fatalf(">> gorm insert err: %v", ie)
	}
	log.Printf(">> gorm insert success: %v", iu)

	// update
	time.Sleep(time.Second)
	up := map[string]interface{}{
		"name":        "jerry",
		"update_time": time.Now().Unix(),
	}
	ue := DB.Model(&GormUser{}).Where("id = ?", iu.ID).Updates(up).Error
	if ue != nil {
		log.Fatalf(">> gorm update err: %v", ue)
	}
	log.Println(">> gorm update success")

	// select
	var su GormUser
	se := DB.Where("id = ?", iu.ID).First(&su).Error
	if se != nil {
		log.Fatalf(">> gorm select err: %v", se)
	}
	log.Printf(">> gorm select success: %v", su)

	// delete
	de := DB.Where("id = ?", iu.ID).Delete(&GormUser{}).Error
	if de != nil {
		log.Fatalf(">> gorm delete err: %v", de)
	}
	log.Println(">> gorm delete success")
}

type GormTime struct {
	ID    int       `json:"id" gorm:"id"`
	TTime time.Time `json:"t_time" gorm:"t_time"`
	DTime time.Time `json:"d_time" gorm:"d_time"`
	BTime int64     `json:"b_time" gorm:"b_time"`
}

func (gt *GormTime) TableName() string {
	return "gorm_time"
}

func TestTime(t *testing.T) {
	initGorm()

	now := time.Now()
	it := &GormTime{
		TTime: now,
		DTime: now,
		BTime: now.UnixMicro(),
	}
	ie := DB.Create(it).Error
	if ie != nil {
		log.Fatalf(">> gorm insert err: %v", ie)
	}
	log.Printf(">> gorm insert success: %v", it)

	var st GormTime
	se := DB.Where("id = ?", it.ID).First(&st).Error
	if se != nil {
		log.Fatalf(">> gorm select err: %v", se)
	}
	log.Printf(">> gorm select success: %v", st)

	// Output:

	// 1) parseTime=True&loc=Local
	// 2023/11/02 23:08:12 >> init gorm success
	//
	// 2023/11/02 23:08:12 /Users/yangbo/Documents/workspace/myproject/go-study-example/third/ch102_mysql/gorm/gorm_test.go:141
	// [1.794ms] [rows:1] INSERT INTO `gorm_time` (`t_time`,`d_time`,`b_time`) VALUES ('2023-11-02 23:08:12.858','2023-11-02 23:08:12.858',1698937692858097)
	// 2023/11/02 23:08:12 >> gorm insert success: &{1 2023-11-02 23:08:12.858097 +0800 CST m=+0.006780501 2023-11-02 23:08:12.858097 +0800 CST m=+0.006780501 1698937692858097}
	//
	// 2023/11/02 23:08:12 /Users/yangbo/Documents/workspace/myproject/go-study-example/third/ch102_mysql/gorm/gorm_test.go:148
	// [0.755ms] [rows:1] SELECT * FROM `gorm_time` WHERE id = 1 ORDER BY `gorm_time`.`id` LIMIT 1
	// 2023/11/02 23:08:12 >> gorm select success: {1 2023-11-02 23:08:13 +0800 CST 2023-11-02 23:08:13 +0800 CST 1698937692858097}

	// 2) parseTime=True
	// 2023/11/02 23:12:05 >> init gorm success
	//
	// 2023/11/02 23:12:05 /Users/yangbo/Documents/workspace/myproject/go-study-example/third/ch102_mysql/gorm/gorm_test.go:142
	// [1.473ms] [rows:1] INSERT INTO `gorm_time` (`t_time`,`d_time`,`b_time`) VALUES ('2023-11-02 23:12:05.166','2023-11-02 23:12:05.166',1698937925166302)
	// 2023/11/02 23:12:05 >> gorm insert success: &{2 2023-11-02 23:12:05.166302 +0800 CST m=+0.012275793 2023-11-02 23:12:05.166302 +0800 CST m=+0.012275793 1698937925166302}
	//
	// 2023/11/02 23:12:05 /Users/yangbo/Documents/workspace/myproject/go-study-example/third/ch102_mysql/gorm/gorm_test.go:149
	// [0.371ms] [rows:1] SELECT * FROM `gorm_time` WHERE id = 2 ORDER BY `gorm_time`.`id` LIMIT 1
	// 2023/11/02 23:12:05 >> gorm select success: {2 2023-11-02 15:12:05 +0000 UTC 2023-11-02 15:12:05 +0000 UTC 1698937925166302}
}
