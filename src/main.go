package main

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "strings"
  "time"
  "strconv"

  "golang.org/x/term"
)

type Alarm func(string)

func setup(reader *bufio.Reader) (string, string) {
  fmt.Print("work time (minutes): ")
  wtIn, _ := reader.ReadString('\n')
  workTime := strings.Replace(wtIn, "\n", "", -1)
  fmt.Print("break time (minutes): ")
  btIn, _ := reader.ReadString('\n')
  breakTime := strings.Replace(btIn, "\n", "", -1)

  return workTime, breakTime
}

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

func manage(wait, resume chan bool, reader *bufio.Reader) {
  old, _ := term.MakeRaw(int(os.Stdin.Fd()))
  defer term.Restore(int(os.Stdin.Fd()), old)
  w := false
  for true {
    var in []byte = make([]byte, 1)
    os.Stdin.Read(in)

    if in[0] == byte(' ') {
      w = !w
      if w {
        wait <- true
      } else {
        resume <- true
      }
    }
    if in[0] == byte('q') {
      return
    }
  }
}

func notifySend(message string) {
  cmd := exec.Command("notify-send", message)
  beep := exec.Command("ffplay", "assets/beep.mp3", "-nodisp")

  cmd.Run()
  go beep.Run()
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  wt, bt := setup(reader)

  wait := make(chan bool)
  resume := make(chan bool)

  go run(wait, resume, wt, bt, notifySend)
  manage(wait, resume, reader)
}
