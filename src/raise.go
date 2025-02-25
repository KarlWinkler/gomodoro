package main

import (
	"fmt"
)

func raise(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}

	return false
}