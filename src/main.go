package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "time"
  "strconv"
)

func setup(reader *bufio.Reader) (string, string) {
  fmt.Println("work time (minutes): ")
  wtIn, _ := reader.ReadString('\n')
  workTime := strings.Replace(wtIn, "\n", "", -1)
  fmt.Println("break time (minutes): ")
  btIn, _ := reader.ReadString('\n')
  breakTime := strings.Replace(btIn, "\n", "", -1)

  return workTime, breakTime
}

func run(wt, bt string) {
  wtInt, wtErr := strconv.Atoi(wt)
  btInt, btErr := strconv.Atoi(bt)

  if wtErr != nil || btErr != nil {
    fmt.Println("one of your inputs was invalid")
    return
  }

  workTimer := wtInt * 60
  breakTimer := btInt * 60
  for true {
    for workTimer > 0 {
      timer := time.NewTimer(1 * time.Second)

      <- timer.C
      fmt.Printf("\r%02d:%02d", int(workTimer / 60), workTimer % 60)
      workTimer--
    }
    for breakTimer > 0 {
      timer := time.NewTimer(1 * time.Second)

      <- timer.C
      fmt.Printf("\r%02d:%02d", int(breakTimer / 60), breakTimer % 60)
      breakTimer--
    }
    break
  }
  fmt.Println(workTimer, breakTimer)
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  wt, bt := setup(reader)
  run(wt, bt)
}
