package main

import (
	"fmt"
	"time"

	"github.com/uniplaces/carbon"
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

	fmt.Println("--------------------------------------------\n")
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
	fmt.Println("------------------- Subtract -------------------------\n")
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

	fmt.Println("------------------- Add -------------------------\n")

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

	fmt.Println("------------------- timezones ------------------------- \n")

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

	india_loc, _ := time.LoadLocation("Asia/Kolkata")
	now = time.Now().In(india_loc)
	fmt.Println("Location : ", india_loc, " Time : ", now) // Asia/Kolkata

	fmt.Println("----------- Convert specific UTC date time to PST, HST, MST and SGT ----------\n")

	t, err = time.Parse("2006 01 02 15 04", "2019 10 11 16 50")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	loc, err = time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)

	t = t.In(loc)
	fmt.Println(t.Format(time.RFC822))

	loc, err = time.LoadLocation("Singapore")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)

	t = t.In(loc)
	fmt.Println(t.Format(time.RFC822))

	loc, err = time.LoadLocation("US/Hawaii")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)

	t = t.In(loc)
	fmt.Println(t.Format(time.RFC822))

	loc, err = time.LoadLocation("EST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)

	t = t.In(loc)
	fmt.Println(t.Format(time.RFC822))

	loc, err = time.LoadLocation("MST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(loc)

	t = t.In(loc)
	fmt.Println(t.Format(time.RFC822))

	fmt.Println("------------------- Carbon date-time package -------------------------\n")

	fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())
	today, _ := carbon.NowInLocation("Asia/Kolkata")
	fmt.Printf("Right now in London is %s\n", today)

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Tomorrow is %s\n", carbon.Now().AddDay())
	fmt.Printf("Yesterday is %s\n", carbon.Now().SubDay())
	fmt.Printf("After 5 Days %s\n", carbon.Now().AddDays(5))
	fmt.Printf("Before 5 Days %s\n", carbon.Now().SubDays(5))

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Next Month is %s\n", carbon.Now().AddMonth())
	fmt.Printf("Last Month is %s\n", carbon.Now().SubMonth())

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Next week is %s\n", carbon.Now().AddWeek())
	fmt.Printf("Last week is %s\n", carbon.Now().SubWeek())

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Next Year %s\n", carbon.Now().AddYear())
	fmt.Printf("Last Year %s\n", carbon.Now().SubYear())
	fmt.Printf("After 5 Years %s\n", carbon.Now().AddYears(5))
	fmt.Printf("After 5 Years and 1 Month %s\n", carbon.Now().AddYears(5).AddMonths(2))
	fmt.Printf("Before 5 Years %s\n", carbon.Now().SubYears(5))

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Next Hour %s\n", carbon.Now().AddHour())
	fmt.Printf("Last Hour %s\n", carbon.Now().SubHour())
	fmt.Printf("After 5 Mins %s\n", carbon.Now().AddMinutes(5))
	fmt.Printf("Before 5 Mins %s\n", carbon.Now().SubMinutes(5))

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Weekday? %t\n", carbon.Now().IsWeekday())
	fmt.Printf("Weekend? %t\n", carbon.Now().IsWeekend())
	fmt.Printf("LeapYear? %t\n", carbon.Now().IsLeapYear())
	fmt.Printf("Past? %t\n", carbon.Now().IsPast())
	fmt.Printf("Future? %t\n", carbon.Now().IsFuture())

	fmt.Printf("\n#######################################\n")
	fmt.Printf("Start of day:   %s\n", today.StartOfDay())
	fmt.Printf("End of day: %s\n", today.EndOfDay())
	fmt.Printf("Start of month: %s\n", today.StartOfMonth())
	fmt.Printf("End of month:   %s\n", today.EndOfMonth())
	fmt.Printf("Start of year:  %s\n", today.StartOfYear())
	fmt.Printf("End of year:    %s\n", today.EndOfYear())
	fmt.Printf("Start of week:  %s\n", today.StartOfWeek())
	fmt.Printf("End of week:    %s\n", today.EndOfWeek())

}
