package helper

import (
	"log"
	"time"
)

func TimeHostNow() time.Time {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Printf("Error get time, cause:%+v\n", err)
	}
	now := time.Now()
	timeInLoc := now.In(location)
	return timeInLoc
}
