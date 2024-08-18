package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Wellcome to study time on golang")
	presentTime := time.Now() //get current time
	fmt.Println("presentTime:", presentTime)
	fmt.Println("Format the presentTime", presentTime.Format("01-02-2006, monday")) //format time

	timeString := "2024-08-15 14:00:00"
	timeFormat := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(timeFormat, timeString)
	if err != nil {
		fmt.Println("Parsing error", err)
		return
	}
	fmt.Println("parsedTime:", parsedTime)

	//Time artihemtic
	currentTime := time.Now()

	//add 5 hours
	futureTime := currentTime.Add(5 * time.Hour)
	fmt.Println("Time 5 hour later", futureTime)

	//sub 2 days
	pastTime := currentTime.AddDate(0, 0, -2)
	fmt.Println("Time 2 days ago:", pastTime)

	//different b/n times
	duration := futureTime.Sub((currentTime))
	fmt.Println("difference", duration)

	//working with durations
	tenDays, _ := time.ParseDuration("10d")
	fmt.Println("10 days later", futureTime.Add(tenDays))

	durations := 2*time.Hour + 30*time.Minute
	fmt.Println("durations", durations)

	//sleep after durations
	// time.Sleep(durations)
	// fmt.Println("woke up after sleeping for", durations)

}
