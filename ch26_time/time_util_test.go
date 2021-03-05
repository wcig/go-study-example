package ch26_time

import (
	"fmt"
	"testing"
	"time"
)

// test
func TestUtil(t *testing.T) {
	fmt.Println(GetNowUnix())
	fmt.Println(GetNowUnixMs())
	fmt.Println(GetNowUnixNs())

	fmt.Println(GetNowTimeStr())
	fmt.Println(GetNowTimeMsStr())

	fmt.Println(GetTodayBeginTime())
	fmt.Println(GetTodayEndTime())

	fmt.Println(GetTimeByStr("2021-03-05 01:02:03"))
	fmt.Println(GetUnixSByStr("2021-03-05 01:02:03"))
	fmt.Println(GetUnixMsByStr("2021-03-05 01:02:03.004"))
}

// output:
// 1614929534
// 1614929534458
// 1614929534458520502
// 2021-03-05 15:32:14
// 2021-03-05 15:32:14.458
// 2021-03-05 00:00:00 +0800 CST
// 2021-03-05 23:59:59.999999999 +0800 CST
// 2021-03-05 01:02:03 +0000 UTC <nil>
// 1614906123 <nil>
// 1614906123004 <nil>

/* ------------------------------------------------------------------- */
const (
	datetimeFormat   = "2006-01-02 15:04:05"
	datetimeMsFormat = "2006-01-02 15:04:05.999"
)

// 当前时间戳(秒)
func GetNowUnix() int64 {
	return time.Now().Unix()
}

// 当前时间戳(毫秒)
func GetNowUnixMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// 当前时间戳(纳秒)
func GetNowUnixNs() int64 {
	return time.Now().UnixNano()
}

// 当前时间(字符串格式)
func GetNowTimeStr() string {
	return time.Now().Format(datetimeFormat)
}

// 当前时间(字符串格式+毫秒)
func GetNowTimeMsStr() string {
	return time.Now().Format(datetimeMsFormat)
}

// 今天起始时间
func GetTodayBeginTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// 今天技术时间
func GetTodayEndTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, now.Location())
}

// 字符串->时间(格式:2006-01-02 15:04:05)
func GetTimeByStr(val string) (tt time.Time, err error) {
	return time.Parse(datetimeFormat, val)
}

// 字符串->时间戳(秒)
func GetUnixSByStr(val string) (ts int64, err error) {
	tt, err := time.Parse(datetimeFormat, val)
	if err != nil {
		return 0, err
	}
	return tt.Unix(), nil
}

// 字符串->时间戳(毫秒)
func GetUnixMsByStr(val string) (ts int64, err error) {
	tt, err := time.Parse(datetimeMsFormat, val)
	if err != nil {
		return 0, err
	}
	return tt.UnixNano() / 1e6, nil
}
