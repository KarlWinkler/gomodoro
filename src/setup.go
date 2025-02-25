package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"errors"
)

const (
	SECOND = 1
	MINUTE = 60
	HOUR = 3600
)

func setup(reader *bufio.Reader) (int, int) {
	wt := get_time_with_retry("work time [minutes]: ", reader)
	bt := get_time_with_retry("break time [minutes]: ", reader)

  return wt, bt
}

func get_time_with_retry(prompt string, reader *bufio.Reader) int {
	for true {
		// promt the user
		in := _prompt_user(prompt, reader)
		number, unit, err := get_inputs(in)
		if raise(err) { continue }

		// parse the time
		time, err := get_time(number)
		if raise(err) { continue }

		// parse the magnitude
		var magnitude int
		if unit != "" {
			magnitude, err = get_magnitude(unit)
			if raise(err) { continue }
		} else {
			magnitude = MINUTE
		}

		return time * magnitude
	}
	panic("Escaped Loop in { setup.go::get_time_with_retry }")
}

func _prompt_user(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	in, _ := reader.ReadString('\n')
	in = strings.Replace(in, "\n", "", -1)

	return in
}

func get_inputs(in string) (string, string, error) {
	inputs := strings.Split(in, " ")
	if len(inputs) > 1 {
		return inputs[0], inputs[1], nil
	} else if len(inputs) > 0 {
		return inputs[0], "", nil
	} else {
		return "", "", errors.New("You must supply an input")
	}
}

func get_time(number string) (int, error) {
	time, err := strconv.Atoi(number)

	if err != nil {
		return -1, errors.New("Not a number! Please try again")
	}

	return time, nil
}

func get_magnitude(unit string) (int, error) {
	units := map [string]int{"s": SECOND, "m": MINUTE, "h": HOUR}
	unit = format_unit(unit)

	magnitude, ok := units[unit]
	if ok {
		return magnitude, nil
	} else {
		return -1, errors.New("You must give a valid unit [second, minute, hour]")
	}
}

func format_unit(unit string) string {
	unit = strings.ToLower(unit)
	unit = unit[:1]

	return unit
}
