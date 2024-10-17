package main

import "fmt"

type Rsvp struct {
  Name, Mail, Phone string
  WillAttend bool
}

var responses = make([]*Rsvp, 0, 10)

func main() {
  fmt.Println("TODO: add some features")
}
