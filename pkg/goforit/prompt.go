package goforit

import (
	"fmt"
	"time"
)

func Prompt() {
	for {
		fmt.Println("Hello world")
		time.Sleep(1*time.Second)
	}

}
