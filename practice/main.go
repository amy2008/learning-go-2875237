package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("I practised the lib time at ", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at ", t)
	fmt.Println(t.Format(time.Kitchen))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format(time.DateTime))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009") //Parse() is used to parse the string to a Time.time object
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)

	var myTime string
	myTime = os.Args[1]
	d, err := time.Parse("22:04", myTime)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}

	t := time.Now()
	fmt.Println(t.Format())

}
