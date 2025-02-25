package main

import (
	"fmt"
	"time"
)

type Alarm func(string)

func run(wait, resume chan bool, wt, bt int, alarm Alarm) {
  for true {
		countdown(wait, resume, "Work", wt)
		alarm("Break Time!")

		countdown(wait, resume, "Break", bt)
		alarm("Break Time!")
  }
}

func countdown(wait, resume chan bool, mode string, currentTime int) {
	printTimer(mode, currentTime)
	for currentTime > 0 {
		timer := time.NewTimer(1 * time.Second)

		select {
			case <- wait:
				printTimer(mode, currentTime)
				fmt.Print(" ||")
				<- resume
				printTimer(mode, currentTime)
			case <- timer.C:
				currentTime--
				printTimer(mode, currentTime)
		}
	}
}

func printTimer(mode string, timer int) {
	fmt.Printf("\033[2K\r%s: %02d:%02d", mode, int(timer / 60), timer % 60)
}
