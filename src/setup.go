package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

)

func setup(reader *bufio.Reader) (int, int) {
	wt := _get_time_with_retry("work time: ", reader)
	bt := _get_time_with_retry("break time: ", reader)

  return wt, bt
}

// work time (minutes): 

func _get_time_with_retry(prompt string, reader *bufio.Reader) int {
	for true {
		time, err := _get_time(prompt, reader)

		if err == nil {
			return time
		} else {
			fmt.Println("Not a number! Please try again,")
		}
	}
	panic("Escaped Loop in { setup.go::_get_time_with_retry }")
}

func _get_time(prompt string, reader *bufio.Reader) (int, error) {
	fmt.Print(prompt)
	in, _ := reader.ReadString('\n')
	in = strings.Replace(in, "\n", "", -1)

	time, err := strconv.Atoi(in)

	return time, err
}
