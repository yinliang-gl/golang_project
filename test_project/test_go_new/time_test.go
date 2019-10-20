package test_go_new

/**
参考 https://blog.csdn.net/wschq/article/details/80114036
*/
import (
	"fmt"
	"math"
	"testing"
	"time"
)

/**
初始化
*/
func TestTime01(t *testing.T) {
	// func Now() Time
	fmt.Println(time.Now())

	// func Parse(layout, value string) (Time, error)
	time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51")

	// func ParseInLocation(layout, value string, loc *Location) (Time, error) (layout已带时区时可直接用Parse)
	time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local)

	// func Unix(sec int64, nsec int64) Time
	time.Unix(1e9, 0)

	// func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	time.Date(2018, 1, 2, 15, 30, 10, 0, time.Local)

	// func (t Time) In(loc *Location) Time 当前时间对应指定时区的时间
	loc, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Now().In(loc))
}

/**
格式化和常用方法

*/
func TestTime02(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // 2018-04-24 10:11:20
	fmt.Println(time.Now().Format(time.UnixDate))         // Tue Apr 24 09:59:02 CST 2018

	fmt.Println(time.Now().Unix())

	// 获取指定日期的时间戳
	dt, _ := time.Parse("2016-01-02 15:04:05", "2018-04-23 12:24:51")
	fmt.Println(dt.Unix())

	fmt.Println(time.Date(2018, 1, 2, 15, 30, 10, 0, time.Local).Unix())
}

/*
时间段(Duartion)

      func (d Duration) Hours() float64
      func (d Duration) Minutes() float64
      func (d Duration) Seconds() float64
      func (d Duration) Nanoseconds() int64
      func (d Duration) Round(m Duration) Duration         // 四舍五入
      func (d Duration) Truncate(m Duration) Duration      // 向下取整
*/
func TestTime03(t *testing.T) {
	tp, _ := time.ParseDuration("1.5s")
	fmt.Println(tp.Truncate(1000), tp.Seconds(), tp.Nanoseconds())
}

/*
时区(Location)
*/
func TestTime04(t *testing.T) {
	// 默认UTC
	_, _ = time.LoadLocation("")
	// 服务器设定的时区，一般为CST
	_, _ = time.LoadLocation("Local")
	// 美国洛杉矶PDT
	_, _ = time.LoadLocation("America/Los_Angeles")

	// 获取指定时区的时间点
	local, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println(time.Date(2018, 1, 1, 12, 0, 0, 0, local))

}

/*
时间运算
  func Until(t Time) Duration     //  等价于 t.Sub(Now())，t与当前时间的间隔
*/
func TestTime05(t *testing.T) {

	// func After(d Duration) <-chan Time  非阻塞,可用于延迟
	time.After(time.Duration(10) * time.Second)

	// func Since(t Time) Duration 两个时间点的间隔
	start := time.Now()

	// func Sleep(d Duration)   休眠多少时间，休眠时处于阻塞状态，后续程序无法执行
	time.Sleep(time.Duration(10) * time.Second)

	fmt.Println(time.Since(start)) // 等价于 Now().Sub(t)， 可用来计算一段业务的消耗时间

	// func (t Time) Add(d Duration) Time
	fmt.Println(start.Add(time.Duration(10) * time.Second)) // 加
	// func (t Time) Sub(u Time) Duration                    // 减

	// func (t Time) AddDate(years int, months int, days int) Time
	fmt.Println(start.AddDate(1, 1, 1))

	// func (t Time) Before(u Time) bool
	// func (t Time) After(u Time) bool
	// func (t Time) Equal(u Time) bool          比较时间点时尽量使用Equal函数

	dt1 := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	dt2 := time.Date(2018, 1, 9, 23, 59, 22, 100, time.Local)
	// 不用关注时区，go会转换成时间戳进行计算
	fmt.Println(dt1.Sub(dt2))

	now := time.Now()

	// 一段时间之后
	fmt.Println(now.Add(time.Duration(10) * time.Minute))

	// 计算两个时间点的相差天数
	dt1 = time.Date(dt1.Year(), dt1.Month(), dt1.Day(), 0, 0, 0, 0, time.Local)
	dt2 = time.Date(dt2.Year(), dt2.Month(), dt2.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(int(math.Ceil(dt1.Sub(dt2).Hours() / 24)))

}

/*
时区转换
*/
func TestTime06(t *testing.T) {
	// time.Local 用来表示当前服务器时区
	// 自定义地区时间
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Println(time.Date(2018, 1, 2, 0, 0, 0, 0, beijing)) // 2018-01-02 00:00:00 +0800 Beijing Time

	// 当前时间转为指定时区时间
	fmt.Println(time.Now().In(beijing))

	// 指定时间转换成指定时区对应的时间
	_, _ = time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local)

	// 当前时间在零时区年月日   时分秒  时区
	_, _, _ = time.Now().UTC().Date()  // 2018 April 24
	_, _, _ = time.Now().UTC().Clock() // 3 47 15
	_, _ = time.Now().UTC().Zone()     // UTC

}

/*
比较两个时间点
*/

func TestTime07(t *testing.T) {
	dt := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	fmt.Println(time.Now().After(dt))  // true
	fmt.Println(time.Now().Before(dt)) // false

	// 是否相等 判断两个时间点是否相等时推荐使用 Equal 函数
	fmt.Println(dt.Equal(time.Now()))
}

/*
设置执行时间，超时限制等
*/
func TestTime08(t *testing.T) {
	c := make(chan interface{})
	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(time.Duration(5) * time.Second):
		fmt.Println("time out")
	}
}
