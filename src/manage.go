package main

import (
	"bufio"
	"os"

	"golang.org/x/term"
)

type State struct {
  Wait chan bool
  Resume chan bool
  Paused bool
  Quit bool
}

func manage(wait, resume chan bool, reader *bufio.Reader) {
  old, _ := term.MakeRaw(int(os.Stdin.Fd()))
  defer term.Restore(int(os.Stdin.Fd()), old)
  
  state := State{Wait: wait, Resume: resume, Paused: false, Quit: false}

  var in []byte = make([]byte, 1)

  for true {
    os.Stdin.Read(in)

    handle_input(&state, in)

    if state.Quit { return }
  }

}

func handle_input(state *State, in []byte) {
  switch in[0] {
    case ' ':
      state.Paused = !state.Paused
      pause(state)
    case 'q':
      state.Quit = true
  }
}

func pause(state *State) {
  if state.Paused {
    state.Wait <- true
  } else {
    state.Resume <- true
  }
}
