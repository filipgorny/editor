package log

import (
	"fmt"
	"log"
)

var LogHistory = make([]string, 0)

func Log(args ...interface{}) {
	log.Println(fmt.Sprint(args...))
	LogHistory = append(LogHistory, fmt.Sprint(args...))
}

func Print() {
	for _, line := range LogHistory {
		fmt.Println(line)
	}
}
