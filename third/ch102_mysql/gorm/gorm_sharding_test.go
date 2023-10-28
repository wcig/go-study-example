package gorm

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/sharding"
)

type GormShardingUser struct {
	ID         string `json:"id" gorm:"id"`
	Name       string `json:"name" gorm:"age"`
	Age        int    `json:"age" gorm:"age"`
	CreateTime int64  `json:"create_time" gorm:"create_time"`
	UpdateTime int64  `json:"updateTime" gorm:"updateTime"`
}

func (gsu *GormShardingUser) TableName() string {
	return "gorm_sharding_user"
}

func NewGid() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}

func TestSharding(t *testing.T) {
	initGorm()
	defer closeGorm()

	err := DB.Use(sharding.Register(sharding.Config{
		ShardingKey:         "id",
		NumberOfShards:      3,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "gorm_sharding_user"))
	if err != nil {
		log.Fatalf(">> gorm sharding config err: %v", err)
	}

	// insert
	log.Println(">> ---------------insert---------------")
	var ids []string
	for i := 0; i < 5; i++ {
		id := NewGid()
		ids = append(ids, id)
		iu := &GormShardingUser{
			ID:         id,
			Name:       fmt.Sprintf("tom-%d", i),
			Age:        12,
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
		}
		ie := DB.Create(iu).Error
		if ie != nil {
			log.Fatalf(">> gorm sharding insert err: %v", err)
		}
		log.Printf(">> gorm sharding insert success: %v", iu)
	}

	// select
	log.Println(">> ---------------select---------------")
	for i := 0; i < 5; i++ {
		id := ids[i]
		var su GormShardingUser
		se := DB.Where("id = ?", id).First(&su).Error
		if se != nil {
			log.Fatalf(">> gorm sharding select err: %v", err)
		}
		log.Printf(">> gorm sharding select success: %v", su)
	}

	// delete
	log.Println(">> ---------------delete---------------")
	for i := 0; i < 5; i++ {
		id := ids[i]
		se := DB.Where("id = ?", id).Delete(&GormShardingUser{}).Error
		if se != nil {
			log.Fatalf(">> gorm sharding delete err: %v", err)
		}
		log.Printf(">> gorm sharding delete success: %s", id)
	}
}
