package ch26_time

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 获取当前时间Time
func TestNowTime(t *testing.T) {
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
	year, week := lt.ISOWeek()
	fmt.Println(year, week)

	// UTC时间
	fmt.Println("当前时间转换为UTC时间:", lt.UTC())
}

// output:
// 当前时间: 2021-03-05 15:12:39.788645297 +0800 CST m=+0.000487483
// 秒时间戳: 1614928359
// 毫秒时间戳: 1614928359788
// 纳秒时间戳: 1614928359788645297
// 位置: Local
// 时区名称:CST, 时区相对标准时间偏移量(单位秒):28800
// 时间: 2021 March 5 64 Friday 15 12 39 788645297
// 2021 9
// 2021-03-05 07:12:39.788645297 +0000 UTC

// 位置-时区
func TestLocation(t *testing.T) {
	utcLc, err := time.LoadLocation("UTC")
	assert.Nil(t, err)
	fmt.Println(utcLc)

	shLc, err := time.LoadLocation("Asia/Shanghai")
	assert.Nil(t, err)
	fmt.Println(shLc)

	lt := time.Now()
	fmt.Println(lt)
	fmt.Println(lt.In(utcLc))
	fmt.Println(lt.In(shLc))

	format := "2006-01-02 15:04:05.999"
	str := "2021-03-05 11:43:11.959"
	tt, err := time.ParseInLocation(format, str, utcLc)
	assert.Nil(t, err)
	fmt.Println(tt)
}

// output:
// UTC
// Asia/Shanghai
// 2021-03-05 11:50:11.237275717 +0800 CST m=+0.000541605
// 2021-03-05 03:50:11.237275717 +0000 UTC
// 2021-03-05 11:50:11.237275717 +0800 CST
// 2021-03-05 11:43:11.959 +0000 UTC

// 位置-时区 (时区信息文件目录: $GOROOT/lib/time/zoneinfo.zip)
func TestLocationWithZone(t *testing.T) {
	utcZoneLc := time.FixedZone("UTC", 0)
	shZoneLc := time.FixedZone("Asia/Shanghai", 8*60*60)
	fmt.Println(utcZoneLc)
	fmt.Println(shZoneLc)

	lt := time.Now()
	fmt.Println(lt)
	fmt.Println(lt.In(utcZoneLc))
	fmt.Println(lt.In(shZoneLc))
}

// output:
// UTC
// Asia/Shanghai
// 2021-03-05 11:49:35.595836736 +0800 CST m=+0.000499919
// 2021-03-05 03:49:35.595836736 +0000 UTC
// 2021-03-05 11:49:35.595836736 +0800 Asia/Shanghai

// 时间计算
func TestTimeCal(t *testing.T) {
	lt := time.Now()
	fmt.Println(lt)

	// 时间加减
	lt = lt.Add(time.Second)
	lt = lt.Add(time.Minute)
	lt = lt.Add(-time.Hour)
	fmt.Println(lt)

	lt = lt.AddDate(-1, 1, 1)
	fmt.Println(lt)

	// 时间取整
	fmt.Println(lt.Truncate(time.Second))                        // 纳秒取零
	fmt.Println(lt.Truncate(time.Minute))                        // 纳秒,秒取零
	fmt.Println(lt.Truncate(time.Hour))                          // 纳秒,秒,分钟取零
	fmt.Println(lt.Truncate(time.Hour * 24).Add(time.Hour * -8)) // 纳秒,秒,分钟,小时取零

	// 时间前后判断
	fmt.Println(lt.Before(lt.Add(time.Second)))
	fmt.Println(lt.After(lt.Add(time.Second)))
	fmt.Println(lt.Equal(lt.Add(time.Second)))

	// now - t -> duration
	oneHourAgoTime := time.Now().Add(time.Hour)
	fmt.Println(int64(time.Since(oneHourAgoTime).Seconds()))

	// t - now
	fmt.Println(int64(time.Until(oneHourAgoTime).Seconds()))

	// t1 - t2
	now := time.Now()
	fmt.Println(int64(now.Sub(now.Add(-time.Hour)).Seconds()))
}

// output:
// 2021-03-05 15:14:01.262368652 +0800 CST m=+0.000533499
// 2021-03-05 14:15:02.262368652 +0800 CST m=-3538.999466501
// 2020-04-06 14:15:02.262368652 +0800 CST
// 2020-04-06 14:15:02 +0800 CST
// 2020-04-06 14:15:00 +0800 CST
// 2020-04-06 14:00:00 +0800 CST
// 2020-04-06 00:00:00 +0800 CST
// true
// false
// false
// -3599
// 3599
// 3600

// Duration
func TestDuration(t *testing.T) {
	// 时长基础单位
	fmt.Println(time.Nanosecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Millisecond)
	fmt.Println(time.Second)
	fmt.Println(time.Minute)
	fmt.Println(time.Hour)

	// 时长换算
	fmt.Println(time.Hour.Hours())
	fmt.Println(time.Hour.Minutes())
	fmt.Println(time.Hour.Seconds())
	fmt.Println(time.Hour.Milliseconds())
	fmt.Println(time.Hour.Microseconds())
	fmt.Println(time.Hour.Nanoseconds())

	// string -> duration
	hd, _ := time.ParseDuration("1h")
	md, _ := time.ParseDuration("2m")
	sd, _ := time.ParseDuration("3s")
	fmt.Println(hd, md, sd)

	// 截断
	td1, _ := time.ParseDuration("1h2m3s4ms")
	fmt.Println(td1)
	fmt.Println(td1.Truncate(time.Second))
	fmt.Println(td1.Truncate(time.Minute))
	fmt.Println(td1.Truncate(time.Hour))

	// 四舍五入
	td2, _ := time.ParseDuration("1h2m30s")
	fmt.Println(td2)
	fmt.Println(td2.Round(time.Minute))
}

// output:
// 1ns
// 1µs
// 1ms
// 1s
// 1m0s
// 1h0m0s
// 1
// 60
// 3600
// 3600000
// 3600000000
// 3600000000000
// 1h0m0s 2m0s 3s
// 1h2m3.004s
// 1h2m3s
// 1h2m0s
// 1h0m0s
// 1h2m30s
// 1h3m0s

// Date -> Time
func TestInitTimeWithDate(t *testing.T) {
	utcLc, _ := time.LoadLocation("UTC")
	tt := time.Date(2021, 3, 5, 13, 43, 1, 2, utcLc)
	fmt.Println(tt) // 2021-03-05 13:43:01.000000002 +0000 UTC
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
