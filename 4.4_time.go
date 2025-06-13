package main

import (
	"fmt"
	"math"
	"time"
)

func timeTest() {
	now := time.Now().AddDate(-8, 0, 0)
	day := now.Format("2006/01/02")
	zone := now.Format("-0700")
	timeStr := fmt.Sprintf("%s 09:30:00 %s", day, zone)
	fmt.Println(timeStr)

	if t, err := time.Parse("2006/01/02 15:04:05 -0700", timeStr); err != nil {
		panic(err)
	} else {
		fmt.Println(1, t.Local().UnixMilli())
		fmt.Println(2, t.Format("2006年01月02日 15时04分05秒 时区-0700"))
		fmt.Println(3, t.Weekday(), int(t.Weekday()))
		fmt.Println(t.ISOWeek())
		d := time.Since(t).Hours() / 24
		fmt.Println(math.Ceil(d))
	}

}
