package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Print("Hello world")
	logrus.WithField("superId", "test").
		Println("Hello dorw")
}
