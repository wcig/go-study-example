package ch42_time

import (
	"fmt"
	"testing"
	"time"
)

var c chan int

func handle(int) {}

func TestAfter(t *testing.T) {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(3 * time.Second):
		fmt.Println("timed out")
	}
	// output:
	// timed out
}

func TestSleep(t *testing.T) {
	fmt.Println(time.Now())
	time.Sleep(100 * time.Millisecond)
	fmt.Println(time.Now())
	// output:
	// 2021-07-04 21:19:10.624008 +0800 CST m=+0.001123017
	// 2021-07-04 21:19:10.725104 +0800 CST m=+0.102218252
}

func TestTick(t *testing.T) {
	c := time.Tick(5 * time.Second)
	for next := range c {
		fmt.Printf("%v\n", next)
	}
}

func TestTypeDuration(t *testing.T) {
	start := time.Now()
	time.Sleep(time.Second)
	end := time.Now()
	d := end.Sub(start)
	fmt.Println(d) // 1.000412678s
}

func TestParseDuration(t *testing.T) {
	hours, _ := time.ParseDuration("10h")
	ccomplex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	micro2, _ := time.ParseDuration("1us") // 支持接受不正确但常见的micro前缀u

	fmt.Println(hours)
	fmt.Println(ccomplex)
	fmt.Printf("There are %.0f seconds in %v.\n", ccomplex.Seconds(), ccomplex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro)
	// output:
	// 10h0m0s
	// 1h10m10s
	// There are 4210 seconds in 1h10m10s.
	// There are 1000 nanoseconds in 1µs.
	// There are 1.00e-06 seconds in 1µs.
}

func TestSince(t *testing.T) {
	t1 := time.Now()
	time.Sleep(time.Second)
	d := time.Since(t1)
	fmt.Println(d) // 1.000057505s
}

func TestUntil(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(time.Second)
	d := time.Until(t2)
	fmt.Println(d) // 999.999658ms
}

func TestTypeDurationUnit(t *testing.T) {
	d, _ := time.ParseDuration("1h")
	fmt.Println("hours:", d.Hours())
	fmt.Println("minutes:", d.Minutes())
	fmt.Println("seconds:", d.Seconds())
	fmt.Println("milliseconds:", d.Milliseconds())
	fmt.Println("microseconds:", d.Microseconds())
	fmt.Println("nanoseconds:", d.Nanoseconds())
	// output:
	// hours: 1
	// minutes: 60
	// seconds: 3600
	// milliseconds: 3600000
	// microseconds: 3600000000
	// nanoseconds: 3600000000000
}

func TestTypeDurationRound(t *testing.T) {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
	}
	// output:
	// d.Round(   1ns) = 1h15m30.918273645s
	// d.Round(   1µs) = 1h15m30.918274s
	// d.Round(   1ms) = 1h15m30.918s
	// d.Round(    1s) = 1h15m31s
	// d.Round(    2s) = 1h15m30s
	// d.Round(  1m0s) = 1h16m0s
	// d.Round( 10m0s) = 1h20m0s
	// d.Round(1h0m0s) = 1h0m0s
}

func TestTypeDurationTruncate(t *testing.T) {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, t := range trunc {
		fmt.Printf("d.Truncate(%6s) = %s\n", t, d.Truncate(t).String())
	}
	// output:
	// d.Truncate(   1ns) = 1h15m30.918273645s
	// d.Truncate(   1µs) = 1h15m30.918273s
	// d.Truncate(   1ms) = 1h15m30.918s
	// d.Truncate(    1s) = 1h15m30s
	// d.Truncate(    2s) = 1h15m30s
	// d.Truncate(  1m0s) = 1h15m0s
	// d.Truncate( 10m0s) = 1h10m0s
	// d.Truncate(1h0m0s) = 1h0m0s
}

func TestTypeLocation(t *testing.T) {
	localLocation := time.Local
	fmt.Println("local location:", localLocation)

	utcLocation := time.UTC
	fmt.Println("utc location:", utcLocation)
	// output:
	// local location: Local
	// utc location: UTC
}

func TestFixedZone(t *testing.T) {
	loc := time.FixedZone("UTC-8", -8*60*60)
	tt := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", tt.Format(time.RFC822)) // The time is: 10 Nov 09 23:00 UTC-8
}

func TestLoadLocation(t *testing.T) {
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}
	timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC.In(location)) // 2018-08-30 05:00:00 -0700 PDT

	location, _ = time.LoadLocation("Asia/Shanghai")
	timeInUTC = time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println(timeInUTC.In(location)) // 2018-08-30 20:00:00 +0800 CST
}

func TestTypeMonth(t *testing.T) {
	m := time.January
	fmt.Println(m.String()) // January
}

func TestTypeTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
	// output:
	// Current time:  2021-07-04 22:01:05.485374 +0800 CST m=+1.003103486
	// Current time:  2021-07-04 22:01:06.485025 +0800 CST m=+2.002746356
	// Current time:  2021-07-04 22:01:07.485018 +0800 CST m=+3.002731133
	// Current time:  2021-07-04 22:01:08.487606 +0800 CST m=+4.005310505
	// Current time:  2021-07-04 22:01:09.486623 +0800 CST m=+5.004318398
	// Current time:  2021-07-04 22:01:10.486134 +0800 CST m=+6.003820585
	// Current time:  2021-07-04 22:01:11.485212 +0800 CST m=+7.002890629
	// Current time:  2021-07-04 22:01:12.486298 +0800 CST m=+8.003967329
	// Current time:  2021-07-04 22:01:13.483397 +0800 CST m=+9.001057507
	// Current time:  2021-07-04 22:01:14.484739 +0800 CST m=+10.002390834
	// Done!
}

func TestTimer(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	timer := time.NewTimer(10 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
	// output:
	// Current time:  2021-07-04 22:10:29.610309 +0800 CST m=+1.004556306
	// Current time:  2021-07-04 22:10:30.612405 +0800 CST m=+2.006644190
	// Current time:  2021-07-04 22:10:31.608567 +0800 CST m=+3.002798059
	// Current time:  2021-07-04 22:10:32.610508 +0800 CST m=+4.004730419
	// Current time:  2021-07-04 22:10:33.609487 +0800 CST m=+5.003700902
	// Current time:  2021-07-04 22:10:34.608419 +0800 CST m=+6.002625695
	// Current time:  2021-07-04 22:10:35.607829 +0800 CST m=+7.002027558
	// Current time:  2021-07-04 22:10:36.60695 +0800 CST m=+8.001140335
	// Current time:  2021-07-04 22:10:37.609684 +0800 CST m=+9.003866342
	// Current time:  2021-07-04 22:10:38.609597 +0800 CST m=+10.003770316
	// Done!
}

func TestTypeWeekDay(t *testing.T) {
	wd := time.Sunday
	fmt.Println(wd) // Sunday
}
