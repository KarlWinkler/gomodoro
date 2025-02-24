package main

import (
  "bufio"
  "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  wt, bt := setup(reader)

  wait := make(chan bool)
  resume := make(chan bool)

  go run(wait, resume, wt, bt, notifySend)
  manage(wait, resume, reader)
}
