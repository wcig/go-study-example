package ch26_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 获取当前时间Time
func TestNow(t *testing.T) {
	lt := time.Now()
	fmt.Println("当前时间:", lt)

	secondTimestamp := lt.Unix()
	MillisecondTimestamp := lt.UnixNano() / 1e6
	nanosecondTimestamp := lt.UnixNano()
	fmt.Println("秒时间戳:", secondTimestamp)
	fmt.Println("毫秒时间戳:", MillisecondTimestamp)
	fmt.Println("纳秒时间戳:", nanosecondTimestamp)

	location := lt.Location()
	fmt.Println("位置:", location)

	zoneName, zoneOffset := lt.Zone()
	fmt.Printf("时区名称:%s, 时区相对标准时间偏移量(单位秒):%d\n", zoneName, zoneOffset)

	fmt.Println("时间:", lt.Year(), lt.Month(), lt.Day(), lt.YearDay(), lt.Weekday(), lt.Hour(), lt.Minute(), lt.Second(), lt.Nanosecond())
}

// Time -> string
func TestTime2String(t *testing.T) {
	now := time.Now()
	fmt.Println(time.Now().Format(time.UnixDate))      // Thu Mar  4 16:54:35 CST 2021
	fmt.Println(time.Now().Format(time.RFC3339Nano))   // 2021-03-04T16:54:35.185979712+08:00
	fmt.Println(time.Now().Format(time.StampMicro))    // Mar  4 16:54:35.185982
	fmt.Println(now.Format("2006-01-02 15:04:05.999")) // 2021-03-04 16:54:35.185
}

// string -> Time
func TestString2Time(t *testing.T) {
	str := "2021-03-04 16:54:35.185"
	format := "2006-01-02 15:04:05.999"
	tt, err := time.Parse(format, str)
	assert.Nil(t, err)
	fmt.Println(tt)
}

// Time -> timestamp
func TestTime2Timestamp(t *testing.T) {
	lt := time.Now()
	secondTimestamp := lt.Unix()
	MillisecondTimestamp := lt.UnixNano() / 1e6
	nanosecondTimestamp := lt.UnixNano()
	fmt.Println("秒时间戳:", secondTimestamp)
	fmt.Println("毫秒时间戳:", MillisecondTimestamp)
	fmt.Println("纳秒时间戳:", nanosecondTimestamp)
}

// timestamp -> Time
func TestTimestamp2Time(t *testing.T) {
	lt := time.Now()
	fmt.Println(lt) // 2021-03-04 18:23:53.783620241 +0800 CST m=+0.000490590

	sts := lt.Unix()
	tt1 := time.Unix(sts, 0)
	fmt.Println(tt1) // 2021-03-04 18:23:53 +0800 CST

	nts := lt.UnixNano()
	tt2 := time.Unix(0, nts)
	fmt.Println(tt2) // 2021-03-04 18:23:53.783620241 +0800 CST
}
