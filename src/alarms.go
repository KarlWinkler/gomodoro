package main

import (
  "os/exec"
)

func notifySend(message string) {
  cmd := exec.Command("notify-send", message)
  beep := exec.Command("ffplay", "assets/beep.mp3", "-nodisp")

  cmd.Run()
  go beep.Run()
}
