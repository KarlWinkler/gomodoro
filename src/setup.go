package main

import (
	"bufio"
	"fmt"
	"strings"
)

func setup(reader *bufio.Reader) (string, string) {
  fmt.Print("work time (minutes): ")
  wtIn, _ := reader.ReadString('\n')
  workTime := strings.Replace(wtIn, "\n", "", -1)
  fmt.Print("break time (minutes): ")
  btIn, _ := reader.ReadString('\n')
  breakTime := strings.Replace(btIn, "\n", "", -1)

  return workTime, breakTime
}