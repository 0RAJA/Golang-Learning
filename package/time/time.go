package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
time包:
	1年=365天，day
	1天=24小时，hour
	1小时=60分钟，minute
	1分钟=60秒，second
	1秒钟=1000毫秒，millisecond
	1毫秒=1000微秒，microsecond-->us
	1微秒=1000纳秒，nanosecond-- >ns
	1纳秒=1000皮秒，picoseconds-->ps
*/
func main() {
	//1.获取当前时间
	t1 := time.Now()
	fmt.Printf("%T\n%v\n", t1, t1)
	fmt.Println("-------------------------------")
	//2.获取指定时间
	t2 := time.Date(2021, 5, 7, 14, 20, 20, 0, time.Local)
	fmt.Println(t2)
	fmt.Println("-------------------------------")
	//3.Time和string的转换
	//格式化时间: Mon Jan 2 15:04:05 -0700 MST 2006
	//Time -> string
	s1 := t1.Format("2006年1月2日 15:04:05 MST Mon Jan") //格式化转换
	fmt.Println(s1)
	fmt.Println(t1.String()) //直接转换
	//string -> Time
	s2 := "2021年3月26日Fri"
	t3, err := time.Parse("2006年1月2日Fri", s2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t3)
	fmt.Println("-------------------------------")
	//4.根据时间Time获取指定内容
	year, month, day := t1.Date() //年月日
	hour, min, sec := t1.Clock()  //时分秒
	fmt.Println(year, month, day, hour, min, sec)
	fmt.Println(t1.Year(), t1.YearDay()) //t1的年份,t1年已经度过的天数
	fmt.Println(t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond())
	fmt.Println("-------------------------------")
	//5.时间戳--距离 1970年1月1日0时0分0秒 的时间差值:秒,纳秒
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	fmt.Println(t4.Unix(), t4.UnixNano()) //3600s
	fmt.Println(t1.Unix())
	fmt.Println("-------------------------------")
	//6.时间间隔
	t5 := t1.Add(time.Minute + time.Hour)
	fmt.Println(t5)
	t6 := t1.AddDate(1, 0, 0)
	fmt.Println(t6)
	fmt.Printf("%v\n", t6.Sub(t5)) //t6 - t5
	fmt.Println("-------------------------------")
	//7.睡眠
	t7 := time.Now()
	time.Sleep(2 * time.Second)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10) + 1
	fmt.Println(num)
	time.Sleep(time.Second * time.Duration(num)) //注意要强制类型转换
	fmt.Println(time.Now().Sub(t7))
	fmt.Println("-------------------------------")
}
