package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := GetWeek()
	fmt.Println("s", s)

}

func GetWeek() (week string) {

	nowTime := time.Now()
	// nowDay := nowTime.Format("2006-01-02 15:04:05")
	str := []string{}
	for i := 1; i < 8; i++ {
		t := nowTime.Unix() - int64(60*24*60*i)

		timeFm := time.Unix(t, 0).Format("2006-1-2")
		str = append(str, "'"+timeFm+"'")
	}
	week = strings.Join(str, ",")

	fmt.Println("week:", week)
	return week
}
