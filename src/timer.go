package main

import (
	"fmt"
	"strconv"
	"time"
)

type Alarm func(string)

func run(wait, resume chan bool, wt, bt string, alarm Alarm) {
  wtInt, wtErr := strconv.Atoi(wt)
  btInt, btErr := strconv.Atoi(bt)

  if wtErr != nil || btErr != nil {
    fmt.Println("one of your inputs was invalid")
    return
  }

  for true {
    workTimer := wtInt * 60
    breakTimer := btInt * 60

    fmt.Printf("\033[2K\rWork: %02d:%02d", int(workTimer / 60), workTimer % 60)
    for workTimer > 0 {
      timer := time.NewTimer(1 * time.Second)

      select {
        case <- wait:
          fmt.Printf("\033[2K\rWork: %02d:%02d ||", int(workTimer / 60), workTimer % 60)
          <- resume
          fmt.Printf("\033[2K\rWork: %02d:%02d", int(workTimer / 60), workTimer % 60)
        case <- timer.C:
          workTimer--
          fmt.Printf("\033[2K\rWork: %02d:%02d", int(workTimer / 60), workTimer % 60)
      }
    }
    alarm("Break Time!")

    fmt.Printf("\033[2K\rBreak: %02d:%02d", int(breakTimer / 60), breakTimer % 60)
    for breakTimer > 0 {
      timer := time.NewTimer(1 * time.Second)

      select {
        case <- wait:
          fmt.Printf("\033[2K\rBreak: %02d:%02d ||", int(breakTimer / 60), breakTimer % 60)
          <- resume
          fmt.Printf("\033[2K\rBreak: %02d:%02d", int(breakTimer / 60), breakTimer % 60)
        case <- timer.C:
          breakTimer--
          fmt.Printf("\033[2K\rBreak: %02d:%02d", int(breakTimer / 60), breakTimer % 60)
      }
    }
    alarm("Work Time!")
  }
}
