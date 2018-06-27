package mydate

import "time"

func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")

}

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}
