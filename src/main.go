package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
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
  for true {
    break
  }
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  wt, bt := setup(reader)
  run(wt, bt)
}
