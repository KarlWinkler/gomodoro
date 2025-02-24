package main

import (
	"bufio"
	"os"

	"golang.org/x/term"
)

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