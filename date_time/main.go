package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Now :", now)

	fmt.Println("Year :", now.Year())
	fmt.Println("Month :", now.Month())
	fmt.Println("Year :", now.Day())
	fmt.Println("Hour :", now.Hour())
	fmt.Println("Minute :", now.Minute())
	fmt.Println("Seconds :", now.Second())
	fmt.Println("Weekday :", now.Weekday())
	fmt.Println("YearDay :", now.YearDay())

	fmt.Println("--------------------------------------------")
	currentTime := time.Now()

	fmt.Println("Current Time in String: ", currentTime.String())

	fmt.Println("MM-DD-YYYY : ", currentTime.Format("1-02-2006"))

	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))

	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))

	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))

	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))

	fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))

	fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000"))

	fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))

	fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))

	fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))

	fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))

	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))

	fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))

	fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm"))
	fmt.Println("------------------- Subtract -------------------------")
	now = time.Now()
	fmt.Println("Today:", now)

	after := now.AddDate(-1, 0, 0)
	fmt.Println("Subtract 1 Year:", after)

	after = now.AddDate(0, -1, 0)
	fmt.Println("Subtract 1 Month:", after)

	after = now.AddDate(0, 0, -1)
	fmt.Println("Subtract 1 Day:", after)

	after = now.AddDate(-2, -2, -5)
	fmt.Println("Subtract multiple values:", after)

	after = now.Add(-10 * time.Minute)
	fmt.Println("Subtract 10 Minutes:", after)

	after = now.Add(-10 * time.Second)
	fmt.Println("Subtract 10 Second:", after)

	after = now.Add(-10 * time.Hour)
	fmt.Println("Subtract 10 Hour:", after)

	after = now.Add(-10 * time.Millisecond)
	fmt.Println("Subtract 10 Millisecond:", after)

	after = now.Add(-10 * time.Microsecond)
	fmt.Println("Subtract 10 Microsecond:", after)

	after = now.Add(-10 * time.Nanosecond)
	fmt.Println("Subtract 10 Nanosecond:", after)

	fmt.Println("------------------- Add -------------------------")

	now = time.Now()
	fmt.Println("\nToday:", now)

	after = now.AddDate(1, 0, 0)
	fmt.Println("\nAdd 1 Year:", after)

	after = now.AddDate(0, 1, 0)
	fmt.Println("\nAdd 1 Month:", after)

	after = now.AddDate(0, 0, 1)
	fmt.Println("\nAdd 1 Day:", after)

	after = now.AddDate(2, 2, 5)
	fmt.Println("\nAdd multiple values:", after)

	after = now.Add(10 * time.Minute)
	fmt.Println("\nAdd 10 Minutes:", after)

	after = now.Add(10 * time.Second)
	fmt.Println("\nAdd 10 Second:", after)

	after = now.Add(10 * time.Hour)
	fmt.Println("\nAdd 10 Hour:", after)

	after = now.Add(10 * time.Millisecond)
	fmt.Println("\nAdd 10 Millisecond:", after)

	after = now.Add(10 * time.Microsecond)
	fmt.Println("\nAdd 10 Microsecond:", after)

	after = now.Add(10 * time.Nanosecond)
	fmt.Println("\nAdd 10 Nanosecond:", after)

	fmt.Println("------------------- timezones -------------------------")

	t := time.Now()
	fmt.Println("Location : ", t.Location(), " Time : ", t) // local time

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Location : ", location, " Time : ", t.In(location)) // America/New_York

	loc, _ := time.LoadLocation("Asia/Shanghai")
	now = time.Now().In(loc)
	fmt.Println("Location : ", loc, " Time : ", now) // Asia/Shanghai

	loc, _ := time.LoadLocation("Asia/Shanghai")
	now = time.Now().In(loc)
	fmt.Println("Location : ", loc, " Time : ", now) // Asia/Shanghai
}
