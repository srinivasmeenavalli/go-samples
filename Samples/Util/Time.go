package main

import (
	"fmt"
	"time"
)

/*
Go offers extensive support for times and durations
*/
func main() {
	p := fmt.Println

	now := time.Now()
	p(now)
	/*
		You can build a time struct by providing the
		year, month, day, etc.
		 Times are always associated
		 with a Location, i.e. time zone
	*/
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("Date: ", then)
	p("Year: ", then.Year())
	p("Month: ", then.Month())
	p("Day: ", then.Day())
	p("Hour: ", then.Hour())
	p("Minute: ", then.Minute())
	p("Second: ", then.Second())
	p("NanoSecond: ", then.Nanosecond())
	p("Location: ", then.Location())

	p("Weekday: ", then.Weekday())
	/*
		These methods compare two times,
		testing if the first occurs before,
		after, or at the same time as the second,
		 respectively
	*/
	p("Before current time ? : ", then.Before(now))
	p("After current time ? :", then.After(now))
	p("Equal to current time ? :", then.Equal(now))
	/*
		The Sub methods returns
		a Duration representing the interval between two times
	*/
	diff := now.Sub(then)
	p(diff)

	p("Hours:", diff.Hours())
	p("Minutes:", diff.Minutes())
	p("Seconds:", diff.Seconds())
	p("Nano Seconds:", diff.Nanoseconds())

	p("Present Time:", then.Add(diff))
	p("Past time travel:", then.Add(-diff))
}
