package main

import (
	"fmt"
	"time"
)

func main() {
	// time_layout := "2006-01-02 15:04:05"
	// a, err := time.Parse(time_layout, "2020-06-10 20:39:55")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(a.Format(time.RFC3339))
	// b := time.Now().Format(time.RFC3339)
	// fmt.Println(b)
	a := time.Now().Unix()
	fmt.Println(a)
	// 1591886530
	fmt.Println(time.Unix(a, 0))
	// 2020-06-11 10:42:10 -0400 EDT
	fmt.Printf("%T\n", a)
	fmt.Println(time.Time{}.Unix())
	var UTC *time.Location
	UTC, err := time.LoadLocation("")
	if err != nil {
		fmt.Println(err)
	}
	
	b, err := time.ParseInLocation("2006-01-02", "2020-06-10", UTC)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b.Unix())
	fmt.Println(time.Unix(b.Unix(), 0))
	fmt.Printf("%T", time.Time{}.Unix())

	fmt.Println()
	fmt.Println(time.Unix(1591718419,0).UTC())
}


// FROM POSTGRES
// "2006-01-02 15:04:05"

// We need to parse, then format

// TO POSTGRES
// time.Now().Format(time.RFC3339), format as 