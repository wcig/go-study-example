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
